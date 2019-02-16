// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"context"
	"log"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

var (
	// ErrNotExist defines an error that the query has no return value in db
	ErrNotExist = errors.New("not exist in db")
)

// CalcBeaconChainHeight calculates the corresponding beacon chain height for an epoch
type CalcBeaconChainHeight func(uint64) (uint64, error)

// KVStore defines the db interface using in committee
type KVStore interface {
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
}

// Config defines the config of the committee
type Config struct {
	NumOfRetries              uint8  `yaml:"numOfRetries"`
	BeaconChainAPI            string `yaml:"beaconChainAPI"`
	BeaconChainHeightInterval uint64 `yaml:"beaconChainHeightInterval"`
	BeaconChainStartHeight    uint64 `yaml:"beaconChainStartHeight"`
	StakingContractAddress    string `yaml:"stakingContractAddress"`
	PaginationSize            uint8  `yaml:"paginationSize"`
	VoteThreshold             uint64 `yaml:"voteThreshold"`
	ScoreThreshold            uint64 `yaml:"scoreThreshold"`
	SelfStakingThreshold      uint64 `yaml:"selfStakingThreshold"`
	CacheSize                 uint8  `yaml:"cacheSize"`
}

type resultCache struct {
	size    uint8
	results []*types.ElectionResult
	heights []uint64
	index   map[uint64]int
	cursor  int
}

func newResultCache(size uint8) *resultCache {
	return &resultCache{
		size:    size,
		results: make([]*types.ElectionResult, size),
		heights: make([]uint64, size),
		index:   map[uint64]int{},
		cursor:  0,
	}
}

func (c *resultCache) insert(height uint64, r *types.ElectionResult) {
	if i, exists := c.index[height]; exists {
		c.results[i] = r
		return
	}
	delete(c.index, c.heights[c.cursor])
	c.results[c.cursor] = r
	c.heights[c.cursor] = height
	c.index[height] = c.cursor
	c.cursor = (c.cursor + 1) % int(c.size)
}

func (c *resultCache) get(height uint64) *types.ElectionResult {
	i, exists := c.index[height]
	if !exists {
		return nil
	}
	return c.results[i]
}

type heightManager struct {
	heights []uint64
	times   []time.Time
}

func newHeightManager() *heightManager {
	return &heightManager{
		heights: []uint64{},
		times:   []time.Time{},
	}
}

func (m *heightManager) nearestHeightBefore(ts time.Time) uint64 {
	l := len(m.heights)
	if l == 0 {
		return 0
	}
	if m.times[0].After(ts) {
		return 0
	}
	head := 0
	tail := l
	for {
		if tail-head <= 1 {
			break
		}
		mid := (head + tail) / 2
		if m.times[mid].After(ts) {
			tail = mid
		} else {
			head = mid
		}
	}
	return m.heights[head]
}

func (m *heightManager) lastestHeight() uint64 {
	l := len(m.heights)
	if l == 0 {
		return 0
	}
	return m.heights[l-1]
}

func (m *heightManager) validate(height uint64, ts time.Time) error {
	l := len(m.heights)
	if l == 0 {
		return nil
	}
	if m.heights[l-1] >= height {
		return errors.Errorf(
			"invalid height %d, current tail is %d",
			height,
			m.heights[l-1],
		)
	}
	if !ts.After(m.times[l-1]) {
		return errors.Errorf(
			"invalid timestamp %s, current tail is %s",
			ts,
			m.times[l-1],
		)
	}
	return nil
}

func (m *heightManager) add(height uint64, ts time.Time) error {
	if err := m.validate(height, ts); err != nil {
		return err
	}
	m.heights = append(m.heights, height)
	m.times = append(m.times, ts)
	return nil
}

// Committee defines an interface of an election committee
// It could be considered as a light state db of beacon chain, that
type Committee interface {
	// Start starts the committee service
	Start(context.Context) error
	// Stop stops the committee service
	Stop(context.Context) error
	// ResultByHeight returns the result on a specific ethereum height
	ResultByHeight(height uint64) (*types.ElectionResult, error)
	// ResultByTime returns the nearest result before time
	ResultByTime(timestamp time.Time) (uint64, *types.ElectionResult, error)
	// OnNewBlock is a callback function which will be called on new block created
	OnNewBlock(height uint64)
	// LatestHeight returns the height with latest result
	LatestHeight() uint64
}

type committee struct {
	db            KVStore
	carrier       carrier.Carrier
	cfg           Config
	cache         *resultCache
	heightManager *heightManager
	nextHeight    uint64
	currentHeight uint64
	terminate     chan bool
	mutex         sync.RWMutex
}

// NewCommittee creates a committee
func NewCommittee(db KVStore, cfg Config) (Committee, error) {
	if !common.IsHexAddress(cfg.StakingContractAddress) {
		return nil, errors.New("Invalid staking contract address")
	}
	carrier, err := carrier.NewEthereumVoteCarrier(
		cfg.BeaconChainAPI,
		common.HexToAddress(cfg.StakingContractAddress),
	)
	if err != nil {
		return nil, err
	}
	return &committee{
		db:            db,
		cache:         newResultCache(cfg.CacheSize),
		heightManager: newHeightManager(),
		carrier:       carrier,
		cfg:           cfg,
		terminate:     make(chan bool),
		currentHeight: 0,
		nextHeight:    cfg.BeaconChainStartHeight,
	}, nil
}

