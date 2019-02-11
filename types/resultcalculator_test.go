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
	candidates := genTestCandidates()
	votes := genTestVotes(now, require)
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
				now.Add(-10*time.Hour),
				mockVoteFilter,
				mockCalcWeight,
				mockCandidateFilter,
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.NoError(calculator.AddVotes(votes))
			require.Error(calculator.AddCandidates([]*Candidate{c}))
		})
		t.Run("fail-to-add-duplicate", func(t *testing.T) {
			calculator := NewResultCalculator(
				now.Add(-10*time.Hour),
				mockVoteFilter,
				mockCalcWeight,
				mockCandidateFilter,
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.Error(calculator.AddCandidates(candidates))
		})
		t.Run("fail-to-add-after-calculating", func(t *testing.T) {
			calculator := NewResultCalculator(
				now.Add(-10*time.Hour),
				mockVoteFilter,
				mockCalcWeight,
				mockCandidateFilter,
			)
			calculator.Calculate()
			require.Error(calculator.AddCandidates(candidates))
		})
		t.Run("success", func(t *testing.T) {
			calculator := NewResultCalculator(
				now.Add(-10*time.Hour),
				mockVoteFilter,
				mockCalcWeight,
				mockCandidateFilter,
			)
			require.NoError(calculator.AddCandidates(candidates))
			require.Equal(len(candidates), len(calculator.candidates))
			require.NoError(calculator.AddCandidates([]*Candidate{c}))
			require.Equal(len(candidates)+1, len(calculator.candidates))
		})
	})
	t.Run("failed-to-add-votes", func(t *testing.T) {
		calculator := NewResultCalculator(
			now.Add(-10*time.Hour),
			mockVoteFilter,
			mockCalcWeight,
			mockCandidateFilter,
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
			calculator.AddVotes([]*Vote{vote})
			require.Equal(int32(0), calculator.totalVotes)
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
			require.Equal(int32(0), calculator.totalVotes)
		})
		t.Run("after-calculated", func(t *testing.T) {
			calculator.Calculate()
			require.Error(calculator.AddVotes(votes))
		})

	})

	//		selfvotes	totalvotes
	// c1	false		false
	// c2	false		true
	// c3	true		false
	// c4	true		true
	/*
		require.NoError(calculator.AddCandidates(candidates))
		require.NoError(calculator.AddVotes(votes))
		result, err := calculator.Calculate()
		b, err := result.Serialize()
		require.NoError(err)
		r := &Result{}
		require.NoError(r.Deserialize(b))
		require.Equal(result.mintTime, r.mintTime)
		require.Equal(4, len(r.votes))
		require.Equal(2, len(r.scores))
		require.Equal(result.scores["candidate1"], r.scores["candidate1"])
		require.Equal(result.scores["candidate2"], r.scores["candidate2"])
		for i := 0; i < 4; i++ {
			expectedV, err := result.votes[0].Serialize()
			require.NoError(err)
			actualV, err := r.votes[0].Serialize()
			require.NoError(err)
			require.True(bytes.Equal(expectedV, actualV))
		}
		rb, err := r.Serialize()
		require.NoError(err)
		require.True(bytes.Equal(b, rb))
	*/
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

func mockCandidateFilter(c *Candidate) bool {
	return c.Score().Cmp(big.NewInt(2000)) < 0 &&
		c.SelfStakingScore().Cmp(big.NewInt(1000)) < 0
}

func mockVoteFilter(v *Vote) bool {
	return v.Amount().Cmp(big.NewInt(10)) < 0
}

func genTestVotes(now time.Time, require *require.Assertions) []*Vote {
	votes := []*Vote{}
	// votes from voter1
	vote, err := NewVote(
		now,
		2*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter1"),
		[]byte("candidate1"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter2
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter2"),
		[]byte("candidate2"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter3
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(4),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = NewVote(
		now.Add(-9*time.Hour),
		2*time.Hour,
		big.NewInt(2),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter4
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		10*time.Hour,
		big.NewInt(20),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		100*time.Hour,
		big.NewInt(10),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter5
	vote, err = NewVote(
		now.Add(-6*time.Hour),
		14*time.Hour,
		big.NewInt(2),
		big.NewInt(19),
		[]byte("voter5"),
		[]byte("candidate1"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter5"),
		[]byte("candidate2"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter6
	vote, err = NewVote(
		now.Add(3*time.Hour),
		0,
		big.NewInt(35),
		big.NewInt(20),
		[]byte("voter6"),
		[]byte("candidate1"),
		false,
	)
	require.NoError(err)
	votes = append(votes, vote)
	// votes from voter7
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate2"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	votes = append(votes, vote)
	vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	)
	require.NoError(err)
	return append(votes, vote)
}
