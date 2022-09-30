package votesync

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
)

type iotexClient struct {
	client iotexapi.APIServiceClient
}

func NewIoTeXClient(client iotexapi.APIServiceClient) *iotexClient {
	return &iotexClient{client: client}
}

func (ic *iotexClient) BlockTime(h uint64) (time.Time, error) {
	resp, err := ic.client.GetBlockMetas(context.Background(), &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: h, Count: 1,
			},
		},
	})
	if err != nil {
		return time.Now(), errors.Wrapf(err, "failed to fetch block meta %v", h)
	}
	bms := resp.GetBlkMetas()
	if len(bms) != 1 {
		return time.Now(), errors.Wrapf(err, "asked 1 block, but got none-1 value %v", h)
	}
	ts := bms[0].GetTimestamp()
	bt, err := ptypes.Timestamp(ts)
	if err != nil {
		return time.Now(), errors.Wrapf(err, "failed to parse timestamp in blockmeta %v", h)
	}
	return bt, nil
}

func (ic *iotexClient) Tip() (uint64, error) {
	response, err := ic.client.GetChainMeta(
		context.Background(),
		&iotexapi.GetChainMetaRequest{},
	)
	if err != nil {
		return 0, err
	}

	return response.ChainMeta.Height, nil
}
