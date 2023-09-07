package votesync

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

const _viewIDOffsite = 10000000

const (
	// PollProtocolID is ID of poll protocol
	PollProtocolID      = "poll"
	protocolID          = "staking"
	readBucketsLimit    = 30000
	readCandidatesLimit = 20000
)

var _fixedAdhocPower = big.NewInt(1)

// VoteSync defines fields used in VoteSync
type VoteSync struct {
	agentContract          *agentContract
	votingPowers           *VotingPowers
	client                 *iotexClient
	fetcher                *VoteFetcher
	vpsContract            *rwvps
	brokerContract         *brokerContract
	clerkContract          *clerkContract
	discord                *discord
	lastViewHeight         uint64
	lastViewTimestamp      time.Time
	lastUpdateHeight       uint64
	lastBrokerUpdateHeight uint64
	lastClerkUpdateHeight  uint64
	lastUpdateTimestamp    time.Time
	timeInternal           time.Duration
	lastNativeEphoch       uint64
	tempLastNativeEphoch   uint64
	terminate              chan bool
	terminated             bool
	dardanellesHeight      uint64
	fairbankHeight         uint64
}

// Config defines the configs for VoteSync
type Config struct {
	GravityChainAPIs          []string      `yaml:"gravityChainAPIs"`
	GravityChainTimeInterval  time.Duration `yaml:"gravityChainTimeInterval"`
	OperatorPrivateKey        string        `yaml:"operatorPrivateKey"`
	IoTeXAPI                  string        `yaml:"ioTeXAPI"`
	IoTeXAPISecure            bool          `yaml:"ioTeXAPISecure"`
	RegisterContractAddress   string        `yaml:"registerContractAddress"`
	StakingContractAddress    string        `yaml:"stakingContractAddress"`
	PaginationSize            uint8         `yaml:"paginationSize"`
	BrokerPaginationSize      uint8         `yaml:"brokerPaginationSize"`
	VitaContractAddress       string        `yaml:"vitaContractAddress"`
	DiscordBotToken           string        `yaml:"discordBotToken"`
	DiscordChannelID          string        `yaml:"discordChannelID"`
	DiscordMsg                string        `yaml:"discordMsg"`
	DiscordReminder           string        `yaml:"discordReminder"`
	DardanellesHeight         uint64        `yaml:"dardanellesHeight"`
	FairBankHeight            uint64        `yaml:"fairbankHeight"`
	NativeCommitteeInitHeight uint64        `yaml:"nativeCommitteeInitHeight"`
	EnableAgentMode           bool          `yaml:"enableAgentMode"`
	AgentContractAddress      string        `yaml:"agentContractAddress"`
}

// WeightedVote defines voter and votes for weighted vote
type WeightedVote struct {
	Voter string
	Votes *big.Int
}

func ioToEthAddress(str string) (common.Address, error) {
	addr, err := address.FromString(str)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(addr.Bytes()), nil
}

