package server

import (
	"github.com/iotexproject/iotex-election/votesync"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type ServerMix struct {
	ess Server
}

type MixConfig struct {
	ElectionConfig  Config              `yaml:"electionConfig"`
	NativeConfig    NativeStakingConfig `yaml:"nativeConfig"`
	VoteSync        votesync.Config     `yaml:"voteSync"`
	EnableVoteSync  bool                `yaml:"enableVoteSync"`
	DummyServerPort int                 `yaml:"dummyServerPort"`
}

func NewServerMix(mCfg MixConfig) (*ServerMix, error) {
	var err error
	var ess Server
	if mCfg.DummyServerPort != 0 {
		ess, err = NewDummyServer(mCfg.DummyServerPort)
		if err != nil {
			return nil, err
		}
		zap.L().Info("New dummy server created")
	} else {
		var vs *votesync.VoteSync
		if mCfg.EnableVoteSync {
			vs, err = votesync.NewVoteSync(mCfg.VoteSync)
			if err != nil {
				return nil, err
			}
		}
		ess, err = NewServer(&mCfg.ElectionConfig, vs)
		if err != nil {
			return nil, err
		}
	}

	return &ServerMix{ess: ess}, nil
}

func (sm *ServerMix) Start(ctx context.Context) error {
	return sm.ess.Start(ctx)
}

func (sm *ServerMix) Stop(ctx context.Context) error {
	return sm.ess.Stop(ctx)
}
