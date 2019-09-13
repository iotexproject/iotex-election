// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package carrier

import (
	"context"
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/types"
)

// TipInfo is the info of a tip block
type TipInfo struct {
	Height    uint64
	BlockTime time.Time
}

// Carrier defines an interfact to fetch votes
type Carrier interface {
	// BlockTimestamp returns the timestamp of a block
	BlockTimestamp(uint64) (time.Time, error)
	// SubscribeNewBlock callbacks on new block created
	SubscribeNewBlock(chan *TipInfo, chan error, chan bool)
	// HasStakingEvents returns true if there is any staking related events or error
	HasStakingEvents(*big.Int, *big.Int) bool
	// Tip returns the latest height and its timestamp
	Tip() (*TipInfo, error)
	// Registrations returns the candidate registrations on height
	Registrations(uint64, *big.Int, uint8) (*big.Int, []*types.Registration, error)
	// Buckets returns the buckets on height
	Buckets(uint64, *big.Int, uint8) (*big.Int, []*types.Bucket, error)
	// Close closes carrier
	Close()
}

// EthClientPool defines a set of ethereum clients with execute interface
type EthClientPool struct {
	clientURLs []string
	client     *ethclient.Client
	lock       sync.RWMutex
}

// NewEthClientPool creates a new pool
func NewEthClientPool(urls []string) *EthClientPool {
	return &EthClientPool{
		clientURLs: urls,
		client:     nil,
	}
}

// Close closes the current client if available
func (pool *EthClientPool) Close() {
	pool.swapClient(nil)
}

func (pool *EthClientPool) swapClient(client *ethclient.Client) {
	pool.lock.Lock()
	defer pool.lock.Unlock()
	if pool.client != client {
		if pool.client != nil {
			pool.client.Close()
		}
		pool.client = client
	}
}

func (pool *EthClientPool) execute(callback func(c *ethclient.Client) error, client *ethclient.Client) error {
	pool.lock.RLock()
	defer pool.lock.RUnlock()
	if client != nil {
		return callback(client)
	}
	if pool.client != nil {
		return callback(pool.client)
	}
	return errors.New("no client available")
}

// Execute executes callback by rotating all client urls
func (pool *EthClientPool) Execute(callback func(c *ethclient.Client) error) (err error) {
	if err = pool.execute(callback, nil); err == nil {
		return
	}
	var client *ethclient.Client
	for i := 0; i < len(pool.clientURLs); i++ {
		if client, err = ethclient.Dial(pool.clientURLs[i]); err != nil {
			zap.L().Error(
				"client is not reachable",
				zap.String("url", pool.clientURLs[i]),
				zap.Error(err),
			)
			continue
		}
		if err = pool.execute(callback, client); err == nil {
			pool.swapClient(client)
			return
		}
	}
	return errors.Wrap(err, "failed to execute callback with any client")
}

type ethereumCarrier struct {
	ethClientPool           *EthClientPool
	stakingContractAddress  common.Address
	registerContractAddress common.Address
}

// NewEthereumVoteCarrier defines a carrier to fetch votes from ethereum contract
func NewEthereumVoteCarrier(
	clientURLs []string,
	registerContractAddress common.Address,
	stakingContractAddress common.Address,
) (Carrier, error) {
	if len(clientURLs) == 0 {
		return nil, errors.New("client URL list is empty")
	}
	return &ethereumCarrier{
		ethClientPool:           NewEthClientPool(clientURLs),
		stakingContractAddress:  stakingContractAddress,
		registerContractAddress: registerContractAddress,
	}, nil
}

func (evc *ethereumCarrier) Close() {
	evc.ethClientPool.Close()
}

func (evc *ethereumCarrier) BlockTimestamp(height uint64) (ts time.Time, err error) {
	err = evc.ethClientPool.Execute(func(client *ethclient.Client) error {
		header, err := client.HeaderByNumber(
			context.Background(),
			big.NewInt(0).SetUint64(height),
		)
		if err == nil {
			ts = time.Unix(int64(header.Time), 0)
		}
		return err
	})
	return
}

