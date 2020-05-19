package votesync

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/wait"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

const _viewIDOffsite = 10000000

const (
	// PollProtocolID is ID of poll protocol
	PollProtocolID      = "poll"
	protocolID          = "staking"
	readBucketsLimit    = 30000
	readCandidatesLimit = 20000
)

//VoteSync defines fields used in VoteSync
type VoteSync struct {
	service                iotex.AuthedClient
	iotexAPI               iotexapi.APIServiceClient
	vpsContract            iotex.Contract
	brokerContract         iotex.Contract
	clerkContract          iotex.Contract
	discordBotToken        string
	discordChannelID       string
	discordMsg             string
	discordReminder        string
	discordReminded        bool
	lastViewHeight         uint64
	lastViewTimestamp      time.Time
	lastUpdateHeight       uint64
	lastBrokerUpdateHeight uint64
	lastClerkUpdateHeight  uint64
	lastUpdateTimestamp    time.Time
	timeInternal           time.Duration
	paginationSize         uint8
	brokerPaginationSize   uint8
	lastNativeEphoch       uint64
	tempLastNativeEphoch   uint64
	terminate              chan bool
	terminated             bool
	dardanellesHeight      uint64
	fairbankHeight         uint64
}

//Config defines the configs for VoteSync
type Config struct {
	GravityChainAPIs          []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval  time.Duration `yaml:"gravityChainTimeInterval"`
	OperatorPrivateKey        string        `yaml:"operatorPrivateKey"`
	IoTeXAPI                  string        `yaml:"ioTeXAPI"`
	IoTeXAPISecure            bool          `yaml:"ioTeXAPISecure"`
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
	FairBankHeight            uint64        `yaml:"fairbankHeight"`
	NativeCommitteeInitHeight uint64        `yaml:"nativeCommitteeInitHeight"`
}

