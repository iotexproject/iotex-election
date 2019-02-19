// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"bytes"
	"errors"
	"math/big"

	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/util"
)

// Candidate defines a delegate candidate
type Candidate struct {
	name              []byte
	beaconPubKey      []byte
	operatorPubKey    []byte
	rewardPubKey      []byte
	score             *big.Int
	selfStakingScore  *big.Int
	selfStakingWeight uint32
}

// NewCandidate creates a new candidate with scores as 0s
func NewCandidate(
	name []byte,
	beaconPubKey []byte,
	operatorPubKey []byte,
	rewardPubKey []byte,
	selfStakingWeight uint32,
) *Candidate {
	return &Candidate{
		name:              util.CopyBytes(name),
		beaconPubKey:      util.CopyBytes(beaconPubKey),
		operatorPubKey:    util.CopyBytes(operatorPubKey),
		rewardPubKey:      util.CopyBytes(rewardPubKey),
		score:             big.NewInt(0),
		selfStakingScore:  big.NewInt(0),
		selfStakingWeight: selfStakingWeight,
	}
}

// Clone clones the candidate
func (c *Candidate) Clone() *Candidate {
	return &Candidate{
		name:              c.Name(),
		beaconPubKey:      c.BeaconPubKey(),
		operatorPubKey:    c.OperatorPubKey(),
		rewardPubKey:      c.RewardPubKey(),
		score:             c.Score(),
		selfStakingScore:  c.SelfStakingScore(),
		selfStakingWeight: c.SelfStakingWeight(),
	}
}

func (c *Candidate) equal(candidate *Candidate) bool {
	if c == candidate {
		return true
	}
	if c == nil || candidate == nil {
		return false
	}
	if !bytes.Equal(c.name, candidate.name) {
		return false
	}
	if !bytes.Equal(c.beaconPubKey, candidate.beaconPubKey) {
		return false
	}
	if !bytes.Equal(c.operatorPubKey, candidate.operatorPubKey) {
		return false
	}
	if !bytes.Equal(c.rewardPubKey, candidate.rewardPubKey) {
		return false
	}
	if c.score.Cmp(candidate.score) != 0 {
		return false
	}
	if c.selfStakingScore.Cmp(candidate.selfStakingScore) != 0 {
		return false
	}
	return c.selfStakingWeight == candidate.selfStakingWeight
}

func (c *Candidate) reset() *Candidate {
	c.selfStakingScore.SetInt64(0)
	c.score.SetInt64(0)
	return c
}

func (c *Candidate) addScore(s *big.Int) error {
	if s.Cmp(big.NewInt(0)) < 0 {
		return errors.New("score cannot be negative")
	}
	c.score.Add(c.score, s)
	return nil
}

func (c *Candidate) addSelfStakingScore(s *big.Int) error {
	if s.Cmp(big.NewInt(0)) < 0 {
		return errors.New("score cannot be negative")
	}
	c.selfStakingScore.Add(c.selfStakingScore, s)
	return nil
}

// Name returns the name of this candidate
func (c *Candidate) Name() []byte {
	return util.CopyBytes(c.name)
}

// BeaconPubKey returns the public key of this candidate on beacon chain
func (c *Candidate) BeaconPubKey() []byte {
	return util.CopyBytes(c.beaconPubKey)
}

// OperatorPubKey returns the public key of the assigned operator on chain
func (c *Candidate) OperatorPubKey() []byte {
	return util.CopyBytes(c.operatorPubKey)
}

// RewardPubKey returns the public key of the assigned benefiter on chain
func (c *Candidate) RewardPubKey() []byte {
	return util.CopyBytes(c.rewardPubKey)
}

// Score returns the total votes (weighted) of this candidate
func (c *Candidate) Score() *big.Int {
	return new(big.Int).Set(c.score)
}

// SelfStakingScore returns the total self votes (weighted)
func (c *Candidate) SelfStakingScore() *big.Int {
	return new(big.Int).Set(c.selfStakingScore)
}

// SelfStakingWeight returns the extra weight for self staking
func (c *Candidate) SelfStakingWeight() uint32 {
	return c.selfStakingWeight
}

// ToProtoMsg converts the instance to a protobuf message
func (c *Candidate) ToProtoMsg() (*pb.Candidate, error) {
	return &pb.Candidate{
		Name:              c.Name(),
		BeaconPubKey:      c.BeaconPubKey(),
		OperatorPubKey:    c.OperatorPubKey(),
		RewardPubKey:      c.RewardPubKey(),
		Score:             c.score.Bytes(),
		SelfStakingScore:  c.selfStakingScore.Bytes(),
		SelfStakingWeight: c.selfStakingWeight,
	}, nil
}

// FromProtoMsg fills the instance with a protobuf message
func (c *Candidate) FromProtoMsg(msg *pb.Candidate) error {
	c.name = util.CopyBytes(msg.GetName())
	c.beaconPubKey = util.CopyBytes(msg.GetBeaconPubKey())
	c.operatorPubKey = util.CopyBytes(msg.GetOperatorPubKey())
	c.rewardPubKey = util.CopyBytes(msg.GetRewardPubKey())
	c.score = new(big.Int).SetBytes(msg.GetScore())
	c.selfStakingScore = new(big.Int).SetBytes(msg.GetSelfStakingScore())
	c.selfStakingWeight = msg.GetSelfStakingWeight()

	return nil
}
