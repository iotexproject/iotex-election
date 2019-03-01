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
