// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/iotex-election/types"

	"github.com/stretchr/testify/require"
)

func TestVoteFilter(t *testing.T) {
	require := require.New(t)
	c := &committee{voteThreshold: big.NewInt(10)}
	vote1, err := types.NewVote(
		time.Now(),
		time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.True(c.voteFilter(vote1))

	vote2, err := types.NewVote(
		time.Now(),
		time.Hour,
		big.NewInt(30),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.False(c.voteFilter(vote2))
}
func TestCandidateFilter(t *testing.T) {
	require := require.New(t)
	c := &committee{selfStakingThreshold: big.NewInt(10), scoreThreshold: big.NewInt(10)}
	candidate1 := &Candidate{score: big.NewInt(5), selfStakingTokens: big.NewInt(5)}
	candidate2 := &Candidate{score: big.NewInt(5), selfStakingTokens: big.NewInt(20)}
	candidate3 := &Candidate{score: big.NewInt(20), selfStakingTokens: big.NewInt(5)}
	candidate4 := &Candidate{score: big.NewInt(20), selfStakingTokens: big.NewInt(20)}
	intface := interface{}(candidate1)
	require.True(c.candidateFilter(intface.(*types.Candidate)))
	intface = interface{}(candidate2)
	require.True(c.candidateFilter(intface.(*types.Candidate)))
	intface = interface{}(candidate3)
	require.True(c.candidateFilter(intface.(*types.Candidate)))
	intface = interface{}(candidate4)
	require.False(c.candidateFilter(intface.(*types.Candidate)))
}

type Candidate struct {
	name              []byte
	address           []byte
	operatorAddress   []byte
	rewardAddress     []byte
	score             *big.Int
	selfStakingTokens *big.Int
	selfStakingWeight uint64
}

func (c *Candidate) SelfStakingTokens() *big.Int {
	return new(big.Int).Set(c.selfStakingTokens)
}
func (c *Candidate) Score() *big.Int {
	return new(big.Int).Set(c.score)
}
