// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"math/big"
	"testing"

	"github.com/iotexproject/go-ethereum/accounts/abi/bind"
	"github.com/iotexproject/go-ethereum/common"
	"github.com/iotexproject/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestIOTXContract(t *testing.T) {
	client, err := ethclient.Dial("https://kovan.infura.io/")
	require.NoError(t, err)
	sc, err := NewIOTXCaller(common.HexToAddress("a45e1d096a5d1b7db32e9309f67b293d3d8de759"), client)
	require.NoError(t, err)
	b, err := sc.BalanceOf(
		&bind.CallOpts{BlockNumber: big.NewInt(10245807)},
		common.HexToAddress("731eae7bEdec1F0A5A52BEb39a4e1dCdb4bb77Ac"),
	)
	require.NoError(t, err)
	require.Equal(t, 0, b.Cmp(big.NewInt(0)))
	b, err = sc.BalanceOf(
		&bind.CallOpts{BlockNumber: big.NewInt(10245808)},
		common.HexToAddress("731eae7bEdec1F0A5A52BEb39a4e1dCdb4bb77Ac"),
	)
	require.NoError(t, err)
	require.Equal(t, 0, b.Cmp(big.NewInt(123456789)))
}
