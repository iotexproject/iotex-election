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
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-election/server"
)

func main() {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.InfoLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)
	var configPath string
	flag.StringVar(&configPath, "config", "server.yaml", "path of server config file")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		zap.L().Fatal("failed to load config file", zap.Error(err))
	}
	var config server.MixConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		zap.L().Fatal("failed to unmarshal config", zap.Error(err))
	}
	sm, err := server.NewServerMix(config)
	if err != nil {
		zap.L().Fatal("failed to create server", zap.Error(err))
	}
	zap.L().Info("New server mix created")
	if err := sm.Start(context.Background()); err != nil {
		zap.L().Fatal("failed to start ranking server", zap.Error(err))
	}
	zap.L().Info("Service started")
	defer sm.Stop(context.Background())
	select {}
}
