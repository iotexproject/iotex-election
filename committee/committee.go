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
	"fmt"
	"math"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

// NextHeightKey defines the constant key of next height
var NextHeightKey = int64(0)

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
		oldDB                 db.KVStoreWithNamespace
		db                    *sql.DB
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
		mintTime          time.Time
		noNewStakingEvent bool
		migration         bool
		buckets           []*types.Bucket
		registrations     []*types.Registration
	}
)

func (ec *committee) createTables() error {
	tableCreations := []string{
		"CREATE TABLE IF NOT EXISTS buckets (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB UNIQUE, start_time TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB)",
		"CREATE TABLE IF NOT EXISTS registrations (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB UNIQUE, name BLOB, address BLOB, operator_address BLOB, reward_address BLOB, self_staking_weight INTEGER)",
		"CREATE TABLE IF NOT EXISTS height_to_registrations (height INTEGER, rid INTEGER REFERENCES registrations(id), CONSTRAINT key PRIMARY KEY (height, rid))",
		"CREATE TABLE IF NOT EXISTS height_to_buckets (height INTEGER, bid INTEGER REFERENCES buckets(id), times INTEGER, CONSTRAINT key PRIMARY KEY (height, bid))",
		"CREATE TABLE IF NOT EXISTS mint_times (height INTEGER PRIMARY KEY, time TIMESTAMP)",
		"CREATE TABLE IF NOT EXISTS next_height (key INTEGER PRIMARY KEY, height integer)",
		"CREATE TABLE IF NOT EXISTS identical_buckets (height INTEGER PRIMARY KEY, identical_to INTEGER)",
		"CREATE TABLE IF NOT EXISTS identical_registrations (height INTEGER PRIMARY KEY, identical_to INTEGER)",
	}
	for _, creation := range tableCreations {
		if _, err := ec.db.Exec(creation); err != nil {
			return err
		}
	}
	return nil
}

