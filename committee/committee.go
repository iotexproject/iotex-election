// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package committee

import (
	"context"
	"database/sql"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	// require sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
)

// CalcGravityChainHeight calculates the corresponding gravity chain height for an epoch
type CalcGravityChainHeight func(uint64) (uint64, error)

// Config defines the config of the committee
type Config struct {
	NumOfRetries               uint8    `yaml:"numOfRetries"`
	GravityChainAPIs           []string `yaml:"gravityChainAPIs"`
	GravityChainHeightInterval uint64   `yaml:"gravityChainHeightInterval"`
	GravityChainStartHeight    uint64   `yaml:"gravityChainStartHeight"`
	RegisterContractAddress    string   `yaml:"registerContractAddress"`
	StakingContractAddress     string   `yaml:"stakingContractAddress"`
	PaginationSize             uint8    `yaml:"paginationSize"`
	VoteThreshold              string   `yaml:"voteThreshold"`
	ScoreThreshold             string   `yaml:"scoreThreshold"`
	SelfStakingThreshold       string   `yaml:"selfStakingThreshold"`
	CacheSize                  uint32   `yaml:"cacheSize"`
	NumOfFetchInParallel       uint8    `yaml:"numOfFetchInParallel"`
	SkipManifiedCandidate      bool     `yaml:"skipManifiedCandidate"`
	GravityChainBatchSize      uint64   `yaml:"gravityChainBatchSize"`
}

// STATUS represents the status of committee
type STATUS uint8

const (
	// STARTING stands for a starting status
	STARTING STATUS = iota
	// ACTIVE stands for an active status
	ACTIVE
	// INACTIVE stands for an inactive status
	INACTIVE
)

type (
	// Committee defines an interface of an election committee
	// It could be considered as a light state db of gravity chain, that
	Committee interface {
		// Start starts the committee service
		Start(context.Context) error
		// Stop stops the committee service
		Stop(context.Context) error
		// ResultByHeight returns the result on a specific ethereum height
		ResultByHeight(height uint64) (*types.ElectionResult, error)
		// FetchResultByHeight returns the buckets
		FetchResultByHeight(height uint64) (*types.ElectionResult, error)
		// HeightByTime returns the nearest result before time
		HeightByTime(timestamp time.Time) (uint64, error)
		// LatestHeight returns the height with latest result
		LatestHeight() uint64
		// Status returns the committee status
		Status() STATUS
	}

	committee struct {
		archive               Archive
		carrier               carrier.Carrier
		retryLimit            uint8
		paginationSize        uint8
		fetchInParallel       uint8
		skipManifiedCandidate bool
		voteThreshold         *big.Int
		scoreThreshold        *big.Int
		selfStakingThreshold  *big.Int
		interval              uint64

		cache *lru.Cache

		startHeight           uint64
		currentHeight         uint64
		lastUpdateTimestamp   int64
		terminate             chan bool
		mutex                 sync.RWMutex
		gravityChainBatchSize uint64
	}

	rawData struct {
		mintTime          time.Time
		noNewStakingEvent bool
		migration         bool
		buckets           []*types.Bucket
		registrations     []*types.Registration
	}
)

