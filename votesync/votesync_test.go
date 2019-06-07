package votesync

import (
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/types"
	"github.com/stretchr/testify/require"
)

var cfg = Config{
	OperatorPrivateKey:       "a000000000000000000000000000000000000000000000000000000000000000",
	IoTeXAPI:                 "api.testnet.iotex.one:443",
	VitaContractAddress:      "io1mxlycqanea52zskuquvun83pvvp28jn5kcz8vn", // TODO: this is vps contract, not vita contract
	GravityChainAPIs:         []string{"https://mainnet.infura.io/v3/b355cae6fafc4302b106b937ee6c15af"},
	RegisterContractAddress:  "0x95724986563028deb58f15c5fac19fa09304f32d",
	StakingContractAddress:   "0x87c9dbff0016af23f5b1ab9b8e072124ab729193",
	PaginationSize:           100,
	GravityChainTimeInterval: 24 * time.Hour,
}

func TestFetchVotesByHeight(t *testing.T) {
	require := require.New(t)
	vs, err := NewVoteSync(cfg)
	require.NoError(err)
	re, err := vs.fetchVotesByHeight(7858000)
	require.NoError(err)
	require.NotZero(len(re))
}

func TestFetchVoteUpdate(t *testing.T) {
	require := require.New(t)
	vs, err := NewVoteSync(cfg)
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

func (*mockCarrier) SubscribeNewBlock(chan *carrier.TipInfo, chan error, chan bool) {}

func (*mockCarrier) Tip() (*carrier.TipInfo, error) { return &carrier.TipInfo{}, nil }

func (*mockCarrier) Candidates(uint64, *big.Int, uint8) (*big.Int, []*types.Candidate, error) {
	return nil, nil, nil
}

func (*mockCarrier) Votes(h uint64, pidx *big.Int, count uint8) (*big.Int, []*types.Vote, error) {
	if pidx.Cmp(big.NewInt(1)) > 0 {
		return nil, nil, nil
	}

	startTime := time.Unix(1559220700, 0)
	v1, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	v2, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte("OldVoter"),
		[]byte("Candidate"),
		true,
	)
	v3, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte("OldVoter2"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}

	v4, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	v5, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(4),
		big.NewInt(4),
		[]byte("Voter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	v6, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
		big.NewInt(3),
		[]byte("NewVoter"),
		[]byte("Candidate"),
		true,
	)
	if err != nil {
		return nil, nil, err
	}
	v7, err := types.NewVote(
		startTime,
		24*7*time.Hour,
		big.NewInt(3),
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
		return nidx, []*types.Vote{v1, v2, v3}, nil
	}
	if h == 2 {
		return nidx, []*types.Vote{v4, v5, v6, v7}, nil
	}
	return nil, nil, nil
}

func (*mockCarrier) Close() {}
