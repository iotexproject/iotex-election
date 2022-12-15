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
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/iotexproject/iotex-address/address"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/iotexproject/iotex-election/committee"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/pb/api"
	electionpb "github.com/iotexproject/iotex-election/pb/election"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
	"github.com/iotexproject/iotex-election/votesync"
)

// Config defines the config for server
type Config struct {
	DB                   db.Config        `yaml:"db"`
	Port                 int              `yaml:"port"`
	HttpPort             int              `yaml:"httpPort"`
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
	api.UnimplementedAPIServiceServer
	port                 int
	electionCommittee    committee.Committee
	grpcServer           *grpc.Server
	selfStakingThreshold *big.Int
	scoreThreshold       *big.Int
	vs                   *votesync.VoteSync
}

// NewServer returns an implementation of ranking server
func NewServer(cfg *Config, vs *votesync.VoteSync) (Server, error) {
	archive, err := committee.NewArchive(cfg.DB.DBPath, cfg.DB.NumOfRetries, cfg.Committee.GravityChainStartHeight, cfg.Committee.GravityChainHeightInterval)
	if err != nil {
		return nil, err
	}
	c, err := committee.NewCommittee(archive, cfg.Committee)
	if err != nil {
		return nil, err
	}
	scoreThreshold, ok := new(big.Int).SetString(cfg.ScoreThreshold, 10)
	if !ok {
		return nil, errors.New("invalid score threshold")
	}
	selfStakingThreshold, ok := new(big.Int).SetString(cfg.SelfStakingThreshold, 10)
	if !ok {
		return nil, errors.New("invalid self staking threshold")
	}
	s := &server{
		electionCommittee:    c,
		port:                 cfg.Port,
		scoreThreshold:       scoreThreshold,
		selfStakingThreshold: selfStakingThreshold,
		vs:                   vs,
	}
	s.grpcServer = grpc.NewServer()
	api.RegisterAPIServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)
	if cfg.HttpPort > 0 {
		go func() {
			gwmux := runtime.NewServeMux()
			if err := api.RegisterAPIServiceHandlerServer(context.Background(), gwmux, s); err != nil {
				zap.L().Panic("failed to register api server")
			}
			gwServer := &http.Server{
				Addr:    fmt.Sprintf(":%d", cfg.HttpPort),
				Handler: gwmux,
			}
			if err := gwServer.ListenAndServe(); err != nil {
				zap.L().Panic("failed to servert api gateway server", zap.Error(err))
			}
		}()
	}

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
		if s.vs != nil {
			s.vs.Start(ctx)
		}
	}()
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			zap.L().Fatal("Failed to serve", zap.Error(err))
		}
	}()
	return s.electionCommittee.Start(ctx)
}

func (s *server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	if s.vs != nil {
		s.vs.Stop(ctx)
	}
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
	// If limit is missing, return all candidates with indices starting from the offset
	if limit == uint32(0) {
		limit = math.MaxUint32
	}
	if len(candidates) < int(offset+limit) {
		limit = uint32(len(candidates)) - offset
	}
	response := &api.CandidateResponse{
		Candidates: make([]*api.Candidate, limit),
	}
	for i := uint32(0); i < limit; i++ {
		candidate := candidates[offset+i]
		var ra string
		var oa string
		if util.IsAllZeros(candidate.RewardAddress()) {
			ra = ""
		} else {
			ra = string(candidate.RewardAddress())
		}
		if util.IsAllZeros(candidate.OperatorAddress()) {
			oa = ""
		} else {
			oa = string(candidate.OperatorAddress())
		}
		response.Candidates[i] = &api.Candidate{
			Name:               hex.EncodeToString(candidate.Name()),
			Address:            hex.EncodeToString(candidate.Address()),
			RewardAddress:      ra,
			OperatorAddress:    oa,
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
	var votes []*types.Vote
	switch len(name) {
	case 0:
		votes = result.Votes()
	case 12:
		votes = result.VotesByDelegate(name)
	default:
		return nil, errors.New("invalid candidate name")
	}
	if votes == nil {
		return nil, errors.New("no buckets for the candidate")
	}
	offset := request.Offset
	if int(offset) >= len(votes) {
		return nil, errors.New("offset is out of range")
	}

	return s.toBucketResponse(votes, offset, request.Limit, result.MintTime()), nil
}

func (s *server) toBucketResponse(votes []*types.Vote, offset uint32, limit uint32, mintTime time.Time) *api.BucketResponse {
	// If limit is missing, return all buckets with indices starting from the offset
	if limit == uint32(0) {
		limit = math.MaxUint32
	}
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
			RemainingDuration: vote.RemainingTime(mintTime).String(),
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
	result, err := s.electionCommittee.ResultByHeight(height)
	if err != nil {
		return nil, err
	}
	votes := result.Votes()
	if votes == nil {
		return nil, errors.New("no buckets available")
	}
	offset := request.Offset
	if int(offset) >= len(votes) {
		return nil, errors.New("offset is out of range")
	}

	return s.toBucketResponse(votes, offset, request.Limit, result.MintTime()), nil
}

func (s *server) GetRawData(ctx context.Context, request *api.GetRawDataRequest) (*api.RawDataResponse, error) {
	height, err := strconv.ParseUint(request.Height, 10, 64)
	if err != nil {
		return nil, err
	}
	buckets, regs, timestamp, err := s.electionCommittee.RawDataByHeight(height)
	if err != nil {
		return nil, err
	}
	return s.toRawDataResponse(timestamp, regs, buckets)
}

func (s *server) toRawDataResponse(mintTime time.Time, regs []*types.Registration, buckets []*types.Bucket) (*api.RawDataResponse, error) {
	response := &api.RawDataResponse{
		Buckets:       make([]*electionpb.Bucket, len(buckets)),
		Registrations: make([]*electionpb.Registration, len(regs)),
	}
	for i := 0; i < len(buckets); i++ {
		bucketPb, err := buckets[i].ToProtoMsg()
		if err != nil {
			return nil, err
		}
		response.Buckets[i] = bucketPb
	}

	for i := 0; i < len(regs); i++ {
		regPb, err := regs[i].ToProtoMsg()
		if err != nil {
			return nil, err
		}
		response.Registrations[i] = regPb
	}
	t, err := ptypes.TimestampProto(mintTime)
	if err != nil {
		return nil, err
	}
	response.Timestamp = t
	return response, nil
}

func (s *server) GetProof(ctx context.Context, request *api.ProofRequest) (*api.ProofResponse, error) {
	if s.vs == nil {
		return nil, errors.New("no vote sync server")
	}
	addr, err := address.FromString(request.Account)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to cast address %s", request.Account)
	}
	cycle, amount, proof, err := s.vs.ProofForAccount(addr)
	if err != nil {
		return nil, err
	}
	if proof == nil {
		return nil, nil
	}

	return &api.ProofResponse{
		Deadline: cycle.String(),
		Amount:   amount.String(),
		Proof:    hex.EncodeToString(proof),
	}, nil
}
