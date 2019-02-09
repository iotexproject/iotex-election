// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewVote(t *testing.T) {
	require := require.New(t)
	startTime := time.Now()
	t.Run("failed-to-new-with-negative-duration", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			-24*7*time.Hour,
			big.NewInt(3),
			big.NewInt(3),
			[]byte{},
			[]byte{},
			true,
		)
		require.Error(err)
		require.Nil(vote)
	})
	t.Run("failed-to-new-with-negative-amount", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			24*7*time.Hour,
			big.NewInt(-3),
			big.NewInt(3),
			[]byte{},
			[]byte{},
			true,
		)
		require.Error(err)
		require.Nil(vote)
	})
	t.Run("failed-to-new-with-negative-weighted-amount", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			24*7*time.Hour,
			big.NewInt(3),
			big.NewInt(-3),
			[]byte{},
			[]byte{},
			true,
		)
		require.Error(err)
		require.Nil(vote)
	})
	t.Run("success-new-a-vote", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			24*7*time.Hour,
			big.NewInt(3),
			big.NewInt(3),
			[]byte("Voter"),
			[]byte("Candidate"),
			true,
		)
		require.NoError(err)
		require.NotNil(vote)
		t.Run("serialize", func(t *testing.T) {
			b, err := vote.Serialize()
			require.NoError(err)
			clone := &Vote{}
			require.NoError(clone.Deserialize(b))
			require.True(vote.equal(clone))
		})
	})
}

func TestRemainingTime(t *testing.T) {
	require := require.New(t)
	startTime := time.Now()
	stakingDuration := 24 * 7 * time.Hour
	t.Run("decay-vote", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			stakingDuration,
			big.NewInt(3),
			big.NewInt(3),
			[]byte{},
			[]byte{},
			true,
		)
		require.NoError(err)
		require.NotNil(vote)
		t.Run("before-start-time", func(t *testing.T) {
			require.Equal(0*time.Hour, vote.RemainingTime(startTime.Add(-time.Hour)))
		})
		t.Run("after-end-time", func(t *testing.T) {
			require.Equal(0*time.Hour, vote.RemainingTime(startTime.Add(stakingDuration+time.Hour)))
		})
		t.Run("valid-time", func(t *testing.T) {
			require.Equal(time.Hour, vote.RemainingTime(startTime.Add(stakingDuration-time.Hour)))
		})
	})
	t.Run("non-decay-vote", func(t *testing.T) {
		vote, err := NewVote(
			startTime,
			stakingDuration,
			big.NewInt(3),
			big.NewInt(3),
			[]byte{},
			[]byte{},
			false,
		)
		require.NoError(err)
		require.NotNil(vote)
		t.Run("before-start-time", func(t *testing.T) {
			require.Equal(0*time.Hour, vote.RemainingTime(startTime.Add(-time.Hour)))
		})
		t.Run("after-end-time", func(t *testing.T) {
			require.Equal(stakingDuration, vote.RemainingTime(startTime.Add(stakingDuration+time.Hour)))
		})
		t.Run("valid-time", func(t *testing.T) {
			require.Equal(stakingDuration, vote.RemainingTime(startTime.Add(stakingDuration-time.Hour)))
		})
	})
}
