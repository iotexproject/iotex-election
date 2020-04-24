// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

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
		12,
		time.Second,
		[]string{"https://kovan.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b"},
		common.HexToAddress("0xb4ca6cf2fe760517a3f92120acbe577311252663"),
		common.HexToAddress("0xdedf0c1610d8a75ca896d8c93a0dc39abf7daff4"),
	)
	require.NoError(err)
	defer carrier.Close()
	t.Run("Registrations", func(t *testing.T) {
		nextIndex, candidates, err := carrier.Registrations(uint64(10454030), big.NewInt(1), uint8(10))
		require.NoError(err)
		require.Equal(0, big.NewInt(10).Cmp(nextIndex))
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
	t.Run("Buckets", func(t *testing.T) {
		lastIndex, votes, err := carrier.Buckets(uint64(10454030), big.NewInt(0), uint8(10))
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
		tipChan := make(chan uint64)
		reportChan := make(chan error)
		unsubscribe := make(chan bool)
		carrier.SubscribeNewBlock(tipChan, reportChan, unsubscribe)
		select {
		case <-tipChan:
			break
		case err := <-reportChan:
			require.NoError(err)
		}
	})
}