// NewVoteSync instantiates new VoteSync
func NewVoteSync(cfg Config) (*VoteSync, error) {
	ctx := context.Background()

	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Second)),
		grpc_retry.WithMax(3),
	}
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	}
	if cfg.IoTeXAPISecure {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}
	conn, err := grpc.DialContext(ctx, cfg.IoTeXAPI, dialOpts...)
	if err != nil {
		return nil, err
	}
	operatorAccount, err := account.HexStringToAccount(cfg.OperatorPrivateKey)
	if err != nil {
		return nil, err
	}
	iotexAPI := iotexapi.NewAPIServiceClient(conn)
	apiClient := NewIoTeXClient(iotexAPI)
	fetcher := &VoteFetcher{iotexAPI: iotexAPI}
	authClient := iotex.NewAuthedClient(iotexAPI, 1, operatorAccount)

	vitaContractAddress, err := address.FromString(cfg.VitaContractAddress)
	if err != nil {
		return nil, err
	}
	vitaContract, err := NewVitaContract(authClient, vitaContractAddress)
	if err != nil {
		return nil, err
	}

	vpsContractAddress, err := vitaContract.VPS()
	if err != nil {
		return nil, err
	}

	brokerContractAddress, err := vitaContract.DonationPoolAddress()
	if err != nil {
		return nil, err
	}

	clerkContractAddress, err := vitaContract.RewardPoolAddress()
	if err != nil {
		return nil, err
	}
	zap.L().Info("vote contracts.", zap.String("brokerContract", brokerContractAddress.String()), zap.String("clerkContract", clerkContractAddress.String()))

	vpsContract, err := NewRotatableWeightedVPS(authClient, vpsContractAddress, cfg.PaginationSize)
	if err != nil {
		return nil, err
	}

	lastUpdateHeight, err := vpsContract.ViewID()
	if err != nil {
		return nil, err
	}
	if lastUpdateHeight.Uint64() == 0 {
		lastUpdateHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	}
	if lastUpdateHeight.Uint64() > _viewIDOffsite {
		lastUpdateHeight.Sub(lastUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}
	lastUpdateTimestamp, err := apiClient.BlockTime(lastUpdateHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastViewHeight, err := vpsContract.InactiveViewID()
	if err != nil {
		return nil, err
	}
	if lastViewHeight.Uint64() > _viewIDOffsite {
		lastViewHeight.Sub(lastViewHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}
	if lastViewHeight.Uint64() < cfg.FairBankHeight {
		lastViewHeight = new(big.Int).SetUint64(cfg.FairBankHeight)
	}
	lastViewTimestamp, err := apiClient.BlockTime(lastViewHeight.Uint64())
	if err != nil {
		return nil, err
	}

	lastBrokerUpdateHeight, err := vitaContract.LastDonationPoolClaimViewID()
	if err != nil {
		return nil, err
	}
	if lastBrokerUpdateHeight.Uint64() > _viewIDOffsite {
		lastBrokerUpdateHeight.Sub(lastBrokerUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}

	lastClerkUpdateHeight, err := vitaContract.LastRewardPoolClaimViewID()
	if err != nil {
		return nil, err
	}
	if lastClerkUpdateHeight.Uint64() > _viewIDOffsite {
		lastClerkUpdateHeight.Sub(lastClerkUpdateHeight, new(big.Int).SetUint64(_viewIDOffsite))
	}

	brokerContract, err := NewBrokerContract(authClient, brokerContractAddress, cfg.BrokerPaginationSize)
	if err != nil {
		return nil, err
	}

	clerkContract, err := NewClerkContract(authClient, clerkContractAddress)
	if err != nil {
		return nil, err
	}

	var agentContract *agentContract
	if cfg.EnableAgentMode {
		agentContractAddress, err := address.FromString(cfg.AgentContractAddress)
		if err != nil {
			return nil, err
		}
		agentContract, err = NewAgentContract(authClient, agentContractAddress)
		if err != nil {
			return nil, err
		}
	}

	return &VoteSync{
		client:                 apiClient,
		agentContract:          agentContract,
		votingPowers:           &VotingPowers{},
		fetcher:                fetcher,
		vpsContract:            vpsContract,
		brokerContract:         brokerContract,
		clerkContract:          clerkContract,
		timeInternal:           cfg.GravityChainTimeInterval,
		lastViewHeight:         lastViewHeight.Uint64(),
		lastViewTimestamp:      lastViewTimestamp,
		lastUpdateHeight:       lastUpdateHeight.Uint64(),
		lastUpdateTimestamp:    lastUpdateTimestamp,
		lastBrokerUpdateHeight: lastBrokerUpdateHeight.Uint64(),
		lastClerkUpdateHeight:  lastClerkUpdateHeight.Uint64(),
		terminate:              make(chan bool),
		terminated:             false,
		discord: &discord{
			botToken:    cfg.DiscordBotToken,
			channelID:   cfg.DiscordChannelID,
			newCycleMsg: cfg.DiscordMsg,
			reminderMsg: cfg.DiscordReminder,
		},
		dardanellesHeight: cfg.DardanellesHeight,
		fairbankHeight:    cfg.FairBankHeight,
	}, nil
}

// Start starts voteSync
func (vc *VoteSync) Start(ctx context.Context) {
	errChan := make(chan error)

	zap.L().Info("Start VoteSync.",
		zap.Uint64("lastUpdateHeight", vc.lastUpdateHeight),
		zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight),
		zap.Uint64("lastClerkUpdateHeight", vc.lastClerkUpdateHeight),
		zap.Uint64("lastViewID", vc.lastViewHeight),
	)
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-vc.terminate:
				return
			case <-ticker.C:
				tip, err := vc.client.Tip()
				if err != nil {
					zap.L().Error("failed to get iotex tip", zap.Error(err))
					continue
				}
				blockTime, err := vc.client.BlockTime(tip)
				if err != nil {
					zap.L().Error("failed to get block time", zap.Error(err))
					continue
				}
				if blockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal)) {
					if err := vc.sync(ctx, vc.lastViewHeight, tip, blockTime); err != nil {
						zap.L().Error("failed to sync votes", zap.Error(err))
						continue
					}
					if err := vc.discord.SendNewCycleMessage(); err != nil {
						zap.L().Error("failed to send discord msg", zap.Error(err))
					}
				}
				if blockTime.After(vc.lastUpdateTimestamp.Add(vc.timeInternal*24/25)) && !vc.discord.Reminded() {
					if err := vc.discord.SendReminder(); err != nil {
						zap.L().Error("failed to send discord reminder", zap.Error(err))
					}
				}

				if vc.lastUpdateHeight > vc.lastBrokerUpdateHeight {
					if err := vc.settle(vc.lastUpdateHeight); err != nil {
						zap.L().Error("failed to settle broker", zap.Error(err))
					}
				}

				if vc.lastUpdateHeight > vc.lastClerkUpdateHeight {
					if err := vc.claimForClerk(); err != nil {
						zap.L().Error("failed to claim for clerk", zap.Error(err))
					}
				}
			case err := <-errChan:
				zap.L().Error("something goes wrong", zap.Error(err))
			}
		}
	}()
}

