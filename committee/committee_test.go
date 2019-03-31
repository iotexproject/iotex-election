// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/iotex-election/types"

	"github.com/stretchr/testify/require"
)

func TestResultCalculator(t *testing.T) {
	require := require.New(t)
	now := time.Now()
	mintTime := now.Add(-3 * time.Hour)
	candidates := genTestCandidates()
	votes := genTestVotes(mintTime, require)

	var cfg Config
	cfg.NumOfRetries = 8
	cfg.BeaconChainAPIs = []string{"https://kovan.infura.io"}
	cfg.BeaconChainHeightInterval = 100
	cfg.BeaconChainStartHeight = 7368630
	cfg.RegisterContractAddress = "0xb4ca6cf2fe760517a3f92120acbe577311252663"
	cfg.StakingContractAddress = "0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"
	cfg.PaginationSize = 100
	cfg.VoteThreshold = "10"
	cfg.ScoreThreshold = "10"
	cfg.SelfStakingThreshold = "0" // must be 0,because cannot set candidate's StakingTokens
	cfg.CacheSize = 100
	commp, err := NewCommittee(nil, cfg)
	require.NoError(err)

	// get latest block from kovan
	rc, err := commp.(*committee).calculator(10662182)
	require.NoError(err)
	require.NotNil(rc)
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	result, err := rc.Calculate()
	require.NoError(err)
	fmt.Println(result.String())

	//votesBy := result.VotesByDelegate([]byte("candidate1"))
	//for _, v := range votesBy {
	//	fmt.Println(v)
	//}
	//delegates := result.Delegates()
	//require.Equal(2, len(delegates))
	//
	//for _, delegate := range delegates {
	//	fmt.Println(string(delegate.Name()))
	//	fmt.Println(string(delegate.Address()))
	//	fmt.Println(delegate.Score().Text(10))
	//	fmt.Println(delegate.SelfStakingTokens().Text(10))
	//	fmt.Println(delegate.SelfStakingWeight())
	//}
	//expectedVotes := [][]*big.Int{
	//	[]*big.Int{big.NewInt(1960), big.NewInt(660), big.NewInt(1135)},
	//	[]*big.Int{big.NewInt(700), big.NewInt(900), big.NewInt(451)},
	//}
	//for i, d := range delegates {
	//	require.NotNil(expectedDelegates[i])
	//	require.NotNil(expectedVotes[i])
	//	require.True(expectedDelegates[i].equal(d))
	//	for j, v := range result.VotesByDelegate(d.Name()) {
	//		require.NotNil(expectedVotes[i][j])
	//		require.Equal(0, expectedVotes[i][j].Cmp(v.WeightedAmount()))
	//	}
	//}
}
func genTestVotes(mintTime time.Time, require *require.Assertions) []*types.Vote {
	votes := []*types.Vote{}
	// 3 types, bigger equal and smaller than VoteThreshold
	// 3 types, bigger equal and smaller than VoteThreshold
	// votes from voter1
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
	// score 11
	vote, err = types.NewVote(
		mintTime.Add(-2*time.Hour),
		3*time.Hour,
		big.NewInt(11),  //amount
		big.NewInt(100), //weight
		[]byte("voter3"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = types.NewVote(
		mintTime.Add(-2*time.Hour),
		3*time.Hour,
		big.NewInt(5),   //amount
		big.NewInt(100), //weight
		[]byte("voter4"),
		[]byte("candidate4"),
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
		types.NewCandidate(
			[]byte("candidate3"),
			[]byte("candidate3addr"),
			[]byte("operatorPubKey3"),
			[]byte("rewardPubKey3"),
			1,
		),
		types.NewCandidate(
			[]byte("candidate4"),
			[]byte("candidate4addr"),
			[]byte("operatorPubKey4"),
			[]byte("rewardPubKey4"),
			1,
		),
	}
}
