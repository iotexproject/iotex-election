// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package explorer

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/explorer/idl/explorer"
)

// Impl implements Explorer interface
type Impl struct {
	c committee.Committee
}

// NewExplorer creates an explorer instance
func NewExplorer(c committee.Committee) explorer.Explorer {
	return &Impl{c: c}
}

// GetMeta returns the meta of the chain
func (e *Impl) GetMeta() (explorer.ChainMeta, error) {
	height := e.c.LatestHeight()
	result, err := e.c.ResultByHeight(height)
	if err != nil {
		return explorer.ChainMeta{}, err
	}
	return explorer.ChainMeta{
		Height:          strconv.FormatUint(height, 10),
		TotalCandidates: int64(len(result.Delegates())),
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (e *Impl) GetCandidates(
	request explorer.GetCandidatesRequest,
) ([]explorer.Candidate, error) {
	offset := request.Offset
	limit := request.Limit
	if offset < 0 || limit < 0 {
		return nil, errors.New("offset and limit should be positive number")
	}
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := e.c.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	candidates := result.Delegates()
	if int64(len(candidates)) <= offset {
		return nil, errors.New("offset is larger than candidate length")
	}
	if int64(len(candidates)) < offset+limit {
		limit = int64(len(candidates)) - offset
	}
	retval := make([]explorer.Candidate, limit)
	for i := offset; i < offset+limit; i++ {
		candidate := candidates[i]
		retval[i].Name = hex.EncodeToString(candidate.Name())
		retval[i].PubKey = hex.EncodeToString(candidate.BeaconPubKey())
		retval[i].TotalWeightedVotes = candidate.Score().Text(10)
	}

	return retval, nil
}

// GetBucketsByCandidate returns the buckets
func (e *Impl) GetBucketsByCandidate(
	request explorer.GetBucketsByCandidateRequest,
) ([]explorer.Bucket, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := e.c.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	votes := result.VotesByDelegate([]byte(request.Name))
	if votes == nil {
		return nil, errors.New("No buckets for the candidate")
	}
	retval := make([]explorer.Bucket, len(votes))
	for i, vote := range votes {
		retval[i].Voter = hex.EncodeToString(vote.Voter())
		retval[i].Votes = vote.Amount().Text(10)
		retval[i].WeightedVotes = vote.WeightedAmount().Text(10)
		retval[i].RemainingDuration = int64(vote.RemainingTime(result.MintTime()) / time.Second)
	}

	return retval, nil
}
