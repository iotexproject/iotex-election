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
		common.HexToAddress("0x67A99F0D2cBde47E1A59F1FabF334db207A48C48"),
		client,
	)
	require.NoError(t, err)
	retval, err := caller.GetAllCandidates(&bind.CallOpts{BlockNumber: big.NewInt(10412500)}, big.NewInt(1), big.NewInt(12))
	require.NoError(t, err)
	require.Equal(t, 10, len(retval.Names))
	name, err := hex.DecodeString("616263000000000000000000")
	require.NoError(t, err)
	require.True(t, bytes.Equal(name, retval.Names[0][:]))
	require.Equal(t, common.HexToAddress("0x4Cd9dE46FED0c91fecc15d8392468f7EFEE34E25"), retval.Addresses[0])
	operator, err := hex.DecodeString("696f3132367863726a6874703237656e6437366163396e6d7836707832303732")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorAddr[0][:])
	operator, err = hex.DecodeString("633376677a367375770000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, operator, retval.IoOperatorAddr[1][:])
	reward, err := hex.DecodeString("696f31333463736b786576797839767264706636376879307032757967397566")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardAddr[0][:])
	reward, err = hex.DecodeString("3268746b3975637a610000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	require.Equal(t, reward, retval.IoRewardAddr[1][:])
}