// Stop stops voteSync
func (vc *VoteSync) Stop(ctx context.Context) {
	if vc.terminated {
		return
	}
	close(vc.terminate)
	vc.terminated = true
}

func (vc *VoteSync) ProofForAccount(acct address.Address) (*big.Int, *big.Int, []byte, error) {
	if vc.agentContract == nil {
		return nil, nil, nil, errors.New("agent mode is not enabled")
	}
	agentCycle, size, claimed, err := vc.agentContract.Claimed(acct)
	if err != nil {
		return nil, nil, nil, err
	}
	if claimed {
		return nil, nil, nil, nil
	}
	cycle, power := vc.votingPowers.VotingPower(common.BytesToAddress(acct.Bytes()))
	if cycle.Cmp(agentCycle) != 0 {
		return nil, nil, nil, errors.New("unexpected status")
	}
	amount := new(big.Int).Div(power.Mul(power, size), vc.votingPowers.Total())
	proof, err := vc.agentContract.Digest(amount, cycle)
	if err != nil {
		return nil, nil, nil, err
	}

	return agentCycle, amount, proof, nil
}

func (vc *VoteSync) brokerSettle() error {
	oldStart := uint64(0)
	for {
		if err := vc.brokerContract.Settle(); err != nil {
			return err
		}

		newStart, err := vc.brokerContract.NextBidToSettle()
		if err != nil {
			return err
		}
		if oldStart == newStart {
			return nil
		}
		oldStart = newStart
	}
}

func (vc *VoteSync) settle(h uint64) error {
	l := zap.L().With(zap.Uint64("lastBrokerUpdateHeight", vc.lastBrokerUpdateHeight))
	l.Info("Start broker settle process.")
	// settle broker
	if err := vc.brokerSettle(); err != nil {
		return errors.Wrap(err, "broker settle error")
	}
	l.Info("Finished broker settle.")
	if err := vc.brokerContract.Reset(); err != nil {
		return errors.Wrap(err, "broker reset error")
	}
	vc.lastBrokerUpdateHeight = h
	l.Info("Finished broker reset.", zap.Uint64("brokerUpdatedHeight", h))
	return nil
}

func (vc *VoteSync) claimForClerk() error {
	zap.L().Info("Start clerk claim process.", zap.Uint64("lastClerkUpdateHeight", vc.lastClerkUpdateHeight))
	if err := vc.clerkContract.Claim(); err != nil {
		return err
	}
	vc.lastClerkUpdateHeight = vc.lastUpdateHeight
	zap.L().Info("Finished clerk.", zap.Uint64("cleerkUpdatedHeight", vc.lastUpdateHeight))
	return nil
}

func (vc *VoteSync) sync(ctx context.Context, prevHeight, currHeight uint64, currTs time.Time) error {
	zap.L().Info("Start VoteSyncing.", zap.Uint64("lastViewID", prevHeight), zap.Uint64("nextViewID", currHeight))
	if vc.agentContract != nil {
		buckets, candidates, err := vc.fetcher.FetchBucketsByHeight(ctx, currHeight)
		if err != nil {
			return err
		}
		totalVotes := big.NewInt(0)
		votingPowers := make(map[common.Address]*big.Int)
		votes := calWeightedVotes(buckets, candidates)
		for _, vote := range votes {
			totalVotes = totalVotes.Add(totalVotes, vote.Votes)
			addr, err := ioToEthAddress(vote.Voter)
			if err != nil {
				return err
			}
			votingPowers[addr] = vote.Votes
		}
		totalPower, err := vc.vpsContract.TotalPower()
		if err != nil {
			return err
		}
		if totalPower.Cmp(_fixedAdhocPower) != 0 {
			votes, err := vc.vpsContract.VoterPowers()
			if err != nil {
				return err
			}
			voters := make([]common.Address, 0, len(votes)+1)
			weights := make([]*big.Int, 0, len(votes)+1)
			for voter := range votes {
				voters = append(voters, voter)
				weights = append(weights, big.NewInt(0))
			}
			voters = append(voters, vc.agentContract.Address())
			weights = append(weights, _fixedAdhocPower)
			if err := vc.vpsContract.UpdateVotingPowers(voters, weights); err != nil {
				return err
			}
		}
		vc.votingPowers.Update(nil, totalVotes, votingPowers)
	} else {
		ret, err := vc.fetchVotesUpdate(ctx, prevHeight, currHeight)
		if err != nil {
			return errors.Wrap(err, "fetch vote error")
		}
		zap.L().Info("Need to sync.", zap.Int("numVoter", len(ret)))

		var (
			addrs   []common.Address
			weights []*big.Int
		)
		for _, vote := range ret {
			addr, err := ioToEthAddress(vote.Voter)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed convert address:%s", vote.Voter))
			}
			addrs = append(addrs, addr)
			weights = append(weights, vote.Votes)
		}
		if err := vc.vpsContract.UpdateVotingPowers(addrs, weights); err != nil {
			return errors.Wrap(err, "update vote error")
		}
	}

	if err := vc.vpsContract.Rotate(new(big.Int).SetUint64(currHeight + _viewIDOffsite)); err != nil {
		return errors.Wrap(err, "failed to execute rotate")
	}

	vc.lastViewHeight = vc.lastUpdateHeight
	vc.lastViewTimestamp = vc.lastUpdateTimestamp
	vc.lastUpdateHeight = currHeight
	vc.lastUpdateTimestamp = currTs
	vc.lastNativeEphoch = vc.tempLastNativeEphoch
	zap.L().Info("Successfully synced votes.", zap.Uint64("lastViewID", vc.lastViewHeight), zap.Uint64("viewID", currHeight))
	return nil
}

