package votesync

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/wait"
	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

//VoteSync defines fields used in VoteSync
type VoteSync struct {
	service                   iotex.AuthedClient
	iotexAPI                  iotexapi.APIServiceClient
	vpsContract               iotex.Contract
	brokerContract            iotex.Contract
	clerkContract             iotex.Contract
	discordBotToken           string
	discordChannelID          string
	discordMsg                string
	discordReminder           string
	discordReminded           bool
	carrier                   carrier.Carrier
	lastViewHeight            uint64
	lastViewTimestamp         time.Time
	lastUpdateHeight          uint64
	lastBrokerUpdateHeight    uint64
	lastClerkUpdateHeight     uint64
	lastUpdateTimestamp       time.Time
	timeInternal              time.Duration
	paginationSize            uint8
	brokerPaginationSize      uint8
	lastNativeEphoch          uint64
	tempLastNativeEphoch      uint64
	terminate                 chan bool
	dardanellesHeight         uint64
	nativeCommittee           *committee.NativeCommittee
	nativeCommitteeInitHeight uint64
}

//Config defines the configs for VoteSync
type Config struct {
	GravityChainAPIs          []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval  time.Duration `yaml:"gravityChainTimeInterval"`
	OperatorPrivateKey        string        `yaml:"operatorPrivateKey"`
	IoTeXAPI                  string        `yaml:"ioTeXAPI"`
	RegisterContractAddress   string        `yaml:"registerContractAddress"`
	StakingContractAddress    string        `yaml:"stakingContractAddress"`
	PaginationSize            uint8         `yaml:"paginationSize"`
	BrokerPaginationSize      uint8         `yaml:"brokerPaginationSize"`
	VitaContractAddress       string        `yaml:"vitaContractAddress"`
	DiscordBotToken           string        `yaml:"discordBotToken"`
	DiscordChannelID          string        `yaml:"discordChannelID"`
	DiscordMsg                string        `yaml:"discordMsg"`
	DiscordReminder           string        `yaml:"discordReminder"`
	DardaenllesHeight         uint64        `yaml:"dardanellesHeight"`
	NativeCommitteeInitHeight uint64        `yaml:"nativeCommitteeInitHeight"`
}

//WeightedVote defines voter and votes for weighted vote
type WeightedVote struct {
	Voter []byte
	Votes *big.Int
}

//toIoAddress converts ethAddress to ioAddress
func toIoAddress(addr common.Address) (address.Address, error) {
	pkhash, err := hexutil.Decode(addr.String())
	if err != nil {
		return nil, err
	}
	return address.FromBytes(pkhash)
}

