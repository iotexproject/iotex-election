// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package server

import (
	"encoding/hex"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/iotexproject/iotex-address/address"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/iotex-election/pb/api"
	"github.com/iotexproject/iotex-election/votesync"
)

// DummyServer defines the interface of the ranking dummy server implementation
type DummyServer interface {
	api.APIServiceServer
	Start(context.Context) error
	Stop(context.Context) error
}

// server implements api.APIServiceServer.
type dummyServer struct {
	api.UnimplementedAPIServiceServer
	port       int
	grpcServer *grpc.Server
	vs         *votesync.VoteSync
}

// NewDummyServer returns an implementation of ranking dummy server
func NewDummyServer(port int, vs *votesync.VoteSync) (DummyServer, error) {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.InfoLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)
	s := &dummyServer{
		port: port,
		vs:   vs,
	}
	s.grpcServer = grpc.NewServer()
	api.RegisterAPIServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)
	return s, nil
}

func (s *dummyServer) Start(ctx context.Context) error {
	zap.L().Info("Start ranking dummpy server")
	zap.L().Info("Listen to port", zap.Int("port", s.port))
	portStr := ":" + strconv.Itoa(s.port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		zap.L().Error("Ranking dummpy server failed to listen port.", zap.Error(err))
		return err
	}
	go func() {
		if s.vs != nil {
			s.vs.Start(ctx)
		}
	}()
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			zap.L().Fatal("Failed to serve", zap.Error(err))
		}
	}()
	return nil
}

func (s *dummyServer) Stop(ctx context.Context) error {
	zap.L().Info("Dummpy server is stopping")
	s.grpcServer.Stop()
	if s.vs != nil {
		s.vs.Stop(ctx)
	}
	return nil
}

// GetMeta returns the meta of the chain
func (s *dummyServer) GetMeta(ctx context.Context, empty *empty.Empty) (*api.ChainMeta, error) {
	zap.L().Info("Dummpy server calls GetMeta func")
	return nil, nil
}

func (s *dummyServer) IsHealth(ctx context.Context, empty *empty.Empty) (*api.HealthCheckResponse, error) {
	zap.L().Info("Dummpy server calls IsHealth func")
	return nil, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *dummyServer) GetCandidates(ctx context.Context, request *api.GetCandidatesRequest) (*api.CandidateResponse, error) {
	zap.L().Info("Dummpy server calls GetCandidates func")
	return nil, nil
}

// GetCandidateByName returns the candidate details
func (s *dummyServer) GetCandidateByName(ctx context.Context, request *api.GetCandidateByNameRequest) (*api.Candidate, error) {
	zap.L().Info("Dummpy server calls GetCandidateByName func")
	return nil, nil
}

// GetBucketsByCandidate returns the buckets
func (s *dummyServer) GetBucketsByCandidate(ctx context.Context, request *api.GetBucketsByCandidateRequest) (*api.BucketResponse, error) {
	zap.L().Info("Dummpy server calls GetBucketsByCandidate func")
	return nil, nil
}

// GetBuckets returns a list of buckets
func (s *dummyServer) GetBuckets(ctx context.Context, request *api.GetBucketsRequest) (*api.BucketResponse, error) {
	zap.L().Info("Dummpy server calls GetBuckets func")
	return nil, nil
}

func (s *dummyServer) GetRawData(ctx context.Context, request *api.GetRawDataRequest) (*api.RawDataResponse, error) {
	zap.L().Info("Dummpy server calls GetRawData func")
	return nil, nil
}

func (s *dummyServer) GetProof(ctx context.Context, request *api.ProofRequest) (*api.ProofResponse, error) {
	zap.L().Info("Dummpy server calls GetProof func")
	if s.vs == nil {
		return nil, errors.New("no vote sync server")
	}
	addr, err := address.FromString(request.Account)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to cast account %s", request.Account)
	}
	proof, err := s.vs.ProofForAccount(addr)
	if err != nil {
		return nil, err
	}
	if proof == nil {
		return nil, nil
	}

	return &api.ProofResponse{Proof: hex.EncodeToString(proof)}, nil
}
