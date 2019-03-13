// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResultCalculator(t *testing.T) {
	require := require.New(t)
	now := time.Now()
	mintTime := now.Add(-10 * time.Hour)
	candidates := genTestCandidates()
	votes := genTestVotes(mintTime, require)
	t.Run("add-candidates", func(t *testing.T) {
		c := NewCandidate(
			[]byte("new candidate"),
			[]byte("new beacon pub key"),
			[]byte("new operator pub key"),
			[]byte("new reward pub key"),
			1,
		)
		t.Run("fail-to-add-after-votes", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.NoError(calculator.AddVotes(votes))
			require.Error(calculator.AddCandidates([]*Candidate{c}))
		})
		t.Run("fail-to-add-duplicate", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.Error(calculator.AddCandidates(candidates))
		})
		t.Run("fail-to-add-after-calculating", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			calculator.Calculate()
			require.Error(calculator.AddCandidates(candidates))
		})
		t.Run("success", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.Equal(len(candidates), len(calculator.candidates))
			require.NoError(calculator.AddCandidates([]*Candidate{c}))
			require.Equal(len(candidates)+1, len(calculator.candidates))
		})
	})
	t.Run("failed-to-add-votes", func(t *testing.T) {
		calculator := NewResultCalculator(
			mintTime,
			mockVoteFilter(10),
			mockCalcWeight,
			mockCandidateFilter(2000, 1000),
		)
		calculator.AddCandidates([]*Candidate{NewCandidate(
			[]byte("candidate1"),
			[]byte("voter1"),
			[]byte("operator"),
			[]byte("reward"),
			2,
		)})
		t.Run("vote-for-candidate-not-in-pool", func(t *testing.T) {
			vote, err := NewVote(
				now.Add(-20*time.Hour),
				time.Hour*3,
				big.NewInt(10),
				big.NewInt(0),
				[]byte("voter2"),
				[]byte("candidate not in list"),
				false,
			)
			require.NoError(err)
			require.NoError(calculator.AddVotes([]*Vote{vote}))
			require.Equal(0, big.NewInt(40).Cmp(calculator.totalVotes))
		})
		t.Run("vote-not-qualified", func(t *testing.T) {
			vote, err := NewVote(
				now.Add(-20*time.Hour),
				time.Hour*3,
				big.NewInt(9),
				big.NewInt(0),
				[]byte("voter2"),
				[]byte("candidate1"),
				false,
			)
			require.NoError(err)
			calculator.AddVotes([]*Vote{vote})
			require.Equal(0, big.NewInt(40).Cmp(calculator.totalVotes))
		})
		t.Run("after-calculated", func(t *testing.T) {
			calculator.Calculate()
			require.Error(calculator.AddVotes(votes))
		})

	})
	t.Run("calculate", func(t *testing.T) {
		//		selfvotes	totalvotes
		// c1	false		false
		// c2	false		true
		// c3	true		false
		// c4	true		true (lower rank)
		// c5	true		true (higher rank)
		calculator := NewResultCalculator(
			mintTime,
			mockVoteFilter(10),
			mockCalcWeight,
			mockCandidateFilter(2000, 1000),
		)
		candidate5 := NewCandidate(
			[]byte("candidate5"),
			[]byte("voter5"),
			[]byte("operatorPubKey5"),
			[]byte("rewardPubKey5"),
			2,
		)
		require.NoError(calculator.AddCandidates(candidates))
		require.NoError(calculator.AddCandidates([]*Candidate{candidate5}))
		require.NoError(calculator.AddVotes(votes))
		result, err := calculator.Calculate()
		require.NoError(err)
		delegates := result.Delegates()
		require.Equal(2, len(delegates))
		ec4 := candidates[2].Clone()
		ec4.score = big.NewInt(2051)
		ec4.selfStakingTokens = big.NewInt(1000)
		ec5 := candidate5.Clone()
		ec5.score = big.NewInt(3755)
		ec5.selfStakingTokens = big.NewInt(1010)
		expectedDelegates := []*Candidate{ec5, ec4}
		expectedVotes := [][]*big.Int{
			[]*big.Int{big.NewInt(1960), big.NewInt(660), big.NewInt(1135)},
			[]*big.Int{big.NewInt(700), big.NewInt(900), big.NewInt(451)},
		}
		for i, d := range delegates {
			require.NotNil(expectedDelegates[i])
			require.NotNil(expectedVotes[i])
			require.True(expectedDelegates[i].equal(d))
			for j, v := range result.VotesByDelegate(d.Name()) {
				require.NotNil(expectedVotes[i][j])
				require.Equal(0, expectedVotes[i][j].Cmp(v.WeightedAmount()))
			}
		}
	})
}

