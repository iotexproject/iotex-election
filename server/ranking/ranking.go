// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package ranking

import (
	"encoding/hex"
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/db"
	pb "github.com/iotexproject/iotex-election/pb/ranking"
)

// Config defines the config for server
type Config struct {
	DB        db.Config        `yaml:"db"`
	Port      int              `yaml:"port"`
	Committee committee.Config `yaml:"committee"`
}

// Server defines the interface of the ranking server implementation
type Server interface {
	pb.RankingServer
	Start(context.Context) error
	Stop(context.Context) error
}

// server is used to implement pb.RankingServer.
type server struct {
	port              int
	electionCommittee committee.Committee
	grpcServer        *grpc.Server
}

// NewServer returns an implementation of ranking server
func NewServer(cfg *Config) (Server, error) {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.InfoLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)

	var c committee.Committee
	if cfg.DB.DBPath != "" {
		c, err = committee.NewCommitteeWithKVStoreWithNamespace(db.NewBoltDB(cfg.DB), cfg.Committee)
	} else {
		c, err = committee.NewCommittee(db.NewInMemKVStore(), cfg.Committee)
	}
	if err != nil {
		return nil, err
	}
	s := &server{electionCommittee: c, port: cfg.Port}
	s.grpcServer = grpc.NewServer()
	pb.RegisterRankingServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)

	return s, nil
}

func (s *server) Start(ctx context.Context) error {
	zap.L().Info("Start ranking server")
	if err := s.electionCommittee.Start(ctx); err != nil {
		return err
	}
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
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	return s.electionCommittee.Stop(ctx)
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *empty.Empty) (*pb.ChainMeta, error) {
	height := s.electionCommittee.LatestHeight()
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return &pb.ChainMeta{}, err
	}

	return &pb.ChainMeta{
		Height:           strconv.FormatUint(height, 10),
		TotalCandidates:  uint64(len(result.Delegates())),
		TotalVotedStakes: result.TotalVotedStakes().Text(10),
		TotalVotes:       result.TotalVotes().Text(10),
	}, nil
}

// GetCandidates returns a list of candidates sorted by weighted votes
func (s *server) GetCandidates(ctx context.Context, request *pb.GetCandidatesRequest) (*pb.CandidateResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	candidates := result.Delegates()
	offset := request.Offset
	if len(candidates) <= int(offset) {
		return nil, errors.New("offset is larger than candidate length")
	}
	limit := request.Limit
	if len(candidates) < int(offset+limit) {
		limit = uint32(len(candidates)) - offset
	}
	response := &pb.CandidateResponse{
		Candidates: make([]*pb.Candidate, limit),
	}
	for i := uint32(0); i < limit; i++ {
		candidate := candidates[offset+i]
		response.Candidates[i] = &pb.Candidate{
			Name:               hex.EncodeToString(candidate.Name()),
			Address:            hex.EncodeToString(candidate.Address()),
			TotalWeightedVotes: candidate.Score().Text(10),
		}
	}

	return response, nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *pb.GetBucketsByCandidateRequest) (*pb.BucketResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	name, err := hex.DecodeString(request.Name)
	if err != nil {
		return nil, err
	}
	if len(name) != 12 {
		return nil, errors.New("invalid candidate name")
	}
	votes := result.VotesByDelegate(name)
	if votes == nil {
		return nil, errors.New("No buckets for the candidate")
	}
	offset := request.Offset
	if int(offset) >= len(votes) {
		return nil, errors.New("offset is out of range")
	}
	limit := request.Limit
	if int(offset+limit) >= len(votes) {
		limit = uint32(len(votes)) - offset
	}
	response := &pb.BucketResponse{
		Buckets: make([]*pb.Bucket, limit),
	}
	for i := uint32(0); i < limit; i++ {
		vote := votes[offset+i]
		response.Buckets[i] = &pb.Bucket{
			Voter:             hex.EncodeToString(vote.Voter()),
			Votes:             vote.Amount().Text(10),
			WeightedVotes:     vote.WeightedAmount().Text(10),
			RemainingDuration: vote.RemainingTime(result.MintTime()).String(),
		}
	}

	return response, nil
}
