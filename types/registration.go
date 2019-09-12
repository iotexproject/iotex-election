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

	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/util"
)

// Registration defines the meta of a candidate
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
func (c *Registration) Clone() *Registration {
	return &Registration{
		name:              c.Name(),
		address:           c.Address(),
		operatorAddress:   c.OperatorAddress(),
		rewardAddress:     c.RewardAddress(),
		selfStakingWeight: c.SelfStakingWeight(),
	}
}

// Name returns the name of this candidate
func (c *Registration) Name() []byte {
	return util.CopyBytes(c.name)
}

// Address returns the address of this candidate on gravity chain
func (c *Registration) Address() []byte {
	return util.CopyBytes(c.address)
}

// OperatorAddress returns the address of the assigned operator on chain
func (c *Registration) OperatorAddress() []byte {
	return util.CopyBytes(c.operatorAddress)
}

// RewardAddress returns the address of the assigned benefiter on chain
func (c *Registration) RewardAddress() []byte {
	return util.CopyBytes(c.rewardAddress)
}

// SelfStakingWeight returns the extra weight for self staking
func (c *Registration) SelfStakingWeight() uint64 {
	return c.selfStakingWeight
}

// ToProtoMsg converts the instance to a protobuf message (CandidateCore)
func (c *Registration) ToProtoMsg() (*pb.Registration, error) {
	return &pb.Registration{
		Name:              c.Name(),
		Address:           c.Address(),
		OperatorAddress:   c.OperatorAddress(),
		RewardAddress:     c.RewardAddress(),
		SelfStakingWeight: c.selfStakingWeight,
	}, nil
}

// Serialize serializes the candidate to bytes
func (c *Registration) Serialize() ([]byte, error) {
	cPb, err := c.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(cPb)
}

// FromProtoMsg fills the instance with a protobuf message (CandidateCore)
func (c *Registration) FromProtoMsg(msg *pb.Registration) error {
	c.name = util.CopyBytes(msg.GetName())
	c.address = util.CopyBytes(msg.GetAddress())
	c.operatorAddress = util.CopyBytes(msg.GetOperatorAddress())
	c.rewardAddress = util.CopyBytes(msg.GetRewardAddress())
	c.selfStakingWeight = msg.GetSelfStakingWeight()

	return nil
}

// Deserialize deserializes a byte array to candidate
func (c *Registration) Deserialize(data []byte) error {
	cPb := &pb.Registration{}
	if err := proto.Unmarshal(data, cPb); err != nil {
		return err
	}

	return c.FromProtoMsg(cPb)
}

// Equal returns true if two candidates are identical
func (c *Registration) Equal(candidate *Registration) bool {
	if c == candidate {
		return true
	}
	if c == nil || candidate == nil {
		return false
	}
	if !bytes.Equal(c.name, candidate.name) {
		return false
	}
	if !bytes.Equal(c.address, candidate.address) {
		return false
	}
	if !bytes.Equal(c.operatorAddress, candidate.operatorAddress) {
		return false
	}
	if !bytes.Equal(c.rewardAddress, candidate.rewardAddress) {
		return false
	}

	return c.selfStakingWeight == candidate.selfStakingWeight
}
