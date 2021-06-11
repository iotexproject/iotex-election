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
	"math/big"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/protobuf/ptypes"
	lru "github.com/hashicorp/golang-lru"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/types"
)

type (
	NativeCommittee struct {
		archive         *BucketArchive
		retryLimit      uint8
		interval        uint64
		fetchInParallel uint8

		cache *lru.Cache

		startHeight         uint64
		currentHeight       uint64
		lastUpdateTimestamp int64
		terminate           chan bool
		mutex               sync.RWMutex

		tipHeight              uint64
		tickerDuration         time.Duration
		client                 iotexapi.APIServiceClient
		stakingContractAddress address.Address
		boundContract          *bind.BoundContract
	}

	NativeCommitteeConfig struct {
		StakingContractAddress string `yaml:"stakingContractAddress"`
		IoTeXAPI               string `yaml:"iotexAPI"`
		StartHeight            uint64 `yaml:"startHeight"`
		Interval               uint64 `yaml:"interval"`
	}

	pyggBucket struct {
		Index     uint64
		CanName   [12]byte
		Amount    *big.Int
		Duration  time.Duration
		StartTime time.Time
		Decay     bool
		Owner     address.Address
	}

	byIndex struct {
		logs []*iotextypes.Log
	}
)

func (bi byIndex) Len() int {
	return len(bi.logs)
}

func (bi byIndex) Less(i, j int) bool {
	return bi.logs[i].Index < bi.logs[j].Index
}
func (bi byIndex) Swap(i, j int) {
	bi.logs[i], bi.logs[j] = bi.logs[j], bi.logs[i]
}

func (pb *pyggBucket) Hash() (hash.Hash256, error) {
	if pb.Amount == nil {
		return hash.ZeroHash256, nil
	}
	b, err := types.NewBucket(
		pb.StartTime,
		pb.Duration,
		pb.Amount,
		pb.Owner.Bytes(),
		pb.CanName[:],
		pb.Decay,
	)
	if err != nil {
		return hash.ZeroHash256, err
	}
	return b.Hash()
}

// TODO: construct from abi
var (
	topicCreated = "d7812fae7f8126d2df0f5449a2cc0744d2e9d3fc8c161de6193bc4df6c68d365"
	topicUpdated = "0b074423c8a0f26c131cd7c88b19ef6adf084b812c97bdd1fb9dcf339ee9a387"
	topicUnstake = "9954bdedc474e937b39bbb080fc136e2edf1cef61f0906d36203267f4930762e"
)

func NewNativeStakingCommittee(
	archive *BucketArchive,
	cfg NativeCommitteeConfig,
) (*NativeCommittee, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.PyggStakingABI))
	if err != nil {
		return nil, err
	}
	cache, err := lru.New(100)
	if err != nil {
		return nil, err
	}
	addr, err := address.FromString(cfg.StakingContractAddress)
	if err != nil {
		return nil, err
	}
	conn, err := iotex.NewDefaultGRPCConn(cfg.IoTeXAPI)
	if err != nil {
		return nil, err
	}
	return &NativeCommittee{
		boundContract:          bind.NewBoundContract(common.BytesToAddress(addr.Bytes()), parsed, nil, nil, nil),
		archive:                archive,
		cache:                  cache,
		client:                 iotexapi.NewAPIServiceClient(conn),
		currentHeight:          0,
		fetchInParallel:        uint8(10),
		interval:               cfg.Interval,
		tickerDuration:         30 * time.Minute,
		lastUpdateTimestamp:    0,
		retryLimit:             uint8(10),
		startHeight:            cfg.StartHeight,
		stakingContractAddress: addr,
		terminate:              make(chan bool),
	}, nil
}

func (nc *NativeCommittee) Start(ctx context.Context) error {
	nc.mutex.Lock()
	defer nc.mutex.Unlock()
	if err := nc.archive.Start(ctx); err != nil {
		return err
	}
	tip, err := nc.tip()
	if err != nil {
		return errors.Wrap(err, "faile to get tip height")
	}
	go func() {
		zap.L().Info("catching up via network", zap.Uint64("tip", tip))
		for h := nc.nextHeight(); h < tip; h += nc.interval * 10 {
			zap.L().Info("catching up to", zap.Uint64("height", h))
			if err := nc.sync(h); err != nil {
				zap.L().Error("failed to fetch data", zap.Error(err))
			}
			time.Sleep(10 * time.Second)
		}
		zap.L().Info("catching up to tip", zap.Uint64("height", tip))
		if err := nc.sync(tip); err != nil {
			zap.L().Error("failed to fetch data", zap.Error(err))
		}
		ticker := time.NewTicker(nc.tickerDuration)
		for {
			select {
			case <-nc.terminate:
				nc.terminate <- true
				return
			case <-ticker.C:
				tip, err := nc.tip()
				if err != nil {
					zap.L().Error("failed to get tip", zap.Error(err))
					break
				}
				if err := nc.sync(tip); err != nil {
					zap.L().Error("failed to sync", zap.Error(err))
				}
			}
		}
	}()
	return nil
}

