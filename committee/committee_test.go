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
	var cfg Config
	cfg.NumOfRetries = 8
	cfg.BeaconChainAPIs = []string{"https://mainnet.infura.io/v3/b355cae6fafc4302b106b937ee6c15af"}
	cfg.BeaconChainHeightInterval = 100
	cfg.BeaconChainStartHeight = 7368630
	cfg.RegisterContractAddress = "0x95724986563028deb58f15c5fac19fa09304f32d"
	cfg.StakingContractAddress = "0x87c9dbff0016af23f5b1ab9b8e072124ab729193"
	cfg.PaginationSize = 100
	cfg.VoteThreshold = "3"
	cfg.ScoreThreshold = "4"
	cfg.SelfStakingThreshold = "5"
	cfg.CacheSize = 100
	commp, err := NewCommittee(nil, cfg)
	require.NoError(err)
	rc, err := commp.(*committee).calculator(10)
	require.NoError(err)
	require.NotNil(rc)
	candidate := types.NewCandidate(
		[]byte("candidate1"),
		[]byte("voter1"),
		[]byte("operatorPubKey1"),
		[]byte("rewardPubKey1"),
		1,
	)
	candidates := []*types.Candidate{candidate}
	vote, err := types.NewVote(
		time.Now(),
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.Error(err)
	require.NotNil(vote)

	votes := []*types.Vote{vote}
	require.NoError(rc.AddCandidates(candidates))
	require.NoError(rc.AddVotes(votes))
	require.NoError(err)
	sorted, err := rc.Calculate()
	require.NoError(err)
	fmt.Println(sorted.String())
}
