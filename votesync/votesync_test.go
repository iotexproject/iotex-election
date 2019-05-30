package votesync

import (
	"testing"
)

func TestVoteSync(t *testing.T) {
	cfg := Config{
		GravityChainAPIs:           []string{"https://mainnet.infura.io/v3/b355cae6fafc4302b106b937ee6c15af"},
		RegisterContractAddress:    "0x95724986563028deb58f15c5fac19fa09304f32d",
		StakingContractAddress:     "0x87c9dbff0016af23f5b1ab9b8e072124ab729193",
		PaginationSize:             100,
		GravityChainHeightInterval: 8640,
	}
	vs, err := NewVoteSync(cfg)
	if err != nil {
		t.Fatal(err)
	}
	re, err := vs.fetchVotesByHeight(7858000)
	if err != nil {
		t.Fatal(err)
	}
	if len(re) == 0 {
		t.Fatal("fail to fetch votes")
	}
}
