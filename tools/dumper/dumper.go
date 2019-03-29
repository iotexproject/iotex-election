// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-election/committee"
)

func main() {
	var configPath string
	var height uint64
	flag.StringVar(&configPath, "config", "committee.yaml", "path of server config file")
	flag.Uint64Var(&height, "height", 0, "ethereuem height")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
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
