package votesync

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/account"
	"github.com/iotexproject/iotex-antenna-go/iotx"
	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

type VoteSync struct {
	operator               string
	vitaContractAddress    string
	vpsContractAddress     string
	brokerContractAddress  string
	service                *iotx.Iotx
	carrier                carrier.Carrier
	lastViewHeight         uint64
	lastViewTimestamp      time.Time
	lastUpdateHeight       uint64
	lastBrokerUpdateHeight uint64
	lastUpdateTimestamp    time.Time
	timeInternal           time.Duration
	paginationSize         uint8
	brokerPaginationSize   uint8
	terminate              chan bool
}

type Config struct {
	GravityChainAPIs         []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval time.Duration `yaml:"gravityChainTimeInterval"`
	OperatorPrivateKey       string        `yaml:"operatorPrivateKey"`
	IoTeXAPI                 string        `yaml:"ioTeXAPI"`
	RegisterContractAddress  string        `yaml:"registerContractAddress"`
	StakingContractAddress   string        `yaml:"stakingContractAddress"`
	PaginationSize           uint8         `yaml:"paginationSize"`
	BrokerPaginationSize     uint8         `yaml:"brokerPaginationSize"`
	VitaContractAddress      string        `yaml:"vitaContractAddress"`
}

type WeightedVote struct {
	Voter []byte
	Votes *big.Int
}

func readContract(
	service *iotx.Iotx,
	contractABI string,
	contractAddr string,
	method string,
	accountAddr string,
	retval interface{},
) error {
	parsed, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return err
	}
	var response string
	err = backoff.Retry(func() error {
		response, err = service.ReadContractByMethod(&iotx.ContractRequest{
			Address:  contractAddr,
			From:     accountAddr,
			Abi:      contractABI,
			Method:   method,
			GasLimit: "5000000",
			GasPrice: "1",
		})
		return err
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return err
	}
	decoded, err := hex.DecodeString(response)
	if err != nil {
		return err
	}
	return parsed.Unpack(retval, method, decoded)
}

func toIoAddress(addr common.Address) (string, error) {
	pkhash, err := hex.DecodeString(strings.TrimLeft(addr.String(), "0x"))
	if err != nil {
		return "", err
	}
	ioaddr, err := address.FromBytes(pkhash)
	if err != nil {
		return "", err
	}
	return ioaddr.String(), nil
}

