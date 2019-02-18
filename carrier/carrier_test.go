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

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestVoteCarrier(t *testing.T) {
	require := require.New(t)
	// TODO: update contract address once finalize it
	carrier, err := NewEthereumVoteCarrier(
		"wss://kovan.infura.io/ws",
		common.HexToAddress("0x57e2b1b258ab80e0b01136409d5b60e7d2debb49"),
		common.HexToAddress("0xf488342896e4ef30022a88d869caaa329d476aa9"),
	)
	require.NoError(err)
	t.Run("Candidates", func(t *testing.T) {
		nextIndex, candidates, err := carrier.Candidates(uint64(10377500), big.NewInt(1), uint8(2))
		require.Equal(0, big.NewInt(3).Cmp(nextIndex))
		require.NoError(err)
		require.Equal(2, len(candidates))
		name, err := hex.DecodeString("726f626f7432000000000000")
		require.NoError(err)
		require.True(bytes.Equal(name, candidates[0].Name()))
		name, err = hex.DecodeString("726f626f7431000000000000")
		require.NoError(err)
		require.True(bytes.Equal(name, candidates[1].Name()))
	})
	t.Run("Votes", func(t *testing.T) {
		lastIndex, votes, err := carrier.Votes(uint64(10377500), big.NewInt(0), uint8(3))
		require.NoError(err)
		require.Equal(0, big.NewInt(1).Cmp(lastIndex))
		require.Equal(1, len(votes))
		require.Equal(int64(1550363360), votes[0].StartTime().Unix())
		require.Equal(24*7*time.Hour, votes[0].Duration())
		amount, ok := new(big.Int).SetString("250000000000000000000", 10)
		require.True(ok)
		require.Equal(0, amount.Cmp(votes[0].Amount()))
		require.Equal(true, votes[0].Decay())
		require.Equal(0, strings.Compare("95a971937f343591352c56eabf04a1d69de18c4e", hex.EncodeToString(votes[0].Voter())))
		canName, err := hex.DecodeString("726f626f7432000000000000")
		require.NoError(err)
		require.True(bytes.Equal(canName, votes[0].Candidate()))
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