//NewVoteSync instantiates new VoteSync
func NewVoteSync(cfg Config, nativeCommittee *committee.NativeCommittee) (*VoteSync, error) {
	ctx := context.Background()
	carrier, err := carrier.NewEthereumVoteCarrier(
		12,
		1*time.Minute,
		cfg.GravityChainAPIs,
		common.HexToAddress(cfg.RegisterContractAddress),
		common.HexToAddress(cfg.StakingContractAddress),
	)
	if err != nil {
		return nil, err
	}

	conn, err := iotex.NewDefaultGRPCConn(cfg.IoTeXAPI)
	if err != nil {
		return nil, err
	}
	operatorPrivateKey, err := crypto.HexStringToPrivateKey(cfg.OperatorPrivateKey)
	if err != nil {
		return nil, err
	}
	operatorAccount, err := account.PrivateKeyToAccount(operatorPrivateKey)
	if err != nil {
		return nil, err
	}
	iotexAPI := iotexapi.NewAPIServiceClient(conn)
	cli := iotex.NewAuthedClient(iotexAPI, operatorAccount)

	vitaABI, err := abi.JSON(strings.NewReader(contract.VitaABI))
	if err != nil {
		return nil, err
	}
	vitaContractAddress, err := address.FromString(cfg.VitaContractAddress)
	if err != nil {
		return nil, err
	}
	vitaContract := cli.Contract(vitaContractAddress, vitaABI)

	var addr common.Address
	d, err := vitaContract.Read("vps").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&addr); err != nil {
		return nil, err
	}
	vpsContractAddress, err := toIoAddress(addr)
	if err != nil {
		return nil, err
	}

	d, err = vitaContract.Read("donationPoolAddress").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&addr); err != nil {
		return nil, err
	}
	brokerContractAddress, err := toIoAddress(addr)
	if err != nil {
		return nil, err
	}

	d, err = vitaContract.Read("rewardPoolAddress").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&addr); err != nil {
		return nil, err
	}
	clerkContractAddress, err := toIoAddress(addr)
	if err != nil {
		return nil, err
	}
	zap.L().Info("vote contracts.", zap.String("brokerContract", brokerContractAddress.String()), zap.String("clerkContract", clerkContractAddress.String()))

	vpsABI, err := abi.JSON(strings.NewReader(contract.RotatableVPSABI))
	if err != nil {
		return nil, err
	}
	vpsContract := cli.Contract(vpsContractAddress, vpsABI)

	lastUpdateHeight := new(big.Int)
	d, err = vpsContract.Read("viewID").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&lastUpdateHeight); err != nil {
		return nil, err
	}
	lastUpdateTimestamp, err := carrier.BlockTimestamp(lastUpdateHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastViewHeight := new(big.Int)
	d, err = vpsContract.Read("inactiveViewID").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&lastViewHeight); err != nil {
		return nil, err
	}
	lastViewTimestamp, err := carrier.BlockTimestamp(lastViewHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastBrokerUpdateHeight := new(big.Int)
	d, err = vitaContract.Read("lastDonationPoolClaimViewID").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&lastBrokerUpdateHeight); err != nil {
		return nil, err
	}

	lastClerkUpdateHeight := new(big.Int)
	d, err = vitaContract.Read("lastRewardPoolClaimViewID").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&lastClerkUpdateHeight); err != nil {
		return nil, err
	}

	brokerABI, err := abi.JSON(strings.NewReader(contract.BrokerABI))
	if err != nil {
		return nil, err
	}
	brokerContract := cli.Contract(brokerContractAddress, brokerABI)

	clerkABI, err := abi.JSON(strings.NewReader(contract.ClerkABI))
	if err != nil {
		return nil, err
	}
	clerkContract := cli.Contract(clerkContractAddress, clerkABI)

	return &VoteSync{
		carrier:                carrier,
		vpsContract:            vpsContract,
		brokerContract:         brokerContract,
		clerkContract:          clerkContract,
		service:                cli,
		iotexAPI:               iotexAPI,
		timeInternal:           cfg.GravityChainTimeInterval,
		paginationSize:         cfg.PaginationSize,
		brokerPaginationSize:   cfg.BrokerPaginationSize,
		lastViewHeight:         lastViewHeight.Uint64(),
		lastViewTimestamp:      lastViewTimestamp,
		lastUpdateHeight:       lastUpdateHeight.Uint64(),
		lastUpdateTimestamp:    lastUpdateTimestamp,
		lastBrokerUpdateHeight: lastBrokerUpdateHeight.Uint64(),
		lastClerkUpdateHeight:  lastClerkUpdateHeight.Uint64(),
		terminate:              make(chan bool),
		discordBotToken:        cfg.DiscordBotToken,
		discordChannelID:       cfg.DiscordChannelID,
		discordMsg:             cfg.DiscordMsg,
		discordReminder:        cfg.DiscordReminder,
		dardanellesHeight:      cfg.DardaenllesHeight,
		nativeCommittee:        nativeCommittee,
	}, nil
}

