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

func TestRegisterContract(t *testing.T) {
	client, err := ethclient.Dial("https://kovan.infura.io")
	require.NoError(t, err)
	caller, err := NewRegisterCaller(
		common.HexToAddress("57e2b1b258ab80e0b01136409d5b60e7d2debb49"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetAllCandidates(&bind.CallOpts{BlockNumber: big.NewInt(10377416)}, big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, 1, len(retval.Names))
	name, err := hex.DecodeString("726f626f7432000000000000")
	require.NoError(t, err)
	require.True(t, bytes.Equal(name, retval.Names[0][:]))
	require.Equal(t, common.HexToAddress("1D23bF4b8c64e4cdaC4448A5aF777FebF9fedE90"), retval.Addresses[0])
	operator, err := hex.DecodeString("0458aff95fc6dd60a16261398937fada39f4f314085c447c2f54fe03b7df68e5")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorPubKeys[0][:])
	operator, err = hex.DecodeString("cce29fee97708250a0aeed5285466c0d5e8e5520933466f8ca4dd7511ab53187")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorPubKeys[1][:])
	operator, err = hex.DecodeString("c100000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorPubKeys[2][:])
	reward, err := hex.DecodeString("0458aff95fc6dd60a16261398937fada39f4f314085c447c2f54fe03b7df68e5")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardPubKeys[0][:])
	reward, err = hex.DecodeString("cce29fee97708250a0aeed5285466c0d5e8e5520933466f8ca4dd7511ab53187")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardPubKeys[1][:])
	reward, err = hex.DecodeString("c100000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardPubKeys[2][:])
}
