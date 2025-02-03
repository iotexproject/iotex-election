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
)

type clerkContract struct {
	contract iotex.Contract
}

func NewClerkContract(cli iotex.AuthedClient, addr address.Address) (*clerkContract, error) {
	clerkABI, err := abi.JSON(strings.NewReader(contract.ClerkABI))
	if err != nil {
		return nil, err
	}
	return &clerkContract{contract: cli.Contract(addr, clerkABI)}, nil
}

func (cc *clerkContract) Claim() error {
	_, err := cc.contract.Execute("claim").SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(5000000).Call(context.Background())
	return err
}