func (evc *ethereumCarrier) SubscribeNewBlock(
	tipChan chan *TipInfo,
	report chan error,
	unsubscribe chan bool,
) {
	ticker := time.NewTicker(60 * time.Second)
	lastHeight := uint64(0)
	go func() {
		for {
			select {
			case <-unsubscribe:
				unsubscribe <- true
				return
			case <-ticker.C:
				if tip, err := evc.tip(lastHeight); err != nil {
					report <- err
				} else {
					tipChan <- tip
				}
			}
		}
	}()
}

func (evc *ethereumCarrier) Tip() (*TipInfo, error) {
	return evc.tip(0)
}

func (evc *ethereumCarrier) tip(lastHeight uint64) (tip *TipInfo, err error) {
	if err = evc.ethClientPool.Execute(func(client *ethclient.Client) error {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err == nil {
			if header.Number.Uint64() > lastHeight {
				tip = &TipInfo{
					Height:    header.Number.Uint64(),
					BlockTime: time.Unix(int64(header.Time), 0),
				}
				return nil
			}
			err = errors.Errorf(
				"client is out of date, client height %d < last height %d",
				header.Number.Uint64(),
				lastHeight,
			)
		}
		return err
	}); err != nil {
		err = errors.Wrap(err, "failed to get tip height")
	}
	return
}

func (evc *ethereumCarrier) candidates(
	opts *bind.CallOpts,
	startIndex *big.Int,
	limit *big.Int,
) (result struct {
	Names          [][12]byte
	Addresses      []common.Address
	IoOperatorAddr [][32]byte
	IoRewardAddr   [][32]byte
	Weights        []*big.Int
}, err error) {
	if err = evc.ethClientPool.Execute(func(client *ethclient.Client) error {
		if caller, err := contract.NewRegisterCaller(evc.registerContractAddress, client); err == nil {
			var count *big.Int
			if count, err = caller.CandidateCount(opts); err != nil {
				return err
			}
			if startIndex.Cmp(count) >= 0 {
				return nil
			}
			result, err = caller.GetAllCandidates(opts, startIndex, limit)
		}
		return err
	}); err != nil {
		err = errors.Wrap(err, "failed to get candidates")
	}
	return
}

func (evc *ethereumCarrier) HasStakingEvents(from *big.Int, to *big.Int) bool {
	retval := true
	if err := evc.ethClientPool.Execute(func(client *ethclient.Client) error {
		logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: from,
			ToBlock:   to,
			Addresses: []common.Address{evc.stakingContractAddress},
			Topics: [][]common.Hash{
				[]common.Hash{
					common.HexToHash("0xbecddf0f61f76a4ac94a507fbc32c036d2fb7c4b466cad82dd9a4a2d76b263fe"), // created
					common.HexToHash("0x004bbbedd0138c223ffed73fdab05a22a5d22770de54bea694d06661d59d1600"), // updated
					common.HexToHash("0xaa192dc938c20fb63756fbd8f4d9f46092c3252f772b2c549c4688c118b6b475"), // unstaked
				},
			},
		})
		if err == nil {
			retval = len(logs) != 0
		}
		return err
	}); err != nil {
		return true
	}
	return retval
}

func (evc *ethereumCarrier) Registrations(
	height uint64,
	startIndex *big.Int,
	count uint8,
) (*big.Int, []*types.Registration, error) {
	if startIndex == nil || startIndex.Cmp(big.NewInt(1)) < 0 {
		startIndex = big.NewInt(1)
	}
	retval, err := evc.candidates(
		&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(height)},
		startIndex,
		big.NewInt(int64(count)),
	)
	if err != nil {
		return nil, nil, err
	}
	num := len(retval.Names)
	if len(retval.Addresses) != num {
		return nil, nil, errors.New("invalid addresses from GetAllCandidates")
	}
	operatorPubKeys, err := decodeAddress(retval.IoOperatorAddr, num)
	if err != nil {
		return nil, nil, err
	}
	rewardPubKeys, err := decodeAddress(retval.IoRewardAddr, num)
	if err != nil {
		return nil, nil, err
	}
	registrations := make([]*types.Registration, num)
	for i := 0; i < num; i++ {
		registrations[i] = types.NewRegistration(
			retval.Names[i][:],
			retval.Addresses[i][:],
			operatorPubKeys[i],
			rewardPubKeys[i],
			retval.Weights[i].Uint64(),
		)
	}
	return new(big.Int).Add(startIndex, big.NewInt(int64(num))), registrations, nil
}

