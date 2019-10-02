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
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testItems = []item{
		{"a", big.NewInt(10), 3},
		{"b", big.NewInt(20), 2},
		{"c", big.NewInt(30), 1},
		{"d", big.NewInt(20), 1},
		{"e", big.NewInt(30), 1},
	}
	sortedItems = []item{
		{"e", big.NewInt(30), 1},
		{"c", big.NewInt(30), 1},
		{"b", big.NewInt(20), 2},
		{"d", big.NewInt(20), 1},
		{"a", big.NewInt(10), 3},
	}
)

func TestListItem(t *testing.T) {
	require := require.New(t)
	il := make(itemList, len(testItems))
	for i, item := range testItems {
		il[i] = item
	}
	sort.Stable(il)
	for i, item := range sortedItems {
		require.Equal(il[i].Key, item.Key)
		require.Equal(il[i].Value, item.Value)
		require.Equal(il[i].Priority, item.Priority)
	}
}