// NewCommittee creates a committee
func NewCommittee(newDB *sql.DB, cfg Config, oldDB db.KVStoreWithNamespace) (Committee, error) {
	if !common.IsHexAddress(cfg.StakingContractAddress) {
		return nil, errors.New("Invalid staking contract address")
	}
	archive, err := NewArchive(newDB, cfg.GravityChainStartHeight, cfg.GravityChainHeightInterval, oldDB)
	if err != nil {
		return nil, err
	}
	carrier, err := carrier.NewEthereumVoteCarrier(
		12,
		time.Minute,
		cfg.GravityChainAPIs,
		common.HexToAddress(cfg.RegisterContractAddress),
		common.HexToAddress(cfg.StakingContractAddress),
	)
	if err != nil {
		return nil, err
	}
	zap.L().Info(
		"Carrier created",
		zap.String("registerContractAddress", cfg.RegisterContractAddress),
		zap.String("stakingContractAddress", cfg.StakingContractAddress),
	)
	voteThreshold, ok := new(big.Int).SetString(cfg.VoteThreshold, 10)
	if !ok {
		return nil, errors.New("Invalid vote threshold")
	}
	scoreThreshold, ok := new(big.Int).SetString(cfg.ScoreThreshold, 10)
	if !ok {
		return nil, errors.New("Invalid score threshold")
	}
	selfStakingThreshold, ok := new(big.Int).SetString(cfg.SelfStakingThreshold, 10)
	if !ok {
		return nil, errors.New("Invalid self staking threshold")
	}
	fetchInParallel := uint8(10)
	if cfg.NumOfFetchInParallel > 0 {
		fetchInParallel = cfg.NumOfFetchInParallel
	}
	gravityChainBatchSize := uint64(10)
	if cfg.GravityChainBatchSize > 0 {
		gravityChainBatchSize = cfg.GravityChainBatchSize
	}
	cache, err := lru.New(int(cfg.CacheSize))
	if err != nil {
		return nil, err
	}
	return &committee{
		archive:               archive,
		cache:                 cache,
		carrier:               carrier,
		retryLimit:            cfg.NumOfRetries,
		paginationSize:        cfg.PaginationSize,
		fetchInParallel:       fetchInParallel,
		skipManifiedCandidate: cfg.SkipManifiedCandidate,
		voteThreshold:         voteThreshold,
		scoreThreshold:        scoreThreshold,
		selfStakingThreshold:  selfStakingThreshold,
		terminate:             make(chan bool),
		startHeight:           cfg.GravityChainStartHeight,
		interval:              cfg.GravityChainHeightInterval,
		currentHeight:         0,
		gravityChainBatchSize: gravityChainBatchSize,
	}, nil
}

func (ec *committee) Start(ctx context.Context) (err error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if err = ec.archive.Start(ctx); err != nil {
		return err
	}

	tip, err := ec.carrier.Tip()
	if err != nil {
		return errors.Wrap(err, "failed to get tip height")
	}
	tipChan := make(chan uint64)
	reportChan := make(chan error)
	go func() {
		zap.L().Info("catching up via network")
		gap := ec.interval * ec.gravityChainBatchSize
		for h := ec.nextHeight() + gap; h < tip; h += gap {
			zap.L().Info("catching up to", zap.Uint64("height", h))
			data, err := ec.fetchInBatch(h)
			if err != nil {
				zap.L().Error("failed to fetch data", zap.Error(err))
			}
			if err := ec.storeInBatch(data); err != nil {
				zap.L().Error("failed to catch up via network", zap.Uint64("height", h), zap.Error(err))
			}
		}
		data, err := ec.fetchInBatch(tip)
		if err != nil {
			zap.L().Error("failed to fetch data", zap.Error(err))
		}
		if err := ec.storeInBatch(data); err != nil {
			zap.L().Error("failed to catch up via network", zap.Error(err))
		}
		zap.L().Info("subscribing to new block")
		ec.carrier.SubscribeNewBlock(tipChan, reportChan, ec.terminate)
		for {
			select {
			case <-ec.terminate:
				ec.terminate <- true
				return
			case tip := <-tipChan:
				zap.L().Info("new ethereum block", zap.Uint64("height", tip))
				if err := ec.Sync(tip); err != nil {
					zap.L().Error("failed to sync", zap.Error(err))
				}
			case err := <-reportChan:
				zap.L().Error("something goes wrong", zap.Error(err))
			}
		}
	}()
	return nil
}

func (ec *committee) Stop(ctx context.Context) error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	ec.terminate <- true
	ec.carrier.Close()

	return ec.archive.Stop(ctx)
}

func (ec *committee) Status() STATUS {
	lastUpdateTimestamp := atomic.LoadInt64(&ec.lastUpdateTimestamp)
	switch {
	case lastUpdateTimestamp == 0:
		return STARTING
	case lastUpdateTimestamp > time.Now().Add(-5*time.Minute).Unix():
		return ACTIVE
	default:
		return INACTIVE
	}
}

func (ec *committee) Sync(tipHeight uint64) error {
	data, err := ec.fetchInBatch(tipHeight)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	return ec.storeInBatch(data)
}

func (ec *committee) nextHeight() uint64 {
	height := ec.latestHeightInArchive()
	if height == 0 {
		return ec.startHeight
	}
	return height + ec.interval
}

