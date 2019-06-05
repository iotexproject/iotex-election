package votesync

import (
	"context"
	"encoding/hex"
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
	operator           string
	vpsContractAddress string
	service            *iotx.Iotx
	carrier            carrier.Carrier
	lastHeight         uint64
	lastTimestamp      time.Time
	timeInternal       time.Duration
	paginationSize     uint8
	terminate          chan bool
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
	lastHeight := new(big.Int)
	if err := parsed.Unpack(&lastHeight, "viewID", decoded); err != nil {
		return nil, err
	}
	lastTimestamp, err := carrier.BlockTimestamp(lastHeight.Uint64())
	if err != nil {
		return nil, err
	}
	return &VoteSync{
		carrier:            carrier,
		vpsContractAddress: cfg.VotingSystemContractAddress,
		operator:           operatorAccount.Address(),
		service:            service,
		timeInternal:       cfg.GravityChainTimeInterval,
		paginationSize:     cfg.PaginationSize,
		lastHeight:         lastHeight.Uint64(),
		lastTimestamp:      lastTimestamp,
		terminate:          make(chan bool),
	}, nil
}

func (vc *VoteSync) Start(ctx context.Context) {
	heightChan := make(chan uint64)
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.", zap.Uint64("viewID", vc.lastHeight))
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
				if t.After(vc.lastTimestamp.Add(vc.timeInternal)) {
					// TODO: add retry and alert
					if err := vc.sync(vc.lastHeight, height, t); err != nil {
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
		GasLimit: "400000",
		GasPrice: "1",
	}, addrs, weights)
	if err != nil {
		return err
	}
	time.Sleep(20 * time.Second)

	return vc.checkExecutionByHash(hash)
}

func (vc *VoteSync) sync(prevHeight, currHeight uint64, currTs time.Time) error {
	ret, err := vc.fetchVotesUpdate(prevHeight, currHeight, currTs)
	if err != nil {
		return errors.Wrap(err, "fetch vote error")
	}

	var addrs []common.Address
	var weights []*big.Int
	for _, vote := range ret {
		addrs = append(addrs, common.BytesToAddress(vote.Voter))
		weights = append(weights, vote.Votes)

		if len(addrs)%int(vc.paginationSize) == 0 {
			if err := vc.updateVotingPowers(addrs, weights); err != nil {
				return errors.Wrap(err, "update vote error")
			}
			addrs = []common.Address{}
			weights = []*big.Int{}
		}
	}
	if len(addrs) > 0 {
		if err := vc.updateVotingPowers(addrs, weights); err != nil {
			return errors.Wrap(err, "update vote error")
		}
	}
	hash, err := vc.service.ExecuteContract(&iotx.ContractRequest{
		Address:  vc.vpsContractAddress,
		From:     vc.operator,
		Abi:      contract.RotatableVPSABI,
		Method:   "rotate",
		Amount:   "0",
		GasLimit: "400000",
		GasPrice: "1",
	}, new(big.Int).SetUint64(currHeight))
	if err != nil {
		return err
	}
	time.Sleep(20 * time.Second)
	if err := vc.checkExecutionByHash(hash); err != nil {
		return err
	}

	vc.lastHeight = currHeight
	vc.lastTimestamp = currTs
	zap.L().Info("Successfully synced votes.", zap.Uint64("viewID", currHeight))
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(prevHeight, currHeight uint64, currTs time.Time) ([]*WeightedVote, error) {
	prev, err := vc.retryFetchResultByHeight(prevHeight)
	if err != nil {
		return nil, err
	}
	curr, err := vc.retryFetchResultByHeight(currHeight)
	if err != nil {
		return nil, err
	}

	r := make(map[string]*WeightedVote)
	for _, v := range curr {
		vs := types.CalcWeightedVotes(v, currTs)
		wv, ok := r[hex.EncodeToString(v.Voter())]
		if ok {
			wv.Votes.Add(wv.Votes, vs)
			continue
		}
		r[hex.EncodeToString(v.Voter())] = &WeightedVote{
			Voter: v.Voter(),
			Votes: vs,
		}
	}

	var ret []*WeightedVote
	for _, v := range r {
		ret = append(ret, v)
	}

	m := make(map[string]*types.Vote)
	for _, v := range prev {
		m[hex.EncodeToString(v.Voter())] = v
	}
	for k, v := range m {
		if _, ok := r[k]; !ok {
			ret = append(ret, &WeightedVote{
				Voter: v.Voter(),
				Votes: big.NewInt(0),
			})
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
