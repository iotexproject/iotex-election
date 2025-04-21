// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package contract

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestStakingContract(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial("https://kovan.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b")
	require.NoError(t, err)
	caller, err := NewStakingCaller(
		common.HexToAddress("0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetActiveBuckets(
		&bind.CallOpts{BlockNumber: big.NewInt(10454030)},
		big.NewInt(0),
		big.NewInt(10),
	)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(10), retval.Count)

	amount, ok := new(big.Int).SetString("500000000000000000000", 10)
	require.True(t, ok)
	require.Equal(t, false, retval.Decays[0])
	require.Equal(t, 0, amount.Cmp(retval.StakedAmounts[0]))
	require.Equal(t, 0, big.NewInt(14).Cmp(retval.StakeDurations[0]))
	require.Equal(t, 0, big.NewInt(1551375520).Cmp(retval.StakeStartTimes[0]))
	canName, err := hex.DecodeString("696f783132336b6b61617363")
	require.NoError(t, err)
	require.True(t, bytes.Equal(canName, retval.CanNames[0][:]))
	require.Equal(t, common.HexToAddress("0x4Cd9dE46FED0c91fecc15d8392468f7EFEE34E25"), retval.Owners[0])
}
