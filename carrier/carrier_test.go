// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package carrier

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/iotexproject/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVoteCarrier(t *testing.T) {
	require := require.New(t)
	// TODO: update contract address once finalize it
	carrier, err := NewEthereumVoteCarrier(
		"wss://kovan.infura.io/ws",
		common.HexToAddress("0x5573c5c69e6bceac4aad14e2c98fbceee8d8c0b8"),
	)
	require.NoError(err)
	t.Run("Votes", func(t *testing.T) {
		lastIndex, votes, err := carrier.Votes(uint64(10246226), big.NewInt(-1), uint8(3))
		require.NoError(err)
		require.Equal(0, big.NewInt(1).Cmp(lastIndex))
		require.Equal(1, len(votes))
		require.Equal(int64(1548986412), votes[0].StartTime().Unix())
		require.Equal(336*time.Hour, votes[0].Duration())
		require.Equal(0, big.NewInt(66).Cmp(votes[0].Amount()))
		require.Equal(false, votes[0].Decay())
		require.Equal(0, strings.Compare("85f8ff7151de8eff96f8ba4190b1fce316a241ab", hex.EncodeToString(votes[0].Voter())))
		require.True(bytes.Equal([]byte(""), votes[0].Candidate()))

	})
	t.Run("BlockTimestamp", func(t *testing.T) {
		ts, err := carrier.BlockTimestamp(uint64(10246228))
		require.NoError(err)
		require.Equal(int64(1548986420), ts.Unix())
	})
	t.Run("SubscribeNewBlock", func(t *testing.T) {
		close := make(chan bool)
		heightChan := make(chan uint64)
		go func() {
			time.Sleep(30 * time.Second)
			close <- true
		}()
		for {
			err := carrier.SubscribeNewBlock(func(height uint64) {
				heightChan <- height
			}, close)
			if err == nil {
				break
			}
		}
		lastHeight := uint64(0)
		var i int
		for i = 0; i < 5; i++ {
			select {
			case height := <-heightChan:
				if i != 0 {
					require.Equal(lastHeight+1, height)
				}
				lastHeight = height
			case <-close:
				close <- true
				break
			}
		}
		require.Equal(5, i)
	})
}
