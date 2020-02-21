package server

import (
	"github.com/iotexproject/iotex-election/votesync"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type ServerMix struct {
	ess      Server
	nss      Server
	voteSync *votesync.VoteSync
}

type MixConfig struct {
	ElectionConfig  Config              `yaml:"electionConfig"`
	NativeConfig    NativeStakingConfig `yaml:"nativeConfig"`
	VoteSync        votesync.Config     `yaml:"voteSync"`
	EnableVoteSync  bool                `yaml:"enableVoteSync"`
	DummyServerPort int                 `yaml:"dummyServerPort"`
}

func NewServerMix(mCfg MixConfig) (*ServerMix, error) {
	var ess Server
	var err error
	if mCfg.DummyServerPort != 0 {
		ess, err = NewDummyServer(mCfg.DummyServerPort)
		if err != nil {
			return nil, err
		}
		zap.L().Info("New dummy server created")
	} else {
		ess, err = NewServer(&mCfg.ElectionConfig)
		if err != nil {
			return nil, err
		}
	}
	nss, err := NewNativeStakingServer(&mCfg.NativeConfig)
	if err != nil {
		return nil, err
	}
	var vs *votesync.VoteSync
	if mCfg.EnableVoteSync {
		vs, err = votesync.NewVoteSync(mCfg.VoteSync, nss.Committee())
		if err != nil {
			return nil, err
		}
	}

	return &ServerMix{
		ess:      ess,
		nss:      nss,
		voteSync: vs,
	}, nil
}

func (sm *ServerMix) Start(ctx context.Context) error {
	if err := sm.ess.Start(ctx); err != nil {
		return err
	}
	if err := sm.nss.Start(ctx); err != nil {
		return err
	}
	if sm.voteSync != nil {
		sm.voteSync.Start(ctx)
	}
	return nil
}

func (sm *ServerMix) Stop(ctx context.Context) error {
	sm.voteSync.Stop(ctx)
	if err := sm.nss.Stop(ctx); err != nil {
		return err
	}
	return sm.ess.Stop(ctx)
}
