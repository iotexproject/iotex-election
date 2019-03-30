// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"fmt"
	"math"
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

	//var cfg Config
	//cfg.NumOfRetries = 8
	//cfg.BeaconChainAPIs = []string{"https://mainnet.infura.io/v3/b355cae6fafc4302b106b937ee6c15af"}
	//cfg.BeaconChainHeightInterval = 100
	//cfg.BeaconChainStartHeight = 7368630
	//cfg.RegisterContractAddress = "0x95724986563028deb58f15c5fac19fa09304f32d"
	//cfg.StakingContractAddress = "0x87c9dbff0016af23f5b1ab9b8e072124ab729193"
	//cfg.PaginationSize = 100
	//cfg.VoteThreshold = "0"
	//cfg.ScoreThreshold = "0"
	//cfg.SelfStakingThreshold = "0"
	//cfg.CacheSize = 100
	//commp, err := NewCommittee(nil, cfg)
	//require.NoError(err)
	//rc, err := commp.(*committee).calculator(10)
	rc := types.NewResultCalculator(
		mintTime,
		mockVoteFilter(8),
		mockCalcWeight,
		mockCandidateFilter(10, 80),
	)
	//require.NoError(err)
	require.NotNil(rc)
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	result, err := rc.Calculate()
	require.NoError(err)
	fmt.Println(result.String())

	votesBy := result.VotesByDelegate([]byte("voter1"))
	for _, v := range votesBy {
		fmt.Println(v)
	}
	delegates := result.Delegates()
	require.Equal(4, len(delegates))

	for _, delegate := range delegates {
		fmt.Println(string(delegate.Name()))
		fmt.Println(string(delegate.Address()))
		fmt.Println(delegate.Score().Text(10))
		fmt.Println(delegate.SelfStakingTokens().Text(10))
		fmt.Println(delegate.SelfStakingWeight())
	}
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
	// votes from voter1
	// (10 -3+ 1) * 100 = 800
	vote, err := types.NewVote(
		mintTime.Add(-3*time.Hour),
		10*time.Hour,
		big.NewInt(100),
		big.NewInt(11),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (3-2 + 1) * 9 = 18?
	vote, err = types.NewVote(
		mintTime.Add(-2*time.Hour),
		3*time.Hour,
		big.NewInt(9), //amount
		big.NewInt(1), //weight
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
			11,
		),
		types.NewCandidate(
			[]byte("candidate2"),
			[]byte("candidate2addr"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			10,
		),
	}
}
func mockCalcWeight(v *types.Vote, t time.Time) *big.Int {
	if t.Before(v.StartTime()) {
		return big.NewInt(0)
	}
	remainingTime := v.RemainingTime(t).Seconds()
	weight := 1.0
	if remainingTime > 0 {
		weight += math.Ceil(remainingTime / time.Hour.Seconds())
	}
	amount := new(big.Float).SetInt(v.Amount())
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount
}

func mockCandidateFilter(
	ScoreThreshold int64,
	SelfStakingTokenThreshold int64,
) types.CandidateFilterFunc {
	return func(c *types.Candidate) bool {
		return c.Score().Cmp(big.NewInt(ScoreThreshold)) < 0 ||
			c.SelfStakingTokens().Cmp(big.NewInt(SelfStakingTokenThreshold)) < 0
	}
}

func mockVoteFilter(VoteThreshold int64) types.VoteFilterFunc {
	return func(v *types.Vote) bool {
		return v.Amount().Cmp(big.NewInt(VoteThreshold)) < 0
	}
}
