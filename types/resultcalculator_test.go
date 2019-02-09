// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

/*
func TestSerialize(t *testing.T) {
	require := require.New(t)
	now := time.Now()

	//	selfvotes	totalvotes
	//	c1	false		false
	//	c2	false		true
	//	c3	true		false
	//	c4	true		true

	candidates := genTestCandidates()
	votes, err := genTestVotes()
	require.NoError(err)
	calculator := NewResultCalculator(
		now.Add(7*24*time.Hour),
		func(v *Vote, t time.Time) *big.Int {
			if t.Before(v.StartTime()) {
				return big.NewInt(0)
			}
			remainingTime := v.RemainingTime(t).Seconds()
			weight := 1.0
			if remainingTime > 0 {
				weight += math.Ceil(remainingTime / 86400)
			}
			amount := new(big.Float).SetInt(v.amount)
			weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
			return weightedAmount
		},
		func(c *Candidate) bool {
			return new(big.Int).SetUint64(2000).Cmp(c.Score()) > 0 &&
				c.SelfStakingScore().Cmp(big.NewInt(1000)) < 0
		},
	)
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
}

func genTestVotes(now time.Time) (votes []*Vote, err error) {
	var vote *Vote
	votes = []*Vote{}

	// votes from voter1
	if vote, err = NewVote(
		now,
		2*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter1"),
		[]byte("candidate1"),
		false,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter2
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(100),
		big.NewInt(0),
		[]byte("voter2"),
		[]byte("candidate2"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter3
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		1*time.Hour,
		big.NewInt(4),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		false,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	if vote, err = NewVote(
		now.Add(-9*time.Hour),
		2*time.Hour,
		big.NewInt(2),
		big.NewInt(0),
		[]byte("voter3"),
		[]byte("candidate3"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter4
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		10*time.Hour,
		big.NewInt(20),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		100*time.Hour,
		big.NewInt(10),
		big.NewInt(0),
		[]byte("voter4"),
		[]byte("candidate4"),
		false,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter5
	if vote, err = NewVote(
		now.Add(-6*time.Hour),
		14*time.Hour,
		big.NewInt(2),
		big.NewInt(19),
		[]byte("voter5"),
		[]byte("candidate1"),
		false,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter5"),
		[]byte("candidate2"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter6
	if vote, err = NewVote(
		now.Add(3*time.Hour),
		0,
		big.NewInt(35),
		big.NewInt(20),
		[]byte("voter6"),
		[]byte("candidate1"),
		false,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	// votes from voter7
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate2"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	if vote, err = NewVote(
		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	if vote, err = NewVote(

		now.Add(-10*time.Hour),
		21*time.Hour,
		big.NewInt(45),
		big.NewInt(21),
		[]byte("voter7"),
		[]byte("candidate3"),
		true,
	); err != nil {
		return
	}
	votes = append(votes, vote)
	return
}
*/
