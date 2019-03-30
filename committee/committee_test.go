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
	mintTime := now.Add(-10 * time.Hour)
	candidates := genTestCandidates()
	votes := genTestVotes(mintTime, require)

	var cfg Config
	cfg.NumOfRetries = 8
	cfg.BeaconChainAPIs = []string{"https://mainnet.infura.io/v3/b355cae6fafc4302b106b937ee6c15af"}
	cfg.BeaconChainHeightInterval = 100
	cfg.BeaconChainStartHeight = 7368630
	cfg.RegisterContractAddress = "0x95724986563028deb58f15c5fac19fa09304f32d"
	cfg.StakingContractAddress = "0x87c9dbff0016af23f5b1ab9b8e072124ab729193"
	cfg.PaginationSize = 100
	cfg.VoteThreshold = "0"
	cfg.ScoreThreshold = "0"
	cfg.SelfStakingThreshold = "0"
	cfg.CacheSize = 100
	commp, err := NewCommittee(nil, cfg)
	require.NoError(err)
	rc, err := commp.(*committee).calculator(10)

	require.NoError(err)
	require.NotNil(rc)
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	result, err := rc.Calculate()
	require.NoError(err)
	delegates := result.Delegates()
	require.Equal(4, len(delegates))
	for _, delegate := range delegates {
		fmt.Println(string(delegate.Name()))
		fmt.Println(string(delegate.Address()))
		fmt.Println(delegate.Score())
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
	// (2 + 1) * 100 = 300
	vote, err := types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
	vote, err = types.NewVote(
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
func genTestCandidates() []*types.Candidate {
	return []*types.Candidate{
		types.NewCandidate(
			[]byte("candidate1"),
			[]byte("voter1"),
			[]byte("operatorPubKey1"),
			[]byte("rewardPubKey1"),
			1,
		),
		types.NewCandidate(
			[]byte("candidate2"),
			[]byte("voter2"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			1,
		),
		types.NewCandidate(
			[]byte("candidate3"),
			[]byte("voter3"),
			[]byte("operatorPubKey3"),
			[]byte("rewardPubKey3"),
			10,
		),
		types.NewCandidate(
			[]byte("candidate4"),
			[]byte("voter4"),
			[]byte("operatorPubKey4"),
			[]byte("rewardPubKey4"),
			1,
		),
	}
}
