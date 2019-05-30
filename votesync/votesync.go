package votesync

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/types"
)

type VoteSync struct {
	carrier        carrier.Carrier
	lastHeight     uint64
	lastTimestamp  time.Time
	timeInternal   time.Duration
	paginationSize uint8
	terminate      chan bool
}

type Config struct {
	GravityChainAPIs         []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval time.Duration `yaml:"gravityChainTimeInterval"`
	RegisterContractAddress  string        `yaml:"registerContractAddress"`
	StakingContractAddress   string        `yaml:"stakingContractAddress"`
	PaginationSize           uint8         `yaml:"paginationSize"`
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

	// TODO get lastHeight from iotex contract
	lastHeight := uint64(7840000)
	lastTimestamp, err := carrier.BlockTimestamp(lastHeight)
	if err != nil {
		return nil, err
	}
	return &VoteSync{
		carrier:        carrier,
		timeInternal:   cfg.GravityChainTimeInterval,
		paginationSize: cfg.PaginationSize,
		lastHeight:     lastHeight,
		lastTimestamp:  lastTimestamp,
		terminate:      make(chan bool),
	}, nil
}

func (vc *VoteSync) Start(ctx context.Context) {
	heightChan := make(chan uint64)
	errChan := make(chan error)
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

func (vc *VoteSync) sync(prevHeight, currHeight uint64, currTs time.Time) error {
	ret, err := vc.fetchVotesUpdate(prevHeight, currHeight, currTs)
	if err != nil {
		return err
	}

	// TODO write results to iotex contract
	log.Println(ret)
	vc.lastHeight = currHeight
	vc.lastTimestamp = currTs
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
