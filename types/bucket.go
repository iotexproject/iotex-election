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
	"math/big"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"

	pb "github.com/iotexproject/iotex-election/pb/election"
)

// Bucket defines a bucket stored in staking contract
type Bucket struct {
	index     *big.Int
	startTime time.Time
	duration  time.Duration
	amount    *big.Int
	decay     bool
	voter     []byte
	candidate []byte
}

// NewBucket creates a new bucket
func NewBucket(
	index     *big.Int,
	startTime time.Time,
	duration time.Duration,
	amount *big.Int,
	voter []byte,
	candidate []byte,
	decay bool,
) (*Bucket, error) {
	if duration < 0 {
		return nil, errors.Errorf("duration %s cannot be negative", duration)
	}
	if amount == nil || big.NewInt(0).Cmp(amount) > 0 {
		return nil, errors.Errorf("amount %s cannot be nil or negative", amount)
	}
	cVoter := make([]byte, len(voter))
	copy(cVoter, voter)
	cCandidate := make([]byte, len(candidate))
	copy(cCandidate, candidate)

	return &Bucket{
		index:	   index,
		startTime: startTime,
		duration:  duration,
		amount:    new(big.Int).Set(amount),
		voter:     cVoter,
		candidate: cCandidate,
		decay:     decay,
	}, nil
}

// Clone clones the vote
func (bucket *Bucket) Clone() *Bucket {
	return &Bucket{
		bucket.Index(),
		bucket.StartTime(),
		bucket.Duration(),
		bucket.Amount(),
		bucket.Decay(),
		bucket.Voter(),
		bucket.Candidate(),
	}
}

func (bucket *Bucket) Index() *big.Int {
	return new(big.Int).Set(bucket.amount)
}

// StartTime returns the start time
func (bucket *Bucket) StartTime() time.Time {
	return bucket.startTime
}

// Duration returns the duration this vote for
func (bucket *Bucket) Duration() time.Duration {
	return bucket.duration
}

// Voter returns the voter address in bytes
func (bucket *Bucket) Voter() []byte {
	voter := make([]byte, len(bucket.voter))
	copy(voter, bucket.voter)

	return voter
}

// Amount returns the amount of vote
func (bucket *Bucket) Amount() *big.Int {
	return new(big.Int).Set(bucket.amount)
}

// Candidate returns the candidate
func (bucket *Bucket) Candidate() []byte {
	candidate := make([]byte, len(bucket.candidate))
	copy(candidate, bucket.candidate)

	return candidate
}

// Decay returns whether this is a decay vote
func (bucket *Bucket) Decay() bool {
	return bucket.decay
}

// RemainingTime returns the remaining time to given time
func (bucket *Bucket) RemainingTime(now time.Time) time.Duration {
	if now.Before(bucket.startTime) {
		return 0
	}
	if bucket.decay {
		endTime := bucket.startTime.Add(bucket.duration)
		if endTime.After(now) {
			return bucket.startTime.Add(bucket.duration).Sub(now)
		}
		return 0
	}
	return bucket.duration
}

// ToProtoMsg converts the vote to protobuf
func (bucket *Bucket) ToProtoMsg() (*pb.Bucket, error) {
	startTime, err := ptypes.TimestampProto(bucket.startTime)
	if err != nil {
		return nil, err
	}
	return &pb.Bucket{
		Index:	   bucket.index.Bytes(),
		Voter:     bucket.Voter(),
		Candidate: bucket.Candidate(),
		Amount:    bucket.amount.Bytes(),
		StartTime: startTime,
		Duration:  ptypes.DurationProto(bucket.duration),
		Decay:     bucket.decay,
	}, nil
}

// Serialize serializes the vote to bytes
func (bucket *Bucket) Serialize() ([]byte, error) {
	vPb, err := bucket.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(vPb)
}

// FromProtoMsg extracts vote details from protobuf message (voteCore)
func (bucket *Bucket) FromProtoMsg(vPb *pb.Bucket) (err error) {
	bucket.index = big.NewInt(0).SetBytes(vPb.Index)
	voter := make([]byte, len(vPb.Voter))
	copy(voter, vPb.Voter)
	bucket.voter = voter
	candidate := make([]byte, len(vPb.Candidate))
	copy(candidate, vPb.Candidate)
	bucket.candidate = candidate
	bucket.amount = big.NewInt(0).SetBytes(vPb.Amount)
	if bucket.startTime, err = ptypes.Timestamp(vPb.StartTime); err != nil {
		return err
	}
	if bucket.duration, err = ptypes.Duration(vPb.Duration); err != nil {
		return err
	}
	if bucket.duration < 0 {
		return errors.Errorf("duration %s cannot be negative", bucket.duration)
	}
	bucket.decay = vPb.Decay
	return nil
}

// Deserialize deserializes a byte array to vote
func (bucket *Bucket) Deserialize(data []byte) error {
	vPb := &pb.Bucket{}
	if err := proto.Unmarshal(data, vPb); err != nil {
		return err
	}

	return bucket.FromProtoMsg(vPb)
}

func (v *Bucket) equal(bucket *Bucket) bool {
	if v == bucket {
		return true
	}
	if v == nil || bucket == nil {
		return false
	}
	if v.index.Cmp(bucket.index) != 0 {
		return false
	}
	if !v.startTime.Equal(bucket.startTime) {
		return false
	}
	if v.duration != bucket.duration {
		return false
	}
	if v.amount.Cmp(bucket.amount) != 0 {
		return false
	}
	if !bytes.Equal(v.voter, bucket.voter) {
		return false
	}
	if !bytes.Equal(v.candidate, bucket.candidate) {
		return false
	}
	return v.decay == bucket.decay
}
