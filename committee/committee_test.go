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
	now := time.Now()
	mintTime := now.Add(-3 * time.Hour)
	candidates := genTestCandidates()
	votes := genTestVotes(mintTime, require)

	c := &committee{selfStakingThreshold: big.NewInt(10), scoreThreshold: big.NewInt(10)}
	rc := types.NewResultCalculator(time.Now(), c.voteFilter, c.calcWeightedVotes, c.candidateFilter)

	require.NotNil(rc)
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	_, err := rc.Calculate()
	require.NoError(err)

	require.False(c.candidateFilter(candidates[0]))
	require.True(c.candidateFilter(candidates[1]))
}
func genTestVotes(mintTime time.Time, require *require.Assertions) []*types.Vote {
	votes := []*types.Vote{}
	// score 100
	vote, err := types.NewVote(
		mintTime.Add(-3*time.Hour),
		10*time.Hour,
		big.NewInt(100),
		big.NewInt(110),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// score 10
	vote, err = types.NewVote(
		mintTime.Add(-2*time.Hour),
		3*time.Hour,
		big.NewInt(10),  //amount
		big.NewInt(100), //weight
		[]byte("voter2"),
		[]byte("candidate2"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)

	return append(votes, vote)
}
func genTestCandidates() []*types.Candidate {
	return []*types.Candidate{
		types.NewCandidate(
			[]byte("candidate1"),
			[]byte("candidate1addr"),
			[]byte("operatorPubKey1"),
			[]byte("rewardPubKey1"),
			1,
		),
		types.NewCandidate(
			[]byte("candidate2"),
			[]byte("candidate2addr"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			20,
		),
	}
}
