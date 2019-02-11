// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResult(t *testing.T) {
	require := require.New(t)
	t.Run("without-data", func(t *testing.T) {
		result := &Result{
			mintTime:   time.Now(),
			candidates: []*Candidate{},
			votes:      map[string][]*Vote{},
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &Result{}
		require.NoError(clone.Deserialize(b))
		require.True(result.mintTime.Equal(clone.mintTime))
		require.Equal(0, len(clone.candidates))
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
		result := &Result{
			mintTime:   time.Now(),
			candidates: candidates,
			votes:      votes,
		}
		b, err := result.Serialize()
		require.NoError(err)
		clone := &Result{}
		require.NoError(clone.Deserialize(b))
		require.True(result.mintTime.Equal(clone.mintTime))
		cs := result.Candidates()
		ccs := clone.Candidates()
		require.Equal(len(cs), len(ccs))
		require.Equal(len(result.votes), len(clone.votes))
		for i, c := range cs {
			require.True(c.equal(ccs[i]))
			vs := result.VotesByCandidate(c.Name())
			cvs := clone.VotesByCandidate(c.Name())
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
			100,
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
