// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"

	pb "github.com/iotexproject/iotex-election/pb/ranking"
	"github.com/iotexproject/iotex-election/server/ranking"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
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
	s := grpc.NewServer()
	pb.RegisterRankingServer(s, rankingServer)

	s.Serve(lis)
}
