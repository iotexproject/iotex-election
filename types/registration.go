// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"

	"github.com/golang/protobuf/proto"

	"github.com/iotexproject/go-pkgs/hash"
	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/util"
)

// Registration defines a registration in contract
type Registration struct {
	name              []byte
	address           []byte
	operatorAddress   []byte
	rewardAddress     []byte
	selfStakingWeight uint64
}

// NewRegistration creates a new registration
func NewRegistration(
	name []byte,
	address []byte,
	operatorAddress []byte,
	rewardPubKey []byte,
	selfStakingWeight uint64,
) *Registration {
	return &Registration{
		name:              util.CopyBytes(name),
		address:           util.CopyBytes(address),
		operatorAddress:   util.CopyBytes(operatorAddress),
		rewardAddress:     util.CopyBytes(rewardPubKey),
		selfStakingWeight: selfStakingWeight,
	}
}

// Clone clones the registration
func (reg *Registration) Clone() *Registration {
	return &Registration{
		name:              reg.Name(),
		address:           reg.Address(),
		operatorAddress:   reg.OperatorAddress(),
		rewardAddress:     reg.RewardAddress(),
		selfStakingWeight: reg.SelfStakingWeight(),
	}
}

// Name returns the name of this candidate
func (reg *Registration) Name() []byte {
	return util.CopyBytes(reg.name)
}

// Address returns the address of this candidate on gravity chain
func (reg *Registration) Address() []byte {
	return util.CopyBytes(reg.address)
}

// OperatorAddress returns the address of the assigned operator on chain
func (reg *Registration) OperatorAddress() []byte {
	return util.CopyBytes(reg.operatorAddress)
}

// RewardAddress returns the address of the assigned benefiter on chain
func (reg *Registration) RewardAddress() []byte {
	return util.CopyBytes(reg.rewardAddress)
}

// SelfStakingWeight returns the extra weight for self staking
func (reg *Registration) SelfStakingWeight() uint64 {
	return reg.selfStakingWeight
}

// ToProtoMsg converts the instance to a protobuf message (CandidateCore)
func (reg *Registration) ToProtoMsg() (*pb.Registration, error) {
	return &pb.Registration{
		Name:              reg.Name(),
		Address:           reg.Address(),
		OperatorAddress:   reg.OperatorAddress(),
		RewardAddress:     reg.RewardAddress(),
		SelfStakingWeight: reg.selfStakingWeight,
	}, nil
}

// Serialize serializes the candidate to bytes
func (reg *Registration) Serialize() ([]byte, error) {
	cPb, err := reg.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(cPb)
}

// Hash returns the hash of serialized data
func (reg *Registration) Hash() (hash.Hash256, error) {
	data, err := reg.Serialize()
	if err != nil {
		return hash.ZeroHash256, err
	}
	return hash.Hash256b(data), nil
}

// FromProtoMsg fills the instance with a protobuf message (CandidateCore)
func (reg *Registration) FromProtoMsg(msg *pb.Registration) error {
	reg.name = util.CopyBytes(msg.GetName())
	reg.address = util.CopyBytes(msg.GetAddress())
	reg.operatorAddress = util.CopyBytes(msg.GetOperatorAddress())
	reg.rewardAddress = util.CopyBytes(msg.GetRewardAddress())
	reg.selfStakingWeight = msg.GetSelfStakingWeight()

	return nil
}

// Deserialize deserializes a byte array to candidate
func (reg *Registration) Deserialize(data []byte) error {
	cPb := &pb.Registration{}
	if err := proto.Unmarshal(data, cPb); err != nil {
		return err
	}

	return reg.FromProtoMsg(cPb)
}

// Equal returns true if two candidates are identical
func (reg *Registration) Equal(r *Registration) bool {
	if reg == r {
		return true
	}
	if reg == nil || r == nil {
		return false
	}
	if !bytes.Equal(reg.name, r.name) {
		return false
	}
	if !bytes.Equal(reg.address, r.address) {
		return false
	}
	if !bytes.Equal(reg.operatorAddress, r.operatorAddress) {
		return false
	}
	if !bytes.Equal(reg.rewardAddress, r.rewardAddress) {
		return false
	}

	return reg.selfStakingWeight == r.selfStakingWeight
}
