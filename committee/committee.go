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
	"crypto/sha256"
	"math"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

const (
	// PollNamespace is the bucket name for election Result
	PollNamespace = "Poll"

	// BlockTimeNamespace is the bucket name for election Time
	BlockTimeNamespace = "BlockTime"

	// BucketNamespace is the bucket name for Bucket
	BucketNamespace = "Bucket"

	// RegistrationNamespace is the bucket name for Registration
	RegistrationNamespace = "Registration"
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
		// FetchResultByHeight returns the votes
		FetchResultByHeight(height uint64) (*types.ElectionResult, error)
		// HeightByTime returns the nearest result before time
		HeightByTime(timestamp time.Time) (uint64, error)
		// LatestHeight returns the height with latest result
		LatestHeight() uint64
		// Status returns the committee status
		Status() STATUS
	}

	committee struct {
		db                    db.KVStoreWithNamespace
		carrier               carrier.Carrier
		retryLimit            uint8
		paginationSize        uint8
		fetchInParallel       uint8
		skipManifiedCandidate bool
		voteThreshold         *big.Int
		scoreThreshold        *big.Int
		selfStakingThreshold  *big.Int
		interval              uint64

		cache         *resultCache
		heightManager *heightManager

		startHeight           uint64
		nextHeight            uint64
		currentHeight         uint64
		lastUpdateTimestamp   int64
		terminate             chan bool
		mutex                 sync.RWMutex
		gravityChainBatchSize uint64
	}

	rawData struct {
		mintTime   time.Time
		candidates [][]byte
		buckets    [][]byte
	}
)

// NewCommittee creates a committee
func NewCommittee(db db.KVStoreWithNamespace, cfg Config) (Committee, error) {
	if !common.IsHexAddress(cfg.StakingContractAddress) {
		return nil, errors.New("Invalid staking contract address")
	}
	carrier, err := carrier.NewEthereumVoteCarrier(
		cfg.GravityChainAPIs,
		common.HexToAddress(cfg.RegisterContractAddress),
		common.HexToAddress(cfg.StakingContractAddress),
	)
	zap.L().Info(
		"Carrier created",
		zap.String("registerContractAddress", cfg.RegisterContractAddress),
		zap.String("stakingContractAddress", cfg.StakingContractAddress),
	)
	if err != nil {
		return nil, err
	}
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
	return &committee{
		db:                    db,
		cache:                 newResultCache(cfg.CacheSize),
		heightManager:         newHeightManager(),
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
		nextHeight:            cfg.GravityChainStartHeight,
		gravityChainBatchSize: gravityChainBatchSize,
	}, nil
}

