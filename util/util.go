// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package util

import (
	"encoding/binary"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/minio/blake2b-simd"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	maxUint = ^uint(0)
	maxInt  = int64(maxUint >> 1)
)

var ErrWrongType = errors.New("wrong data type held by interface{}")

//Uint64ToInt64 converts uint64 to int64
func Uint64ToInt64(u uint64) int64 {
	if u > uint64(maxInt) {
		zap.L().Panic("Height can't be converted to int64")
	}
	return int64(u)
}

//TimeToBytes converts time to []byte
func TimeToBytes(t time.Time) ([]byte, error) {
	return t.MarshalBinary()
}

//BytesToTime converts []byte to time
func BytesToTime(b []byte) (time.Time, error) {
	var t time.Time
	if err := t.UnmarshalBinary(b); err != nil {
		return t, err
	}
	return t, nil
}

//Uint64ToBytes converts uint64 to []byte
func Uint64ToBytes(u uint64) []byte {
	retval := make([]byte, 8)
	binary.LittleEndian.PutUint64(retval, u)

	return retval
}

//BytesToUint64 converts []byte to uint64
func BytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

//CopyBytes copy []byte to another []byte
func CopyBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)

	return c
}

// IsAllZeros return true if all bytes are zero
func IsAllZeros(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}

func ToEtherAddress(v interface{}) (common.Address, error) {
	if addr, ok := v.(common.Address); ok {
		return addr, nil
	}
	return common.Address{}, ErrWrongType
}

func ToBigInt(v interface{}) (*big.Int, error) {
	if bn, ok := v.(*big.Int); ok {
		return bn, nil
	}
	return nil, ErrWrongType
}

type item struct {
	Key      string
	Value    *big.Int
	Priority uint64
}

type itemList []item

func (p itemList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p itemList) Len() int      { return len(p) }
func (p itemList) Less(i, j int) bool {
	switch p[i].Value.Cmp(p[j].Value) {
	case -1:
		return false
	case 1:
		return true
	}
	switch {
	case p[i].Priority < p[j].Priority:
		return false
	case p[i].Priority > p[j].Priority:
		return true
	}
	// This is a corner case, which rarely happens.
	return strings.Compare(p[i].Key, p[j].Key) > 0
}

// Sort sorts candidates, with ts to resolve the equal case
func Sort(candidates map[string]*big.Int, seed uint64) []string {
	p := make(itemList, 0, len(candidates))
	suffix := Uint64ToBytes(seed)
	for name, score := range candidates {
		priority := blake2b.Sum256(append([]byte(name), suffix...))
		p = append(p, item{
			Key:      name,
			Value:    score,
			Priority: BytesToUint64(priority[:8]),
		})
	}
	sort.Stable(p)
	qualifiers := make([]string, len(p))
	for i, item := range p {
		qualifiers[i] = item.Key
	}
	return qualifiers
}
