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
	"errors"
	"math/big"
)

// Candidate defines a delegate candidate
type Candidate struct {
	Registration
	score             *big.Int
	selfStakingTokens *big.Int
}

// NewCandidate creates a new candidate with scores as 0s
func NewCandidate(
	reg *Registration,
	score *big.Int,
	selfStakingTokens *big.Int,
) *Candidate {
	return &Candidate{
		*reg.Clone(),
		score,
		selfStakingTokens,
	}
}

// Clone clones the candidate
func (c *Candidate) Clone() *Candidate {
	return &Candidate{
		*c.Registration.Clone(),
		c.Score(),
		c.SelfStakingTokens(),
	}
}

// Equal returns true if two candidates are identical
func (c *Candidate) Equal(candidate *Candidate) bool {
	if c == candidate {
		return true
	}
	if c == nil || candidate == nil {
		return false
	}
	if !c.Registration.Equal(&candidate.Registration) {
		return false
	}
	if c.score.Cmp(candidate.score) != 0 {
		return false
	}
	if c.selfStakingTokens.Cmp(candidate.selfStakingTokens) != 0 {
		return false
	}
	return c.selfStakingWeight == candidate.selfStakingWeight
}

func (c *Candidate) reset() *Candidate {
	c.selfStakingTokens.SetInt64(0)
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

func (c *Candidate) addSelfStakingTokens(s *big.Int) error {
	if s.Cmp(big.NewInt(0)) < 0 {
		return errors.New("score cannot be negative")
	}
	c.selfStakingTokens.Add(c.selfStakingTokens, s)
	return nil
}

// Score returns the total votes (weighted) of this candidate
func (c *Candidate) Score() *big.Int {
	return new(big.Int).Set(c.score)
}

// SetScore set score value in Candidate
func (c *Candidate) SetScore(score *big.Int) {
	c.score = score
}

// SelfStakingTokens returns the total self votes (weighted)
func (c *Candidate) SelfStakingTokens() *big.Int {
	return new(big.Int).Set(c.selfStakingTokens)
}

// SetSelfStakingTokens set selfStakingTokens value in Candidate
func (c *Candidate) SetSelfStakingTokens(selfStakingTokens *big.Int) {
	c.selfStakingTokens = selfStakingTokens
}
