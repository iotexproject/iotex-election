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
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/util"
)

// ErrInvalidProto indicates a format error of an election proto
var ErrInvalidProto = errors.New("Invalid election proto")

// Poll defines the struct stores election result in db
type Poll struct {
	buckets [][]byte
	regs    [][]byte
}

// NewPoll creates a new electionresultMeta
func NewPoll(
	buckets [][]byte,
	regs [][]byte,
) *Poll {
	return &Poll{buckets, regs}
}

// Registrations returns a list of Registrations
func (m *Poll) Registrations() [][]byte {
	return m.regs
}

// Buckets returns all buckets
func (m *Poll) Buckets() [][]byte {
	return m.buckets
}

// ToProtoMsg converts the electionresultMeta to protobuf
func (m *Poll) ToProtoMsg() (*pb.Poll, error) {
	regKeys := make([][]byte, len(m.regs))
	bucketKeys := make([][]byte, len(m.buckets))

	for i, reg := range m.regs {
		regKeys[i] = util.CopyBytes(reg)
	}
	for i, bucket := range m.buckets {
		bucketKeys[i] = util.CopyBytes(bucket)
	}
	return &pb.Poll{
		Registrations: regKeys,
		Buckets:       bucketKeys,
	}, nil
}

// Serialize converts result to byte array
func (m *Poll) Serialize() ([]byte, error) {
	rPb, err := m.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(rPb)
}

// FromProtoMsg extracts result details from protobuf message
func (m *Poll) FromProtoMsg(rPb *pb.Poll) (err error) {
	m.buckets = make([][]byte, len(rPb.Buckets))
	for i, bucket := range rPb.Buckets {
		m.buckets[i] = util.CopyBytes(bucket)
	}

	m.regs = make([][]byte, len(rPb.Registrations))
	for i, reg := range rPb.Registrations {
		m.regs[i] = util.CopyBytes(reg)
	}
	return nil
}

// Deserialize converts a byte array to election result
func (m *Poll) Deserialize(data []byte) error {
	pb := &pb.Poll{}
	if err := proto.Unmarshal(data, pb); err != nil {
		return err
	}
	return m.FromProtoMsg(pb)
}