func NewVoteSync(cfg Config) (*VoteSync, error) {
	carrier, err := carrier.NewEthereumVoteCarrier(
		cfg.GravityChainAPIs,
		common.HexToAddress(cfg.RegisterContractAddress),
		common.HexToAddress(cfg.StakingContractAddress),
	)
	if err != nil {
		return nil, err
	}
	service, err := iotx.NewIotx(cfg.IoTeXAPI, true)
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
	if service.Accounts.AddAccount(operatorAccount); err != nil {
		return nil, err
	}
	var addr common.Address
	if err = readContract(
		service,
		contract.VitaABI,
		cfg.VitaContractAddress,
		"vps",
		operatorAccount.Address(),
		&addr,
	); err != nil {
		return nil, err
	}
	vpsContractAddress, err := toIoAddress(addr)
	if err != nil {
		return nil, err
	}
	if err = readContract(
		service,
		contract.VitaABI,
		cfg.VitaContractAddress,
		"donatePoolAddress",
		operatorAccount.Address(),
		&addr,
	); err != nil {
		return nil, err
	}
	brokerContractAddress, err := toIoAddress(addr)
	if err != nil {
		return nil, err
	}

	lastUpdateHeight := new(big.Int)
	if err = readContract(
		service,
		contract.RotatableVPSABI,
		vpsContractAddress,
		"viewID",
		operatorAccount.Address(),
		&lastUpdateHeight,
	); err != nil {
		return nil, err
	}
	lastUpdateTimestamp, err := carrier.BlockTimestamp(lastUpdateHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastViewHeight := new(big.Int)
	if err = readContract(
		service,
		contract.RotatableVPSABI,
		vpsContractAddress,
		"inactiveViewID",
		operatorAccount.Address(),
		&lastViewHeight,
	); err != nil {
		return nil, err
	}
	lastViewTimestamp, err := carrier.BlockTimestamp(lastViewHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastBrokerUpdateHeight := new(big.Int)
	if err = readContract(
		service,
		contract.VitaABI,
		cfg.VitaContractAddress,
		"lastDonatePoolClaimViewID",
		operatorAccount.Address(),
		&lastBrokerUpdateHeight,
	); err != nil {
		return nil, err
	}
	return &VoteSync{
		carrier:                carrier,
		vitaContractAddress:    cfg.VitaContractAddress,
		vpsContractAddress:     vpsContractAddress,
		brokerContractAddress:  brokerContractAddress,
		operator:               operatorAccount.Address(),
		service:                service,
		timeInternal:           cfg.GravityChainTimeInterval,
		paginationSize:         cfg.PaginationSize,
		brokerPaginationSize:   cfg.BrokerPaginationSize,
		lastViewHeight:         lastViewHeight.Uint64(),
		lastViewTimestamp:      lastViewTimestamp,
		lastUpdateHeight:       lastUpdateHeight.Uint64(),
		lastUpdateTimestamp:    lastUpdateTimestamp,
		lastBrokerUpdateHeight: lastBrokerUpdateHeight.Uint64(),
		terminate:              make(chan bool),
	}, nil
}

func (vc *VoteSync) Start(ctx context.Context) {
	tipChan := make(chan *carrier.TipInfo)
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.",
		zap.Uint64("lastUpdateHeight", vc.lastUpdateHeight),
		zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight),
		zap.Uint64("lastViewID", vc.lastViewHeight),
	)
	go func() {
		for {
			select {
			case <-vc.terminate:
				vc.terminate <- true
				return
			case tip := <-tipChan:
				if tip.BlockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal)) {
					if err := vc.sync(vc.lastViewHeight, tip.Height, vc.lastViewTimestamp, tip.BlockTime); err != nil {
						zap.L().Error("failed to sync votes", zap.Error(err))
						continue
					}
				}

				if vc.lastUpdateHeight > vc.lastBrokerUpdateHeight {
					if err := vc.settle(vc.lastUpdateHeight); err != nil {
						zap.L().Error("failed to settle broker", zap.Error(err))
					}
				}
			case err := <-errChan:
				zap.L().Error("something goes wrong", zap.Error(err))
			}
		}
	}()
	vc.carrier.SubscribeNewBlock(tipChan, errChan, vc.terminate)
}

func (vc *VoteSync) Stop(ctx context.Context) {
	vc.terminate <- true
	vc.carrier.Close()
	return
}

func (vc *VoteSync) checkExecutionByHash(hash string) error {
	response, err := vc.service.GetReceiptByAction(&iotexapi.GetReceiptByActionRequest{
		ActionHash: hash,
	})
	if err != nil {
		return err
	}
	if response.ReceiptInfo.Receipt.Status != 1 {
		return errors.Errorf("execution failed: %s", hash)
	}

	return err
}

func (vc *VoteSync) brokerReset() error {
	return backoff.Retry(func() error {
		hash, err := vc.service.ExecuteContract(&iotx.ContractRequest{
			Address:  vc.brokerContractAddress,
			From:     vc.operator,
			Abi:      contract.BrokerABI,
			Method:   "reset",
			Amount:   "0",
			GasLimit: "5000000",
			GasPrice: "1",
		})
		if err != nil {
			return err
		}
		time.Sleep(20 * time.Second)
		return vc.checkExecutionByHash(hash)
	}, backoff.NewExponentialBackOff())
}

func (vc *VoteSync) brokerNextBidToSettle() (uint64, error) {
	nextBidToSettle := new(big.Int)
	if err := readContract(
		vc.service,
		contract.BrokerABI,
		vc.brokerContractAddress,
		"nextBidToSettle",
		vc.operator,
		&nextBidToSettle,
	); err != nil {
		return 0, err
	}
	return nextBidToSettle.Uint64(), nil
}

