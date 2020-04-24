package votesync

import (
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/iotex-election/types"
	"github.com/stretchr/testify/require"
)

var cfg = Config{
	OperatorPrivateKey:       "a000000000000000000000000000000000000000000000000000000000000000",
	IoTeXAPI:                 "api.testnet.iotex.one:443",
	VitaContractAddress:      "io1l9eflyzsmt9pyaud05wk8rajfanxp24xr5vm8d",
	GravityChainAPIs:         []string{"https://mainnet.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b"},
	RegisterContractAddress:  "0x92adef0e5e0c2b4f64a1ac79823f7ad3bc1662c4",
	StakingContractAddress:   "0x3bbe2346c40d34fc3f66ab059f75a6caece2c3b3",
	PaginationSize:           100,
	BrokerPaginationSize:     20,
	GravityChainTimeInterval: 24 * time.Hour,
}

func TestFetchVotesByHeight(t *testing.T) {
	require := require.New(t)
	vs, err := NewVoteSync(cfg, nil)
	require.NoError(err)
	re, err := vs.fetchBucketsByHeight(7858000)
	require.NoError(err)
	require.NotZero(len(re))
}

func TestFetchVoteUpdate(t *testing.T) {
	t.Skip() // native contract seems stuck on test. won't fix. will move to nsV2
	require := require.New(t)
	vs, err := NewVoteSync(cfg, nil)
	require.NoError(err)
	vs.carrier = &mockCarrier{}
	ts, err := vs.carrier.BlockTimestamp(2)
	require.NoError(err)
	re, err := vs.fetchVotesUpdate(1, 2, ts, ts)
	require.NoError(err)
	require.Equal(3, len(re))
	for _, r := range re {
		switch string(r.Voter) {
		case "Voter":
			require.Zero(r.Votes.Cmp(big.NewInt(7)))
		case "NewVoter":
			require.Zero(r.Votes.Cmp(big.NewInt(3)))
		case "OldVoter":
			require.Zero(r.Votes.Cmp(big.NewInt(0)))
		}
	}
}

type mockCarrier struct{}

func (*mockCarrier) BlockTimestamp(uint64) (time.Time, error) {
	return time.Unix(1559240700, 0), nil
}

func (*mockCarrier) SubscribeNewBlock(chan uint64, chan error, chan bool) {}

func (*mockCarrier) HasStakingEvents(*big.Int, *big.Int) bool {
	return true
}

func (*mockCarrier) Tip() (uint64, error) { return 0, nil }

func (*mockCarrier) Registrations(uint64, *big.Int, uint8) (*big.Int, []*types.Registration, error) {
	return nil, nil, nil
}

func (*mockCarrier) Buckets(h uint64, pidx *big.Int, count uint8) (*big.Int, []*types.Bucket, error) {
	if pidx.Cmp(big.NewInt(1)) > 0 {
		return nil, nil, nil
	}

	startTime := time.Unix(1559220700, 0)
	b1, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b2, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("OldVoter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b3, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("OldVoter2"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b4, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b5, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(4),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b6, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("NewVoter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	b7, err := types.NewBucket(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		[]byte("OldVoter2"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}

	nidx := pidx.Add(pidx, big.NewInt(1))
	if h == 1 {
		return nidx, []*types.Bucket{b1, b2, b3}, nil
	}
	if h == 2 {
		return nidx, []*types.Bucket{b4, b5, b6, b7}, nil
	}
	return nil, nil, nil
}

func (*mockCarrier) Close() {}
