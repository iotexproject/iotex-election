// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCandidate(t *testing.T) {
	require := require.New(t)
	c := NewCandidate(
		[]byte("candidate1"),
		[]byte("voter1"),
		[]byte("operatorPubKey1"),
		[]byte("rewardPubKey1"),
		1,
	)
	t.Run("protobuf", func(t *testing.T) {
		cPb, err := c.ToProtoMsg()
		require.NoError(err)
		cc := &Candidate{}
		require.NoError(cc.FromProtoMsg(cPb))
		require.True(c.equal(cc))
	})
	t.Run("clone", func(t *testing.T) {
		cc := c.Clone()
		require.True(c.equal(cc))
	})
	t.Run("add-scores", func(t *testing.T) {
		require.Error(c.addScore(big.NewInt(-1)))
		require.Error(c.addSelfStakingScore(big.NewInt(-1)))
		require.NoError(c.addScore(big.NewInt(1)))
		require.Equal(0, c.Score().Cmp(big.NewInt(1)))
		require.NoError(c.addSelfStakingScore(big.NewInt(1)))
		require.Equal(0, c.SelfStakingScore().Cmp(big.NewInt(1)))
		c.reset()
	})
}