func (vc *VoteSync) brokerSettle() error {
	oldStart := uint64(0)
	for {
		if err := backoff.Retry(func() error {
			hash, err := vc.service.ExecuteContract(&iotx.ContractRequest{
				Address:  vc.brokerContractAddress,
				From:     vc.operator,
				Abi:      contract.BrokerABI,
				Method:   "settle",
				Amount:   "0",
				GasLimit: "5000000",
				GasPrice: "1",
			}, big.NewInt(0).SetUint64(uint64(vc.brokerPaginationSize)))
			if err != nil {
				return err
			}
			time.Sleep(20 * time.Second)
			return vc.checkExecutionByHash(hash)
		}, backoff.NewExponentialBackOff()); err != nil {
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

func (vc *VoteSync) updateVotingPowers(addrs []common.Address, weights []*big.Int) error {
	return backoff.Retry(func() error {
		hash, err := vc.service.ExecuteContract(&iotx.ContractRequest{
			Address:  vc.vpsContractAddress,
			From:     vc.operator,
			Abi:      contract.RotatableVPSABI,
			Method:   "updateVotingPowers",
			Amount:   "0",
			GasLimit: "5000000",
			GasPrice: "1",
		}, addrs, weights)
		if err != nil {
			return err
		}
		time.Sleep(20 * time.Second)
		return vc.checkExecutionByHash(hash)
	}, backoff.NewExponentialBackOff())
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

		if len(addrs)%int(vc.paginationSize/2) == 0 {
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
	err = backoff.Retry(func() error {
		hash, err := vc.service.ExecuteContract(&iotx.ContractRequest{
			Address:  vc.vpsContractAddress,
			From:     vc.operator,
			Abi:      contract.RotatableVPSABI,
			Method:   "rotate",
			Amount:   "0",
			GasLimit: "4000000",
			GasPrice: "1",
		}, new(big.Int).SetUint64(currHeight))
		if err != nil {
			return err
		}
		time.Sleep(20 * time.Second)
		return vc.checkExecutionByHash(hash)
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return errors.Wrap(err, "failed to execute rotate")
	}

	vc.lastViewHeight = vc.lastUpdateHeight
	vc.lastViewTimestamp = vc.lastUpdateTimestamp
	vc.lastUpdateHeight = currHeight
	vc.lastUpdateTimestamp = currTs
	zap.L().Info("Successfully synced votes.", zap.Uint64("lastViewID", vc.lastViewHeight), zap.Uint64("viewID", currHeight))
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(prevHeight, currHeight uint64, prevTs, currTs time.Time) ([]*WeightedVote, error) {
	// prevHeight == 0, only run at first 2 time. get all votes from currHeight
	if prevHeight == 0 {
		curr, err := vc.retryFetchResultByHeight(currHeight)
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

	prev, err := vc.retryFetchResultByHeight(prevHeight)
	if err != nil {
		return nil, err
	}
	curr, err := vc.retryFetchResultByHeight(currHeight)
	if err != nil {
		return nil, err
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

func (vc *VoteSync) retryFetchResultByHeight(h uint64) ([]*types.Vote, error) {
	var (
		ret []*types.Vote
		err error
	)
	nerr := backoff.Retry(func() error {
		ret, err = vc.fetchVotesByHeight(h)
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

func (vc *VoteSync) fetchVotesByHeight(h uint64) ([]*types.Vote, error) {
	var allVotes []*types.Vote
	idx := big.NewInt(0)
	for {
		var (
			votes []*types.Vote
			err   error
		)
		if idx, votes, err = vc.carrier.Votes(h, idx, vc.paginationSize); err != nil {
			return nil, err
		}
		allVotes = append(allVotes, votes...)
		if len(votes) < int(vc.paginationSize) {
			break
		}
	}
	return allVotes, nil
}

func calWeightedVotes(curr []*types.Vote, currTs time.Time) map[string]*WeightedVote {
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