//Start starts voteSync
func (vc *VoteSync) Start(ctx context.Context) {
	tipChan := make(chan uint64)
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.",
		zap.Uint64("lastUpdateHeight", vc.lastUpdateHeight),
		zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight),
		zap.Uint64("lastClerkUpdateHeight", vc.lastClerkUpdateHeight),
		zap.Uint64("lastViewID", vc.lastViewHeight),
	)
	go func() {
		for {
			select {
			case <-vc.terminate:
				vc.terminate <- true
				return
			case tip := <-tipChan:
				blockTime, err := vc.carrier.BlockTimestamp(tip)
				if err != nil {
					zap.L().Error("failed to get block time", zap.Error(err))
					continue
				}
				if blockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal)) {
					if err := vc.sync(vc.lastViewHeight, tip, vc.lastViewTimestamp, blockTime); err != nil {
						zap.L().Error("failed to sync votes", zap.Error(err))
						continue
					}
					if err := vc.sendDiscordMsg(vc.discordMsg); err != nil {
						zap.L().Error("failed to send discord msg", zap.Error(err))
					}
					vc.discordReminded = false
				}
				if blockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal*24/25)) && !vc.discordReminded {
					if err := vc.sendDiscordMsg(vc.discordReminder); err != nil {
						zap.L().Error("failed to send discord reminder", zap.Error(err))
					}
					vc.discordReminded = true
				}

				if vc.lastUpdateHeight > vc.lastBrokerUpdateHeight {
					if err := vc.settle(vc.lastUpdateHeight); err != nil {
						zap.L().Error("failed to settle broker", zap.Error(err))
					}
				}

				if vc.lastUpdateHeight > vc.lastClerkUpdateHeight {
					if err := vc.claimForClerk(); err != nil {
						zap.L().Error("failed to claim for clerk", zap.Error(err))
					}
				}
			case err := <-errChan:
				zap.L().Error("something goes wrong", zap.Error(err))
			}
		}
	}()
	vc.carrier.SubscribeNewBlock(tipChan, errChan, vc.terminate)
}

//Stop stops voteSync
func (vc *VoteSync) Stop(ctx context.Context) {
	vc.terminate <- true
	vc.carrier.Close()
	return
}

func (vc *VoteSync) brokerReset() error {
	caller := vc.brokerContract.Execute("reset").SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000)
	return wait.Wait(context.Background(), caller)
}

func (vc *VoteSync) brokerNextBidToSettle() (uint64, error) {
	nextBidToSettle := new(big.Int)
	d, err := vc.brokerContract.Read("nextBidToSettle").Call(context.Background())
	if err != nil {
		return 0, err
	}
	if err := d.Unmarshal(&nextBidToSettle); err != nil {
		return 0, err
	}
	return nextBidToSettle.Uint64(), nil
}

func (vc *VoteSync) brokerSettle() error {
	oldStart := uint64(0)
	for {
		caller := vc.brokerContract.Execute(
			"settle", big.NewInt(0).SetUint64(uint64(vc.brokerPaginationSize))).
			SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000)
		if err := wait.Wait(context.Background(), caller); err != nil {
			return err
		}

		newStart, err := vc.brokerNextBidToSettle()
		if err != nil {
			return err
		}
		if oldStart == newStart {
			return nil
		}
		oldStart = newStart
	}
}

func (vc *VoteSync) settle(h uint64) error {
	l := zap.L().With(zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight))
	l.Info("Start broker settle process.")
	// settle broker
	if err := vc.brokerSettle(); err != nil {
		return errors.Wrap(err, "broker settle error")
	}
	l.Info("Finished broker settle.")
	if err := vc.brokerReset(); err != nil {
		return errors.Wrap(err, "broker reset error")
	}
	vc.lastBrokerUpdateHeight = h
	l.Info("Finished broker reset.", zap.Uint64("brokerUpdatedHeight", h))
	return nil
}

func (vc *VoteSync) claimForClerk() error {
	zap.L().Info("Start clerk claim process.", zap.Uint64("lastClerkUpdateHeight", vc.lastClerkUpdateHeight))
	caller := vc.clerkContract.Execute("claim").SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000)
	if err := wait.Wait(context.Background(), caller); err != nil {
		return err
	}
	vc.lastClerkUpdateHeight = vc.lastUpdateHeight
	zap.L().Info("Finished clerk.", zap.Uint64("cleerkUpdatedHeight", vc.lastUpdateHeight))
	return nil
}

func (vc *VoteSync) updateVotingPowers(addrs []common.Address, weights []*big.Int) error {
	caller := vc.vpsContract.Execute("updateVotingPowers", addrs, weights).
		SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(7000000)
	return wait.Wait(context.Background(), caller)
}

