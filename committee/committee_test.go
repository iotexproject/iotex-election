// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package committee

import (
	"database/sql"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-election/types"
)

func TestReadBlockTime(t *testing.T) {
	req := require.New(t)
	sqlDB, err := sql.Open("sqlite3", "/Users/dustin/iotexproject/iotex-core/data/poll.db")
	req.NoError(err)

	tQ := "SELECT MAX(height) FROM mint_time WHERE ? >= time AND EXISTS (SELECT * FROM mint_time WHERE ? <= time)"
	ts := time.Unix(1555898790, 0).UTC()
	println("time =", ts.String())
	var height uint64
	err = sqlDB.QueryRow(tQ, ts.String(), ts.String()).Scan(&height)
	req.NoError(err)
	println("height =", height)

	hQ := "SELECT time FROM mint_time WHERE height = ?"
	var timestamp string
	err = sqlDB.QueryRow(hQ, 7614600).Scan(&timestamp)
	req.NoError(err)
	println("ts =", timestamp)
	ts, err = time.Parse(time.RFC3339, timestamp)
	req.NoError(err)
	println("parse =", ts.String())
	var ts1 time.Time
	err = sqlDB.QueryRow(hQ, 7614600).Scan(&ts1)
	req.NoError(err)
	println("ts =", ts1.String())
	println("UTC =", ts1.UTC().String())
}

func TestCalcWeightedVotes(t *testing.T) {
	require := require.New(t)
	committee := &committee{}
	startTime := time.Now()
	duration := time.Hour * 24 * 14
	bucket1, err := types.NewBucket(
		startTime,
		duration,
		big.NewInt(3000000),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	// now.Before(v.StartTime()),return 0
	require.Equal(0, committee.calcWeightedVotes(bucket1, startTime.Add(-1*time.Hour)).Cmp(big.NewInt(0)))

	// decay is true,startTime+duration is after now,remainingTime is 24*14-24=13*24 hours,weight is ~1.140,ret is 3422048
	require.Equal(0, committee.calcWeightedVotes(bucket1, startTime.Add(time.Hour*24)).Cmp(big.NewInt(3422048)))

	// decay is true,startTime+duration is before now,remainingTime is 0 hours,weight is 1,ret is 3000000
	require.Equal(0, committee.calcWeightedVotes(bucket1, time.Now().Add(24*15*time.Hour)).Cmp(big.NewInt(3000000)))

	bucket2, err := types.NewBucket(
		startTime,
		duration,
		big.NewInt(3000000),
		[]byte{},
		[]byte{},
		false,
	)
	require.NoError(err)
	// decay is false,remainingTime is duration,weight ~1.144，ret is 3434242,whatever now is
	require.Equal(0, committee.calcWeightedVotes(bucket2, startTime.Add(time.Hour*24)).Cmp(big.NewInt(3434242)))
	require.Equal(0, committee.calcWeightedVotes(bucket2, startTime.Add(24*15*time.Hour)).Cmp(big.NewInt(3434242)))
}

func TestVoteFilter(t *testing.T) {
	require := require.New(t)
	c := &committee{voteThreshold: big.NewInt(10)}
	bucket1, err := types.NewBucket(
		time.Now(),
		time.Hour,
		big.NewInt(3),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.True(c.bucketFilter(bucket1))

	bucket2, err := types.NewBucket(
		time.Now(),
		time.Hour,
		big.NewInt(30),
		[]byte{},
		[]byte{},
		true,
	)
	require.NoError(err)
	require.False(c.bucketFilter(bucket2))
}

func TestCandidateFilter(t *testing.T) {
	require := require.New(t)
	committee := &committee{
		scoreThreshold:       big.NewInt(10),
		selfStakingThreshold: big.NewInt(10),
	}
	// candidate1 selfStaking and score both smaller than committee's threshold
	candidate1 := types.NewCandidate(
		types.NewRegistration(
			[]byte("candidate1"),
			[]byte("candidate1addr"),
			[]byte("operatorPubKey1"),
			[]byte("rewardPubKey1"),
			1,
		),
		big.NewInt(9),
		big.NewInt(9),
	)
	require.True(committee.candidateFilter(candidate1))
	// candidate2 selfStaking is below committee's threshold,score is bigger than committee's threshold
	candidate2 := types.NewCandidate(
		types.NewRegistration(
			[]byte("candidate2"),
			[]byte("candidate2addr"),
			[]byte("operatorPubKey2"),
			[]byte("rewardPubKey2"),
			1,
		),
		big.NewInt(11),
		big.NewInt(9),
	)
	require.True(committee.candidateFilter(candidate2))
	// candidate3 selfStaking is bigger than committee's threshold,score is smaller than committee's threshold
	candidate3 := types.NewCandidate(
		types.NewRegistration(
			[]byte("candidate3"),
			[]byte("candidate3addr"),
			[]byte("operatorPubKey3"),
			[]byte("rewardPubKey3"),
			1,
		),
		big.NewInt(9),
		big.NewInt(11),
	)
	require.True(committee.candidateFilter(candidate3))
	// candidate3 selfStaking and score both bigger than committee's threshold
	candidate4 := types.NewCandidate(
		types.NewRegistration(
			[]byte("candidate4"),
			[]byte("candidate4addr"),
			[]byte("operatorPubKey4"),
			[]byte("rewardPubKey4"),
			1,
		),
		big.NewInt(11),
		big.NewInt(11),
	)
	require.False(committee.candidateFilter(candidate4))
}
