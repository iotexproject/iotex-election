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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestIOTXContract(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial("https://kovan.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b")
	require.NoError(t, err)
	contractAddr := common.HexToAddress("51ca23c98b7481951d0904d3f134889713306c75")
	sc, err := NewIOTXCaller(contractAddr, client)
	require.NoError(t, err)
	accountAddr := common.HexToAddress("daa75db17e57ede45cfa589b103ce91566770563")
	transferHeight := int64(10369800)
	b, err := sc.BalanceOf(
		&bind.CallOpts{BlockNumber: big.NewInt(transferHeight - 1)},
		accountAddr,
	)
	require.NoError(t, err)
	require.Equal(t, 0, b.Cmp(big.NewInt(0)))
	b, err = sc.BalanceOf(
		&bind.CallOpts{BlockNumber: big.NewInt(transferHeight)},
		accountAddr,
	)
	require.NoError(t, err)
	balance, ok := new(big.Int).SetString("100000000000000000000000", 10)
	require.True(t, ok)
	require.Equal(t, 0, b.Cmp(balance))
}
