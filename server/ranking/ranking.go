// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package ranking

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"github.com/iotexproject/iotex-election/committee"
	pb "github.com/iotexproject/iotex-election/pb/ranking"
)

// Config defines the config for server
type Config struct {
	DBPath    string           `yaml:"dbPath"`
	Committee committee.Config `yaml:"committee"`
}

// server is used to implement pb.RankingServer.
type server struct {
	c committee.Committee
}

// NewServer returns an implementation of ranking server
func NewServer() (pb.RankingServer, error) {
	// TODO: read config from yaml file
	cfg := &Config{}
	c, err := committee.NewCommittee(nil, cfg.Committee)
	if err != nil {
		return nil, err
	}

	return &server{c: c}, nil
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *empty.Empty) (*pb.ChainMeta, error) {
	height := s.c.LatestHeight()
	result, err := s.c.ResultByHeight(height)
	if err != nil {
		return &pb.ChainMeta{}, err
	}

	return &pb.ChainMeta{
		Height:          strconv.FormatUint(height, 10),
		TotalCandidates: uint64(len(result.Delegates())),
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *server) GetCandidates(ctx context.Context, request *pb.GetCandidatesRequest) (*pb.CandidateResponse, error) {
	offset := request.Offset
	limit := request.Limit
	if offset < 0 || limit < 0 {
		return nil, errors.New("offset and limit should be positive number")
	}
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.c.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	candidates := result.Delegates()
	if uint64(len(candidates)) <= offset {
		return nil, errors.New("offset is larger than candidate length")
	}
	if uint64(len(candidates)) < offset+limit {
		limit = uint64(len(candidates)) - offset
	}
	response := &pb.CandidateResponse{
		Candidates: make([]*pb.Candidate, limit),
	}
	for i := offset; i < offset+limit; i++ {
		candidate := candidates[i]
		response.Candidates[i].Name = hex.EncodeToString(candidate.Name())
		response.Candidates[i].PubKey = hex.EncodeToString(candidate.BeaconPubKey())
		response.Candidates[i].TotalWeightedVotes = candidate.Score().Text(10)
	}

	return response, nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *pb.GetBucketsByCandidateRequest) (*pb.BucketResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.c.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	votes := result.VotesByDelegate([]byte(request.Name))
	if votes == nil {
		return nil, errors.New("No buckets for the candidate")
	}
	response := &pb.BucketResponse{
		Buckets: make([]*pb.Bucket, len(votes)),
	}
	for i, vote := range votes {
		response.Buckets[i].Voter = hex.EncodeToString(vote.Voter())
		response.Buckets[i].Votes = vote.Amount().Text(10)
		response.Buckets[i].WeightedVotes = vote.WeightedAmount().Text(10)
		response.Buckets[i].RemainingDuration = uint64(vote.RemainingTime(result.MintTime()) / time.Second)
	}

	return response, nil
}
