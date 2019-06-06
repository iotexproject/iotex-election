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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type arg struct {
	height uint64
	time   time.Time
}

var (
	args = []arg{
		{0, time.Unix(int64(1546272000), 0)},
		{1, time.Unix(int64(1546272015), 0)},
		{2, time.Unix(int64(1546272030), 0)},
		{3, time.Unix(int64(1546272045), 0)},
		{4, time.Unix(int64(1546272060), 0)},
	}
	validArgs = []arg{
		{30, time.Unix(int64(1546272061), 0)},
		{20, time.Unix(int64(1546272065), 0)},
	}
	invalidArgs = []arg{
		{
			// check for unique
			2, time.Unix(int64(1546272030), 0),
		},
		{
			// check for height in args,but time is not
			3, time.Unix(int64(1546272061), 0),
		},
		{
			// check for height in args,time also in args
			4, time.Unix(int64(1546272030), 0),
		},
		{
			// check for height not in args,but time in args
			20, time.Unix(int64(1546272030), 0),
		},
	}
	beforeTime = []time.Time{
		// time after first time in args
		time.Unix(int64(1546272001), 0),
		// time after second time in args,following is the same
		time.Unix(int64(1546272016), 0),
		time.Unix(int64(1546272031), 0),
		time.Unix(int64(1546272046), 0),
		time.Unix(int64(1546272061), 0),
	}
)

func TestNewHeightManager(t *testing.T) {
	require := require.New(t)
	require.NotNil(newHeightManager())
}

func TestAdd(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	for _, arg := range args {
		require.NoError(hm.add(arg.height, arg.time))
	}
	// test args cannot add
	for _, arg := range invalidArgs {
		require.Error(hm.add(arg.height, arg.time))
	}
	// check for if p1.height > p2.height, then p1.time > p2.time
	for i := 1; i < 5; i++ {
		require.True(hm.heights[i]-hm.heights[i-1] > 0)
		require.True(hm.times[i].After(hm.times[i-1]))
	}
	// check len(height)==len(time)
	require.Equal(len(hm.heights), len(hm.times))
}

func TestValidate(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	for _, arg := range args {
		require.NoError(hm.add(arg.height, arg.time))
	}
	// height and time both valid
	for _, arg := range validArgs {
		require.NoError(hm.validate(arg.height, arg.time))
	}
	// 4 different invalid combinations
	for _, arg := range invalidArgs {
		require.Error(hm.validate(arg.height, arg.time))
	}
}

func TestNearestHeightBefore(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()

	// check len(m.heights)==0
	require.Equal(uint64(0), hm.nearestHeightBefore(time.Unix(int64(1546271000), 0)))
	for _, arg := range args {
		require.NoError(hm.add(arg.height, arg.time))
	}
	// check m.times[0].After(ts)
	require.Equal(uint64(0), hm.nearestHeightBefore(time.Unix(int64(1546271000), 0)))

	// test every height
	for i, ti := range beforeTime {
		require.Equal(args[i].height, hm.nearestHeightBefore(ti))
	}
}

func TestLatestHeight(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	for i, arg := range args {
		// add and then check lastest height
		require.NoError(hm.add(arg.height, arg.time))
		require.Equal(args[i].height, hm.latestHeight())
	}
}