func (nc *NativeCommittee) Stop(ctx context.Context) error {
	nc.mutex.Lock()
	defer nc.mutex.Unlock()
	nc.terminate <- true

	return nc.archive.Stop(ctx)
}

func (nc *NativeCommittee) mintTime(height uint64) (time.Time, error) {
	response, err := nc.client.GetBlockMetas(
		context.Background(),
		&iotexapi.GetBlockMetasRequest{
			Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
				ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
					Start: height,
					Count: 1,
				},
			},
		},
	)
	if err != nil {
		return time.Time{}, err
	}
	return ptypes.Timestamp(response.BlkMetas[0].Timestamp)
}

func (nc *NativeCommittee) fetchDataByHeight(height uint64) (*timeAndBuckets, error) {
	zap.L().Info("fetch data", zap.Uint64("from", height), zap.Uint64("to", height+nc.interval-1))
	mintTime, err := nc.mintTime(height + nc.interval - 1)
	if err != nil {
		return nil, err
	}
	buckets, err := nc.delta(height, nc.interval)
	if err != nil {
		return nil, err
	}

	return &timeAndBuckets{
		mintTime: mintTime,
		buckets:  buckets,
	}, nil
}

func (nc *NativeCommittee) retryFetchDataByHeight(height uint64) (data *timeAndBuckets, err error) {
	for i := uint8(0); i < nc.retryLimit; i++ {
		if data, err = nc.fetchDataByHeight(height); err == nil {
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

type timeAndBuckets struct {
	mintTime time.Time
	buckets  []*pyggBucket
}

func (nc *NativeCommittee) fetchInBatch(tipHeight uint64) (retval map[uint64]*timeAndBuckets, err error) {
	if nc.currentHeight < tipHeight {
		nc.currentHeight = tipHeight
	}
	retval = map[uint64]*timeAndBuckets{}
	var wg sync.WaitGroup
	var lock sync.RWMutex
	limiter := make(chan bool, nc.fetchInParallel)
	for nextHeight := nc.nextHeight(); nextHeight <= nc.currentHeight; nextHeight += nc.interval {
		wg.Add(1)
		go func(height uint64) {
			defer func() {
				<-limiter
				wg.Done()
			}()
			limiter <- true
			data, e := nc.retryFetchDataByHeight(height - nc.interval + 1)
			lock.Lock()
			defer lock.Unlock()
			retval[height] = data
			if e != nil {
				err = e
			}
		}(nextHeight)
	}
	wg.Wait()

	return retval, err
}

func (nc *NativeCommittee) storeInBatch(data map[uint64]*timeAndBuckets) error {
	heights := make([]uint64, 0, len(data))
	mintTimes := make([]time.Time, 0, len(data))
	arrOfBuckets := make([][]*pyggBucket, 0, len(data))
	for height := range data {
		heights = append(heights, height)
		mintTimes = append(mintTimes, data[height].mintTime)
		arrOfBuckets = append(arrOfBuckets, data[height].buckets)
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
		if err := nc.archive.PutDelta(height, mintTimes[index], arrOfBuckets[index]); err != nil {
			return err
		}
	}
	atomic.StoreInt64(&nc.lastUpdateTimestamp, time.Now().Unix())
	return nil
}

func (nc *NativeCommittee) sync(tipHeight uint64) error {
	data, err := nc.fetchInBatch(tipHeight)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	nc.mutex.Lock()
	defer nc.mutex.Unlock()

	return nc.storeInBatch(data)
}

func (nc *NativeCommittee) DataByHeight(height uint64) (time.Time, []*types.Bucket, error) {
	nc.mutex.RLock()
	defer nc.mutex.RUnlock()
	if height < nc.startHeight {
		return time.Time{}, nil, nil
	}
	mintTime, err := nc.archive.MintTime(height)
	if err != nil {
		return time.Time{}, nil, err
	}
	buckets, err := nc.archive.Buckets(height)

	return mintTime, buckets, err
}

func (nc *NativeCommittee) nextHeight() uint64 {
	height := nc.latestHeightInArchive()
	if height == 0 {
		return nc.startHeight
	}

	return height + nc.interval
}

func (nc *NativeCommittee) latestHeightInArchive() uint64 {
	tipHeight, err := nc.archive.TipHeight()
	if err != nil {
		return 0
	}
	return tipHeight
}

func (nc *NativeCommittee) TipHeight() uint64 {
	nc.mutex.RLock()
	defer nc.mutex.RUnlock()

	return nc.latestHeightInArchive()
}

func (nc *NativeCommittee) Status() STATUS {
	lastUpdateTime := atomic.LoadInt64(&nc.lastUpdateTimestamp)
	switch {
	case lastUpdateTime == 0:
		return STARTING
	case lastUpdateTime > time.Now().Add(-61*time.Minute).Unix():
		return ACTIVE
	default:
		return INACTIVE
	}
}

func (nc *NativeCommittee) tip() (tip uint64, err error) {
	response, err := nc.client.GetChainMeta(
		context.Background(),
		&iotexapi.GetChainMetaRequest{},
	)
	if err != nil {
		return 0, err
	}

	return response.ChainMeta.Height, nil
}

func (nc *NativeCommittee) delta(
	from uint64,
	count uint64,
) ([]*pyggBucket, error) {
	topics := make([][]byte, 0)
	if t, err := hex.DecodeString(topicCreated); err == nil {
		topics = append(topics, t)
	} else {
		return nil, err
	}
	if t, err := hex.DecodeString(topicUpdated); err == nil {
		topics = append(topics, t)
	} else {
		return nil, err
	}
	if t, err := hex.DecodeString(topicUnstake); err == nil {
		topics = append(topics, t)
	} else {
		return nil, err
	}
	response, err := nc.client.GetLogs(context.Background(), &iotexapi.GetLogsRequest{
		Filter: &iotexapi.LogsFilter{
			Address: []string{nc.stakingContractAddress.String()},
			Topics:  []*iotexapi.Topics{&iotexapi.Topics{Topic: topics}},
		},
		Lookup: &iotexapi.GetLogsRequest_ByRange{
			ByRange: &iotexapi.GetLogsByRange{
				FromBlock: from,
				ToBlock:   from + count - 1,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	buckets := make([]*pyggBucket, 0)
	sort.Sort(byIndex{response.Logs})
	// update buckets correspondingly
	for _, l := range response.Logs {
		bucket := &pyggBucket{}
		etherLog := toEtherLog(l)
		switch hex.EncodeToString(l.Topics[0]) {
		case topicCreated:
			event := new(contract.PyggStakingPyggCreated)
			if err := nc.boundContract.UnpackLog(event, "PyggCreated", etherLog); err != nil {
				return nil, err
			}
			bucket.Index = event.PyggIndex.Uint64()
			bucket.CanName = event.CanName
			bucket.StartTime = time.Unix(event.StakeStartTime.Int64(), 0)
			bucket.Decay = !event.NonDecay
			bucket.Duration = time.Duration(event.StakeDuration.Uint64()*24) * time.Hour
			bucket.Amount = event.Amount
			bucket.Owner = event.PyggOwner
		case topicUpdated:
			event := new(contract.PyggStakingPyggUpdated)
			if err := nc.boundContract.UnpackLog(event, "PyggUpdated", etherLog); err != nil {
				return nil, err
			}
			bucket.Index = event.PyggIndex.Uint64()
			bucket.CanName = event.CanName
			bucket.StartTime = time.Unix(event.StakeStartTime.Int64(), 0)
			bucket.Decay = !event.NonDecay
			bucket.Duration = time.Duration(event.StakeDuration.Uint64()*24) * time.Hour
			bucket.Amount = event.Amount
			bucket.Owner = event.PyggOwner
		case topicUnstake:
			event := new(contract.PyggStakingPyggUnstake)
			if err := nc.boundContract.UnpackLog(event, "PyggUnstake", etherLog); err != nil {
				return nil, err
			}
			bucket.Index = event.PyggIndex.Uint64()
			bucket.Amount = nil
		default:
			return nil, errors.Errorf("Invalid topic %x", l.Topics[0])
		}
		buckets = append(buckets, bucket)
	}

	return buckets, nil
}

func toEtherLog(log *iotextypes.Log) ethtypes.Log {
	etherLog := ethtypes.Log{
		Data:        log.Data,
		BlockNumber: log.BlkHeight,
		TxHash:      common.BytesToHash(log.BlkHash),
		BlockHash:   common.BytesToHash(log.BlkHash),
		Index:       uint(log.Index),
	}

	for i := range log.Topics {
		etherLog.Topics = append(etherLog.Topics, common.BytesToHash(log.Topics[i]))
	}
	return etherLog
}
