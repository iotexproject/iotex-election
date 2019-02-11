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

	"github.com/iotexproject/go-ethereum/accounts/abi/bind"
	"github.com/iotexproject/go-ethereum/common"
	"github.com/iotexproject/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestRegisterContract(t *testing.T) {
	client, err := ethclient.Dial("https://kovan.infura.io")
	require.NoError(t, err)
	caller, err := NewRegisterCaller(
		common.HexToAddress("0x0f3e55d8a7f45b3e2054b5068e9c5d5a42016766"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetAllCandidates(&bind.CallOpts{BlockNumber: big.NewInt(10335842)}, big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, 1, len(retval.Names))
	name, err := hex.DecodeString("6265737464656c6567617465")
	require.NoError(t, err)
	require.True(t, bytes.Equal(name, retval.Names[0][:]))
	require.Equal(t, common.HexToAddress("85F8Ff7151DE8EFf96F8bA4190b1FCE316a241aB"), retval.Addresses[0])
	operator, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorPubKeys[0][:])
	reward, err := hex.DecodeString("0111111111111111111111111111111111111111111111111111111111111111")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardPubKeys[0][:])
}
