package server

import (
	"strconv"

	pb "github.com/ashishsnigam/iotex-election/explorer_pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iotexproject/iotex-election/committee"

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

func (s *server) GetCandidates(ctx context.Context, cr *pb.GetCandidatesRequest) (*pb.CandidateResponse, error) {
	return new(pb.CandidateResponse), nil
}

func (s *server) GetBucketsByCandidate(ctx context.Context, br *pb.GetBucketsByCandidateRequest) (*pb.BucketResponse, error) {
	return new(pb.BucketResponse), nil
}
