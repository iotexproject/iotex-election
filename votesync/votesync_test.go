package votesync

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var cfg = Config{
	OperatorPrivateKey:       "a000000000000000000000000000000000000000000000000000000000000000",
	IoTeXAPI:                 "api.testnet.iotex.one:80",
	VitaContractAddress:      "io1l9eflyzsmt9pyaud05wk8rajfanxp24xr5vm8d",
	GravityChainAPIs:         []string{"https://mainnet.infura.io/v3/e1f5217dc75d4b77bfede00ca895635b"},
	RegisterContractAddress:  "0x92adef0e5e0c2b4f64a1ac79823f7ad3bc1662c4",
	StakingContractAddress:   "0x3bbe2346c40d34fc3f66ab059f75a6caece2c3b3",
	PaginationSize:           100,
	BrokerPaginationSize:     20,
	FairBankHeight:           3252241,
	GravityChainTimeInterval: 24 * time.Hour,
}

func TestFetchVotesByHeight(t *testing.T) {
	require := require.New(t)
	vs, err := NewVoteSync(cfg)
	require.NoError(err)
	re1, re2, err := vs.retryFetchBucketsByHeight(context.Background(), 3459523)
	require.NoError(err)
	// TODO: this is due to incomplete staking index db, fix later
	require.Zero(len(re1.GetBuckets()))
	require.Zero(len(re2.GetCandidates()))
}

func TestFetchVoteUpdate(t *testing.T) {
	require := require.New(t)
	vs, err := NewVoteSync(cfg)
	require.NoError(err)
	re, err := vs.fetchVotesUpdate(context.Background(), 3252250, 3459523)
	require.NoError(err)
	// TODO: this is due to incomplete staking index db, fix later
	require.Zero(len(re))
}
