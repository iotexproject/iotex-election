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
		c := NewRegistration(
			[]byte("new candidate"),
			[]byte("new gravity pub key"),
			[]byte("new operator pub key"),
			[]byte("new reward pub key"),
			1,
		)
		t.Run("fail-to-add-after-votes", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				false,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddRegistrations(candidates))
			require.NoError(calculator.AddBuckets(votes))
			require.Error(calculator.AddRegistrations([]*Registration{c}))
		})
		t.Run("fail-to-add-duplicate", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				false,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddRegistrations(candidates))
			require.Error(calculator.AddRegistrations(candidates))
		})
		t.Run("fail-to-add-after-calculating", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				false,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			calculator.Calculate()
			require.Error(calculator.AddRegistrations(candidates))
		})
		t.Run("success", func(t *testing.T) {
			calculator := NewResultCalculator(
				mintTime,
				false,
				mockVoteFilter(10),
				mockCalcWeight,
				mockCandidateFilter(2000, 1000),
			)
			require.NoError(calculator.AddRegistrations(candidates))
			require.Equal(len(candidates), len(calculator.candidates))
			require.NoError(calculator.AddRegistrations([]*Registration{c}))
			require.Equal(len(candidates)+1, len(calculator.candidates))
		})
	})
	t.Run("failed-to-add-votes", func(t *testing.T) {
		calculator := NewResultCalculator(
			mintTime,
			false,
			mockVoteFilter(10),
			mockCalcWeight,
			mockCandidateFilter(2000, 1000),
		)
		calculator.AddRegistrations([]*Registration{NewRegistration(
			[]byte("candidate1"),
			[]byte("voter1"),
			[]byte("operator"),
			[]byte("reward"),
			2,
		)})
		t.Run("vote-for-candidate-not-in-pool", func(t *testing.T) {
			vote, err := NewBucket(
				now.Add(-20*time.Hour),
				time.Hour*3,
				big.NewInt(10),
				[]byte("voter2"),
				[]byte("candidate not in list"),
				false,
			)
			require.NoError(err)
			require.NoError(calculator.AddBuckets([]*Bucket{vote}))
			require.Equal(0, big.NewInt(40).Cmp(calculator.totalVotes))
		})
		t.Run("vote-not-qualified", func(t *testing.T) {
			vote, err := NewBucket(
				now.Add(-20*time.Hour),
				time.Hour*3,
				big.NewInt(9),
				[]byte("voter2"),
				[]byte("candidate1"),
				false,
			)
			require.NoError(err)
			calculator.AddBuckets([]*Bucket{vote})
			require.Equal(0, big.NewInt(40).Cmp(calculator.totalVotes))
		})
		t.Run("after-calculated", func(t *testing.T) {
			calculator.Calculate()
			require.Error(calculator.AddBuckets(votes))
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
			false,
			mockVoteFilter(10),
			mockCalcWeight,
			mockCandidateFilter(2000, 1000),
		)
		candidate5 := NewRegistration(
			[]byte("candidate5"),
			[]byte("voter5"),
			[]byte("operatorPubKey5"),
			[]byte("rewardPubKey5"),
			2,
		)
		require.NoError(calculator.AddRegistrations(candidates))
		require.NoError(calculator.AddRegistrations([]*Registration{candidate5}))
		require.NoError(calculator.AddBuckets(votes))
		result, err := calculator.Calculate()
		require.NoError(err)
		delegates := result.Delegates()
		require.Equal(2, len(delegates))
		ec4 := NewCandidate(
			candidates[2],
			big.NewInt(2051),
			big.NewInt(1000),
		)
		ec5 := NewCandidate(
			candidate5,
			big.NewInt(3755),
			big.NewInt(1010),
		)
		expectedDelegates := []*Candidate{ec5, ec4}
		expectedVotes := [][]*big.Int{
			[]*big.Int{big.NewInt(1960), big.NewInt(660), big.NewInt(1135)},
			[]*big.Int{big.NewInt(700), big.NewInt(900), big.NewInt(451)},
		}
		for i, d := range delegates {
			require.NotNil(expectedDelegates[i])
			require.NotNil(expectedVotes[i])
			require.True(expectedDelegates[i].Equal(d))
			for j, v := range result.VotesByDelegate(d.Name()) {
				require.NotNil(expectedVotes[i][j])
				require.Equal(0, expectedVotes[i][j].Cmp(v.WeightedAmount()))
			}
		}
	})
}

