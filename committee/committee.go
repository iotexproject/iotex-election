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
	"encoding/hex"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"

	// require sqlite3 driver
	"github.com/pkg/errors"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"

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
	GravityChainCeilingHeight  uint64   `yaml:"gravityChainCeilingHeight"`
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

// EthHardForkHeight stands for the height of ethereum hard fork
const EthHardForkHeight = 8581700

type (
	// Committee defines an interface of an election committee
	// It could be considered as a light state db of gravity chain, that
	Committee interface {
		// Start starts the committee service
		Start(context.Context) error
		// Stop stops the committee service
		Stop(context.Context) error
		// ResultByHeight returns the result on a specific ethereum height
		ResultByHeight(uint64) (*types.ElectionResult, error)
		//RawDataByHeight returns the bucket list and registration list and mintTime
		RawDataByHeight(uint64) ([]*types.Bucket, []*types.Registration, time.Time, error)
		// HeightByTime returns the nearest result before time
		HeightByTime(time.Time) (uint64, error)
		// LatestHeight returns the height with latest result
		LatestHeight() uint64
		// Status returns the committee status
		Status() STATUS
		// PutNativePollByEpoch puts one native poll record on IoTeX chain
		PutNativePollByEpoch(uint64, time.Time, []*types.Bucket) error
		// NativeBucketsByEpoch returns a list of Bucket of a given epoch number
		NativeBucketsByEpoch(uint64) ([]*types.Bucket, error)
	}

	committee struct {
		archive               PollArchive
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
		terminatedCarrier     bool
		terminatedArchive     bool
		mutex                 sync.RWMutex
		gravityChainBatchSize uint64
		ceilingHeight         uint64
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
func NewCommittee(archive PollArchive, cfg Config) (Committee, error) {
	if !common.IsHexAddress(cfg.StakingContractAddress) {
		return nil, errors.New("Invalid staking contract address")
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
		terminatedCarrier:     false,
		terminatedArchive:     false,
		startHeight:           cfg.GravityChainStartHeight,
		ceilingHeight:         cfg.GravityChainCeilingHeight,
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
	ceilingHeight := ec.ceilingHeight
	if ceilingHeight >= ec.interval {
		ceilingHeight -= ec.interval
	}
	if ec.latestHeightInArchive() >= ceilingHeight && ec.ceilingHeight != 0 {
		zap.L().Info("stop syncing")
		ec.terminateCarrier(ctx)
		return nil
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
				return
			case tip := <-tipChan:
				if ec.currentHeight >= ec.ceilingHeight && ec.ceilingHeight != 0 {
					zap.L().Info("stop syncing")
					ec.terminateCarrier(ctx)
					return
				}

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

func (ec *committee) terminateCarrier(ctx context.Context) {
	if !ec.terminatedCarrier {
		close(ec.terminate)
		ec.carrier.Close()
		ec.terminatedCarrier = true
	}
}

func (ec *committee) terminateArchive(ctx context.Context) error {
	if !ec.terminatedArchive {
		ec.terminatedArchive = true
		return ec.archive.Stop(ctx)
	}
	return nil
}

func (ec *committee) Stop(ctx context.Context) error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	ec.terminateCarrier(ctx)
	return ec.terminateArchive(ctx)

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

func (ec *committee) PutNativePollByEpoch(epochNum uint64, mintTime time.Time, buckets []*types.Bucket) error {
	return ec.archive.PutNativePoll(epochNum, mintTime, buckets)
}

func (ec *committee) NativeBucketsByEpoch(epochNum uint64) ([]*types.Bucket, error) {
	return ec.archive.NativeBuckets(epochNum)
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
	indice := map[uint64]int{}
	for i, height := range heights {
		if _, ok := indice[height]; ok {
			return errors.Errorf("duplicate height %d", height)
		}
		indice[height] = i
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	for _, height := range heights {
		index := indice[height]
		if err := ec.archive.PutPoll(height, mintTimes[index], arrOfRegs[index], arrOfBuckets[index]); err != nil {
			return err
		}
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

func (ec *committee) RawDataByHeight(height uint64) ([]*types.Bucket, []*types.Registration, time.Time, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.rawDataByHeight(height)
}

func (ec *committee) rawDataByHeight(height uint64) ([]*types.Bucket, []*types.Registration, time.Time, error) {
	timestamp, err := ec.archive.MintTime(height)
	if err != nil {
		return nil, nil, time.Time{}, err
	}
	regs, err := ec.archive.Registrations(height)
	if err != nil {
		return nil, nil, time.Time{}, err
	}
	buckets, err := ec.archive.Buckets(height)
	if err != nil {
		return nil, nil, time.Time{}, err
	}
	return buckets, regs, timestamp, nil
}

func (ec *committee) ResultByHeight(height uint64) (*types.ElectionResult, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.resultByHeight(height)
}

func (ec *committee) resultByHeight(height uint64) (*types.ElectionResult, error) {
	zap.L().Debug("fetch result from DB and calculate", zap.Uint64("height", height))
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
	if err := ec.handleEthereumHardFork(height, result); err != nil {
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
	res, err := calculator.Calculate()
	if err != nil {
		return nil, err
	}
	if err := ec.handleEthereumHardFork(height, res); err != nil {
		return nil, err
	}
	return res, nil
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

func (ec *committee) handleEthereumHardFork(height uint64, result *types.ElectionResult) error {
	if height != EthHardForkHeight {
		return nil
	}
	for _, delegate := range result.Delegates() {
		name := hex.EncodeToString(delegate.Name())
		switch name {
		case "000000696f746578636f7265":
			score, ok := new(big.Int).SetString("85373235544231218078559584", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "00000000006d6574616e7978":
			score, ok := new(big.Int).SetString("59632560935643656968902530", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "67616d6566616e7461737900":
			score, ok := new(big.Int).SetString("53200765151851838442704552", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "00000000707265616e67656c":
			score, ok := new(big.Int).SetString("50419789330925706718338211", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "00007976616c696461746f72":
			score, ok := new(big.Int).SetString("49956076291800440218188229", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "000000636f696e6765636b6f":
			score, ok := new(big.Int).SetString("45541967921325925112783154", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "000000696f7465787465616d":
			score, ok := new(big.Int).SetString("42048368188254741149181523", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "0000626c6f636b666f6c696f":
			score, ok := new(big.Int).SetString("40118847473353343676639849", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "696f7478706c6f726572696f":
			score, ok := new(big.Int).SetString("38637407472542613934244717", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "0068756f626977616c6c6574":
			score, ok := new(big.Int).SetString("19826333897499304850764901", 10)
			if ok {
				delegate.SetScore(score)
			}
		case "636f6e73656e7375736e6574":
			score, ok := new(big.Int).SetString("4562820963931216603918007", 10)
			if ok {
				delegate.SetScore(score)
			}
		default:
			continue
		}
	}
	return nil
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