func (vc *VoteSync) sync(prevHeight, currHeight uint64, prevTs, currTs time.Time) error {
	zap.L().Info("Start VoteSyncing.", zap.Uint64("lastViewID", prevHeight), zap.Uint64("nextViewID", currHeight))
	ret, err := vc.fetchVotesUpdate(prevHeight, currHeight, prevTs, currTs)
	if err != nil {
		return errors.Wrap(err, "fetch vote error")
	}
	zap.L().Info("Need to sync.", zap.Int("numVoter", len(ret)))

	var (
		addrs   []common.Address
		weights []*big.Int
		reqNum  int
	)
	for _, vote := range ret {
		addrs = append(addrs, common.BytesToAddress(vote.Voter))
		weights = append(weights, vote.Votes)

		if len(addrs)%int(vc.paginationSize/4) == 0 {
			if err := vc.updateVotingPowers(addrs, weights); err != nil {
				return errors.Wrap(err, fmt.Sprintf("update vote error, reqNum:%d", reqNum))
			}
			reqNum++
			addrs = []common.Address{}
			weights = []*big.Int{}
		}
	}
	if len(addrs) > 0 {
		if err := vc.updateVotingPowers(addrs, weights); err != nil {
			return errors.Wrap(err, fmt.Sprintf("update vote error, reqNum:%d", reqNum))
		}
	}
	caller := vc.vpsContract.Execute("rotate", new(big.Int).SetUint64(currHeight)).
		SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(4000000)
	if err := wait.Wait(context.Background(), caller); err != nil {
		return errors.Wrap(err, "failed to execute rotate")
	}

	vc.lastViewHeight = vc.lastUpdateHeight
	vc.lastViewTimestamp = vc.lastUpdateTimestamp
	vc.lastUpdateHeight = currHeight
	vc.lastUpdateTimestamp = currTs
	vc.lastNativeEphoch = vc.tempLastNativeEphoch
	zap.L().Info("Successfully synced votes.", zap.Uint64("lastViewID", vc.lastViewHeight), zap.Uint64("viewID", currHeight))
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(prevHeight, currHeight uint64, prevTs, currTs time.Time) ([]*WeightedVote, error) {
	// prevHeight == 0, only run at first 2 time. get all votes from currHeight
	if prevHeight == 0 {
		curr, err := vc.retryFetchBucketsByHeight(currHeight)
		if err != nil {
			return nil, err
		}

		n := calWeightedVotes(curr, currTs)
		var ret []*WeightedVote
		for _, nv := range n {
			ret = append(ret, nv)
		}
		return ret, nil
	}

	prev, err := vc.retryFetchBucketsByHeight(prevHeight)
	if err != nil {
		return nil, err
	}

	curr, err := vc.retryFetchBucketsByHeight(currHeight)
	if err != nil {
		return nil, err
	}

	if prevNative, currNative, err := vc.fetchNativeBuckets(prevTs, currTs); err != nil {
		// log on error first
		zap.L().Error("failed to fetch native buckets", zap.Error(err))
	} else {
		prev = append(prev, prevNative...)
		curr = append(curr, currNative...)
	}

	p := calWeightedVotes(prev, prevTs)
	n := calWeightedVotes(curr, currTs)

	var ret []*WeightedVote
	// check for all voters in old view
	// if they don't exist in new view map, append 0 value for them
	// if they do exisit in new view map, append only if the vote value is different
	for k, pv := range p {
		nv, ok := n[k]
		if !ok {
			ret = append(ret, &WeightedVote{
				Voter: pv.Voter,
				Votes: big.NewInt(0),
			})
		} else {
			if nv.Votes.Cmp(pv.Votes) != 0 {
				ret = append(ret, nv)
			}
		}
	}
	// check for all new voters in new view
	// if they don't exist in old view map, append
	for k, nv := range n {
		if _, ok := p[k]; !ok {
			ret = append(ret, nv)
		}
	}
	return ret, nil
}

func (vc *VoteSync) retryFetchBucketsByHeight(h uint64) ([]*types.Bucket, error) {
	var (
		ret []*types.Bucket
		err error
	)
	nerr := backoff.Retry(func() error {
		ret, err = vc.fetchBucketsByHeight(h)
		return err
	}, backoff.NewExponentialBackOff())
	if nerr != nil {
		zap.L().Error(
			"failed to fetch vote result by height",
			zap.Error(nerr),
			zap.Uint64("height", h),
		)
	}
	return ret, nerr
}

func (vc *VoteSync) fetchBucketsByHeight(h uint64) ([]*types.Bucket, error) {
	var allVotes []*types.Bucket
	idx := big.NewInt(0)
	for {
		var (
			votes []*types.Bucket
			err   error
		)
		if idx, votes, err = vc.carrier.Buckets(h, idx, vc.paginationSize); err != nil {
			return nil, err
		}
		allVotes = append(allVotes, votes...)
		if len(votes) < int(vc.paginationSize) {
			break
		}
	}
	return allVotes, nil
}

func (vc *VoteSync) fetchNativeBuckets(prevTs, currTs time.Time) ([]*types.Bucket, []*types.Bucket, error) {
	ctx := context.Background()
	resp, err := vc.iotexAPI.GetChainMeta(ctx, &iotexapi.GetChainMetaRequest{})
	if err != nil {
		return nil, nil, err
	}
	h := resp.ChainMeta.GetHeight()
	var retryCounter int
	// get a iotex block which timestamp is after currTs
	for {
		time.Sleep(15 * time.Second)
		bt, err := vc.getIotexBlockTime(ctx, h)
		if err != nil {
			if retryCounter > 20 {
				return nil, nil, errors.Wrap(err, "already tried many times to get block on iotex")
			}
			retryCounter++
			continue
		}
		if bt.After(currTs) {
			break
		}
		h++
	}

	findEpoch := func(ctx context.Context, en uint64, t time.Time) (uint64, error) {
		retryCounter := 0
		for {
			h := vc.getEpochHeight(en)
			bt, err := vc.getIotexBlockTime(ctx, h)
			if err != nil {
				if retryCounter > 10 {
					return en, errors.Wrap(err, "already tried many times to get block on iotex")
				}
				retryCounter++
				time.Sleep(5 * time.Second)
				continue
			}
			en--
			if bt.Before(t) {
				break
			}
		}
		return en, nil
	}

	// get epoch number of the block we got
	en := vc.getEpochNum(h)
	// find a epoch start block which is the first one timestamp before currTs
	currEn, err := findEpoch(ctx, en, currTs)
	if err != nil {
		return nil, nil, err
	}

	// check if there is a cached prevEpochNum
	prevEn := vc.lastNativeEphoch
	if prevEn == 0 {
		// no cached prevEpochNum, find a epoch start block which is the first one timestamp before prevTs
		prevEn, err = findEpoch(ctx, en, prevTs)
		if err != nil {
			return nil, nil, err
		}
	}

	// get currBuckets and prevBuckets
	pbkts, err := vc.getNativeElectionBuckets(ctx, prevEn)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to fetch prev epoch buckets, epoch number %v", prevEn)
	}

	cbkts, err := vc.getNativeElectionBuckets(ctx, currEn)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to fetch curr epoch buckets, epoch number %v", currEn)
	}
	vc.tempLastNativeEphoch = currEn
	return pbkts, cbkts, nil
}

