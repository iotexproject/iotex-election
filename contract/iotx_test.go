// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

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
	client, err := ethclient.Dial("https://kovan.infura.io/")
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