// NewCommittee creates a committee
func NewCommittee(newDB *sql.DB, cfg Config, oldDB db.KVStoreWithNamespace) (Committee, error) {
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
		oldDB:                 oldDB,
		db:                    newDB,
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

func (ec *committee) heightWithIdenticalRegs(height uint64) (uint64, error) {
	var val int64
	row := ec.db.QueryRow("SELECT identical_to FROM identical_registrations WHERE height = ?", util.Uint64ToInt64(height))
	err := row.Scan(&val)
	switch err {
	case nil:
		return uint64(val), nil
	case sql.ErrNoRows:
		return height, nil
	default:
		return 0, err
	}
}

func (ec *committee) heightWithIdenticalBuckets(height uint64) (uint64, error) {
	var val int64
	row := ec.db.QueryRow("SELECT identical_to FROM identical_buckets WHERE height = ?", util.Uint64ToInt64(height))
	err := row.Scan(&val)
	switch err {
	case nil:
		return uint64(val), nil
	case sql.ErrNoRows:
		return height, nil
	default:
		return 0, err
	}
}

func (ec *committee) hasIdenticalRegistrations(
	regs [][]byte,
	heightToCompare uint64,
) bool {
	if heightToCompare < ec.startHeight {
		return false
	}
	lastRegHashes, err := ec.registrationHashes(heightToCompare)
	if err != nil {
		return false
	}
	if len(regs) != len(lastRegHashes) {
		return false
	}
	rhs := map[hash.Hash256]bool{}
	for _, h := range regs {
		if err != nil {
			return false
		}
		if _, ok := rhs[hash.Hash256b(h)]; ok {
			return false
		}
		rhs[hash.Hash256b(h)] = true
	}
	for _, h := range lastRegHashes {
		if _, ok := rhs[h]; !ok {
			return false
		}
		rhs[h] = false
	}
	return true
}

func (ec *committee) hasIdenticalBuckets(
	buckets map[hash.Hash256]int,
	heightToCompare uint64,
) bool {
	if heightToCompare < ec.startHeight {
		return false
	}
	lastBucketHashes, err := ec.bucketHashes(heightToCompare)
	if err != nil {
		return false
	}
	if len(buckets) != len(lastBucketHashes) {
		return false
	}
	for h, times := range buckets {
		last, ok := lastBucketHashes[h]
		if !ok {
			return false
		}
		if last != times {
			return false
		}
	}
	return true
}

func (ec *committee) migrateResult(height uint64, r *types.ElectionResult) error {
	candidates := r.Delegates()
	regs := make([]*types.Registration, 0, len(candidates))
	for _, candidate := range candidates {
		regs = append(regs, &candidate.Registration)
	}
	votes := r.Votes()
	buckets := make([]*types.Bucket, 0, len(votes))
	for _, vote := range votes {
		buckets = append(buckets, &vote.Bucket)
	}
	return ec.storeData(height, &rawData{
		mintTime:      r.MintTime(),
		registrations: regs,
		buckets:       buckets,
		migration:     true,
	})
}

func (ec *committee) migrate(ctx context.Context) error {
	if ec.oldDB == nil {
		return nil
	}
	kvstore := db.NewKVStoreWithNamespaceWrapper("electionNS", ec.oldDB)
	if err := kvstore.Start(ctx); err != nil {
		return err
	}
	nextHeightHash, err := kvstore.Get(db.NextHeightKey)
	if err != nil {
		return err
	}
	nextHeight := util.BytesToUint64(nextHeightHash)
	for height := ec.startHeight; height < nextHeight; height += ec.interval {
		data, err := kvstore.Get(util.Uint64ToBytes(height))
		if err != nil {
			return err
		}
		r := &types.ElectionResult{}
		if err := r.Deserialize(data); err != nil {
			return err
		}
		fmt.Println("migrate result", height)
		if err := ec.migrateResult(height, r); err != nil {
			return err
		}
	}
	return nil
}

func (ec *committee) Start(ctx context.Context) (err error) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if err = ec.createTables(); err != nil {
		return err
	}
	if err = ec.migrate(ctx); err != nil {
		return errors.Wrap(err, "failed to migrate")
	}
	if nextHeight, err := ec.loadNextHeight(); err == nil {
		zap.L().Info("restoring from db")
		ec.nextHeight = nextHeight
		for height := ec.startHeight; height < ec.nextHeight; height += ec.interval {
			zap.L().Info("loading", zap.Uint64("height", height))
			mintTime, err := ec.mintTime(height)
			if err != nil {
				return err
			}
			if err := ec.heightManager.add(height, mintTime); err != nil {
				return err
			}
		}
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
		for h := ec.nextHeight + gap; h < tip; h += gap {
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

	return nil
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

func (ec *committee) Sync(tipHeight uint64) error {
	data, err := ec.fetchInBatch(tipHeight)
	if err != nil {
		return err
	}
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	return ec.storeInBatch(data)
}

func (ec *committee) fetchInBatch(tipHeight uint64) (retval map[uint64]*rawData, err error) {
	if ec.currentHeight < tipHeight {
		ec.currentHeight = tipHeight
	}
	retval = map[uint64]*rawData{}
	var wg sync.WaitGroup
	var lock sync.RWMutex
	limiter := make(chan bool, ec.fetchInParallel)
	for nextHeight := ec.nextHeight; nextHeight <= ec.currentHeight; nextHeight += ec.interval {
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
	var heights []uint64
	for height := range data {
		heights = append(heights, height)
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	var latestBlockTime time.Time
	for _, height := range heights {
		if err := ec.storeData(height, data[height]); err != nil {
			return errors.Wrapf(err, "failed to store result of height %d", height)
		}
		ec.nextHeight = height + ec.interval
		latestBlockTime = data[height].mintTime
	}
	zap.L().Info("synced to", zap.Time("block time", latestBlockTime))
	atomic.StoreInt64(&ec.lastUpdateTimestamp, latestBlockTime.Unix())

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

	//calculate the result from DB
	calculator, err := ec.calculator(height)
	if err != nil {
		return nil, err
	}
	regs, err := ec.registrations(height)
	if err != nil {
		return nil, err
	}
	if err := calculator.AddRegistrations(regs); err != nil {
		return nil, err
	}

	buckets, err := ec.buckets(height)
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
	calculator, err := ec.calculator(height)
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

func (ec *committee) loadNextHeight() (uint64, error) {
	var val int64
	row := ec.db.QueryRow("SELECT height FROM next_height WHERE key = ?", NextHeightKey)
	err := row.Scan(&val)
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func (ec *committee) mintTime(height uint64) (time.Time, error) {
	var val time.Time
	row := ec.db.QueryRow("SELECT time FROM mint_times WHERE height = ?", util.Uint64ToInt64(height))
	err := row.Scan(&val)
	if err != nil {
		return time.Time{}, err
	}
	return val, nil
}

func (ec *committee) storeRegistrationsAndBuckets(height uint64, regs []*types.Registration, buckets []*types.Bucket) error {
	tx, err := ec.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	regStmt, err := tx.Prepare("INSERT OR IGNORE INTO registrations (hash, name, address, operator_address, reward_address, self_staking_weight) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer regStmt.Close()
	for _, reg := range regs {
		h, err := reg.Hash()
		if err != nil {
			return err
		}
		if _, err := regStmt.Exec(h[:], reg.Name(), reg.Address(), reg.OperatorAddress(), reg.RewardAddress(), util.Uint64ToInt64(reg.SelfStakingWeight())); err != nil {
			return err
		}
	}
	bucketStmt, err := tx.Prepare("INSERT OR IGNORE INTO buckets (hash, start_time, duration, amount, decay, voter, candidate) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer bucketStmt.Close()
	for _, bucket := range buckets {
		h, err := bucket.Hash()
		if err != nil {
			return err
		}
		if _, err := bucketStmt.Exec(h[:], bucket.StartTime(), bucket.Duration().String(), bucket.Amount().Bytes(), bucket.Decay(), bucket.Voter(), bucket.Candidate()); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (ec *committee) storeData(height uint64, data *rawData) error {
	fmt.Println("store regs and buckets")
	if err := ec.storeRegistrationsAndBuckets(height, data.registrations, data.buckets); err != nil {
		return err
	}
	fmt.Println("store others")
	tx, err := ec.db.Begin()
	if err != nil {
		return err
	}
	irh, err := ec.heightWithIdenticalRegs(height - ec.interval)
	if err != nil {
		return err
	}
	regHashes := make([][]byte, 0, len(data.registrations))
	for _, reg := range data.registrations {
		h, err := reg.Hash()
		if err != nil {
			return err
		}
		regHashes = append(regHashes, h[:])
	}
	fmt.Println("handle registrations")
	if ec.hasIdenticalRegistrations(regHashes, irh) {
		if _, err := tx.Exec("INSERT OR IGNORE INTO identical_registrations (height, identical_to) VALUES (?, ?)", height, irh); err != nil {
			return err
		}
	} else {
		fmt.Println("insert into height to registrations")
		//result, err := tx.Exec("INSERT INTO height_to_registrations (height, rid) VALUES (SELECT ?, id FROM registrations WHERE hash IN (?))", height, pq.Array(regHashes))
		if _, err := tx.Exec("CREATE TABLE temp.regs (height INTEGER, hash BLOB PRIMARY KEY)"); err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO temp.regs (height, hash) VALUES (?, ?")
		if err != nil {
			return err
		}
		defer stmt.Close() 
		for key, value := range regHashes {
			if _, err := stmt.Exec(height, value[:]); err != nil {
				return err
			}
		}
		result, err := tx.Exec(`INSERT OR IGNORE INTO height_to_registrations (height, rid) VALUES (
			SELECT temp.regs.height, registrations.id FROM registrations INNER JOIN temp.regs WHERE registrations.hash = temp.regs.hash
		)`)
		if err != nil {
			return err
		}

		//result, err := tx.Exec("INSERT INTO height_to_registrations (height, rid) VALUES (SELECT ?, id FROM registrations WHERE hash IN (?)", height, regHashes)
		//if err != nil {
		//	return err
		//}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows != int64(len(regHashes)) {
			return errors.New("wrong number of registration records")
		}
	}

	ibh, err := ec.heightWithIdenticalBuckets(height - ec.interval)
	if err != nil {
		return err
	}
	fmt.Println("handle buckets")
	bucketHashes := make(map[hash.Hash256]int)
	for _, bucket := range data.buckets {
		h, err := bucket.Hash()
		if err != nil {
			return err
		}
		if times, ok := bucketHashes[h]; ok {
			bucketHashes[h] = times + 1
		} else {
			bucketHashes[h] = 1
		}
	}
	if !data.migration && data.noNewStakingEvent || data.migration && ec.hasIdenticalBuckets(bucketHashes, ibh) {
		if _, err := tx.Exec("INSERT OR IGNORE INTO identical_buckets (height, identical_to) VALUES (?, ?)", height, ibh); err != nil {
			return err
		}
	} else {
		fmt.Println("insert into height to buckets", bucketHashes)
		if _, err := tx.Exec("DROP TABLE IF EXISTS temp.buckets"); err != nil {
			return err
		}
		if _, err := tx.Exec("CREATE TABLE temp.buckets (height INTEGER PRIMARY KEY, hash BLOB, times INTEGER)"); err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO temp.buckets (height, hash, times) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		for key, value := range bucketHashes {
			if _, err := stmt.Exec(height, key, value); err != nil {
				return err
			}
		}
		result, err := tx.Exec(`INSERT OR IGNORE INTO height_to_buckets (height, bid, times) VALUES (
			SELECT temp.buckets.height, buckets.id, temp.buckets.times FROM buckets INNER JOIN temp.buckets WHERE buckets.hash = temp.buckets.hash
		)`)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows != int64(len(bucketHashes)) {
			return errors.New("wrong number of bucket records")
		}
		if _, err := tx.Exec("DROP TABLE temp.buckets"); err != nil {
			return err
		}
	}
	if _, err := tx.Exec("INSERT INTO mint_times (height, time) VALUES (?, ?)", height, data.mintTime); err != nil {
		return err
	}

	if _, err := tx.Exec("INSERT OR REPLACE INTO nextHeight (key, height) VALUES (?, ?)", NextHeightKey, util.Uint64ToInt64(height)); err != nil {
		return err
	}

	return tx.Commit()
}

func (ec *committee) bucketHashes(height uint64) (map[hash.Hash256]int, error) {
	var hashes map[hash.Hash256]int
	rows, err := ec.db.Query(`
        SELECT b.hash, hb.times as times
        FROM buckets as b INNER JOIN height_to_buckets as hb
        WHERE hb.height = ? AND b.id = hb.bid
    `, util.Uint64ToInt64(height))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var val []byte
		var time int
		if err := rows.Scan(&val, &time); err != nil {
			return nil, err
		}
		hashes[hash.BytesToHash256(val)] = time
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return hashes, nil
}

func (ec *committee) buckets(height uint64) ([]*types.Bucket, error) {
	var buckets []*types.Bucket
	var decay, times int64
	var startTime time.Time
	var rawDuration string
	var amount, voter, candidate []byte

	rows, err := ec.db.Query(`
		SELECT b.start_time, b.duration, b.amount, b.decay, b.voter, b.candidate, hb.times as times
		FROM buckets as b INNER JOIN height_to_buckets as hb
		WHERE hb.height = ? AND buckets.id = hb.bid
	`, util.Uint64ToInt64(height))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// repeated "times"
		if err := rows.Scan(&startTime, &rawDuration, &amount, &voter, &candidate, &times); err != nil {
			return nil, err
		}
		duration, err := time.ParseDuration(rawDuration)
		bucket, err := types.NewBucket(startTime, duration, big.NewInt(0).SetBytes(amount), voter, candidate, decay != 0)
		if err != nil {
			return nil, err
		}
		for i := int64(0); i < times; i++ {
			buckets = append(buckets, bucket)
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return nil, nil
}

func (ec *committee) registrationHashes(height uint64) ([]hash.Hash256, error) {
	var hashes []hash.Hash256
	rows, err := ec.db.Query(`
        SELECT r.hash
        FROM registrations as r INNER JOIN height_to_registrations as hr
        WHERE hr.height = ? AND r.id = hr.rid
    `, util.Uint64ToInt64(height))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var val []byte
		if err := rows.Scan(&val); err != nil {
			return nil, err
		}
		hashes = append(hashes, hash.BytesToHash256(val))
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return hashes, nil
}

func (ec *committee) registrations(height uint64) ([]*types.Registration, error) {
	var registrations []*types.Registration
	var name, address, operatorAddress, rewardAddress []byte
	var selfStakingWeight int64
	rows, err := ec.db.Query(`
        SELECT r.name, r.address, r.operator_address, r.reward_address, r.self_staking_weight
        FROM registrations as r INNER JOIN height_to_registrations as hr
        WHERE hr.height = ? AND r.id = hr.rid
    `, util.Uint64ToInt64(height))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&name, &address, &operatorAddress, &rewardAddress, &selfStakingWeight); err != nil {
			zap.L().Error("failed to scan registration data")
			return nil, err
		}
		reg := types.NewRegistration(name, address, operatorAddress, rewardAddress, uint64(selfStakingWeight))
		registrations = append(registrations, reg)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return registrations, nil
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