func (vc *VoteSync) fetchVotesUpdate(ctx context.Context, prevHeight, currHeight uint64) ([]*WeightedVote, error) {
	// prevHeight == 0, only run at first 2 time. get all votes from currHeight
	if prevHeight == 0 {
		currB, currC, err := vc.fetcher.FetchBucketsByHeight(ctx, currHeight)
		if err != nil {
			return nil, err
		}
		n := calWeightedVotes(currB, currC)
		var ret []*WeightedVote
		for _, nv := range n {
			ret = append(ret, nv)
		}
		return ret, nil
	}

	prevB, prevC, err := vc.fetcher.FetchBucketsByHeight(ctx, prevHeight)
	if err != nil {
		return nil, err
	}

	currB, currC, err := vc.fetcher.FetchBucketsByHeight(ctx, currHeight)
	if err != nil {
		return nil, err
	}

	p := calWeightedVotes(prevB, prevC)
	n := calWeightedVotes(currB, currC)

	var ret []*WeightedVote
	// check for all voters in old view
	// if they don't exist in new view map, append 0 value for them
	// if they do exist in new view map, append only if the vote value is different
	for k, pv := range p {
		nv, ok := n[k]
		if !ok {
			ret = append(ret, &WeightedVote{
				Voter: pv.Voter,
				Votes: big.NewInt(0),
			})
		} else {
			if nv.Votes.Cmp(pv.Votes) != 0 {
				ret = append(ret, nv)
			}
		}
	}
	// check for all new voters in new view
	// if they don't exist in old view map, append
	for k, nv := range n {
		if _, ok := p[k]; !ok {
			ret = append(ret, nv)
		}
	}
	return ret, nil
}

func calWeightedVotes(bs *iotextypes.VoteBucketList, cs *iotextypes.CandidateListV2) map[string]*WeightedVote {
	n := make(map[string]*WeightedVote)
	for _, v := range bs.GetBuckets() {
		var selfStake bool
		for _, c := range cs.GetCandidates() {
			if c.GetSelfStakeBucketIdx() == v.GetIndex() {
				selfStake = true
				break
			}
		}

		vs := calculateVoteWeight(v, selfStake)
		wv, ok := n[v.GetOwner()]
		if ok {
			wv.Votes.Add(wv.Votes, vs)
			continue
		}
		n[v.GetOwner()] = &WeightedVote{
			Voter: v.GetOwner(),
			Votes: vs,
		}
	}
	return n
}

func calculateVoteWeight(v *iotextypes.VoteBucket, selfStake bool) *big.Int {
	durationLg := 1.2
	autoStake := 1.0
	selfStakeR := 1.06
	remainingTime := float64(v.StakedDuration)
	weight := float64(1)
	var m float64
	if v.AutoStake {
		m = autoStake
	}
	if remainingTime > 0 {
		weight += math.Log(math.Ceil(remainingTime/86400)*(1+m)) / math.Log(durationLg) / 100
	}
	if selfStake && v.AutoStake && v.StakedDuration >= 91 {
		// self-stake extra bonus requires enable auto-stake for at least 3 months
		weight *= selfStakeR
	}

	a, _ := new(big.Int).SetString(v.StakedAmount, 10)
	amount := new(big.Float).SetInt(a)
	weightedAmount, _ := amount.Mul(amount, big.NewFloat(weight)).Int(nil)
	return weightedAmount
}
