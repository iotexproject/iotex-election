// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

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
var(
	args = []arg{
		{
			0, time.Unix(int64(1546272000), 0),
		},
		{
			1, time.Unix(int64(1546272015), 0),
		},
		{
			2, time.Unix(int64(1546272030), 0),
		},
		{
			3, time.Unix(int64(1546272045), 0),
		},
		{
			4, time.Unix(int64(1546272060), 0),
		},
	}
	validArgs=[]arg{
		{
			30, time.Unix(int64(1546272061), 0),
		},
		{
			20, time.Unix(int64(1546272065), 0),
		},
	}
	invalidArgs=[]arg{
		{
			// check for unique
			2, time.Unix(int64(1546272030), 0),
		},
		{
			3, time.Unix(int64(1546272061), 0),
		},
		{
			4, time.Unix(int64(1546272030), 0),
		},
		{
			20, time.Unix(int64(1546272030), 0),
		},
	}
	beforeTime=[]time.Time{
		time.Unix(int64(1546272001), 0),
		time.Unix(int64(1546272016), 0),
		time.Unix(int64(1546272031), 0),
		time.Unix(int64(1546272046), 0),
		time.Unix(int64(1546272061), 0),
	}
)
func TestNewHeightManager(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	require.NotNil(hm)
}

func TestAdd(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	var err error
	for _,arg:=range args{
		err=hm.add(arg.height,arg.time)
		require.NoError(err)
	}
	// test args cannot add
	for _,arg:=range invalidArgs{
		err=hm.add(arg.height,arg.time)
		require.Error(err)
	}
	// check for if p1.height > p2.height, then p1.time > p2.time
	for i:=1;i<5;i++{
		require.True(hm.heights[i]-hm.heights[i-1]>0)
		require.True(hm.times[i].After(hm.times[i-1]))
	}
	// check len(height)==len(time)
	require.Equal(len(hm.heights),len(hm.times))
}

func TestValidate(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	var err error
	for _,arg:=range args{
		err=hm.add(arg.height,arg.time)
		require.NoError(err)
	}
	// height and time both valid
	for _,arg:=range validArgs{
		err=hm.validate(arg.height,arg.time)
		require.NoError(err)
	}
	// 4 different invalid combinations
	for _,arg:=range invalidArgs{
		err=hm.validate(arg.height,arg.time)
		require.Error(err)
	}
}

func TestNearestHeightBefore(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()

	var hei uint64
	// check len(m.heights)==0
	hei=hm.nearestHeightBefore(time.Unix(int64(1546271000), 0))
	require.Equal(uint64(0),hei)

	var err error
	for _,arg:=range args{
		err=hm.add(arg.height,arg.time)
		require.NoError(err)
	}
	// check m.times[0].After(ts)
	ts:=time.Unix(int64(1546271000), 0)
	hei=hm.nearestHeightBefore(ts)
	require.Equal(uint64(0),hei)

	// test every height
	for i,ti:=range beforeTime{
		hei=hm.nearestHeightBefore(ti)
		require.Equal(args[i].height,hei)
	}
}

func TestLastestHeight(t *testing.T) {
	require := require.New(t)
	hm := newHeightManager()
	var hei uint64
	var err error
	for i,arg:=range args{
		err=hm.add(arg.height,arg.time)
		require.NoError(err)
		hei=hm.lastestHeight()
		require.Equal(args[i].height,hei)
	}
}