func (vc *VoteSync) getIotexBlockTime(ctx context.Context, h uint64) (time.Time, error) {
	resp, err := vc.iotexAPI.GetBlockMetas(ctx, &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: h, Count: 1,
			},
		},
	})
	if err != nil {
		return time.Now(), errors.Wrapf(err, "failed to fetch block meta %v", h)
	}
	bms := resp.GetBlkMetas()
	if len(bms) != 1 {
		return time.Now(), errors.Wrapf(err, "asked 1 block, but got none-1 value %v", h)
	}
	ts := bms[0].GetTimestamp()
	bt, err := ptypes.Timestamp(ts)
	if err != nil {
		return time.Now(), errors.Wrapf(err, "failed to parse timestamp in blockmeta %v", h)
	}
	return bt, nil
}

func (vc *VoteSync) getNativeElectionBuckets(ctx context.Context, en uint64) ([]*types.Bucket, error) {
	resp, err := vc.iotexAPI.GetElectionBuckets(ctx, &iotexapi.GetElectionBucketsRequest{EpochNum: en})
	if err != nil {
		return nil, errors.Wrap(err, "failed to GetElectionBuckets")
	}
	bkts := make([]*types.Bucket, 0, len(resp.GetBuckets()))
	for _, bkt := range resp.GetBuckets() {
		b, err := bucketfromIotexProtoMsg(bkt)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert iotextypes bucket to election bucket")
		}
		bkts = append(bkts, b)
	}
	if vc.nativeCommittee != nil {
		epochHeight := vc.getEpochHeight(en)
		if epochHeight >= vc.nativeCommitteeInitHeight {
			zap.L().Info("get buckets from native committee", zap.Uint64("epoch height", epochHeight))
			_, buckets, err := vc.nativeCommittee.DataByHeight(epochHeight)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get buckets from native committee")
			}
			if buckets != nil && len(buckets) > 0 {
				bkts = append(bkts, buckets...)
			}
		}
	}
	return bkts, nil
}

