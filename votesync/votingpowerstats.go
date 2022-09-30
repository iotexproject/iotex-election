package votesync

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type VotingPowers struct {
	lock         sync.RWMutex
	cycle        *big.Int
	totalVotes   *big.Int
	votingPowers map[common.Address]*big.Int
}

func (vp *VotingPowers) Total() *big.Int {
	return vp.totalVotes
}

func (vp *VotingPowers) VotingPower(acct common.Address) (*big.Int, *big.Int) {
	vp.lock.RLock()
	defer vp.lock.RUnlock()

	return vp.cycle, vp.votingPowers[acct]
}

func (vp *VotingPowers) Update(cycle *big.Int, total *big.Int, votingPowers map[common.Address]*big.Int) {
	vp.lock.Lock()
	defer vp.lock.Unlock()

	vp.cycle = cycle
	vp.totalVotes = total
	vp.votingPowers = votingPowers
}
