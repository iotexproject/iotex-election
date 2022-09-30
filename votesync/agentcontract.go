package votesync

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/util"
)

type agentContract struct {
	contract iotex.Contract
	addr     common.Address
}

func NewAgentContract(cli iotex.AuthedClient, addr address.Address) (*agentContract, error) {
	agentABI, err := abi.JSON(strings.NewReader(contract.AgentABI))
	if err != nil {
		return nil, err
	}
	return &agentContract{contract: cli.Contract(addr, agentABI), addr: common.BytesToAddress(addr.Bytes())}, nil
}

func (ac *agentContract) Address() common.Address {
	return ac.addr
}

func (ac *agentContract) Claimed(acct address.Address) (*big.Int, *big.Int, bool, error) {
	data, err := ac.contract.Read("claimed", acct).Call(context.Background())
	if err != nil {
		return nil, nil, false, err
	}
	ret, err := data.Unmarshal()
	if err != nil {
		return nil, nil, false, err
	}
	cycle, err := util.ToBigInt(ret[0])
	if err != nil {
		return nil, nil, false, err
	}
	size, err := util.ToBigInt(ret[1])
	if err != nil {
		return nil, nil, false, err
	}
	return cycle, size, ret[2].(bool), nil
}

func (ac *agentContract) Digest(power *big.Int, cycle *big.Int) ([]byte, error) {
	data, err := ac.contract.Read("digest", power, cycle).Call(context.Background())
	if err != nil {
		return nil, err
	}
	ret, err := data.Unmarshal()
	if err != nil {
		return nil, err
	}
	return ret[0].([]byte), nil
}
