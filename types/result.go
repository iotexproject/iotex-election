// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package types

import (
	"encoding/hex"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-election/pb"
)

// ErrInvalidProto indicates a format error of an election proto
var ErrInvalidProto = errors.New("Invalid election proto")

// Result defines the collection of voting result on a height
type Result struct {
	mintTime   time.Time
	candidates []*Candidate
	votes      map[string][]*Vote
}

// MintTime returns the mint time of the corresponding beacon chain block
func (r *Result) MintTime() time.Time {
	return r.mintTime
}

// Candidates returns a list of sorted candidates
func (r *Result) Candidates() []*Candidate {
	return r.candidates
}

// VotesByCandidate returns a list of votes for a given candidate
func (r *Result) VotesByCandidate(name []byte) []*Vote {
	return r.votes[hex.EncodeToString(name)]
}

// ToProtoMsg converts the vote to protobuf
func (r *Result) ToProtoMsg() (*pb.ElectionResult, error) {
	candidates := make([]*pb.Candidate, len(r.candidates))
	candidateVotes := make([]*pb.VoteList, len(r.votes))
	var err error
	for i := 0; i < len(r.candidates); i++ {
		candidate := r.candidates[i]
		if candidates[i], err = candidate.ToProtoMsg(); err != nil {
			return nil, err
		}
		name := hex.EncodeToString(candidate.Name())
		votes, ok := r.votes[name]
		if !ok {
			return nil, errors.Errorf("Cannot find votes for candidate %s", name)
		}
		voteList := make([]*pb.Vote, len(votes))
		for j := 0; j < len(votes); j++ {
			if voteList[j], err = votes[j].ToProtoMsg(); err != nil {
				return nil, err
			}
		}
		candidateVotes[i] = &pb.VoteList{Votes: voteList}
	}
	t, err := ptypes.TimestampProto(r.mintTime)
	if err != nil {
		return nil, err
	}

	return &pb.ElectionResult{
		Timestamp:      t,
		Candidates:     candidates,
		CandidateVotes: candidateVotes,
	}, nil
}

// Serialize converts result to byte array
func (r *Result) Serialize() ([]byte, error) {
	rPb, err := r.ToProtoMsg()
	if err != nil {
		return nil, err
	}
	return proto.Marshal(rPb)
}

// FromProtoMsg extracts result details from protobuf message
func (r *Result) FromProtoMsg(rPb *pb.ElectionResult) (err error) {
	if len(rPb.Candidates) != len(rPb.CandidateVotes) {
		return errors.Wrapf(
			ErrInvalidProto,
			"size of candidate list %d is different from score list %d",
			len(rPb.Candidates),
			len(rPb.CandidateVotes),
		)
	}
	r.votes = map[string][]*Vote{}
	r.candidates = make([]*Candidate, len(rPb.Candidates))
	for i, cPb := range rPb.Candidates {
		r.candidates[i] = &Candidate{}
		if err := r.candidates[i].FromProtoMsg(cPb); err != nil {
			return err
		}
		name := hex.EncodeToString(r.candidates[i].Name())
		if _, ok := r.votes[name]; ok {
			return errors.Wrapf(
				ErrInvalidProto,
				"duplicate candidate %s",
				name,
			)
		}
		voteList := rPb.CandidateVotes[i]
		r.votes[name] = make([]*Vote, len(voteList.Votes))
		for j, vPb := range voteList.Votes {
			r.votes[name][j] = &Vote{}
			if err := r.votes[name][j].FromProtoMsg(vPb); err != nil {
				return err
			}
		}
	}
	if r.mintTime, err = ptypes.Timestamp(rPb.Timestamp); err != nil {
		return err
	}

	return nil
}

// Deserialize converts a byte array to election result
func (r *Result) Deserialize(data []byte) error {
	pb := &pb.ElectionResult{}
	if err := proto.Unmarshal(data, pb); err != nil {
		return err
	}

	return r.FromProtoMsg(pb)
}
