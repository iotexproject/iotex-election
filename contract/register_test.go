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

func TestRegisterContract(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial("https://kovan.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b")
	require.NoError(t, err)
	caller, err := NewRegisterCaller(
		common.HexToAddress("0xb4ca6cf2fe760517a3f92120acbe577311252663"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetAllCandidates(&bind.CallOpts{BlockNumber: big.NewInt(10453950)}, big.NewInt(1), big.NewInt(12))
	require.NoError(t, err)
	require.Equal(t, 9, len(retval.Names))
	name, err := hex.DecodeString("323233343536373839306131")
	require.NoError(t, err)
	require.True(t, bytes.Equal(name, retval.Names[0][:]))
	require.Equal(t, common.HexToAddress("0x10c7F115EB6EFcf55483D63E6FB78Fa39B5f02de"), retval.Addresses[0])
	operator, err := hex.DecodeString("696f317a7033743376677076757a70737573736b70326b383075356a6e6e7434")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorAddr[0][:])
	operator, err = hex.DecodeString("663578637a6d6b61300000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorAddr[1][:])
	reward, err := hex.DecodeString("696f317a7033743376677076757a70737573736b70326b383075356a6e6e7434")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardAddr[0][:])
	reward, err = hex.DecodeString("663578637a6d6b61300000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardAddr[1][:])
	require.Equal(t, 0, big.NewInt(100000).Cmp(retval.Weights[0]))
}
