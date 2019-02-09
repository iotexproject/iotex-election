// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"

	"github.com/iotexproject/iotex-election/util"
)

type item struct {
	Key      string
	Value    *big.Int
	Priority uint64
}

type itemList []item

func (p itemList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p itemList) Len() int      { return len(p) }
func (p itemList) Less(i, j int) bool {
	switch p[i].Value.Cmp(p[j].Value) {
	case -1:
		return false
	case 1:
		return true
	}
	switch {
	case p[i].Priority > p[j].Priority:
		return true
	case p[i].Priority < p[j].Priority:
		return false
	}
	// This is a corner case, which rarely happens.
	return strings.Compare(p[i].Key, p[j].Key) > 0
}

// ResultCalculator defines a calculator for a set of votes
type ResultCalculator struct {
	calcScore      func(*Vote, time.Time) *big.Int
	filter         func(*Candidate) bool
	mintTime       time.Time
	candidates     map[string]*Candidate
	candidateVotes map[string][]*Vote
	totalVotes     int32
	calculated     bool
	mutex          sync.RWMutex
}

// NewResultCalculator creates a result builder
func NewResultCalculator(
	mintTime time.Time,
	calcScore func(*Vote, time.Time) *big.Int,
	filter func(*Candidate) bool,
) *ResultCalculator {
	return &ResultCalculator{
		calcScore:      calcScore,
		filter:         filter,
		mintTime:       mintTime.UTC(),
		candidates:     map[string]*Candidate{},
		candidateVotes: map[string][]*Vote{},
		totalVotes:     0,
		calculated:     false,
	}
}

// AddCandidates adds candidates to result
func (builder *ResultCalculator) AddCandidates(candidates []*Candidate) error {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	if builder.calculated {
		return errors.New("Cannot modify a calculated result")
	}
	if builder.totalVotes > 0 {
		return errors.New("Candidates should be added before any votes")
	}
	for _, c := range candidates {
		name := builder.hex(c.Name())
		if _, exists := builder.candidates[name]; exists {
			return errors.Errorf("Duplicate candidate %s", name)
		}
		builder.candidates[name] = c.Clone().reset()
		builder.candidateVotes[name] = []*Vote{}
	}
	return nil
}

// AddVotes adds votes to result
func (builder *ResultCalculator) AddVotes(votes []*Vote) error {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	if builder.calculated {
		return errors.New("Cannot modify a calculated result")
	}
	for _, v := range votes {
		name := v.Candidate()
		if name == nil {
			continue
		}
		nameHex := builder.hex(name)
		candidate, exists := builder.candidates[nameHex]
		if !exists {
			continue
		}
		score := builder.calcScore(v, builder.mintTime)
		if bytes.Equal(v.Voter(), candidate.beaconPubKey) {
			score.Mul(score, big.NewInt(int64(candidate.selfStakingWeight)))
			candidate.addSelfStakingScore(score)
		}
		cVote := v.Clone()
		if err := cVote.SetWeightedAmount(score); err != nil {
			return err
		}
		candidate.addScore(score)
		builder.candidateVotes[nameHex] = append(builder.candidateVotes[nameHex], cVote)
		builder.totalVotes++
	}
	return nil
}

// Calculate summaries the result with candidates and votes added
func (builder *ResultCalculator) Calculate() (*Result, error) {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	if builder.calculated {
		return nil, errors.New("Cannot modify a calculated result")
	}
	qualifiers := builder.filterAndSortCandidates()
	candidates := make([]*Candidate, len(qualifiers))
	votes := map[string][]*Vote{}
	for i, name := range qualifiers {
		candidates[i] = builder.candidates[name]
		votes[name] = builder.candidateVotes[name]
	}
	builder.calculated = true

	return &Result{
		mintTime:   builder.mintTime,
		candidates: candidates,
		votes:      votes,
	}, nil
}

func (builder *ResultCalculator) filterAndSortCandidates() []string {
	p := make(itemList, len(builder.candidates))
	num := 0
	tsBytes := util.Uint64ToBytes(uint64(builder.mintTime.Unix()))
	for name, candidate := range builder.candidates {
		if !builder.filter(candidate) {
			priority := blake2b.Sum256(append([]byte(name), tsBytes...))
			p[num] = item{
				Key:      name,
				Value:    candidate.selfStakingScore,
				Priority: util.BytesToUint64(priority[:8]),
			}
			num++
		}
	}
	sort.Stable(p[:num])
	qualifiers := make([]string, num)
	for i := 0; i < num; i++ {
		qualifiers[i] = p[i].Key
	}
	return qualifiers
}

func (builder *ResultCalculator) hex(name []byte) string {
	return hex.EncodeToString(name)
}
