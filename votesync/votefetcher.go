package votesync

import (
	"context"
	"strconv"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

// VoteFetcher defines fields used to fetch votes
type VoteFetcher struct {
	iotexAPI iotexapi.APIServiceClient
}

func (fetcher *VoteFetcher) FetchBucketsByHeight(ctx context.Context, h uint64) (*iotextypes.VoteBucketList, *iotextypes.CandidateListV2, error) {
	var (
		ret1 *iotextypes.VoteBucketList
		ret2 *iotextypes.CandidateListV2
		err  error
	)
	nerr := backoff.Retry(func() error {
		ret1, err = getAllStakingBuckets(ctx, fetcher.iotexAPI, h)
		if err != nil {
			return err
		}
		ret2, err = getAllStakingCandidates(ctx, fetcher.iotexAPI, h)
		return err
	}, backoff.NewExponentialBackOff())
	if nerr != nil {
		zap.L().Error(
			"failed to fetch vote result by height",
			zap.Error(nerr),
			zap.Uint64("height", h),
		)
	}
	return ret1, ret2, nerr
}

func getAllStakingBuckets(ctx context.Context, chainClient iotexapi.APIServiceClient, height uint64) (voteBucketListAll *iotextypes.VoteBucketList, err error) {
	voteBucketListAll = &iotextypes.VoteBucketList{}
	for i := uint32(0); ; i++ {
		offset := i * readBucketsLimit
		size := uint32(readBucketsLimit)
		voteBucketList, err := getStakingBuckets(ctx, chainClient, offset, size, height)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get bucket")
		}
		voteBucketListAll.Buckets = append(voteBucketListAll.Buckets, voteBucketList.Buckets...)
		if len(voteBucketList.Buckets) < readBucketsLimit {
			break
		}
	}
	return
}

func getAllStakingCandidates(ctx context.Context, chainClient iotexapi.APIServiceClient, height uint64) (candidateListAll *iotextypes.CandidateListV2, err error) {
	candidateListAll = &iotextypes.CandidateListV2{}
	for i := uint32(0); ; i++ {
		offset := i * readCandidatesLimit
		size := uint32(readCandidatesLimit)
		candidateList, err := getStakingCandidates(ctx, chainClient, offset, size, height)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get candidates")
		}
		candidateListAll.Candidates = append(candidateListAll.Candidates, candidateList.Candidates...)
		if len(candidateList.Candidates) < readCandidatesLimit {
			break
		}
	}
	return
}

// getStakingBuckets get specific buckets by height
func getStakingBuckets(ctx context.Context, chainClient iotexapi.APIServiceClient, offset, limit uint32, height uint64) (voteBucketList *iotextypes.VoteBucketList, err error) {
	methodName, err := proto.Marshal(&iotexapi.ReadStakingDataMethod{
		Method: iotexapi.ReadStakingDataMethod_BUCKETS,
	})
	if err != nil {
		return nil, err
	}
	arg, err := proto.Marshal(&iotexapi.ReadStakingDataRequest{
		Request: &iotexapi.ReadStakingDataRequest_Buckets{
			Buckets: &iotexapi.ReadStakingDataRequest_VoteBuckets{
				Pagination: &iotexapi.PaginationParam{
					Offset: offset,
					Limit:  limit,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	readStateRequest := &iotexapi.ReadStateRequest{
		ProtocolID: []byte(protocolID),
		MethodName: methodName,
		Arguments:  [][]byte{arg},
		Height:     strconv.FormatUint(height, 10),
	}
	readStateRes, err := chainClient.ReadState(ctx, readStateRequest)
	if err != nil {
		return
	}
	voteBucketList = &iotextypes.VoteBucketList{}
	if err := proto.Unmarshal(readStateRes.GetData(), voteBucketList); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal VoteBucketList")
	}
	return
}

// getStakingCandidates get specific candidates by height
func getStakingCandidates(ctx context.Context, chainClient iotexapi.APIServiceClient, offset, limit uint32, height uint64) (candidateList *iotextypes.CandidateListV2, err error) {
	methodName, err := proto.Marshal(&iotexapi.ReadStakingDataMethod{
		Method: iotexapi.ReadStakingDataMethod_CANDIDATES,
	})
	if err != nil {
		return nil, err
	}
	arg, err := proto.Marshal(&iotexapi.ReadStakingDataRequest{
		Request: &iotexapi.ReadStakingDataRequest_Candidates_{
			Candidates: &iotexapi.ReadStakingDataRequest_Candidates{
				Pagination: &iotexapi.PaginationParam{
					Offset: offset,
					Limit:  limit,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	readStateRequest := &iotexapi.ReadStateRequest{
		ProtocolID: []byte(protocolID),
		MethodName: methodName,
		Arguments:  [][]byte{arg},
		Height:     strconv.FormatUint(height, 10),
	}
	readStateRes, err := chainClient.ReadState(ctx, readStateRequest)
	if err != nil {
		return
	}
	candidateList = &iotextypes.CandidateListV2{}
	if err := proto.Unmarshal(readStateRes.GetData(), candidateList); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal VoteBucketList")
	}
	return
}
