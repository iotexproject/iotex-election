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

type vitaContract struct {
	contract iotex.Contract
}

func NewVitaContract(cli iotex.AuthedClient, addr address.Address) (*vitaContract, error) {
	vitaABI, err := abi.JSON(strings.NewReader(contract.VitaABI))
	if err != nil {
		return nil, err
	}
	return &vitaContract{contract: cli.Contract(addr, vitaABI)}, nil
}

func (vc *vitaContract) readAddress(fn string) (address.Address, error) {
	d, err := vc.contract.Read(fn).Call(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	ret, err := d.Unmarshal()
	if err != nil {
		return common.Address{}, err
	}
	addr, err := util.ToEtherAddress(ret[0])
	if err != nil {
		return common.Address{}, err
	}
	return address.FromBytes(addr.Bytes())
}

func (vc *vitaContract) readViewID(name string) (*big.Int, error) {
	d, err := vc.contract.Read("lastDonationPoolClaimViewID").Call(context.Background())
	if err != nil {
		return nil, err
	}
	ret, err := d.Unmarshal()
	if err != nil {
		return nil, err
	}
	return util.ToBigInt(ret[0])
}

func (vc *vitaContract) VPS() (address.Address, error) {
	return vc.readAddress("vps")
}

func (vc *vitaContract) DonationPoolAddress() (address.Address, error) {
	return vc.readAddress("donationPoolAddress")
}

func (vc *vitaContract) RewardPoolAddress() (address.Address, error) {
	return vc.readAddress("rewardPoolAddress")
}

func (vc *vitaContract) LastDonationPoolClaimViewID() (*big.Int, error) {
	return vc.readViewID("lastDonationPoolClaimViewID")
}

func (vc *vitaContract) LastRewardPoolClaimViewID() (*big.Int, error) {
	return vc.readViewID("lastRewardPoolClaimViewID")
}
