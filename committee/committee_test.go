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
	testCandidateFilter("0", require)
	//testCandidateFilter("1", require)
}
func testCandidateFilter(SelfStakingThreshold string, require *require.Assertions) {
	now := time.Now()
	mintTime := now.Add(-3 * time.Hour)
	candidates := genTestCandidates()
	votes := genTestVotes(mintTime, require)
	cfg := getCfg(SelfStakingThreshold)
	commp, err := NewCommittee(nil, cfg)
	require.NoError(err)
	rc, err := commp.(*committee).calculator(10662182) //should be new from kovan
	require.NoError(err)
	require.NotNil(rc)
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	_, err = rc.Calculate()
	require.NoError(err)
	if SelfStakingThreshold == "0" {
		require.True(commp.(*committee).candidateFilter(candidates[0]))
		require.True(commp.(*committee).candidateFilter(candidates[1]))
	} else {
		require.True(commp.(*committee).candidateFilter(candidates[0]))
		require.True(commp.(*committee).candidateFilter(candidates[1]))
	}

}
func genTestVotes(mintTime time.Time, require *require.Assertions) []*types.Vote {
	votes := []*types.Vote{}
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
			1,
		),
	}
}
func getCfg(SelfStakingThreshold string) (cfg Config) {
	cfg.NumOfRetries = 8
	cfg.BeaconChainAPIs = []string{"https://kovan.infura.io"}
	cfg.BeaconChainHeightInterval = 100
	cfg.BeaconChainStartHeight = 7368630
	cfg.RegisterContractAddress = "0xb4ca6cf2fe760517a3f92120acbe577311252663"
	cfg.StakingContractAddress = "0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"
	cfg.PaginationSize = 100
	cfg.VoteThreshold = "10"
	cfg.ScoreThreshold = "10"
	cfg.SelfStakingThreshold = SelfStakingThreshold // must be 0,because cannot set candidate's StakingTokens
	cfg.CacheSize = 100
	return
}
