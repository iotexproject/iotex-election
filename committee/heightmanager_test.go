// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package committee

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewHeightManager(t *testing.T) {
	hm := newHeightManager()
	require.NotNil(t, hm)
}

// 0-1546272000
// 1-1546272015
// 2-1546272030
// 3-1546272045
// 4-1546272060
func TestAdd(t *testing.T) {
	type args struct {
		height uint64
		ts     time.Time
	}
	type test struct {
		name    string
		m       *heightManager
		args    args
		wantErr bool
	}

	tests := []test{}
	baseTime := time.Unix(int64(1546272000), 0) // 2019-01-01 00:00:00
	hm := newHeightManager()

	for i := 0; i < 5; i++ {
		ts := baseTime.Add(time.Second * time.Duration(15*i)) // increments in 15 seconds
		name := strconv.Itoa(i) + "-" + strconv.FormatInt(ts.Unix(), 10)
		te := test{name, hm, args{uint64(i), ts}, false}
		tests = append(tests, te)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.m.add(tt.args.height, tt.args.ts)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	type args struct {
		height uint64
		ts     time.Time
	}

	baseTime := time.Unix(int64(1546272000), 0) // 2019-01-01 00:00:00
	hm := newHeightManager()
	for i := 0; i < 5; i++ {
		ts := baseTime.Add(time.Second * time.Duration(15*i)) // increments in 15 seconds
		hm.add(uint64(i), ts)
	}

	tests := []struct {
		name    string
		m       *heightManager
		args    args
		wantErr bool
	}{
		{"0-1546272000", hm, args{0, baseTime}, true},                        // h < latest, t < latest
		{"0-1546272060", hm, args{0, baseTime.Add(time.Second * 60)}, true},  // h < latest, t = latest
		{"0-1546272061", hm, args{0, baseTime.Add(time.Second * 61)}, true},  // h < latest, t > latest
		{"4-1546272000", hm, args{4, baseTime}, true},                        // h = latest, t < latest
		{"4-1546272060", hm, args{4, baseTime.Add(time.Second * 60)}, true},  // h = latest, t = latest
		{"4-1546272061", hm, args{4, baseTime.Add(time.Second * 61)}, true},  // h = latest, t > latest
		{"5-1546272000", hm, args{5, baseTime}, true},                        // h > latest, t < latest
		{"5-1546272060", hm, args{5, baseTime.Add(time.Second * 60)}, true},  // h > latest, t = latest
		{"5-1546272061", hm, args{5, baseTime.Add(time.Second * 61)}, false}, // h > latest, t > latest

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.m.validate(tt.args.height, tt.args.ts)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNearestHeightBefore(t *testing.T) {
	type args struct {
		ts time.Time
	}

	baseTime := time.Unix(int64(1546272000), 0) // 2019-01-01 00:00:00
	hm := newHeightManager()
	for i := 0; i < 5; i++ {
		ts := baseTime.Add(time.Second * time.Duration(15*i)) // increments in 15 seconds
		hm.add(uint64(i), ts)
	}

	tests := []struct {
		name string
		m    *heightManager
		args args
		want uint64
	}{
		{"1546271970", hm, args{baseTime.Add(time.Second * -30)}, 0},
		{"1546271999", hm, args{baseTime.Add(time.Second * -1)}, 0},
		{"1546272000", hm, args{baseTime}, 0},
		{"1546272001", hm, args{baseTime.Add(time.Second * 1)}, 0},
		{"1546272014", hm, args{baseTime.Add(time.Second * 14)}, 0},
		{"1546272015", hm, args{baseTime.Add(time.Second * 15)}, 1},
		{"1546272016", hm, args{baseTime.Add(time.Second * 16)}, 1},
		{"1546272029", hm, args{baseTime.Add(time.Second * 29)}, 1},
		{"1546272030", hm, args{baseTime.Add(time.Second * 30)}, 2},
		{"1546272031", hm, args{baseTime.Add(time.Second * 31)}, 2},
		{"1546272059", hm, args{baseTime.Add(time.Second * 59)}, 3},
		{"1546272060", hm, args{baseTime.Add(time.Second * 60)}, 4},
		{"1546272061", hm, args{baseTime.Add(time.Second * 61)}, 4},
		{"1546272090", hm, args{baseTime.Add(time.Second * 90)}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.nearestHeightBefore(tt.args.ts); got != tt.want {
				t.Errorf("heightManager.nearestHeightBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastestHeight(t *testing.T) {
	baseTime := time.Unix(int64(1546272000), 0) // 2019-01-01 00:00:00
	hm := newHeightManager()
	require.Equal(t, uint64(0), hm.lastestHeight())
	for i := 0; i < 5; i++ {
		ts := baseTime.Add(time.Second * time.Duration(15*i)) // increments in 15 seconds
		hm.add(uint64(i), ts)
		require.Equal(t, uint64(i), hm.lastestHeight())
	}
}
