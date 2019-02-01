// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package vote

import (
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVoteCarrier(t *testing.T) {
	require := require.New(t)
	// TODO: update contract address once finalize it
	carrier := NewEthereumVoteCarrier(
		"https://kovan.infura.io",
		common.HexToAddress("0x5573c5c69e6bceac4aad14e2c98fbceee8d8c0b8"),
	)

	lastIndex, votes, err := carrier.Votes(uint64(10246226), big.NewInt(-1), uint8(3))
	require.NoError(err)
	require.Equal(0, big.NewInt(1).Cmp(lastIndex))
	require.Equal(1, len(votes))
	require.Equal(int64(1548986412), votes[0].startTime.Unix())
	require.Equal(336*time.Hour, votes[0].duration)
	require.Equal(0, big.NewInt(66).Cmp(votes[0].amount))
	require.Equal(false, votes[0].decay)
	require.Equal(common.HexToAddress("0x85F8Ff7151DE8EFf96F8bA4190b1FCE316a241aB"), votes[0].voter)
	require.Equal("", votes[0].candidate)
}
