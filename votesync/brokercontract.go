package votesync

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/util"
)

type brokerContract struct {
	contract  iotex.Contract
	batchSize uint64
}

func NewBrokerContract(cli iotex.AuthedClient, addr address.Address, batchSize uint8) (*brokerContract, error) {
	brokerABI, err := abi.JSON(strings.NewReader(contract.BrokerABI))
	if err != nil {
		return nil, err
	}
	return &brokerContract{contract: cli.Contract(addr, brokerABI), batchSize: uint64(batchSize)}, nil
}

func (bc *brokerContract) Reset() error {
	_, err := bc.contract.Execute("reset").SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000).Call(context.Background())
	return err
}

func (bc *brokerContract) NextBidToSettle() (uint64, error) {
	d, err := bc.contract.Read("nextBidToSettle").Call(context.Background())
	if err != nil {
		return 0, err
	}
	ret, err := d.Unmarshal()
	if err != nil {
		return 0, err
	}
	nextBidToSettle, err := util.ToBigInt(ret[0])
	if err != nil {
		return 0, err
	}
	return nextBidToSettle.Uint64(), nil
}

func (bc *brokerContract) Settle() error {
	_, err := bc.contract.Execute(
		"settle", big.NewInt(0).SetUint64(bc.batchSize)).
		SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000).Call(context.Background())
	return err
}
