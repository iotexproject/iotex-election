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

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-election/types"
)

func TestCalcWeightedVotes(t *testing.T) {
	require := require.New(t)
	committee := &committee{}
	startTime := time.Now()
	duration := time.Hour * 24 * 14
	vote1, err := types.NewVote(
		startTime,
		duration,
		big.NewInt(3000000),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	// now.Before(v.StartTime()),return 0
	require.Equal(0, committee.calcWeightedVotes(vote1, startTime.Add(-1*time.Hour)).Cmp(big.NewInt(0)))

	// decay is true,startTime+duration is after now,remainingTime is 24*14-24=13*24 hours,weight is ~1.140,ret is 3422048
	require.Equal(0, committee.calcWeightedVotes(vote1, startTime.Add(time.Hour*24)).Cmp(big.NewInt(3422048)))

	// decay is true,startTime+duration is before now,remainingTime is 0 hours,weight is 1,ret is 3000000
	require.Equal(0, committee.calcWeightedVotes(vote1, time.Now().Add(24*15*time.Hour)).Cmp(big.NewInt(3000000)))

	vote2, err := types.NewVote(
		startTime,
		duration,
		big.NewInt(3000000),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		false,
	)

	// decay is false,remainingTime is duration,weight ~1.144，ret is 3434242,whatever now is
	require.Equal(0, committee.calcWeightedVotes(vote2, startTime.Add(time.Hour*24)).Cmp(big.NewInt(3434242)))
	require.Equal(0, committee.calcWeightedVotes(vote2, startTime.Add(24*15*time.Hour)).Cmp(big.NewInt(3434242)))
}

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
	committee := &committee{
		scoreThreshold:       big.NewInt(10),
		selfStakingThreshold: big.NewInt(10),
	}
	// candidate1 selfStaking and score both smaller than committee's threshold
	candidate1 := types.NewCandidate(
		[]byte("candidate1"),
		[]byte("candidate1addr"),
		[]byte("operatorPubKey1"),
		[]byte("rewardPubKey1"),
		1,
	)
	candidate1.SetScore(big.NewInt(9))
	candidate1.SetSelfStakingTokens(big.NewInt(9))
	require.True(committee.candidateFilter(candidate1))
	// candidate2 selfStaking is below committee's threshold,score is bigger than committee's threshold
	candidate2 := types.NewCandidate(
		[]byte("candidate2"),
		[]byte("candidate2addr"),
		[]byte("operatorPubKey2"),
		[]byte("rewardPubKey2"),
		1,
	)
	candidate2.SetScore(big.NewInt(11))
	candidate2.SetSelfStakingTokens(big.NewInt(9))
	require.True(committee.candidateFilter(candidate2))
	// candidate3 selfStaking is bigger than committee's threshold,score is smaller than committee's threshold
	candidate3 := types.NewCandidate(
		[]byte("candidate3"),
		[]byte("candidate3addr"),
		[]byte("operatorPubKey3"),
		[]byte("rewardPubKey3"),
		1,
	)
	candidate3.SetScore(big.NewInt(9))
	candidate3.SetSelfStakingTokens(big.NewInt(11))
	require.True(committee.candidateFilter(candidate3))
	// candidate3 selfStaking and score both bigger than committee's threshold
	candidate4 := types.NewCandidate(
		[]byte("candidate4"),
		[]byte("candidate4addr"),
		[]byte("operatorPubKey4"),
		[]byte("rewardPubKey4"),
		1,
	)
	candidate4.SetScore(big.NewInt(11))
	candidate4.SetSelfStakingTokens(big.NewInt(11))
	require.False(committee.candidateFilter(candidate4))
}
