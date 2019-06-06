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
	"github.com/iotexproject/iotex-antenna-go/account"
	"github.com/iotexproject/iotex-antenna-go/iotx"
	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

type VoteSync struct {
	operator            string
	vpsContractAddress  string
	service             *iotx.Iotx
	carrier             carrier.Carrier
	lastViewHeight      uint64
	lastViewTimestamp   time.Time
	lastUpdateHeight    uint64
	lastUpdateTimestamp time.Time
	timeInternal        time.Duration
	paginationSize      uint8
	terminate           chan bool
}

type Config struct {
	GravityChainAPIs            []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval    time.Duration `yaml:"gravityChainTimeInterval"`
	OperatorPrivateKey          string        `yaml:"operatorPrivateKey"`
	IoTeXAPI                    string        `yaml:"ioTeXAPI"`
	IoTeXAPIInSecure            bool          `yaml:"ioTeXAPIInSecure"`
	RegisterContractAddress     string        `yaml:"registerContractAddress"`
	StakingContractAddress      string        `yaml:"stakingContractAddress"`
	PaginationSize              uint8         `yaml:"paginationSize"`
	VotingSystemContractAddress string        `yaml:"votingSystemContractAddress"`
}

type WeightedVote struct {
	Voter []byte
	Votes *big.Int
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
	service, err := iotx.NewIotx(cfg.IoTeXAPI, cfg.IoTeXAPIInSecure)
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
	parsed, err := abi.JSON(strings.NewReader(contract.RotatableVPSABI))
	if err != nil {
		return nil, err
	}
	response, err := service.ReadContractByMethod(&iotx.ContractRequest{
		Address:  cfg.VotingSystemContractAddress,
		From:     operatorAccount.Address(),
		Abi:      contract.RotatableVPSABI,
		Method:   "viewID",
		GasLimit: "400000",
		GasPrice: "1",
	})
	if err != nil {
		return nil, err
	}
	decoded, err := hex.DecodeString(response)
	if err != nil {
		return nil, err
	}
	lastUpdateHeight := new(big.Int)
	if err := parsed.Unpack(&lastUpdateHeight, "viewID", decoded); err != nil {
		return nil, err
	}
	lastUpdateTimestamp, err := carrier.BlockTimestamp(lastUpdateHeight.Uint64())
	if err != nil {
		return nil, err
	}

	// TODO get lastViewHeight
	lastViewHeight := lastUpdateHeight
	lastViewTimestamp, err := carrier.BlockTimestamp(lastViewHeight.Uint64())
	if err != nil {
		return nil, err
	}
	return &VoteSync{
		carrier:             carrier,
		vpsContractAddress:  cfg.VotingSystemContractAddress,
		operator:            operatorAccount.Address(),
		service:             service,
		timeInternal:        cfg.GravityChainTimeInterval,
		paginationSize:      cfg.PaginationSize,
		lastViewHeight:      lastViewHeight.Uint64(),
		lastViewTimestamp:   lastViewTimestamp,
		lastUpdateHeight:    lastUpdateHeight.Uint64(),
		lastUpdateTimestamp: lastUpdateTimestamp,
		terminate:           make(chan bool),
	}, nil
}

func (vc *VoteSync) Start(ctx context.Context) {
	heightChan := make(chan uint64)
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.", zap.Uint64("viewID", vc.lastUpdateHeight))
	go func() {
		for {
			select {
			case <-vc.terminate:
				vc.terminate <- true
				return
			case height := <-heightChan:
				t, err := vc.carrier.BlockTimestamp(height)
				if err != nil {
					zap.L().Error("failed to get eth block time stamp", zap.Error(err))
					continue
				}
				if t.After(vc.lastUpdateTimestamp.Add(vc.timeInternal)) {
					// TODO: add retry and alert
					if err := vc.sync(vc.lastViewHeight, height, vc.lastViewTimestamp, t); err != nil {
						zap.L().Error("failed to sync votes", zap.Error(err))
					}
				}
			case err := <-errChan:
				zap.L().Error("something goes wrong", zap.Error(err))
			}
		}
	}()
	vc.carrier.SubscribeNewBlock(heightChan, errChan, vc.terminate)
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

func (vc *VoteSync) updateVotingPowers(addrs []common.Address, weights []*big.Int) error {
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
}

func (vc *VoteSync) sync(prevHeight, currHeight uint64, prevTs, currTs time.Time) error {
	zap.L().Info("Start syncing today's votes.", zap.Uint64("viewID", currHeight))
	ret, err := vc.fetchVotesUpdate(prevHeight, currHeight, prevTs, currTs)
	if err != nil {
		return errors.Wrap(err, "fetch vote error")
	}

	var (
		addrs   []common.Address
		weights []*big.Int
		reqNum  int
	)
	for _, vote := range ret {
		addrs = append(addrs, common.BytesToAddress(vote.Voter))
		weights = append(weights, vote.Votes)

		//fmt.Println(common.BytesToAddress(vote.Voter).String(), vote.Votes.String())
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
	if err := vc.checkExecutionByHash(hash); err != nil {
		return err
	}

	vc.lastViewHeight = vc.lastUpdateHeight
	vc.lastViewTimestamp = vc.lastUpdateTimestamp
	vc.lastUpdateHeight = currHeight
	vc.lastUpdateTimestamp = currTs
	zap.L().Info("Successfully synced votes.", zap.Uint64("viewID", currHeight))
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(prevHeight, currHeight uint64, prevTs, currTs time.Time) ([]*WeightedVote, error) {
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
