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
	"crypto/sha256"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPoll(t *testing.T) {
	require := require.New(t)
	t.Run("without-data", func(t *testing.T) {
		result := &Poll{
			regs:    [][]byte{},
			buckets: [][]byte{},
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &Poll{}
		require.NoError(clone.Deserialize(b))
		require.Equal(0, len(clone.regs))
		require.Equal(0, len(clone.buckets))
	})
	t.Run("with-data", func(t *testing.T) {
		candidates := genTestCandidates()
		votes := []*Vote{}
		for _, c := range candidates {
			b1, err := NewBucket(
				big.NewInt(int64(rand.Intn(10000))),
				time.Now(),
				2*time.Hour,
				big.NewInt(int64(rand.Intn(10000000))),
				[]byte("voter1"),
				c.name,
				false,
			)
			require.NoError(err)
			v1, err := NewVote(
				b1,
				big.NewInt(int64(rand.Intn(10000000))),
			)
			require.NoError(err)
			b2, err := NewBucket(
				big.NewInt(int64(rand.Intn(10000))),
				time.Now(),
				5*time.Hour,
				big.NewInt(int64(rand.Intn(10000000))),
				[]byte("voter2"),
				c.name,
				true,
			)
			require.NoError(err)
			v2, err := NewVote(
				b2,
				big.NewInt(int64(rand.Intn(10000000))),
			)
			require.NoError(err)
			votes = append(votes, v1)
			votes = append(votes, v2)
		}

		votesKey := [][]byte{}
		for _, v := range votes {
			data, err := v.Serialize()
			require.NoError(err)
			hashval := sha256.Sum256(data)
			hashbytes := make([]byte, 32)
			for _, num := range hashval {
				hashbytes = append(hashbytes, num)
			}
			votesKey = append(votesKey, hashbytes)
		}
		candidatesKey := [][]byte{}
		for _, c := range candidates {
			data, err := c.Serialize()
			require.NoError(err)
			hashval := sha256.Sum256(data)
			hashbytes := make([]byte, 32)
			for _, num := range hashval {
				hashbytes = append(hashbytes, num)
			}
			candidatesKey = append(candidatesKey, hashbytes)
		}

		result := &Poll{
			regs:    candidatesKey,
			buckets: votesKey,
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &Poll{}
		require.NoError(clone.Deserialize(b))
		require.Equal(len(result.Registrations()), 4)
		require.Equal(len(clone.Registrations()), 4)
		require.Equal(len(result.Buckets()), 8)
		require.Equal(len(clone.Buckets()), 8)

	})
}

func genTestCandidates() []*Registration {
	return []*Registration{
		NewRegistration(
			[]byte("candidate1"),
			[]byte("voter1"),
			[]byte("operatorPubKey1"),
			[]byte("rewardPubKey1"),
			1,
		),
		NewRegistration(
			[]byte("candidate2"),
			[]byte("voter2"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			1,
		),
		NewRegistration(
			[]byte("candidate3"),
			[]byte("voter3"),
			[]byte("operatorPubKey3"),
			[]byte("rewardPubKey3"),
			10,
		),
		NewRegistration(
			[]byte("candidate4"),
			[]byte("voter4"),
			[]byte("operatorPubKey4"),
			[]byte("rewardPubKey4"),
			1,
		),
	}
}