func (vc *VoteSync) sendDiscordMsg(msg string) error {
	if vc.discordBotToken == "" || msg == "" {
		return nil
	}

	dg, err := discordgo.New("Bot " + vc.discordBotToken)
	if err != nil {
		return err
	}
	if err := dg.Open(); err != nil {
		return err
	}
	defer dg.Close()

	_, err = dg.ChannelMessageSend(vc.discordChannelID, msg)
	return err
}

func calWeightedVotes(curr []*types.Bucket, currTs time.Time) map[string]*WeightedVote {
	n := make(map[string]*WeightedVote)
	for _, v := range curr {
		vs := types.CalcWeightedVotes(v, currTs)
		wv, ok := n[hex.EncodeToString(v.Voter())]
		if ok {
			wv.Votes.Add(wv.Votes, vs)
			continue
		}
		n[hex.EncodeToString(v.Voter())] = &WeightedVote{
			Voter: v.Voter(),
			Votes: vs,
		}
	}
	return n
}

func (vc *VoteSync) getEpochNum(height uint64) uint64 {
	if height == 0 {
		return 0
	}
	if vc.dardanellesHeight == 0 || height <= vc.dardanellesHeight {
		return (height-1)/360 + 1
	}
	dardanellesEpoch := vc.getEpochNum(vc.dardanellesHeight)
	dardanellesEpochHeight := vc.getEpochHeight(dardanellesEpoch)
	return dardanellesEpoch + (height-dardanellesEpochHeight)/720
}

func (vc *VoteSync) getEpochHeight(epochNum uint64) uint64 {
	if epochNum == 0 {
		return 0
	}
	dardanellesEpoch := vc.getEpochNum(vc.dardanellesHeight)
	if vc.dardanellesHeight == 0 || epochNum <= dardanellesEpoch {
		return (epochNum-1)*360 + 1
	}
	dardanellesEpochHeight := vc.getEpochHeight(dardanellesEpoch)
	return dardanellesEpochHeight + (epochNum-dardanellesEpoch)*720
}

func bucketfromIotexProtoMsg(vPb *iotextypes.ElectionBucket) (*types.Bucket, error) {
	voter := make([]byte, len(vPb.GetVoter()))
	copy(voter, vPb.GetVoter())
	candidate := make([]byte, len(vPb.GetCandidate()))
	copy(candidate, vPb.GetCandidate())
	amount := big.NewInt(0).SetBytes(vPb.GetAmount())
	startTime, err := ptypes.Timestamp(vPb.GetStartTime())
	if err != nil {
		return nil, err
	}
	duration, err := ptypes.Duration(vPb.GetDuration())
	if err != nil {
		return nil, err
	}
	if duration < 0 {
		return nil, errors.Errorf("duration %s cannot be negative", duration)
	}

	return types.NewBucket(startTime, duration, amount, voter, candidate, vPb.GetDecay())
}
