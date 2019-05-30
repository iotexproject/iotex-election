package votesync

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/types"
)

type VoteSync struct {
	carrier        carrier.Carrier
	lastHeight     uint64
	heightInternal uint64
	paginationSize uint8
	terminate      chan bool
}

type Config struct {
	GravityChainAPIs           []string `yaml:"gravityChainAPIs"`
	GravityChainHeightInterval uint64   `yaml:"gravityChainHeightInterval"`
	RegisterContractAddress    string   `yaml:"registerContractAddress"`
	StakingContractAddress     string   `yaml:"stakingContractAddress"`
	PaginationSize             uint8    `yaml:"paginationSize"`
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
	return &VoteSync{
		carrier:        carrier,
		heightInternal: cfg.GravityChainHeightInterval,
		paginationSize: cfg.PaginationSize,
		lastHeight:     lastHeight,
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
				if vc.heightInternal+vc.lastHeight <= height {
					if err := vc.sync(vc.lastHeight, height); err != nil {
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

func (vc *VoteSync) sync(prevHeight, currHeight uint64) error {
	a, b, err := vc.fetchVotesUpdate(prevHeight, currHeight)
	if err != nil {
		return err
	}

	// TODO write results to iotex contract
	log.Println(a, b)
	vc.lastHeight = currHeight
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(prevHeight, currHeight uint64) ([]*types.Vote, [][]byte, error) {
	prev, err := vc.retryFetchResultByHeight(prevHeight)
	if err != nil {
		return nil, nil, err
	}
	curr, err := vc.retryFetchResultByHeight(currHeight)
	if err != nil {
		return nil, nil, err
	}
	m := make(map[string]interface{})
	var cleaned [][]byte
	for _, v := range prev {
		m[hex.EncodeToString(v.Voter())] = v
	}
	for _, v := range curr {
		if _, ok := m[hex.EncodeToString(v.Voter())]; !ok {
			cleaned = append(cleaned, v.Voter())
		}
	}
	return curr, cleaned, nil
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
