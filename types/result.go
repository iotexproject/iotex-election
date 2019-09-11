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
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"


	pb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/util"

)

// ErrInvalidProto indicates a format error of an election proto
var ErrInvalidProto = errors.New("Invalid election proto")


type ElectionResultMeta struct {
	mintTime 		time.Time
	candidates 		[][]byte
	votes 			[][]byte
}


// NewElectionResultMeta creates a new electionresultMeta 
func NewElectionResultMeta(
	mintTime time.Time,
	candidates [][]byte,
	votes [][]byte,
) *ElectionResultMeta {
	return &ElectionResultMeta{
		mintTime:		mintTime,
		candidates:     candidates,
		votes:   		votes,
	}
}

// MintTime returns the mint time of the corresponding gravity chain block
func (m *ElectionResultMeta) MintTime() time.Time {
	return m.mintTime
}

// Candidates returns a list of candidates
func (m *ElectionResultMeta) Candidates() [][]byte {
	return m.candidates
}

// Votes returns all votes
func (m *ElectionResultMeta) Votes() [][]byte {
	return m.votes
}

// ToProtoMsg converts the electionresultMeta to protobuf
func (m *ElectionResultMeta) ToProtoMsg() (*pb.ElectionResultMeta, error) {
	candidatesKey := make([][]byte, len(m.candidates))
	votesKey := make([][]byte, len(m.votes))

	for i, cand := range m.candidates {
		candidatesKey[i] = util.CopyBytes(cand)
	}
	for i, vote := range m.votes {
		votesKey[i] = util.CopyBytes(vote)
	}
	t, err := ptypes.TimestampProto(m.mintTime)
	if err != nil {
		return nil, err
	}
	return &pb.ElectionResultMeta{
		Timestamp:				t,
		CandidatesKey:			candidatesKey,
		VotesKey:    			votesKey,
	}, nil
}

// Serialize converts result to byte array
func (m *ElectionResultMeta) Serialize() ([]byte, error) {
	rPb, err := m.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(rPb)
}



// FromProtoMsg extracts result details from protobuf message
func (m *ElectionResultMeta) FromProtoMsg(rPb *pb.ElectionResultMeta) (err error) {
	if m.mintTime, err = ptypes.Timestamp(rPb.Timestamp); err != nil {
		return err
	}
	m.votes = make([][]byte, len(rPb.VotesKey))
	for i, vote := range rPb.VotesKey {
		m.votes[i] = util.CopyBytes(vote)
	}

	m.candidates = make([][]byte, len(rPb.CandidatesKey))
	for i, cand := range rPb.CandidatesKey {
		m.candidates[i] = util.CopyBytes(cand)
	}
	return nil
}

// Deserialize converts a byte array to election result
func (m *ElectionResultMeta) Deserialize(data []byte) error {
	pb := &pb.ElectionResultMeta{}
	if err := proto.Unmarshal(data, pb); err != nil {
		return err
	}
	return m.FromProtoMsg(pb)
}


// ElectionResult defines the collection of voting result on a height
type ElectionResult struct {
	mintTime         time.Time
	delegates        []*Candidate
	votes            map[string][]*Vote
	totalVotes       *big.Int
	totalVotedStakes *big.Int
}

// MintTime returns the mint time of the corresponding gravity chain block
func (r *ElectionResult) MintTime() time.Time {
	return r.mintTime
}

// Delegates returns a list of sorted delegates
func (r *ElectionResult) Delegates() []*Candidate {
	return r.delegates
}

// VotesByDelegate returns a list of votes for a given delegate
func (r *ElectionResult) VotesByDelegate(name []byte) []*Vote {
	return r.votes[hex.EncodeToString(name)]
}

// Votes returns all votes
func (r *ElectionResult) Votes() []*Vote {
	votes := []*Vote{}
	for _, vs := range r.votes {
		votes = append(votes, vs...)
	}
	return votes
}

// DelegateByName returns the candidate details
func (r *ElectionResult) DelegateByName(name []byte) *Candidate {
	for _, candidate := range r.delegates {
		if bytes.Equal(candidate.Name(), name) {
			return candidate
		}
	}
	return nil
}

// TotalVotes returns the total votes in the result
func (r *ElectionResult) TotalVotes() *big.Int {
	return new(big.Int).Set(r.totalVotes)
}

// TotalVotedStakes returns the total amount of stakings which has been voted
func (r *ElectionResult) TotalVotedStakes() *big.Int {
	return new(big.Int).Set(r.totalVotedStakes)
}

func (r *ElectionResult) String() string {
	var builder strings.Builder
	fmt.Fprintf(
		&builder,
		"Timestamp: %s\nTotal Voted Stakes: %d\nTotal Votes: %d\n",
		r.mintTime,
		r.totalVotedStakes,
		r.totalVotes,
	)
	for i, d := range r.delegates {
		fmt.Fprintf(
			&builder,
			"%d: %s %x\n\toperator address: %s\n\treward: %s\n\tvotes: %s\n",
			i,
			string(d.name),
			d.name,
			string(d.operatorAddress),
			string(d.rewardAddress),
			d.score,
		)
	}
	return builder.String()
}


func (r *ElectionResult) Equal(result *ElectionResult) bool {
	if r == result {
		return true
	}
	if r == nil || result == nil {
		return false
	}
	if !r.mintTime.Equal(result.mintTime) {
		return false
	}
	if r.totalVotedStakes.Cmp(result.totalVotedStakes) != 0 {
		return false
	}
	if r.totalVotes.Cmp(result.totalVotes) != 0 {
		return false
	}
	if len(r.delegates) != len(result.delegates){
		return false
	}
	if len(r.votes) != len(result.votes) {
		return false
	}
	for i, delegate := range r.delegates {
		if !delegate.Equal(result.delegates[i]) {
			return false 
		}
	}
	return true 
}

// NewElectionResultForTest creates an election result for test purpose only
func NewElectionResultForTest(
	mintTime time.Time,
) *ElectionResult {
	return &ElectionResult{
		mintTime: mintTime,
		delegates: []*Candidate{
			&Candidate{
				name:            []byte("name1"),
				address:         []byte("address1"),
				operatorAddress: []byte("io1kfpsvefk74cqxd245j2h5t2pld2wtxzyg6tqrt"),
				rewardAddress:   []byte("io1kfpsvefk74cqxd245j2h5t2pld2wtxzyg6tqrt"),
				score:           big.NewInt(15),
			},
			&Candidate{
				name:            []byte("name2"),
				address:         []byte("address2"),
				operatorAddress: []byte("io1llr6zs37gxrwmvnczexpg35dptta2mdvjv6w2q"),
				rewardAddress:   []byte("io1llr6zs37gxrwmvnczexpg35dptta2mdvjv6w2q"),
				score:           big.NewInt(14),
			},
		},
		votes: map[string][]*Vote{
			"name1": []*Vote{},
			"name2": []*Vote{},
		},
	}
}
