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
		require.Error(c.addSelfStakingTokens(big.NewInt(-1)))
		require.NoError(c.addScore(big.NewInt(1)))
		require.Equal(0, c.Score().Cmp(big.NewInt(1)))
		require.NoError(c.addSelfStakingTokens(big.NewInt(1)))
		require.Equal(0, c.SelfStakingTokens().Cmp(big.NewInt(1)))
		c.reset()
	})
}