// EthereumBucketsResult defines the data structure the buckets api returns
type EthereumBucketsResult struct {
	Count           *big.Int
	Indexes         []*big.Int
	StakeStartTimes []*big.Int
	StakeDurations  []*big.Int
	Decays          []bool
	StakedAmounts   []*big.Int
	CanNames        [][12]byte
	Owners          []common.Address
}

func (evc *ethereumCarrier) buckets(
	opts *bind.CallOpts,
	previousIndex *big.Int,
	limit *big.Int,
) (result EthereumBucketsResult, err error) {
	if err = evc.ethClientPool.Execute(func(client *ethclient.Client) error {
		caller, err := contract.NewStakingCaller(evc.stakingContractAddress, client)
		if err != nil {
			return err
		}
		var bucket struct {
			CanName          [12]byte
			StakedAmount     *big.Int
			StakeDuration    *big.Int
			StakeStartTime   *big.Int
			NonDecay         bool
			UnstakeStartTime *big.Int
			BucketOwner      common.Address
			CreateTime       *big.Int
			Prev             *big.Int
			Next             *big.Int
		}
		if bucket, err = caller.Buckets(opts, previousIndex); err == nil {
			if bucket.Next.Cmp(big.NewInt(0)) <= 0 {
				return nil
			}
			result, err = caller.GetActiveBuckets(opts, previousIndex, limit)
		}
		return err
	}); err != nil {
		err = errors.Wrap(err, "failed to get votes")
	}
	return
}

func (evc *ethereumCarrier) Buckets(
	height uint64,
	previousIndex *big.Int,
	count uint8,
) (*big.Int, []*types.Bucket, error) {
	if previousIndex == nil || previousIndex.Cmp(big.NewInt(0)) < 0 {
		previousIndex = big.NewInt(0)
	}
	buckets, err := evc.buckets(
		&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(height)},
		previousIndex,
		big.NewInt(int64(count)),
	)
	if err != nil {
		return nil, nil, err
	}
	bs := []*types.Bucket{}
	if buckets.Count == nil || buckets.Count.Cmp(big.NewInt(0)) == 0 || len(buckets.Indexes) == 0 {
		return previousIndex, bs, nil
	}
	for i, index := range buckets.Indexes {
		if big.NewInt(0).Cmp(index) == 0 { // back to start, this is a redundant condition
			break
		}
		v, err := types.NewBucket(
			time.Unix(buckets.StakeStartTimes[i].Int64(), 0),
			time.Duration(buckets.StakeDurations[i].Uint64()*24)*time.Hour,
			buckets.StakedAmounts[i],
			buckets.Owners[i].Bytes(),
			buckets.CanNames[i][:],
			buckets.Decays[i],
		)
		if err != nil {
			return nil, nil, err
		}
		bs = append(bs, v)
		if index.Cmp(previousIndex) > 0 {
			previousIndex = index
		}
	}

	return previousIndex, bs, nil
}

func decodeAddress(data [][32]byte, num int) ([][]byte, error) {
	if len(data) != 2*num {
		return nil, errors.New("the length of address array is not as expected")
	}
	keys := [][]byte{}
	for i := 0; i < num; i++ {
		key := append(data[2*i][:], data[2*i+1][:9]...)
		keys = append(keys, key)
	}

	return keys, nil
}
