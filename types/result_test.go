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
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestElectionResult(t *testing.T) {
	require := require.New(t)
	t.Run("without-data", func(t *testing.T) {
		result := &ElectionResult{
			mintTime:         time.Now(),
			delegates:        []*Candidate{},
			votes:            map[string][]*Vote{},
			totalVotes:       new(big.Int),
			totalVotedStakes: new(big.Int),
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &ElectionResult{}
		require.NoError(clone.Deserialize(b))
		require.True(result.mintTime.Equal(clone.mintTime))
		require.Equal(0, len(clone.delegates))
		require.Equal(0, len(clone.votes))
	})
	t.Run("with-data", func(t *testing.T) {
		candidates := genTestCandidates()
		votes := map[string][]*Vote{}
		for _, c := range candidates {
			v1, err := NewVote(
				time.Now(),
				2*time.Hour,
				big.NewInt(int64(rand.Intn(10000000))),
				big.NewInt(int64(rand.Intn(10000000))),
				[]byte("voter1"),
				c.name,
				false,
			)
			require.NoError(err)
			v2, err := NewVote(
				time.Now(),
				5*time.Hour,
				big.NewInt(int64(rand.Intn(10000000))),
				big.NewInt(int64(rand.Intn(10000000))),
				[]byte("voter2"),
				c.name,
				true,
			)
			require.NoError(err)
			votes[hex.EncodeToString(c.name)] = []*Vote{v1, v2}
		}
		result := &ElectionResult{
			mintTime:         time.Now(),
			delegates:        candidates,
			votes:            votes,
			totalVotes:       new(big.Int),
			totalVotedStakes: new(big.Int),
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &ElectionResult{}
		require.NoError(clone.Deserialize(b))
		require.True(result.mintTime.Equal(clone.mintTime))
		cs := result.Delegates()
		ccs := clone.Delegates()
		require.Equal(len(cs), len(ccs))
		require.Equal(len(result.votes), len(clone.votes))
		for i, c := range cs {
			require.True(c.equal(ccs[i]))
			vs := result.VotesByDelegate(c.Name())
			cvs := clone.VotesByDelegate(c.Name())
			require.Equal(len(vs), len(cvs))
			for j, v := range vs {
				require.True(v.equal(cvs[j]))
			}
		}
	})
}

func genTestCandidates() []*Candidate {
	return []*Candidate{
		NewCandidate(
			[]byte("candidate1"),
			[]byte("voter1"),
			[]byte("operatorPubKey1"),
			[]byte("rewardPubKey1"),
			1,
		),
		NewCandidate(
			[]byte("candidate2"),
			[]byte("voter2"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			1,
		),
		NewCandidate(
			[]byte("candidate3"),
			[]byte("voter3"),
			[]byte("operatorPubKey3"),
			[]byte("rewardPubKey3"),
			10,
		),
		NewCandidate(
			[]byte("candidate4"),
			[]byte("voter4"),
			[]byte("operatorPubKey4"),
			[]byte("rewardPubKey4"),
			1,
		),
	}
}
