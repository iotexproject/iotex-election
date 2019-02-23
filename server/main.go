// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"io/ioutil"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-election/server/ranking"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "server.yaml", "path of server config file")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var config ranking.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	rankingServer, err := ranking.NewServer(&config)
	if err != nil {
		zap.L().Fatal("failed to create ranking server", zap.Error(err))
	}
	if err := rankingServer.Start(context.Background()); err != nil {
		zap.L().Fatal("failed to start ranking server", zap.Error(err))
	}
	zap.L().Info("Service started")
	defer rankingServer.Stop(context.Background())
	select {}
}
