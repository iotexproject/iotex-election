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
	"log"

	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-election/server/ranking"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "server.yaml", "path of server config file")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("failed to load config file %v", err)
	}
	var config ranking.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("failed to unmarshal config %v", err)
	}
	rankingServer, err := ranking.NewServer(&config)
	if err != nil {
		log.Fatalf("failed to create ranking server %v", err)
	}
	if err := rankingServer.Start(context.Background()); err != nil {
		log.Fatalf("failed to start ranking server %v", err)
	}
	defer rankingServer.Stop(context.Background())
	select {}
}
