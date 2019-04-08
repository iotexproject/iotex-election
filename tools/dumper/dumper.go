// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/csv"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-core/protogen/iotexapi"
	"github.com/iotexproject/iotex-election/committee"
)

func main() {
	var configPath string
	var epoch uint64
	var height uint64
	endpoint := "api.iotex.one:80"
	flag.StringVar(&configPath, "config", "committee.yaml", "path of server config file")
	flag.Uint64Var(&epoch, "epoch", 0, "iotex epoch")
	flag.Uint64Var(&height, "height", 0, "ethereuem height")
	flag.Parse()
	data, err := ioutil.ReadFile(configPath)
	zap.L().Fatal("failed to load config file", zap.Error(err))
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var config committee.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	committee, err := committee.NewCommittee(nil, config)
	if err != nil {
		zap.L().Fatal("failed to create committee", zap.Error(err))
	}
	if epoch != 0 {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			zap.L().Fatal("failed to connect endpoint", zap.Error(err))
		}
		defer conn.Close()
		cli := iotexapi.NewAPIServiceClient(conn)
		request := iotexapi.GetEpochMetaRequest{EpochNumber: epoch}
		ctx := context.Background()
		response, err := cli.GetEpochMeta(ctx, &request)
		if err != nil {
			zap.L().Fatal("failed to get epoch meta", zap.Error(err))
		}
		height = response.EpochData.GravityChainStartHeight
	}
	result, err := committee.FetchResultByHeight(height)
	if err != nil {
		zap.L().Fatal("failed to fetch result", zap.Uint64("height", height))
	}
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{
		"voter",
		"startTime",
		"duration",
		"decay",
		"tokens",
		"votes",
		"votee",
		"voterIoAddr",
	})
	for _, delegate := range result.Delegates() {
		for _, vote := range result.VotesByDelegate(delegate.Name()) {
			ioAddr, _ := address.FromBytes(vote.Voter())
			ioAddrStr := ioAddr.String()
			if err := writer.Write([]string{
				hex.EncodeToString(vote.Voter()),
				vote.StartTime().String(),
				vote.Duration().String(),
				strconv.FormatBool(vote.Decay()),
				vote.Amount().String(),
				vote.WeightedAmount().String(),
				string(vote.Candidate()),
				ioAddrStr,
			}); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}
	}
	writer.Flush()
}

func init() {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.InfoLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)
}
