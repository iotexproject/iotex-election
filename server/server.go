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
	"errors"
	"log"
	"math/big"
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
	"github.com/iotexproject/iotex-election/pb/api"
)

// Config defines the config for server
type Config struct {
	DB                   db.Config        `yaml:"db"`
	Port                 int              `yaml:"port"`
	Committee            committee.Config `yaml:"committee"`
	SelfStakingThreshold string           `yaml:"selfStakingThreshold"`
	ScoreThreshold       string           `yaml:"scoreThreshold"`
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
	electionCommittee    committee.Committee
	grpcServer           *grpc.Server
	selfStakingThreshold *big.Int
	scoreThreshold       *big.Int
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
	scoreThreshold, ok := new(big.Int).SetString(cfg.ScoreThreshold, 10)
	if !ok {
		return nil, errors.New("Invalid score threshold")
	}
	selfStakingThreshold, ok := new(big.Int).SetString(cfg.SelfStakingThreshold, 10)
	if !ok {
		return nil, errors.New("Invalid self staking threshold")
	}
	s := &server{
		electionCommittee:    c,
		port:                 cfg.Port,
		scoreThreshold:       scoreThreshold,
		selfStakingThreshold: selfStakingThreshold,
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
	if err := s.electionCommittee.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	return s.electionCommittee.Stop(ctx)
}

// GetMeta returns the meta of the chain
func (s *server) GetMeta(ctx context.Context, empty *empty.Empty) (*api.ChainMeta, error) {
	height := s.electionCommittee.LatestHeight()
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return &api.ChainMeta{}, err
	}
	numOfCandidates := uint64(0)
	for _, d := range result.Delegates() {
		if d.Score().Cmp(s.scoreThreshold) >= 0 && d.SelfStakingTokens().Cmp(s.selfStakingThreshold) >= 0 {
			numOfCandidates++
		}
	}

	return &api.ChainMeta{
		Height:           strconv.FormatUint(height, 10),
		TotalCandidates:  numOfCandidates,
		TotalVotedStakes: result.TotalVotedStakes().Text(10),
		TotalVotes:       result.TotalVotes().Text(10),
	}, nil
}

func (s *server) IsHealth(ctx context.Context, empty *empty.Empty) (*api.HealthCheckResponse, error) {
	var status api.HealthCheckResponse_Status
	switch s.electionCommittee.Status() {
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
	response := &api.CandidateResponse{
		Candidates: make([]*api.Candidate, limit),
	}
	for i := uint32(0); i < limit; i++ {
		candidate := candidates[offset+i]
		response.Candidates[i] = &api.Candidate{
			Name:               hex.EncodeToString(candidate.Name()),
			Address:            hex.EncodeToString(candidate.Address()),
			TotalWeightedVotes: candidate.Score().Text(10),
			SelfStakingTokens:  candidate.SelfStakingTokens().Text(10),
		}
	}

	return response, nil
}

// GetCandidateByName returns the candidate details
func (s *server) GetCandidateByName(ctx context.Context, request *api.GetCandidateByNameRequest) (*api.Candidate, error) {
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
	candidate := result.DelegateByName(name)
	if candidate == nil {
		return nil, errors.New("Cannot find candidate details")
	}
	return &api.Candidate{
		Address:            hex.EncodeToString(candidate.Address()),
		Name:               request.Name,
		OperatorAddress:    string(candidate.OperatorAddress()),
		RewardAddress:      string(candidate.RewardAddress()),
		TotalWeightedVotes: candidate.Score().String(),
		SelfStakingTokens:  candidate.SelfStakingTokens().String(),
	}, nil
}

// GetBucketsByCandidate returns the buckets
func (s *server) GetBucketsByCandidate(ctx context.Context, request *api.GetBucketsByCandidateRequest) (*api.BucketResponse, error) {
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
	response := &api.BucketResponse{
		Buckets: make([]*api.Bucket, limit),
	}
	for i := uint32(0); i < limit; i++ {
		vote := votes[offset+i]
		response.Buckets[i] = &api.Bucket{
			Voter:             hex.EncodeToString(vote.Voter()),
			Votes:             vote.Amount().Text(10),
			WeightedVotes:     vote.WeightedAmount().Text(10),
			RemainingDuration: vote.RemainingTime(result.MintTime()).String(),
		}
	}

	return response, nil
}
