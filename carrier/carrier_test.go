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
		[]string{"wss://kovan.infura.io/ws"},
		common.HexToAddress("0xb4ca6cf2fe760517a3f92120acbe577311252663"),
		common.HexToAddress("0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"),
	)
	require.NoError(err)
	defer carrier.Close()
	t.Run("Candidates", func(t *testing.T) {
		nextIndex, candidates, err := carrier.Candidates(uint64(10454030), big.NewInt(1), uint8(10))
		require.Equal(0, big.NewInt(10).Cmp(nextIndex))
		require.NoError(err)
		require.Equal(9, len(candidates))
		names := []string{
			"323233343536373839306131", "323233343536373839306134", "323233343536373839306139",
			"323233343536373839306133", "323233343536373839306138", "323233343536373839306136",
			"323233343536373839306132", "323233343536373839306137", "323233343536373839306135",
		}
		for i, name := range names {
			bname, err := hex.DecodeString(name)
			require.NoError(err)
			require.True(bytes.Equal(bname, candidates[i].Name()))
		}
	})
	t.Run("Votes", func(t *testing.T) {
		lastIndex, votes, err := carrier.Votes(uint64(10454030), big.NewInt(0), uint8(10))
		require.NoError(err)
		require.Equal(0, big.NewInt(11).Cmp(lastIndex))
		require.Equal(10, len(votes))
		require.Equal(int64(1551375520), votes[0].StartTime().Unix())
		require.Equal(24*14*time.Hour, votes[0].Duration())
		amount, ok := new(big.Int).SetString("500000000000000000000", 10)
		require.True(ok)
		require.Equal(0, amount.Cmp(votes[0].Amount()))
		require.Equal(false, votes[0].Decay())

		require.Equal(0, strings.Compare("4cd9de46fed0c91fecc15d8392468f7efee34e25", hex.EncodeToString(votes[0].Voter())))
		canName, err := hex.DecodeString("696f783132336b6b61617363")
		require.NoError(err)
		require.True(bytes.Equal(canName, votes[0].Candidate()))
	})
	t.Run("BlockTimestamp", func(t *testing.T) {
		ts, err := carrier.BlockTimestamp(uint64(10246228))
		require.NoError(err)
		require.Equal(int64(1548986420), ts.Unix())
	})
	t.Run("SubscribeNewBlock", func(t *testing.T) {
		heightChan := make(chan uint64)
		reportChan := make(chan error)
		unsubscribe := make(chan bool)
		carrier.SubscribeNewBlock(heightChan, reportChan, unsubscribe)
		select {
		case <-heightChan:
			break
		case err := <-reportChan:
			require.NoError(err)
		}
	})
}