func mockCalcWeight(v *Vote, t time.Time) *big.Int {
	if t.Before(v.StartTime()) {
		return big.NewInt(0)
	}
	remainingTime := v.RemainingTime(t).Seconds()
	weight := 1.0
	if remainingTime > 0 {
		weight += math.Ceil(remainingTime / time.Hour.Seconds())
	}
	amount := new(big.Float).SetInt(v.amount)
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount
}

func mockCandidateFilter(
	ScoreThreshold int64,
	SelfStakingTokenThreshold int64,
) CandidateFilterFunc {
	return func(c *Candidate) bool {
		return c.Score().Cmp(big.NewInt(ScoreThreshold)) < 0 ||
			c.SelfStakingTokens().Cmp(big.NewInt(SelfStakingTokenThreshold)) < 0
	}
}

func mockVoteFilter(VoteThreshold int64) VoteFilterFunc {
	return func(v *Vote) bool {
		return v.Amount().Cmp(big.NewInt(VoteThreshold)) < 0
	}
}

func genTestVotes(mintTime time.Time, require *require.Assertions) []*Vote {
	votes := []*Vote{}
	// votes from voter1
	// (2 + 1) * 100 = 300
	vote, err := NewVote(
		mintTime.Add(-2*time.Hour),
		4*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// will be filtered with low amount
	vote, err = NewVote(
		mintTime.Add(-85*time.Hour),
		100*time.Hour,
		big.NewInt(9),
		big.NewInt(0),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter2
	// (1 + 1) * 100 = 200
	vote, err = NewVote(
		mintTime.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter2"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter3
	// 70 * 10 = 700
	// 1 * 70 * 10 = 700
	vote, err = NewVote(
		mintTime.Add(-3*time.Hour),
		1*time.Hour,
		big.NewInt(70),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// 30 * 10 = 300
	// (2 + 1) * 30 * 10 = 600
	vote, err = NewVote(
		mintTime.Add(-1*time.Hour),
		2*time.Hour,
		big.NewInt(30),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter4
	// (9 + 1) * 50 = 500
	vote, err = NewVote(
		mintTime.Add(-10*time.Hour),
		9*time.Hour,
		big.NewInt(50),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (40 + 1) * 20 = 820
	vote, err = NewVote(
		mintTime.Add(-60*time.Hour),
		100*time.Hour,
		big.NewInt(20),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter5
	// 490 * 2 = 980
	// (1 + 1) * 490 * 2 = 1960
	vote, err = NewVote(
		mintTime.Add(-6*time.Hour),
		7*time.Hour,
		big.NewInt(490),
		big.NewInt(0),
		[]byte("voter5"),
		[]byte("candidate5"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// 15 * 2 = 30
	// (21 + 1) * 15 * 2 = 660
	vote, err = NewVote(
		mintTime.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(15),
		big.NewInt(0),
		[]byte("voter5"),
		[]byte("candidate5"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter6
	// 1 * 35 = 35
	vote, err = NewVote(
		mintTime.Add(-3*time.Hour),
		0,
		big.NewInt(1135),
		big.NewInt(0),
		[]byte("voter6"),
		[]byte("candidate5"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter7
	// start time > mint time, filtered
	vote, err = NewVote(
		mintTime.Add(1*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(0),
		[]byte("voter7"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (21 + 1) * 90 = 1980
	vote, err = NewVote(
		mintTime.Add(-1*time.Hour),
		21*time.Hour,
		big.NewInt(90),
		big.NewInt(0),
		[]byte("voter7"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (11 + 1) * 41 = 492
	vote, err = NewVote(
		mintTime.Add(-1*time.Hour),
		11*time.Hour,
		big.NewInt(41),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (10 + 1) * 70 = 770
	vote, err = NewVote(
		mintTime.Add(-10*time.Hour),
		20*time.Hour,
		big.NewInt(70),
		big.NewInt(0),
		[]byte("voter7"),
		[]byte("candidate4"),
		true,
	)
	require.NoError(err)
	return append(votes, vote)
}