//WeightedVote defines voter and votes for weighted vote
type WeightedVote struct {
	Voter string
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

func ioToEthAddress(str string) (common.Address, error) {
	addr, err := address.FromString(str)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(addr.Bytes()), nil
}

//NewVoteSync instantiates new VoteSync
func NewVoteSync(cfg Config) (*VoteSync, error) {
	ctx := context.Background()

	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Second)),
		grpc_retry.WithMax(3),
	}
	var conn *grpc.ClientConn
	var err error
	if cfg.IoTeXAPISecure {
		conn, err = grpc.DialContext(ctx, cfg.IoTeXAPI,
			grpc.WithBlock(),
			grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
			grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		conn, err = grpc.DialContext(ctx, cfg.IoTeXAPI,
			grpc.WithBlock(),
			grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
			grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
			grpc.WithInsecure())
	}
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

	if lastUpdateHeight.Uint64() == 0 {
		lastUpdateHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	}
	if lastUpdateHeight.Uint64() > _viewIDOffsite {
		lastUpdateHeight.Sub(lastUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}
	lastUpdateTimestamp, err := iotexBlockTime(context.Background(), iotexAPI, lastUpdateHeight.Uint64())
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
	// TODO fix this
	lastViewHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	if lastViewHeight.Uint64() > _viewIDOffsite {
		lastViewHeight.Sub(lastViewHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}
	lastViewTimestamp, err := iotexBlockTime(context.Background(), iotexAPI, lastViewHeight.Uint64())
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

	// TODO fix this
	lastBrokerUpdateHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	if lastBrokerUpdateHeight.Uint64() > _viewIDOffsite {
		lastBrokerUpdateHeight.Sub(lastBrokerUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}

	lastClerkUpdateHeight := new(big.Int)
	d, err = vitaContract.Read("lastRewardPoolClaimViewID").Call(ctx)
	if err != nil {
		return nil, err
	}
	if err := d.Unmarshal(&lastClerkUpdateHeight); err != nil {
		return nil, err
	}
	// TODO fix this
	lastClerkUpdateHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	if lastClerkUpdateHeight.Uint64() > _viewIDOffsite {
		lastClerkUpdateHeight.Sub(lastClerkUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
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
		terminated:             false,
		discordBotToken:        cfg.DiscordBotToken,
		discordChannelID:       cfg.DiscordChannelID,
		discordMsg:             cfg.DiscordMsg,
		discordReminder:        cfg.DiscordReminder,
		dardanellesHeight:      cfg.DardaenllesHeight,
		fairbankHeight:         cfg.FairBankHeight,
	}, nil
}

//Start starts voteSync
func (vc *VoteSync) Start(ctx context.Context) {
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.",
		zap.Uint64("lastUpdateHeight", vc.lastUpdateHeight),
		zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight),
		zap.Uint64("lastClerkUpdateHeight", vc.lastClerkUpdateHeight),
		zap.Uint64("lastViewID", vc.lastViewHeight),
	)
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-vc.terminate:
				return
			case <-ticker.C:
				tip, err := iotexTip(ctx, vc.iotexAPI)
				if err != nil {
					zap.L().Error("failed to get iotex tip", zap.Error(err))
					continue
				}
				blockTime, err := iotexBlockTime(ctx, vc.iotexAPI, tip)
				if err != nil {
					zap.L().Error("failed to get block time", zap.Error(err))
					continue
				}
				if blockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal)) {
					if err := vc.sync(ctx, vc.lastViewHeight, tip, blockTime); err != nil {
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
}

//Stop stops voteSync
func (vc *VoteSync) Stop(ctx context.Context) {
	if vc.terminated {
		return
	}
	close(vc.terminate)
	vc.terminated = true
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

func (vc *VoteSync) sync(ctx context.Context, prevHeight, currHeight uint64, currTs time.Time) error {
	zap.L().Info("Start VoteSyncing.", zap.Uint64("lastViewID", prevHeight), zap.Uint64("nextViewID", currHeight))
	ret, err := vc.fetchVotesUpdate(ctx, prevHeight, currHeight)
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
		addr, err := ioToEthAddress(vote.Voter)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed convert address:%s", vote.Voter))
		}
		addrs = append(addrs, addr)
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
	caller := vc.vpsContract.Execute("rotate", new(big.Int).SetUint64(currHeight+_viewIDOffsite)).
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

func (vc *VoteSync) fetchVotesUpdate(ctx context.Context, prevHeight, currHeight uint64) ([]*WeightedVote, error) {
	// prevHeight == 0, only run at first 2 time. get all votes from currHeight
	if prevHeight == 0 {
		currB, currC, err := vc.retryFetchBucketsByHeight(ctx, currHeight)
		if err != nil {
			return nil, err
		}

		n := calWeightedVotes(currB, currC)
		var ret []*WeightedVote
		for _, nv := range n {
			ret = append(ret, nv)
		}
		return ret, nil
	}

	prevB, prevC, err := vc.retryFetchBucketsByHeight(ctx, prevHeight)
	if err != nil {
		return nil, err
	}

	currB, currC, err := vc.retryFetchBucketsByHeight(ctx, currHeight)
	if err != nil {
		return nil, err
	}

	p := calWeightedVotes(prevB, prevC)
	n := calWeightedVotes(currB, currC)

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

func (vc *VoteSync) retryFetchBucketsByHeight(ctx context.Context, h uint64) (*iotextypes.VoteBucketList, *iotextypes.CandidateListV2, error) {
	var (
		ret1 *iotextypes.VoteBucketList
		ret2 *iotextypes.CandidateListV2
		err  error
	)
	nerr := backoff.Retry(func() error {
		ret1, err = getAllStakingBuckets(ctx, vc.iotexAPI, h)
		if err != nil {
			return err
		}
		ret2, err = getAllStakingCandidates(ctx, vc.iotexAPI, h)
		return err
	}, backoff.NewExponentialBackOff())
	if nerr != nil {
		zap.L().Error(
			"failed to fetch vote result by height",
			zap.Error(nerr),
			zap.Uint64("height", h),
		)
	}
	return ret1, ret2, nerr
}

func iotexTip(ctx context.Context, iotexAPI iotexapi.APIServiceClient) (uint64, error) {
	response, err := iotexAPI.GetChainMeta(
		context.Background(),
		&iotexapi.GetChainMetaRequest{},
	)
	if err != nil {
		return 0, err
	}

	return response.ChainMeta.Height, nil
}

func iotexBlockTime(ctx context.Context, iotexAPI iotexapi.APIServiceClient, h uint64) (time.Time, error) {
	resp, err := iotexAPI.GetBlockMetas(ctx, &iotexapi.GetBlockMetasRequest{
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

func calWeightedVotes(bs *iotextypes.VoteBucketList, cs *iotextypes.CandidateListV2) map[string]*WeightedVote {
	n := make(map[string]*WeightedVote)
	for _, v := range bs.GetBuckets() {
		var selfStake bool
		for _, c := range cs.GetCandidates() {
			if c.GetSelfStakeBucketIdx() == v.GetIndex() {
				selfStake = true
				break
			}
		}

		vs := calculateVoteWeight(v, selfStake)
		wv, ok := n[v.GetOwner()]
		if ok {
			wv.Votes.Add(wv.Votes, vs)
			continue
		}
		n[v.GetOwner()] = &WeightedVote{
			Voter: v.GetOwner(),
			Votes: vs,
		}
	}
	return n
}

func calculateVoteWeight(v *iotextypes.VoteBucket, selfStake bool) *big.Int {
	durationLg := 1.2
	autoStake := 1.0
	selfStakeR := 1.06
	remainingTime := float64(v.StakedDuration)
	weight := float64(1)
	var m float64
	if v.AutoStake {
		m = autoStake
	}
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)*(1+m)) / math.Log(durationLg) / 100
	}
	if selfStake && v.AutoStake && v.StakedDuration >= 91 {
		// self-stake extra bonus requires enable auto-stake for at least 3 months
		weight *= selfStakeR
	}

	a, _ := new(big.Int).SetString(v.StakedAmount, 10)
	amount := new(big.Float).SetInt(a)
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount
}

func getAllStakingBuckets(ctx context.Context, chainClient iotexapi.APIServiceClient, height uint64) (voteBucketListAll *iotextypes.VoteBucketList, err error) {
	voteBucketListAll = &iotextypes.VoteBucketList{}
	for i := uint32(0); ; i++ {
		offset := i * readBucketsLimit
		size := uint32(readBucketsLimit)
		voteBucketList, err := getStakingBuckets(ctx, chainClient, offset, size, height)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get bucket")
		}
		voteBucketListAll.Buckets = append(voteBucketListAll.Buckets, voteBucketList.Buckets...)
		if len(voteBucketList.Buckets) < readBucketsLimit {
			break
		}
	}
	return
}

// getStakingBuckets get specific buckets by height
func getStakingBuckets(ctx context.Context, chainClient iotexapi.APIServiceClient, offset, limit uint32, height uint64) (voteBucketList *iotextypes.VoteBucketList, err error) {
	methodName, err := proto.Marshal(&iotexapi.ReadStakingDataMethod{
		Method: iotexapi.ReadStakingDataMethod_BUCKETS,
	})
	if err != nil {
		return nil, err
	}
	arg, err := proto.Marshal(&iotexapi.ReadStakingDataRequest{
		Request: &iotexapi.ReadStakingDataRequest_Buckets{
			Buckets: &iotexapi.ReadStakingDataRequest_VoteBuckets{
				Pagination: &iotexapi.PaginationParam{
					Offset: offset,
					Limit:  limit,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	readStateRequest := &iotexapi.ReadStateRequest{
		ProtocolID: []byte(protocolID),
		MethodName: methodName,
		Arguments:  [][]byte{arg, []byte(strconv.FormatUint(height, 10))},
	}
	readStateRes, err := chainClient.ReadState(ctx, readStateRequest)
	if err != nil {
		return
	}
	voteBucketList = &iotextypes.VoteBucketList{}
	if err := proto.Unmarshal(readStateRes.GetData(), voteBucketList); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal VoteBucketList")
	}
	return
}

func getAllStakingCandidates(ctx context.Context, chainClient iotexapi.APIServiceClient, height uint64) (candidateListAll *iotextypes.CandidateListV2, err error) {
	candidateListAll = &iotextypes.CandidateListV2{}
	for i := uint32(0); ; i++ {
		offset := i * readCandidatesLimit
		size := uint32(readCandidatesLimit)
		candidateList, err := getStakingCandidates(ctx, chainClient, offset, size, height)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get candidates")
		}
		candidateListAll.Candidates = append(candidateListAll.Candidates, candidateList.Candidates...)
		if len(candidateList.Candidates) < readCandidatesLimit {
			break
		}
	}
	return
}

// getStakingCandidates get specific candidates by height
func getStakingCandidates(ctx context.Context, chainClient iotexapi.APIServiceClient, offset, limit uint32, height uint64) (candidateList *iotextypes.CandidateListV2, err error) {
	methodName, err := proto.Marshal(&iotexapi.ReadStakingDataMethod{
		Method: iotexapi.ReadStakingDataMethod_CANDIDATES,
	})
	if err != nil {
		return nil, err
	}
	arg, err := proto.Marshal(&iotexapi.ReadStakingDataRequest{
		Request: &iotexapi.ReadStakingDataRequest_Candidates_{
			Candidates: &iotexapi.ReadStakingDataRequest_Candidates{
				Pagination: &iotexapi.PaginationParam{
					Offset: offset,
					Limit:  limit,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	readStateRequest := &iotexapi.ReadStateRequest{
		ProtocolID: []byte(protocolID),
		MethodName: methodName,
		Arguments:  [][]byte{arg, []byte(strconv.FormatUint(height, 10))},
	}
	readStateRes, err := chainClient.ReadState(ctx, readStateRequest)
	if err != nil {
		return
	}
	candidateList = &iotextypes.CandidateListV2{}
	if err := proto.Unmarshal(readStateRes.GetData(), candidateList); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal VoteBucketList")
	}
	return
}
