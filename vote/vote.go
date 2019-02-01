// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package vote

import (
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/iotexproject/go-ethereum/accounts/abi/bind"
	"github.com/iotexproject/go-ethereum/common"
	"github.com/iotexproject/go-ethereum/ethclient"
	"github.com/iotexproject/iotex-election/contract"
)

// Vote defines the structure of a vote
type Vote struct {
	startTime time.Time
	duration  time.Duration
	amount    *big.Int
	voter     common.Address
	candidate string
	decay     bool
}

func (v *Vote) StartTime() time.Time {
	return v.startTime
}

func (v *Vote) Duration() time.Duration {
	return v.duration
}

func (v *Vote) Voter() common.Address {
	return v.voter
}

func (v *Vote) Amount() *big.Int {
	return v.amount
}

func (v *Vote) Candidate() string {
	return v.candidate
}

func (v *Vote) Decay() bool {
	return v.decay
}

func (v *Vote) RemainingTime(now time.Time) time.Duration {
	// TODO: validate that duration is a positive value
	if v.decay {
		return v.startTime.Add(v.duration).Sub(now)
	}
	return v.duration
}

// Carrier defines an interfact to fetch votes
type Carrier interface {
	Votes(uint64, *big.Int, uint8) (*big.Int, []*Vote, error)
}

type ethereumCarrier struct {
	url          string
	contractAddr common.Address
}

// NewEthereumVoteCarrier defines a carrier to fetch votes from ethereum contract
func NewEthereumVoteCarrier(url string, contractAddr common.Address) Carrier {
	return &ethereumCarrier{
		contractAddr: contractAddr,
		url:          url,
	}
}

func (evc *ethereumCarrier) Votes(
	height uint64,
	previousIndex *big.Int,
	count uint8,
) (*big.Int, []*Vote, error) {
	if previousIndex == nil || previousIndex.Cmp(big.NewInt(0)) < 0 {
		previousIndex = big.NewInt(0)
	}
	client, err := ethclient.Dial(evc.url)
	if err != nil {
		return nil, nil, err
	}
	caller, err := contract.NewStakingCaller(evc.contractAddr, client)
	if err != nil {
		return nil, nil, err
	}
	buckets, err := caller.GetActiveBuckets(
		&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(height)},
		previousIndex,
		big.NewInt(int64(count)),
	)
	if err != nil {
		return nil, nil, err
	}
	votes := []*Vote{}
	num := len(buckets.Indexes)
	if num == 0 {
		return previousIndex, votes, nil
	}
	candidates, err := decodeCandidates(buckets.Candidates, num)
	if err != nil {
		return nil, nil, errors.New("invalid candidates return value")
	}
	for i, index := range buckets.Indexes {
		if big.NewInt(0).Cmp(index) == 0 { // back to start
			break
		}
		votes = append(votes, &Vote{
			startTime: time.Unix(buckets.StakeStartTimes[i].Int64(), 0),
			duration:  time.Duration(buckets.StakeDurations[i].Uint64()*24) * time.Hour,
			amount:    buckets.StakedAmounts[i],
			voter:     buckets.Owners[i],
			candidate: candidates[i],
			decay:     buckets.Decays[i],
		})
		if index.Cmp(previousIndex) > 0 {
			previousIndex = index
		}
	}

	return previousIndex, votes, nil
}

func decodeCandidates(data []byte, num int) ([]string, error) {
	if len(data) != 73*num {
		return nil, errors.New("the length of candidates is not as expected")
	}
	candidates := []string{}
	for i := 0; i < num; i++ {
		candidates = append(
			candidates,
			hex.EncodeToString(data[i*73+1:i*73+1+int(data[i*73])]),
		)
	}

	return candidates, nil
}
