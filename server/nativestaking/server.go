// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/pb/api"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/votesync"
)

var ErrNotSupported = errors.New("Not supported")

// Config defines the config for server
type Config struct {
	DB        db.Config                       `yaml:"db"`
	Port      int                             `yaml:"port"`
	Committee committee.NativeCommitteeConfig `yaml:"committee"`
}

// Server defines the interface of the ranking server implementation
type Server interface {
	api.APIServiceServer
	Start(context.Context) error
	Stop(context.Context) error
}

// server implements api.APIServiceServer.
type server struct {
	port                 int
	nativeCommittee      *committee.NativeCommittee
	grpcServer           *grpc.Server
	selfStakingThreshold *big.Int
	scoreThreshold       *big.Int
	voteSync             *votesync.VoteSync
}

// NewServer returns an implementation of ranking server
func NewServer(cfg *Config) (Server, error) {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.DebugLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)
	archive, err := committee.NewBucketArchive(cfg.DB.DBPath, cfg.DB.NumOfRetries, cfg.Committee.StartHeight, cfg.Committee.Interval)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create archive")
	}
	c, err := committee.NewNativeStaingCommittee(archive, cfg.Committee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create committee")
	}
	s := &server{
		nativeCommittee: c,
		port:            cfg.Port,
	}
	s.grpcServer = grpc.NewServer()
	api.RegisterAPIServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	return s, nil
}

func (s *server) Start(ctx context.Context) error {
	zap.L().Info("Start ranking server")
	zap.L().Info("Listen to port", zap.Int("port", s.port))
	portStr := ":" + strconv.Itoa(s.port)
	lis, err := net.Listen("tcp", portStr)
	if err != nil {
		zap.L().Error("Ranking server failed to listen port.", zap.Error(err))
		return err
	}
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			zap.L().Fatal("Failed to serve", zap.Error(err))
		}
	}()
	if err := s.nativeCommittee.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	return s.nativeCommittee.Stop(ctx)
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *empty.Empty) (*api.ChainMeta, error) {
	height := s.nativeCommittee.TipHeight()

	return &api.ChainMeta{
		Height: strconv.FormatUint(height, 10),
	}, nil
}

func (s *server) IsHealth(ctx context.Context, empty *empty.Empty) (*api.HealthCheckResponse, error) {
	var status api.HealthCheckResponse_Status
	switch s.nativeCommittee.Status() {
	case committee.STARTING:
		status = api.HealthCheckResponse_STARTING
	case committee.ACTIVE:
		status = api.HealthCheckResponse_ACTIVE
	case committee.INACTIVE:
		status = api.HealthCheckResponse_INACTIVE
	}
	return &api.HealthCheckResponse{
		Status: status,
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *server) GetCandidates(ctx context.Context, request *api.GetCandidatesRequest) (*api.CandidateResponse, error) {
	return nil, ErrNotSupported
}

// GetCandidateByName returns the candidate details
func (s *server) GetCandidateByName(ctx context.Context, request *api.GetCandidateByNameRequest) (*api.Candidate, error) {
	return nil, ErrNotSupported
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *api.GetBucketsByCandidateRequest) (*api.BucketResponse, error) {
	return nil, ErrNotSupported
}

func (s *server) toBucketResponse(mintTime time.Time, buckets []*types.Bucket, offset uint32, limit uint32) *api.BucketResponse {
	// If limit is missing, return all buckets with indices starting from the offset
	if limit == uint32(0) {
		limit = math.MaxUint32
	}
	if int(offset+limit) >= len(buckets) {
		limit = uint32(len(buckets)) - offset
	}
	response := &api.BucketResponse{
		Buckets: make([]*api.Bucket, limit),
	}
	for i := uint32(0); i < limit; i++ {
		bucket := buckets[offset+i]
		response.Buckets[i] = &api.Bucket{
			Voter:         hex.EncodeToString(bucket.Voter()),
			Votes:         bucket.Amount().Text(10),
			WeightedVotes: types.CalcWeightedVotes(bucket, mintTime).Text(10),
		}
	}
	return response
}

// GetBuckets returns a list of buckets
func (s *server) GetBuckets(ctx context.Context, request *api.GetBucketsRequest) (*api.BucketResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	mintTime, buckets, err := s.nativeCommittee.RawDataByHeight(height)
	if err != nil {
		return nil, err
	}
	offset := request.Offset
	if int(offset) >= len(buckets) {
		return nil, errors.New("offset is out of range")
	}

	return s.toBucketResponse(mintTime, buckets, offset, request.Limit), nil
}

func (s *server) GetRawData(ctx context.Context, request *api.GetRawDataRequest) (*api.RawDataResponse, error) {
	return nil, ErrNotSupported
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "native_committee_server.yaml", "path of server config file")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	s, err := NewServer(&cfg)
	if err != nil {
		zap.L().Fatal("failed to create server", zap.Error(err))
	}
	if err := s.Start(context.Background()); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
	zap.L().Info("Service started")
	defer s.Stop(context.Background())
	select {}
}