func (ec *committee) Start(ctx context.Context) (err error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if err := ec.db.Start(ctx); err != nil {
		return errors.Wrap(err, "error when starting db")
	}
	if startHeight, err := ec.getTimeDB(db.NextHeightKey); err == nil {
		zap.L().Info("restoring from db")
		ec.nextHeight = util.BytesToUint64(startHeight)
		for height := ec.startHeight; height < ec.nextHeight; height += ec.interval {
			zap.L().Info("loading", zap.Uint64("height", height))
			heightKey := ec.dbKey(height)
			data, err := ec.getTimeDB(heightKey)
			if err != nil {
				return err
			}
			time, err := util.BytesToTime(data)
			if err != nil {
				return err
			}
			if err := ec.heightManager.add(height, time); err != nil {
				return err
			}
		}
	}

	tip, err := ec.carrier.Tip()
	if err != nil {
		return errors.Wrap(err, "failed to get tip height")
	}
	tipChan := make(chan *carrier.TipInfo)
	reportChan := make(chan error)
	go func() {
		zap.L().Info("catching up via network")
		gap := ec.interval * ec.gravityChainBatchSize
		for h := ec.nextHeight + gap; h < tip.Height; h += gap {
			zap.L().Info("catching up to", zap.Uint64("height", h))
			data, err := ec.fetchInBatch(h)
			if err != nil {
				zap.L().Error("failed to fetch data", zap.Error(err))
			}
			t, err := ec.carrier.BlockTimestamp(h)
			if err != nil {
				zap.L().Error("failed to get block timestamp", zap.Uint64("height", h), zap.Error(err))
			}
			if err := ec.storeInBatch(data, t); err != nil {
				zap.L().Error("failed to catch up via network", zap.Uint64("height", h), zap.Error(err))
			}
		}
		data, err := ec.fetchInBatch(tip.Height)
		if err != nil {
			zap.L().Error("failed to fetch data", zap.Error(err))
		}
		if err := ec.storeInBatch(data, tip.BlockTime); err != nil {
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
				zap.L().Info("new ethereum block", zap.Uint64("height", tip.Height))
				if err := ec.Sync(tip.Height, tip.BlockTime); err != nil {
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

	return ec.db.Stop(ctx)
}

func (ec *committee) Status() STATUS {
	lastUpdateTimestamp := atomic.LoadInt64(&ec.lastUpdateTimestamp)
	switch {
	case lastUpdateTimestamp == 0:
		return STARTING
	case lastUpdateTimestamp > time.Now().Add(-60*time.Second).Unix():
		return ACTIVE
	default:
		return INACTIVE
	}
}

func (ec *committee) Sync(tipHeight uint64, tipTime time.Time) error {
	data, err := ec.fetchInBatch(tipHeight)
	if err != nil {
		return err
	}
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	return ec.storeInBatch(data, tipTime)
}

func (ec *committee) fetchInBatch(tipHeight uint64) (retval map[uint64]*rawData, err error) {
	if ec.currentHeight < tipHeight {
		ec.currentHeight = tipHeight
	}
	retval = map[uint64]*rawData{}
	var wg sync.WaitGroup
	var lock sync.RWMutex
	limiter := make(chan bool, ec.fetchInParallel)
	for nextHeight := ec.nextHeight; nextHeight <= ec.currentHeight-12; nextHeight += ec.interval {
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

func (ec *committee) storeInBatch(
	data map[uint64]*rawData,
	tipTime time.Time,
) error {
	var heights []uint64
	for height := range data {
		heights = append(heights, height)
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	for _, height := range heights {
		if err := ec.storeData(height, data[height]); err != nil {
			return errors.Wrapf(err, "failed to store result of height %d", height)
		}
		ec.nextHeight = height + ec.interval
	}
	zap.L().Info("synced to", zap.Time("block time", tipTime))
	atomic.StoreInt64(&ec.lastUpdateTimestamp, tipTime.Unix())

	return nil
}

func (ec *committee) LatestHeight() uint64 {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	l := len(ec.heightManager.heights)
	if l == 0 {
		return 0
	}
	return ec.heightManager.heights[l-1]
}

func (ec *committee) HeightByTime(ts time.Time) (uint64, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	// Make sure that we already got a block after the timestamp, such that the height
	// we return here is the last one before ts
	lastUpdateTimestamp := atomic.LoadInt64(&ec.lastUpdateTimestamp)
	if !time.Unix(lastUpdateTimestamp, 0).After(ts) {
		return 0, db.ErrNotExist
	}
	height := ec.heightManager.nearestHeightBefore(ts)
	if height == 0 {
		return 0, db.ErrNotExist
	}

	return height, nil
}

func (ec *committee) ResultByHeight(height uint64) (*types.ElectionResult, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	return ec.resultByHeight(height)
}

func (ec *committee) resultByHeight(height uint64) (*types.ElectionResult, error) {
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
	result := ec.cache.get(height)
	if result != nil {
		return result, nil
	}

	// if cache doesn't have corresponding result, read the resultMeta from db
	heightKey := ec.dbKey(height)
	data, err := ec.getResultDB(heightKey)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, db.ErrNotExist
	}
	poll := &types.Poll{}
	if err := poll.Deserialize(data); err != nil {
		return nil, err
	}

	//calculate the result from resultMeta
	calculator, err := ec.calculator(height)
	if err != nil {
		return nil, err
	}

	regs, err := ec.registrations(poll.Registrations())
	if err != nil {
		return nil, err
	}
	if err := calculator.AddRegistrations(regs); err != nil {
		return nil, err
	}

	buckets, err := ec.buckets(poll.Buckets())
	if err != nil {
		return nil, err
	}
	if err := calculator.AddBuckets(buckets); err != nil {
		return nil, err
	}

	result, err = calculator.Calculate()
	if err != nil {
		return nil, err
	}
	ec.cache.insert(height, result)

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

func (ec *committee) fetchBucketsByHeight(height uint64) ([]*types.Bucket, error) {
	var allVotes []*types.Bucket
	previousIndex := big.NewInt(0)
	for {
		var votes []*types.Bucket
		var err error
		if previousIndex, votes, err = ec.carrier.Buckets(
			height,
			previousIndex,
			ec.paginationSize,
		); err != nil {
			return nil, err
		}
		allVotes = append(allVotes, votes...)
		if len(votes) < int(ec.paginationSize) {
			break
		}
	}
	zap.L().Debug("fetch buckets by height from ethereum", zap.Int("number of buckets", len(allVotes)))
	return allVotes, nil
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

func (ec *committee) calculator(height uint64) (*types.ResultCalculator, error) {
	timestamp, err := ec.getMintTimeByHeight(height)
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
		tip, err := ec.carrier.Tip()
		if err != nil {
			return nil, err
		}
		height = tip.Height
	}
	return ec.fetchResultByHeight(height)
}

func (ec *committee) fetchResultByHeight(height uint64) (*types.ElectionResult, error) {
	zap.L().Info("fetch result from ethereum", zap.Uint64("height", height))
	calculator, err := ec.calculator(height)
	if err != nil {
		return nil, err
	}
	candidates, err := ec.fetchRegistrationsByHeight(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddRegistrations(candidates); err != nil {
		return nil, err
	}
	votes, err := ec.fetchBucketsByHeight(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddBuckets(votes); err != nil {
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
	rhs, err := ec.storeRegistrations(regs)
	if err != nil {
		return nil, err
	}
	buckets, err := ec.fetchBucketsByHeight(height)
	if err != nil {
		return nil, err
	}
	bhs, err := ec.storeBuckets(buckets)
	if err != nil {
		return nil, err
	}
	mintTime, err := ec.getMintTimeByHeight(height)
	if err != nil {
		return nil, err
	}

	return &rawData{mintTime, rhs, bhs}, nil
}

func (ec *committee) dbKey(height uint64) []byte {
	return util.Uint64ToBytes(height)
}

func (ec *committee) getTimeDB(key []byte) ([]byte, error) {
	return ec.db.Get(BlockTimeNamespace, key)
}

func (ec *committee) putTimeDB(key []byte, value []byte) error {
	return ec.db.Put(BlockTimeNamespace, key, value)
}

func (ec *committee) getResultDB(key []byte) ([]byte, error) {
	return ec.db.Get(PollNamespace, key)
}

func (ec *committee) getVoteDB(key []byte) ([]byte, error) {
	return ec.db.Get(BucketNamespace, key)
}

func (ec *committee) getCandidateDB(key []byte) ([]byte, error) {
	return ec.db.Get(RegistrationNamespace, key)
}

func (ec *committee) storeData(height uint64, data *rawData) error {
	if err := ec.heightManager.validate(height, data.mintTime); err != nil {
		zap.L().Fatal(
			"Unexpected status that the upcoming block height or time is invalid",
			zap.Error(err),
		)
	}
	ser, err := types.NewPoll(data.buckets, data.candidates).Serialize()
	if err != nil {
		return err
	}
	timeData, err := util.TimeToBytes(data.mintTime)
	if err != nil {
		return err
	}
	heightKey := ec.dbKey(height)
	batchForResult := db.NewBatch()
	batchForResult.Put(PollNamespace, heightKey, ser, "failed to put election result into db")
	batchForResult.Put(BlockTimeNamespace, heightKey, timeData, "failed to put election time into db")
	batchForResult.Put(BlockTimeNamespace, db.NextHeightKey, util.Uint64ToBytes(height+ec.interval), "failed to put next height into db")
	if err := ec.db.Commit(batchForResult); err != nil {
		return err
	}
	return ec.heightManager.add(height, data.mintTime)
}

func (ec *committee) storeBuckets(buckets []*types.Bucket) ([][]byte, error) {
	batchForVote := db.NewBatch()
	hashes := make([][]byte, len(buckets))
	for i, v := range buckets {
		data, err := v.Serialize()
		if err != nil {
			return nil, err
		}
		hashval := sha256.Sum256(data)
		hashbytes := hashval[:]

		if _, err := ec.getVoteDB(hashbytes); err != nil {
			batchForVote.Put(BucketNamespace, hashbytes, data, "Failed to put a vote into DB")
		}
		hashes[i] = hashbytes
	}
	return hashes, ec.db.Commit(batchForVote)
}

func (ec *committee) storeRegistrations(registrations []*types.Registration) ([][]byte, error) {
	batchForCandidate := db.NewBatch()
	hashes := make([][]byte, len(registrations))
	for i, c := range registrations {
		data, err := c.Serialize()
		if err != nil {
			return nil, err
		}
		hashval := sha256.Sum256(data)
		hashbytes := hashval[:]
		if _, err := ec.getCandidateDB(hashbytes); err != nil {
			batchForCandidate.Put(RegistrationNamespace, hashbytes, data, "Failed to put a candidate into DB")
		}
		hashes[i] = hashbytes
	}
	return hashes, ec.db.Commit(batchForCandidate)
}

func (ec *committee) buckets(keys [][]byte) ([]*types.Bucket, error) {
	var votes []*types.Bucket
	for _, hash := range keys {
		if len(hash) != 32 {
			zap.L().Error("The length of the hash value should be 32")
		}
		data, err := ec.getVoteDB(hash)
		if err != nil {
			return nil, err
		}
		vote := &types.Bucket{}
		if err := vote.Deserialize(data); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func (ec *committee) registrations(keys [][]byte) ([]*types.Registration, error) {
	var cands []*types.Registration
	for _, hash := range keys {
		if len(hash) != 32 {
			zap.L().Error("The length of the hash value should be 32")
		}
		data, err := ec.getCandidateDB(hash)
		if err != nil {
			return nil, err
		}
		cand := &types.Registration{}
		if err := cand.Deserialize(data); err != nil {
			return nil, err
		}
		cands = append(cands, cand)
	}
	return cands, nil
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
