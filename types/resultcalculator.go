// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-election/util"
)

const candidateZero = "000000000000000000000000"

// BucketFilterFunc defines the function to filter vote
type BucketFilterFunc func(*Bucket) bool

// CandidateFilterFunc defines the function to filter candidate
type CandidateFilterFunc func(*Candidate) bool

// ResultCalculator defines a calculator for a set of votes
type ResultCalculator struct {
	calcScore        func(*Bucket, time.Time) *big.Int
	candidateFilter  func(*Candidate) bool
	bucketFilter     func(*Bucket) bool
	mintTime         time.Time
	candidates       map[string]*Candidate
	candidateVotes   map[string][]*Vote
	totalVotes       *big.Int
	totalVotedStakes *big.Int
	calculated       bool
	mutex            sync.RWMutex
	skipManified     bool
}

// NewResultCalculator creates a result calculator
func NewResultCalculator(
	mintTime time.Time,
	skipManified bool,
	bucketFilter BucketFilterFunc, // filter buckets before calculating
	calcScore func(*Bucket, time.Time) *big.Int,
	candidateFilter CandidateFilterFunc, // filter candidates during calculating
) *ResultCalculator {
	return &ResultCalculator{
		calcScore:        calcScore,
		candidateFilter:  candidateFilter,
		bucketFilter:     bucketFilter,
		mintTime:         mintTime.UTC(),
		candidates:       map[string]*Candidate{},
		candidateVotes:   map[string][]*Vote{},
		totalVotedStakes: big.NewInt(0),
		totalVotes:       big.NewInt(0),
		calculated:       false,
		skipManified:     skipManified,
	}
}

// AddRegistrations adds candidates to result
func (calculator *ResultCalculator) AddRegistrations(candidates []*Registration) error {
	calculator.mutex.Lock()
	defer calculator.mutex.Unlock()
	if calculator.calculated {
		return errors.New("Cannot modify a calculated result")
	}
	if calculator.totalVotes.Cmp(big.NewInt(0)) > 0 {
		return errors.New("Candidates should be added before any votes")
	}
	for _, c := range candidates {
		name := calculator.hex(c.Name())
		if _, exists := calculator.candidates[name]; exists {
			return errors.Errorf("Duplicate candidate %s", name)
		}
		if c.SelfStakingWeight() > uint64(1) && calculator.skipManified {
			continue
		}
		calculator.candidates[name] = NewCandidate(c, big.NewInt(0), big.NewInt(0))
		calculator.candidateVotes[name] = []*Vote{}
	}
	return nil
}

// AddBuckets adds bucket to result
func (calculator *ResultCalculator) AddBuckets(buckets []*Bucket) error {
	calculator.mutex.Lock()
	defer calculator.mutex.Unlock()
	if calculator.calculated {
		return errors.New("Cannot modify a calculated result")
	}
	for _, bucket := range buckets {
		if calculator.bucketFilter(bucket) {
			continue
		}
		name := bucket.Candidate()
		if name == nil {
			continue
		}
		nameHex := calculator.hex(name)
		if strings.Compare(nameHex, candidateZero) == 0 {
			continue
		}
		amount := bucket.Amount()
		score := calculator.calcScore(bucket, calculator.mintTime)
		if candidate, exists := calculator.candidates[nameHex]; exists {
			if bytes.Equal(bucket.Voter(), candidate.address) {
				selfStakingWeight := new(big.Int).SetUint64(candidate.selfStakingWeight)
				amount.Mul(amount, selfStakingWeight)
				if err := candidate.addSelfStakingTokens(amount); err != nil {
					return err
				}
				score.Mul(score, selfStakingWeight)
			}
			cVote, err := NewVote(bucket, score)
			if err != nil {
				return err
			}
			if err := candidate.addScore(score); err != nil {
				return err
			}
			calculator.candidateVotes[nameHex] = append(calculator.candidateVotes[nameHex], cVote)
		}
		calculator.totalVotedStakes.Add(calculator.totalVotedStakes, amount)
		calculator.totalVotes.Add(calculator.totalVotes, score)
	}
	return nil
}

// Calculate summaries the result with candidates and votes added
func (calculator *ResultCalculator) Calculate() (*ElectionResult, error) {
	calculator.mutex.Lock()
	defer calculator.mutex.Unlock()
	if calculator.calculated {
		return nil, errors.New("Cannot modify a calculated result")
	}
	qualifiers := calculator.filterAndSortCandidates()
	candidates := make([]*Candidate, len(qualifiers))
	votes := map[string][]*Vote{}
	for i, name := range qualifiers {
		candidates[i] = calculator.candidates[name]
		votes[name] = calculator.candidateVotes[name]
	}
	calculator.calculated = true

	return &ElectionResult{
		mintTime:         calculator.mintTime,
		delegates:        candidates,
		votes:            votes,
		totalVotedStakes: calculator.totalVotedStakes,
		totalVotes:       calculator.totalVotes,
	}, nil
}

func (calculator *ResultCalculator) filterAndSortCandidates() []string {
	candidates := make(map[string]*big.Int, len(calculator.candidates))
	for name, candidate := range calculator.candidates {
		candidates[name] = candidate.Score()
	}
	sorted := util.Sort(candidates, uint64(calculator.mintTime.Unix()))
	var qualifiers []string
	for i, name := range sorted {
		if !calculator.candidateFilter(calculator.candidates[name]) {
			qualifiers = append(qualifiers, sorted[i])
		}
	}

	return qualifiers
}

func (calculator *ResultCalculator) hex(name []byte) string {
	return hex.EncodeToString(name)
}
