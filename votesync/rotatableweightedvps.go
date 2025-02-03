package votesync

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-election/contract"
	"github.com/iotexproject/iotex-election/util"
	"github.com/pkg/errors"
)

type rwvps struct {
	batchSize int
	contract  iotex.Contract
}

func NewRotatableWeightedVPS(cli iotex.AuthedClient, addr address.Address, batchSize uint8) (*rwvps, error) {
	vpsABI, err := abi.JSON(strings.NewReader(contract.RotatableVPSABI))
	if err != nil {
		return nil, err
	}
	return &rwvps{contract: cli.Contract(addr, vpsABI), batchSize: int(batchSize)}, nil
}

func (vps *rwvps) ViewID() (*big.Int, error) {
	d, err := vps.contract.Read("viewID").Call(context.Background())
	if err != nil {
		return nil, err
	}
	ret, err := d.Unmarshal()
	if err != nil {
		return nil, err
	}
	return util.ToBigInt(ret[0])
}

func (vps *rwvps) InactiveViewID() (*big.Int, error) {
	d, err := vps.contract.Read("inactiveViewID").Call(context.Background())
	if err != nil {
		return nil, err
	}
	ret, err := d.Unmarshal()
	if err != nil {
		return nil, err
	}
	return util.ToBigInt(ret[0])
}

func (vps *rwvps) Rotate(viewID *big.Int) error {
	_, err := vps.contract.Execute("rotate", viewID).
		SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(4000000).Call(context.Background())

	return err
}

func (vps *rwvps) UpdateVotingPowers(addrs []common.Address, weights []*big.Int) error {
	if len(addrs) != len(weights) {
		return errors.Errorf("addrs and weights are of different lengths, %d vs %d", len(addrs), len(weights))
	}
	if len(addrs) == 0 {
		return vps.updateVotingPowers(addrs, weights)
	}
	paginationSize := vps.batchSize / 5
	as := []common.Address{}
	ws := []*big.Int{}
	for i := range addrs {
		as = append(as, addrs[i])
		ws = append(ws, weights[i])
		if i%paginationSize == 0 {
			if err := vps.updateVotingPowers(as, ws); err != nil {
				return err
			}
			as = []common.Address{}
			ws = []*big.Int{}
		}
	}
	if len(as) > 0 {
		return vps.updateVotingPowers(as, ws)
	}
	return nil
}

func (vps *rwvps) updateVotingPowers(addrs []common.Address, weights []*big.Int) error {
	_, err := vps.contract.Execute("updateVotingPowers", addrs, weights).
		SetGasPrice(big.NewInt(int64(1 * unit.Qev))).SetGasLimit(7000000).Call(context.Background())
	return err
}

func (vps *rwvps) TotalPower() (*big.Int, error) {
	tpResult, err := vps.contract.Read("totalPower").Call(context.Background())
	if err != nil {
		return nil, err
	}
	ret, err := tpResult.Unmarshal()
	if err != nil {
		return nil, err
	}

	return util.ToBigInt(ret[0])
}

func (vps *rwvps) VoterPowers() (map[common.Address]*big.Int, error) {
	offset := 0
	votingPowers := make(map[common.Address]*big.Int)
	for {
		d, err := vps.contract.Read("voters", offset, vps.batchSize).Call(context.Background())
		if err != nil {
			return nil, err
		}
		ret, err := d.Unmarshal()
		if err != nil {
			return nil, err
		}
		voters := ret[0].([]common.Address)
		d, err = vps.contract.Read("powersOf", voters).Call(context.Background())
		if err != nil {
			return nil, err
		}
		ret, err = d.Unmarshal()
		if err != nil {
			return nil, err
		}
		zero := big.NewInt(0)
		powers := ret[0].([]*big.Int)
		for i := range powers {
			if powers[i].Cmp(zero) != 0 {
				votingPowers[voters[i]] = powers[i]
			}
		}
		if len(voters) < vps.batchSize {
			break
		}
		offset += vps.batchSize
	}

	return votingPowers, nil
}