func (ec *committee) fetchInBatch(tipHeight uint64) (retval map[uint64]*rawData, err error) {
	if ec.currentHeight < tipHeight {
		ec.currentHeight = tipHeight
	}
	retval = map[uint64]*rawData{}
	var wg sync.WaitGroup
	var lock sync.RWMutex
	limiter := make(chan bool, ec.fetchInParallel)
	for nextHeight := ec.nextHeight(); nextHeight <= ec.currentHeight; nextHeight += ec.interval {
		wg.Add(1)
		go func(height uint64) {
			defer func() {
				<-limiter
				wg.Done()
			}()
			limiter <- true
			data, e := ec.retryFetchDataByHeight(height)
			lock.Lock()
			defer lock.Unlock()
			retval[height] = data
			if e != nil {
				err = e
			}
		}(nextHeight)
	}
	wg.Wait()

	return
}

func (ec *committee) storeInBatch(data map[uint64]*rawData) error {
	heights := make([]uint64, 0, len(data))
	mintTimes := make([]time.Time, 0, len(data))
	arrOfRegs := make([][]*types.Registration, 0, len(data))
	arrOfBuckets := make([][]*types.Bucket, 0, len(data))
	for height := range data {
		heights = append(heights, height)
		mintTimes = append(mintTimes, data[height].mintTime)
		arrOfRegs = append(arrOfRegs, data[height].registrations)
		if data[height].noNewStakingEvent {
			arrOfBuckets = append(arrOfBuckets, nil)
		} else {
			arrOfBuckets = append(arrOfBuckets, data[height].buckets)
		}
	}
	if err := ec.archive.PutPolls(heights, mintTimes, arrOfRegs, arrOfBuckets); err != nil {
		return err
	}
	atomic.StoreInt64(&ec.lastUpdateTimestamp, time.Now().Unix())
	return nil
}

func (ec *committee) LatestHeight() uint64 {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.latestHeightInArchive()
}

func (ec *committee) latestHeightInArchive() uint64 {
	tipHeight, err := ec.archive.TipHeight()
	if err != nil {
		return 0
	}
	return tipHeight
}

func (ec *committee) HeightByTime(ts time.Time) (uint64, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	// Make sure that we already got a block after the timestamp, such that the height
	// we return here is the last one before ts
	return ec.archive.HeightBefore(ts)
}

func (ec *committee) ResultByHeight(height uint64) (*types.ElectionResult, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.resultByHeight(height)
}

func (ec *committee) resultByHeight(height uint64) (*types.ElectionResult, error) {
	zap.L().Info("fetch result from DB and calculate", zap.Uint64("height", height))
	if height < ec.startHeight {
		return nil, errors.Errorf(
			"height %d is lower than start height %d",
			height,
			ec.startHeight,
		)
	}
	if (height-ec.startHeight)%ec.interval != 0 {
		return nil, errors.Errorf(
			"height %d is an invalid height",
			height,
		)
	}

	if cacheResult, ok := ec.cache.Get(height); ok {
		if result, as := cacheResult.(*types.ElectionResult); as {
			return result, nil
		}
		return nil, errors.Errorf(
			"lru cache type assertion has error",
		)
	}

	//calculate the result from DB
	calculator, err := ec.calculator(height, true)
	if err != nil {
		return nil, err
	}
	regs, err := ec.archive.Registrations(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddRegistrations(regs); err != nil {
		return nil, err
	}
	buckets, err := ec.archive.Buckets(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddBuckets(buckets); err != nil {
		return nil, err
	}
	result, err := calculator.Calculate()
	if err != nil {
		return nil, err
	}
	ec.cache.Add(height, result)

	return result, nil
}

func (ec *committee) calcWeightedVotes(v *types.Bucket, now time.Time) *big.Int {
	if now.Before(v.StartTime()) {
		return big.NewInt(0)
	}
	remainingTime := v.RemainingTime(now).Seconds()
	weight := float64(1)
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)) / math.Log(1.2) / 100
	}
	amount := new(big.Float).SetInt(v.Amount())
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)

	return weightedAmount
}

func (ec *committee) fetchBucketsByHeight(height uint64, force bool) (bool, []*types.Bucket, error) {
	if height > ec.interval && height != ec.startHeight && !force {
		if !ec.carrier.HasStakingEvents(new(big.Int).SetUint64(height-ec.interval+1), new(big.Int).SetUint64(height)) {
			return true, nil, nil
		}
	}
	buckets, err := ec.fetchBucketsFromEthereum(height)

	return false, buckets, err
}

