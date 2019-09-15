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
	"math"
	"math/big"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/pkg/errors"
)

// Vote defines the structure of a vote
type Vote struct {
	Bucket
	weighted *big.Int
}

// NewVote creates a new vote
func NewVote(
	bucket *Bucket,
	weighted *big.Int,
) (*Vote, error) {
	if weighted == nil || big.NewInt(0).Cmp(weighted) > 0 {
		return nil, errors.Errorf("weighted amount %s cannot be nil or negative", weighted)
	}
	return &Vote{
		*bucket.Clone(),
		new(big.Int).Set(weighted),
	}, nil
}

// Clone clones the vote
func (v *Vote) Clone() *Vote {
	return &Vote{
		*v.Bucket.Clone(),
		v.WeightedAmount(),
	}
}

// SetWeightedAmount sets the weighted amount for the vote
func (v *Vote) SetWeightedAmount(w *big.Int) error {
	if w == nil || big.NewInt(0).Cmp(w) > 0 {
		return errors.New("weighted amount cannot be negative")
	}
	v.weighted = new(big.Int).Set(w)

	return nil
}

// WeightedAmount returns the weighted amount of vote
func (v *Vote) WeightedAmount() *big.Int {
	return new(big.Int).Set(v.weighted)
}

// CalcWeightedVotes calculates the weighted votes based on time
func CalcWeightedVotes(v *Bucket, now time.Time) *big.Int {
	if now.Before(v.StartTime()) {
		return big.NewInt(0)
	}
	remainingTime := v.RemainingTime(now).Seconds()
	weight := float64(1)
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)) / math.Log(1.2) / 100
	}
	amount := new(big.Float).SetInt(v.Amount())
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)

	return weightedAmount
}

// ToProtoMsg converts the vote to protobuf
func (v *Vote) ToProtoMsg() (*pb.Vote, error) {
	startTime, err := ptypes.TimestampProto(v.startTime)
	if err != nil {
		return nil, err
	}
	return &pb.Vote{
		Voter:          v.Voter(),
		Candidate:      v.Candidate(),
		Amount:         v.amount.Bytes(),
		WeightedAmount: v.weighted.Bytes(),
		StartTime:      startTime,
		Duration:       ptypes.DurationProto(v.duration),
		Decay:          v.decay,
	}, nil
}

// Serialize serializes the vote to bytes
func (v *Vote) Serialize() ([]byte, error) {
	vPb, err := v.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(vPb)
}

// FromProtoMsg extracts vote details from protobuf message
func (v *Vote) FromProtoMsg(vPb *pb.Vote) (err error) {
	voter := make([]byte, len(vPb.Voter))
	copy(voter, vPb.Voter)
	v.voter = voter
	candidate := make([]byte, len(vPb.Candidate))
	copy(candidate, vPb.Candidate)
	v.candidate = candidate
	v.amount = big.NewInt(0).SetBytes(vPb.Amount)
	v.weighted = new(big.Int).SetBytes(vPb.WeightedAmount)
	if v.startTime, err = ptypes.Timestamp(vPb.StartTime); err != nil {
		return err
	}
	if v.duration, err = ptypes.Duration(vPb.Duration); err != nil {
		return err
	}
	if v.duration < 0 {
		return errors.Errorf("duration %s cannot be negative", v.duration)
	}
	v.decay = vPb.Decay

	return nil
}

// Deserialize deserializes a byte array to vote
func (v *Vote) Deserialize(data []byte) error {
	vPb := &pb.Vote{}
	if err := proto.Unmarshal(data, vPb); err != nil {
		return err
	}

	return v.FromProtoMsg(vPb)
}
