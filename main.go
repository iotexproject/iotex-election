// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"flag"
	"io/ioutil"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-election/server"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "server.yaml", "path of server config file")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var config server.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	rankingServer, err := server.NewServer(&config)
	if err != nil {
		zap.L().Fatal("failed to create server", zap.Error(err))
	}
	zap.L().Info("New server created")
	if err := rankingServer.Start(context.Background()); err != nil {
		zap.L().Fatal("failed to start ranking server", zap.Error(err))
	}
	zap.L().Info("Service started")
	defer rankingServer.Stop(context.Background())
	select {}
}
