// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

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
	client, err := ethclient.Dial("https://kovan.infura.io")
	require.NoError(t, err)
	caller, err := NewStakingCaller(
		common.HexToAddress("0xf488342896e4ef30022a88d869caaa329d476aa9"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetActiveBuckets(
		&bind.CallOpts{BlockNumber: big.NewInt(10377538)},
		big.NewInt(0),
		big.NewInt(10),
	)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(2), retval.Count)

	amount, ok := new(big.Int).SetString("250000000000000000000", 10)
	require.True(t, ok)
	require.Equal(t, true, retval.Decays[0])
	require.Equal(t, 0, amount.Cmp(retval.StakedAmounts[0]))
	require.Equal(t, 0, big.NewInt(7).Cmp(retval.StakeDurations[0]))
	require.Equal(t, 0, big.NewInt(1550363360).Cmp(retval.StakeStartTimes[0]))
	canName, err := hex.DecodeString("726f626f7432000000000000")
	require.NoError(t, err)
	require.True(t, bytes.Equal(canName, retval.CanNames[0][:]))
	require.Equal(t, common.HexToAddress("0x95a971937F343591352c56EABf04a1D69DE18c4E"), retval.Owners[0])

	amount, ok = new(big.Int).SetString("100000000000000000000", 10)
	require.True(t, ok)
	require.Equal(t, false, retval.Decays[1])
	require.Equal(t, 0, amount.Cmp(retval.StakedAmounts[1]))
	require.Equal(t, 0, big.NewInt(350).Cmp(retval.StakeDurations[1]))
	require.Equal(t, 0, big.NewInt(1550363864).Cmp(retval.StakeStartTimes[1]))
	canName, err = hex.DecodeString("726f626f7431000000000000")
	require.NoError(t, err)
	require.True(t, bytes.Equal(canName, retval.CanNames[1][:]))
	require.Equal(t, common.HexToAddress("1D23bF4b8c64e4cdaC4448A5aF777FebF9fedE90"), retval.Owners[1])
}
