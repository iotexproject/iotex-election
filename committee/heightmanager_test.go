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
)

const baseTime = time.Unix(int64(1546272000), 0) // 2019-01-01 00:00:00

var hm *heightManager

func TestNewHeightManager(t *testing.T) {
	hm = newHeightManager()
	if hm == nil {
		t.Error("newHeightManager() failed")
	}
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

	for i := 0; i < 5; i++ {
		ts := baseTime.Add(time.Second * time.Duration(15*i)) // increments in 15 seconds
		name := strconv.Itoa(i) + "-" + strconv.FormatInt(ts.Unix(), 10)
		te := test{name, hm, args{uint64(i), ts}, false}
		tests = append(tests, te)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.add(tt.args.height, tt.args.ts); (err != nil) != tt.wantErr {
				t.Errorf("heightManager.add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	type args struct {
		height uint64
		ts     time.Time
	}
	tests := []struct {
		name    string
		m       *heightManager
		args    args
		wantErr bool
	}{
		{"0-1546272000", hm, args{0, baseTime}, true},
		{"4-1546272060", hm, args{4, baseTime.Add(time.Second * 60)}, true},
		{"5-1546272060", hm, args{5, baseTime.Add(time.Second * 60)}, true},
		{"5-1546272075", hm, args{5, baseTime.Add(time.Second * 75)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.validate(tt.args.height, tt.args.ts); (err != nil) != tt.wantErr {
				t.Errorf("heightManager.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNearestHeightBefore(t *testing.T) {
	type args struct {
		ts time.Time
	}
	tests := []struct {
		name string
		m    *heightManager
		args args
		want uint64
	}{
		{"1546271999", hm, args{baseTime.Add(time.Second * -1)}, 0},
		{"1546272001", hm, args{baseTime.Add(time.Second * 1)}, 0},
		{"1546272035", hm, args{baseTime.Add(time.Second * 35)}, 2},
		{"1546272065", hm, args{baseTime.Add(time.Second * 65)}, 4},
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
	tests := []struct {
		name string
		m    *heightManager
		want uint64
	}{
		{"LastHeight", hm, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.lastestHeight(); got != tt.want {
				t.Errorf("heightManager.lastestHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
