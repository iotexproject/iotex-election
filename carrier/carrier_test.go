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
		common.HexToAddress("0x1c0e85f8804f1448073aac395453c9e52835d968"),
		common.HexToAddress("0xd546bb3fc2db18618b7c16f155800c9897688f3e"),
	)
	require.NoError(err)
	defer carrier.Close()
	t.Run("Candidates", func(t *testing.T) {
		nextIndex, candidates, err := carrier.Candidates(uint64(10439100), big.NewInt(1), uint8(26))
		require.Equal(0, big.NewInt(25).Cmp(nextIndex))
		require.NoError(err)
		require.Equal(24, len(candidates))
		names := []string{
			"726f626f7462703030303030", "726f626f7462703030303032", "726f626f7462703030303036",
			"726f626f7462703030303031", "726f626f7462703030303037", "726f626f7462703030303035",
			"726f626f7462703030303038", "726f626f7462703030303034", "726f626f7462703030303033",
			"726f626f7462703030303039", "726f626f7462703030303131", "726f626f7462703030303130",
			"726f626f7462703030303135", "726f626f7462703030303133", "726f626f7462703030303137",
			"726f626f7462703030303136", "726f626f7462703030303132", "726f626f7462703030303134",
			"726f626f7462703030303230", "726f626f7462703030303231", "726f626f7462703030303138",
			"726f626f7462703030303232", "726f626f7462703030303233", "726f626f7462703030303139",
		}
		for i, name := range names {
			bname, err := hex.DecodeString(name)
			require.NoError(err)
			require.True(bytes.Equal(bname, candidates[i].Name()))
		}
	})
	t.Run("Votes", func(t *testing.T) {
		lastIndex, votes, err := carrier.Votes(uint64(10439100), big.NewInt(0), uint8(3))
		require.NoError(err)
		require.Equal(0, big.NewInt(3).Cmp(lastIndex))
		require.Equal(3, len(votes))
		require.Equal(int64(1551138764), votes[0].StartTime().Unix())
		require.Equal(24*7*time.Hour, votes[0].Duration())
		amount, ok := new(big.Int).SetString("1200000000000000000000000", 10)
		require.True(ok)
		require.Equal(0, amount.Cmp(votes[0].Amount()))
		require.Equal(true, votes[0].Decay())
		require.Equal(0, strings.Compare("10c7f115eb6efcf55483d63e6fb78fa39b5f02de", hex.EncodeToString(votes[0].Voter())))
		canName, err := hex.DecodeString("726f626f7462703030303030")
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
