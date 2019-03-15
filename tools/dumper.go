// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/iotex-election/carrier"
	"github.com/iotexproject/iotex-election/committee"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "committee.yaml", "path of server config file")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var config committee.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	c, err := carrier.NewEthereumVoteCarrier(
		config.BeaconChainAPIs,
		common.HexToAddress(config.RegisterContractAddress),
		common.HexToAddress(config.StakingContractAddress),
	)
	if err != nil {
		zap.L().Fatal("failed create carrier", zap.Error(err))
	}
	ecarrier, ok := c.(*carrier.EthereumCarrier)
	if !ok {
		zap.L().Fatal("failed to cast ethereum vote carrier")
	}
	height, err := ecarrier.TipHeight()
	if err != nil {
		zap.L().Fatal("failed to get tip height", zap.Error(err))
	}
	writer := csv.NewWriter(os.Stdout)
	previousIndex := big.NewInt(0)
	for {
		result, err := ecarrier.Buckets(
			&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(height)},
			previousIndex,
			big.NewInt(int64(config.PaginationSize)),
		)
		if err != nil {
			zap.L().Fatal("failed to fetch votes", zap.Error(err))
		}
		if result.Count == nil || result.Count.Cmp(big.NewInt(0)) == 0 || len(result.Indexes) == 0 {
			break
		}
		for i, index := range result.Indexes {
			if big.NewInt(0).Cmp(index) == 0 { // back to start, this is a redundant condition
				break
			}
			duration := time.Duration(result.StakeDurations[i].Uint64()*24) * time.Hour
			record := []string{
				index.String(),
				time.Unix(result.StakeStartTimes[i].Int64(), 0).String(),
				duration.String(),
				result.StakedAmounts[i].String(),
				string(result.CanNames[i][:]),
				result.Owners[i].String(),
				strconv.FormatBool(result.Decays[i]),
			}
			if err := writer.Write(record); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
			if index.Cmp(previousIndex) > 0 {
				previousIndex = index
			}
		}
	}
	writer.Flush()
}