func mockCalcWeight(v *Bucket, t time.Time) *big.Int {
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

func mockVoteFilter(VoteThreshold int64) BucketFilterFunc {
	return func(v *Bucket) bool {
		return v.Amount().Cmp(big.NewInt(VoteThreshold)) < 0
	}
}

func genTestVotes(mintTime time.Time, require *require.Assertions) []*Bucket {
	votes := []*Bucket{}
	// votes from voter1
	// (2 + 1) * 100 = 300
	vote, err := NewBucket(
		mintTime.Add(-2*time.Hour),
		4*time.Hour,
		big.NewInt(100),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// will be filtered with low amount
	vote, err = NewBucket(
		mintTime.Add(-85*time.Hour),
		100*time.Hour,
		big.NewInt(9),
		[]byte("voter1"),
		[]byte("candidate1"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter2
	// (1 + 1) * 100 = 200
	vote, err = NewBucket(
		mintTime.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(100),
		[]byte("voter2"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter3
	// 70 * 10 = 700
	// 1 * 70 * 10 = 700
	vote, err = NewBucket(
		mintTime.Add(-3*time.Hour),
		1*time.Hour,
		big.NewInt(70),
		[]byte("voter3"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// 30 * 10 = 300
	// (2 + 1) * 30 * 10 = 600
	vote, err = NewBucket(
		mintTime.Add(-1*time.Hour),
		2*time.Hour,
		big.NewInt(30),
		[]byte("voter3"),
		[]byte("candidate3"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter4
	// (9 + 1) * 50 = 500
	vote, err = NewBucket(
		mintTime.Add(-10*time.Hour),
		9*time.Hour,
		big.NewInt(50),
		[]byte("voter4"),
		[]byte("candidate4"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (40 + 1) * 20 = 820
	vote, err = NewBucket(
		mintTime.Add(-60*time.Hour),
		100*time.Hour,
		big.NewInt(20),
		[]byte("voter4"),
		[]byte("candidate4"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter5
	// 490 * 2 = 980
	// (1 + 1) * 490 * 2 = 1960
	vote, err = NewBucket(
		mintTime.Add(-6*time.Hour),
		7*time.Hour,
		big.NewInt(490),
		[]byte("voter5"),
		[]byte("candidate5"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// 15 * 2 = 30
	// (21 + 1) * 15 * 2 = 660
	vote, err = NewBucket(
		mintTime.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(15),
		[]byte("voter5"),
		[]byte("candidate5"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter6
	// 1 * 35 = 35
	vote, err = NewBucket(
		mintTime.Add(-3*time.Hour),
		0,
		big.NewInt(1135),
		[]byte("voter6"),
		[]byte("candidate5"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter7
	// start time > mint time, filtered
	vote, err = NewBucket(
		mintTime.Add(1*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		[]byte("voter7"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (21 + 1) * 90 = 1980
	vote, err = NewBucket(
		mintTime.Add(-1*time.Hour),
		21*time.Hour,
		big.NewInt(90),
		[]byte("voter7"),
		[]byte("candidate2"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (11 + 1) * 41 = 492
	vote, err = NewBucket(
		mintTime.Add(-1*time.Hour),
		11*time.Hour,
		big.NewInt(41),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// (10 + 1) * 70 = 770
	vote, err = NewBucket(
		mintTime.Add(-10*time.Hour),
		20*time.Hour,
		big.NewInt(70),
		[]byte("voter7"),
		[]byte("candidate4"),
		true,
	)
	require.NoError(err)
	return append(votes, vote)
}
