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

func TestStakingContract(t *testing.T) {
	client, err := ethclient.Dial("https://kovan.infura.io")
	require.NoError(t, err)
	caller, err := NewStakingCaller(
		common.HexToAddress("0x5573c5c69e6bceac4aad14e2c98fbceee8d8c0b8"),
		client,
	)
	require.NoError(t, err)
	bucket, err := caller.Buckets(&bind.CallOpts{BlockNumber: big.NewInt(10246226)}, big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, true, bucket.NonDecay)
	require.Equal(t, 0, big.NewInt(66).Cmp(bucket.StakedAmount))
	require.Equal(t, 0, big.NewInt(14).Cmp(bucket.StakeDuration))
	require.Equal(t, 0, big.NewInt(1548986412).Cmp(bucket.StakeStartTime))
	require.Equal(t, 0, len(bucket.CanPubKey))
	require.Equal(t, 0, big.NewInt(0).Cmp(bucket.UnstakeStartTime))
	require.Equal(t, common.HexToAddress("0x85F8Ff7151DE8EFf96F8bA4190b1FCE316a241aB"), bucket.BucketOwner)
}
