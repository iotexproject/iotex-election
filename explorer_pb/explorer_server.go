package explorer_pb

import (
	"github.com/ashishsnigam/iotex-election/committee"
	"strconv"

	"golang.org/x/net/context"
)

// server is used to implement pb.ExplorerServer.
type server struct {
	savedCustomers []*ExplorerServer
	Impl
}

// Impl implements Explorer interface
type Impl struct {
	c committee.Committee
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *Empty) (*ChainMeta, error) {
	height := s.Impl.c.LatestHeight()
	result, err := s.Impl.c.ResultByHeight(height)
	if err != nil {
		return &ChainMeta{}, err
	}
	return &ChainMeta{
		Height:          strconv.FormatUint(height, 10),
		TotalCandidates: int64(len(result.Candidates())),
	}, nil
}

func (s *server) GetCandidates(ctx context.Context, cr *GetCandidatesRequest) (*CandidateResponse, error) {
	return new(CandidateResponse), nil
}

func (s *server) GetBucketsByCandidate(ctx context.Context, br *GetBucketsByCandidateRequest) (*BucketResponse, error) {
	return new(BucketResponse), nil
}