func (ec *committee) fetchBucketsFromEthereum(height uint64) ([]*types.Bucket, error) {
	var allBuckets []*types.Bucket
	previousIndex := big.NewInt(0)
	for {
		var buckets []*types.Bucket
		var err error
		if previousIndex, buckets, err = ec.carrier.Buckets(
			height,
			previousIndex,
			ec.paginationSize,
		); err != nil {
			return nil, err
		}
		allBuckets = append(allBuckets, buckets...)
		if len(buckets) < int(ec.paginationSize) {
			break
		}
	}
	zap.L().Debug("fetch buckets by height from ethereum", zap.Int("number of buckets", len(allBuckets)))
	return allBuckets, nil
}

func (ec *committee) bucketFilter(v *types.Bucket) bool {
	return ec.voteThreshold.Cmp(v.Amount()) > 0
}

func (ec *committee) candidateFilter(c *types.Candidate) bool {
	return ec.selfStakingThreshold.Cmp(c.SelfStakingTokens()) > 0 ||
		ec.scoreThreshold.Cmp(c.Score()) > 0
}

func (ec *committee) getMintTimeByHeight(height uint64) (time.Time, error) {
	mintTime, err := ec.carrier.BlockTimestamp(height)
	switch errors.Cause(err) {
	case nil:
		break
	case ethereum.NotFound:
		return mintTime, db.ErrNotExist
	default:
		return mintTime, err
	}
	return mintTime, nil
}

func (ec *committee) calculator(height uint64, dbflag bool) (*types.ResultCalculator, error) {
	var timestamp time.Time
	var err error
	if dbflag {
		timestamp, err = ec.archive.MintTime(height)
	} else {
		timestamp, err = ec.getMintTimeByHeight(height)
	}
	if err != nil {
		return nil, err
	}

	return types.NewResultCalculator(
		timestamp,
		ec.skipManifiedCandidate,
		ec.bucketFilter,
		ec.calcWeightedVotes,
		ec.candidateFilter,
	), nil
}

func (ec *committee) fetchRegistrationsByHeight(height uint64) ([]*types.Registration, error) {
	var allCandidates []*types.Registration
	previousIndex := big.NewInt(1)
	for {
		var candidates []*types.Registration
		var err error
		if previousIndex, candidates, err = ec.carrier.Registrations(
			height,
			previousIndex,
			ec.paginationSize,
		); err != nil {
			return nil, err
		}
		allCandidates = append(allCandidates, candidates...)
		if len(candidates) < int(ec.paginationSize) {
			break
		}
	}
	zap.L().Debug("fetch registrations by height from ethereum", zap.Int("number of registrations", len(allCandidates)))
	return allCandidates, nil
}

func (ec *committee) FetchResultByHeight(height uint64) (*types.ElectionResult, error) {
	if height == 0 {
		var err error
		height, err = ec.carrier.Tip()
		if err != nil {
			return nil, err
		}
	}
	return ec.fetchResultByHeight(height)
}

func (ec *committee) fetchResultByHeight(height uint64) (*types.ElectionResult, error) {
	zap.L().Info("fetch result from ethereum", zap.Uint64("height", height))
	calculator, err := ec.calculator(height, false)
	if err != nil {
		return nil, err
	}
	regs, err := ec.fetchRegistrationsByHeight(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddRegistrations(regs); err != nil {
		return nil, err
	}
	_, buckets, err := ec.fetchBucketsByHeight(height, true)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddBuckets(buckets); err != nil {
		return nil, err
	}

	return calculator.Calculate()
}

func (ec *committee) fetchDataByHeight(height uint64) (*rawData, error) {
	zap.L().Info("fetch from ethereum", zap.Uint64("height", height))
	regs, err := ec.fetchRegistrationsByHeight(height)
	if err != nil {
		return nil, err
	}
	noChange, buckets, err := ec.fetchBucketsByHeight(height, false)
	if err != nil {
		return nil, err
	}
	mintTime, err := ec.getMintTimeByHeight(height)
	if err != nil {
		return nil, err
	}

	return &rawData{
		mintTime:          mintTime,
		noNewStakingEvent: noChange,
		registrations:     regs,
		buckets:           buckets,
	}, nil
}

func atos(a []int64) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(b, ",")
}

func (ec *committee) retryFetchDataByHeight(height uint64) (data *rawData, err error) {
	for i := uint8(0); i < ec.retryLimit; i++ {
		if data, err = ec.fetchDataByHeight(height); err == nil {
			break
		}
		zap.L().Error(
			"failed to fetch result by height",
			zap.Error(err),
			zap.Uint64("height", height),
			zap.Uint8("tried", i+1),
		)
	}
	return
}
