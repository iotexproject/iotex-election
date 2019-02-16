package server

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iotexproject/iotex-election/committee"
	pb "github.com/iotexproject/iotex-election/ranking_pb"

	"golang.org/x/net/context"
)

// server is used to implement pb.RankingServer.
type server struct {
	savedCustomers []*pb.RankingServer
	c              committee.Committee
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty empty.Empty) (*pb.ChainMeta, error) {
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
func (s *server) GetCandidates(ctx context.Context, request *pb.GetCandidatesRequest) ([]*pb.CandidateResponse, error) {
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
	retval := make([]*pb.CandidateResponse, limit)
	for i := offset; i < offset+limit; i++ {
		candidate := candidates[i]
		retval[i].Candidates[i].Name = hex.EncodeToString(candidate.Name())
		retval[i].Candidates[i].PubKey = hex.EncodeToString(candidate.BeaconPubKey())
		retval[i].Candidates[i].TotalWeightedVotes = candidate.Score().Text(10)
	}

	return retval, nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *pb.GetBucketsByCandidateRequest) ([]*pb.BucketResponse, error) {
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
	retval := make([]*pb.BucketResponse, len(votes))
	for i, vote := range votes {
		retval[i].Buckets[i].Voter = hex.EncodeToString(vote.Voter())
		retval[i].Buckets[i].Votes = vote.Amount().Text(10)
		retval[i].Buckets[i].WeightedVotes = vote.WeightedAmount().Text(10)
		retval[i].Buckets[i].RemainingDuration = uint64(vote.RemainingTime(result.MintTime()) / time.Second)
	}

	return retval, nil
}