func (ec *committee) Start(ctx context.Context) (err error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	for {
		result, err := ec.resultByHeight(ec.nextHeight)
		if err == ErrNotExist {
			break
		}
		if err == nil {
			if err := ec.heightManager.add(ec.nextHeight, result.MintTime()); err != nil {
				return err
			}
			ec.cache.insert(ec.nextHeight, result)
			ec.currentHeight = ec.nextHeight
			ec.nextHeight += ec.cfg.BeaconChainHeightInterval
			continue
		}
		return err
	}
	for i := uint8(0); i < ec.cfg.NumOfRetries; i++ {
		if err = ec.carrier.SubscribeNewBlock(ec.OnNewBlock, ec.terminate); err == nil {
			break
		}
	}
	return
}

func (ec *committee) Stop(ctx context.Context) error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	ec.terminate <- true
	return nil
}

func (ec *committee) OnNewBlock(tipHeight uint64) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if ec.currentHeight < tipHeight {
		ec.currentHeight = tipHeight
	}
	for {
		if ec.nextHeight > ec.currentHeight {
			break
		}
		var result *types.ElectionResult
		var err error
		for i := uint8(0); i < ec.cfg.NumOfRetries; i++ {
			if result, err = ec.fetchResultByHeight(ec.nextHeight); err != nil {
				log.Println(err)
				continue
			}
			break
		}
		if result == nil {
			log.Printf("failed to fetch result for %d\n", ec.nextHeight)
			return
		}
		if err = ec.heightManager.validate(ec.nextHeight, result.MintTime()); err != nil {
			log.Fatalln(
				"Unexpected status that the upcoming block height or time is invalid",
				err,
			)
			return
		}
		if err = ec.storeResult(ec.nextHeight, result); err != nil {
			log.Println("failed to store result into db", err)
			return
		}
		ec.heightManager.add(ec.nextHeight, result.MintTime())
		ec.cache.insert(ec.nextHeight, result)
		ec.nextHeight += ec.cfg.BeaconChainHeightInterval
	}
}

func (ec *committee) LatestHeight() uint64 {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.heightManager.lastestHeight()
}

func (ec *committee) ResultByTime(ts time.Time) (uint64, *types.ElectionResult, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	height := ec.heightManager.nearestHeightBefore(ts)
	if height == 0 {
		return 0, nil, ErrNotExist
	}
	result, err := ec.resultByHeight(height)

	return height, result, err
}

func (ec *committee) ResultByHeight(height uint64) (*types.ElectionResult, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.resultByHeight(height)
}

func (ec *committee) resultByHeight(height uint64) (*types.ElectionResult, error) {
	if height < ec.cfg.BeaconChainStartHeight {
		return nil, errors.Errorf(
			"height %d is higher than start height %d",
			height,
			ec.cfg.BeaconChainStartHeight,
		)
	}
	if (height-ec.cfg.BeaconChainStartHeight)%ec.cfg.BeaconChainHeightInterval != 0 {
		return nil, errors.Errorf(
			"height %d is an invalid height",
			height,
		)
	}
	result := ec.cache.get(height)
	if result != nil {
		return result, nil
	}
	data, err := ec.db.Get(ec.dbKey(height))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotExist
	}
	result = &types.ElectionResult{}

	return result, result.Deserialize(data)
}

func (ec *committee) calcWeightedVotes(v *types.Vote, now time.Time) *big.Int {
	if now.Before(v.StartTime()) {
		return big.NewInt(0)
	}
	remainingTime := v.RemainingTime(now).Seconds()
	weight := float64(1)
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)) / math.Log(1.2)
	}
	amount := new(big.Float).SetInt(v.Amount())
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)

	return weightedAmount
}

func (ec *committee) fetchResultByHeight(height uint64) (*types.ElectionResult, error) {
	mintTime, err := ec.carrier.BlockTimestamp(height)
	if err != nil {
		return nil, err
	}
	calculator := types.NewResultCalculator(
		mintTime,
		func(v *types.Vote) bool {
			return new(big.Int).SetUint64(ec.cfg.VoteThreshold).Cmp(v.Amount()) > 0
		},
		ec.calcWeightedVotes,
		func(c *types.Candidate) bool {
			return new(big.Int).SetUint64(ec.cfg.SelfStakingThreshold).Cmp(c.SelfStakingScore()) > 0 &&
				new(big.Int).SetUint64(ec.cfg.ScoreThreshold).Cmp(c.Score()) > 0
		},
	)
	previousIndex := big.NewInt(1)
	for {
		var candidates []*types.Candidate
		var err error
		if previousIndex, candidates, err = ec.carrier.Candidates(
			height,
			previousIndex,
			ec.cfg.PaginationSize,
		); err != nil {
			return nil, err
		}
		calculator.AddCandidates(candidates)
		if len(candidates) < int(ec.cfg.PaginationSize) {
			break
		}
	}
	previousIndex = big.NewInt(0)
	for {
		var votes []*types.Vote
		var err error
		if previousIndex, votes, err = ec.carrier.Votes(
			height,
			previousIndex,
			ec.cfg.PaginationSize,
		); err != nil {
			return nil, err
		}
		calculator.AddVotes(votes)
		if len(votes) < int(ec.cfg.PaginationSize) {
			break
		}
	}
	return calculator.Calculate()
}

func (ec *committee) dbKey(height uint64) []byte {
	return util.Uint64ToBytes(height)
}

func (ec *committee) storeResult(height uint64, result *types.ElectionResult) error {
	data, err := result.Serialize()
	if err != nil {
		return err
	}

	return ec.db.Put(ec.dbKey(height), data)
}
