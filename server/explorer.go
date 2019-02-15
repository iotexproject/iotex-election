package server

import (
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iotexproject/iotex-election/committee"
	pb "github.com/iotexproject/iotex-election/explorer_pb"

	"golang.org/x/net/context"
)

// server is used to implement pb.ExplorerServer.
type server struct {
	savedCustomers []*pb.ExplorerServer
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
		TotalCandidates: int64(len(result.Candidates())),
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *server) GetCandidates(ctx context.Context, cr *pb.GetCandidatesRequest) (*pb.CandidateResponse, error) {
	return new(pb.CandidateResponse), nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, br *pb.GetBucketsByCandidateRequest) (*pb.BucketResponse, error) {
	return new(pb.BucketResponse), nil
}
