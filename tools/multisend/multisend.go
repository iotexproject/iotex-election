// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-core/address"
	"github.com/iotexproject/iotex-core/cli/ioctl/util"
	"github.com/iotexproject/iotex-core/protogen/iotexapi"
	"github.com/iotexproject/iotex-election/committee"
)

// Bucket of votes
type Bucket struct {
	ethAddr string
	stakes  string
	bpname  string
}

var abiJSON = `[{"constant":false,"inputs":[{"name":"recipients","type":"address[]"},
{"name":"amounts","type":"uint256[]"},{"name":"payload","type":"string"}],
"name":"multiSend","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},
{"anonymous":false,"inputs":[{"indexed":false,"name":"recipient","type":"address"},
{"indexed":false,"name":"amount","type":"uint256"}],"name":"Transfer","type":"event"},
{"anonymous":false,"inputs":[{"indexed":false,"name":"refund","type":"uint256"}],
"name":"Refund","type":"event"},{"anonymous":false,
"inputs":[{"indexed":false,"name":"payload","type":"string"}],"name":"Payload","type":"event"}]`
var abiFunc = "multiSend"

func main() {
	bpHexMap := map[string]string{
		"iotxplorerio": "696f7478706c6f726572696f",
		"longz":        "000000000000006c6f6e677a",
		"iotextrader":  "00696f746578747261646572",
		"gamefantasy":  "67616d6566616e7461737900",
		"superiotex":   "00007375706572696f746578",
		"iotexhub":     "00000000696f746578687562",
		"consensusnet": "636f6e73656e7375736e6574",
		"keysiotex":    "0000006b657973696f746578",
		"slowmist":     "00000000736c6f776d697374",
		"cryptolions":  "0063727970746f6c696f6e73",
		"iotexteam":    "000000696f7465787465616d",
		"droute":       "00000000000064726f757465",
		"hashbuy":      "000000000068617368627579",
		"cobo":         "0000000000000000636f626f",
		"blockboost":   "0000626c6f636b626f6f7374",
		"lanhu":        "000000000000006c616e6875",
		"cpc":          "000000000000000000637063",
		"capitmu":      "000000000063617069746d75",
		"whales":       "0000000000007768616c6573",
		"draperdragon": "647261706572647261676f6e",
		"alphacoin":    "000000616c706861636f696e",
		"airfoil":      "0000000000616972666f696c",
		"infstones":    "000000696e6673746f6e6573",
		"metanyx":      "00000000006d6574616e7978",
		"iotexbgogo":   "0000696f74657862676f676f",
		"royalland":    "000000726f79616c6c616e64",
		"preangel":     "00000000707265616e67656c",
		"blockvc":      "0000000000626c6f636b7663",
		"iosg":         "0000000000000000696f7367",
		"zhcapital":    "0000007a686361706974616c",
		"meter":        "000000000000006d65746572",
		"":             "000000000000000000000000",
		"pubxpayments": "707562787061796d656e7473",
		"coingecko":    "000000636f696e6765636b6f",
		"iotexmainnet": "696f7465786d61696e6e6574",
		"rkt8":         "0000000000000000726b7438",
		"yvalidator":   "00007976616c696461746f72",
		"wannodes":     "0000000077616e6e6f646573",
		"eon":          "000000000000000000656f6e",
		"iotask":       "000000000000696f7461736b",
		"iotexcore":    "000000696f746578636f7265",
		"iotexgeeks":   "0000696f7465786765656b73",
		"iotexlab":     "00000000696f7465786c6162",
		"raketat8":     "0000000072616b6574617438",
		"iotexunion":   "0000696f746578756e696f6e",
		"cryptolionsx": "63727970746f6c696f6e7378",
		"ducapital":    "00000064756361706974616c",
		"applytoday":   "6170706c79746f6461790000",
		"piexgo":       "00000000000070696578676f",
		"iotexicu":     "00000000696f746578696375",
		"thebottoken":  "746865626f74746f6b656e00",
		"mrtrump":      "00000000006d727472756d70",
		"enlightiv":    "000000656e6c696768746976",
		"iotextech":    "000000696f74657874656368",
		"ratels":       "000000000000726174656c73",
		"wyvalidator":  "00777976616c696461746f72",
		"rosemary0":    "000000726f73656d61727930",
		"rosemary1":    "000000726f73656d61727931",
		"rosemary2":    "000000726f73656d61727932",
		"rosemary3":    "000000726f73656d61727933",
		"rosemary4":    "000000726f73656d61727934",
		"rosemary5":    "000000726f73656d61727935",
		"rosemary6":    "000000726f73656d61727936",
		"rosemary7":    "000000726f73656d61727937",
		"rosemary8":    "000000726f73656d61727938",
		"rosemary9":    "000000726f73656d61727939",
		"rosemary10":   "0000726f73656d6172793130",
		"rosemary11":   "0000726f73656d6172793131",
		"rosemary12":   "0000726f73656d6172793132",
		"rosemary13":   "0000726f73656d6172793133",
		"rosemary14":   "0000726f73656d6172793134",
		"rosemary15":   "0000726f73656d6172793135",
		"rosemary16":   "0000726f73656d6172793136",
		"rosemary17":   "0000726f73656d6172793137",
		"rosemary18":   "0000726f73656d6172793138",
		"rosemary19":   "0000726f73656d6172793139",
		"rosemary20":   "0000726f73656d6172793230",
		"rosemary21":   "0000726f73656d6172793231",
		"rosemary22":   "0000726f73656d6172793232",
		"rosemary23":   "0000726f73656d6172793233",
		"bitwires":     "000000006269747769726573",
		"snzholding":   "0000736e7a686f6c64696e67",
		"iotime":       "000000000000696f74696d65",
		"laomao":       "0000000000006c616f6d616f",
		"wetez":        "00000000000000776574657a",
	}

	var configPath string
	var epoch uint64
	var height uint64
	var bp string
	var bpHex string
	var amount string
	var endpoint string
	flag.StringVar(&configPath, "config", "committee.yaml", "path of server config file")
	flag.Uint64Var(&epoch, "epoch", 0, "iotex epoch")
	flag.Uint64Var(&height, "height", 0, "ethereuem height")
	flag.StringVar(&bp, "bp", "", "bp name")
	flag.StringVar(&bpHex, "bp-hex", "", "bp hex name")
	flag.StringVar(&amount, "amount", "", "amount of IOTX")
	flag.StringVar(&endpoint, "ednpoint", "api.iotex.one:80", "set endpoint")
	flag.Parse()

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panic("failed to load config file", zap.Error(err))
	}
	var config committee.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Panic("failed to unmarshal config", zap.Error(err))
	}
	committee, err := committee.NewCommittee(nil, config)
	if err != nil {
		log.Panic("failed to create committee", zap.Error(err))
	}
	if epoch != 0 {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			log.Panic("failed to connect endpoint", zap.Error(err))
		}
		defer conn.Close()
		cli := iotexapi.NewAPIServiceClient(conn)
		request := iotexapi.GetEpochMetaRequest{EpochNumber: epoch}
		ctx := context.Background()
		response, err := cli.GetEpochMeta(ctx, &request)
		if err != nil {
			log.Panic("failed to get epoch meta", zap.Error(err))
		}
		height = response.EpochData.GravityChainStartHeight
	}
	result, err := committee.FetchResultByHeight(height)
	if err != nil {
		log.Panic("failed to fetch result", zap.Uint64("height", height))
	}
	var buckets []Bucket
	for _, delegate := range result.Delegates() {
		for _, vote := range result.VotesByDelegate(delegate.Name()) {
			buckets = append(buckets, Bucket{
				ethAddr: hex.EncodeToString(vote.Voter()),
				stakes:  vote.WeightedAmount().String(),
				bpname:  string(vote.Candidate()),
			})
		}
	}
	bps := process(buckets)
	var bpByte []byte
	var ok bool
	if len(bpHex) == 0 {
		bpHex, ok = bpHexMap[bp]
		if !ok {
			zeroByte := []byte{}
			for i := 0; i < 12-len(bp); i++ {
				zeroByte = append(zeroByte, byte(0))
			}
			bpByte = append(zeroByte, []byte(bp)...)
		}
	}
	if len(bpHex) != 0 {
		bpByte, err = hex.DecodeString(bpHex)
		if err != nil {
			log.Panic(err.Error())
		}
	}
	bp = string(bpByte)
	bp1, ok := bps[bp]
	if !ok {
		log.Panic("invalid bp name: " + bp)
	}
	totalVotes := big.NewInt(0)
	var keys []string
	fmt.Printf("%-41s\t%-40s\t%-32s%s", "IOAddr", "ETHAddr", "Votes", "Reward(IOTX)\n")
	for k, v := range bp1 {
		votes, _ := new(big.Int).SetString(v, 10)
		totalVotes.Add(totalVotes, votes)
		keys = append(keys, k)
	}
	recipients := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	payload := bp
	totalAmount, err := util.StringToRau(amount, util.IotxDecimalNum)
	if err != nil {
		log.Panic("invalid amount")
	}
	sort.Strings(keys)
	for _, k := range keys {
		ioAddr := toIoAddr(k)
		recipient, err := util.IoAddrToEvmAddr(ioAddr)
		if err != nil {
			log.Panic(err.Error())
		}
		if err != nil {
			log.Panic(err.Error())
		}
		recipients = append(recipients, recipient)
		votes, _ := new(big.Int).SetString(bp1[k], 10)
		amountPerVoter := new(big.Int).Div(new(big.Int).Mul(votes, totalAmount), totalVotes)
		amounts = append(amounts, amountPerVoter)
		fmt.Printf("%s\t%s\t%-32s%s\n", toIoAddr(k), k, bp1[k],
			util.RauToString(amountPerVoter, util.IotxDecimalNum))
	}
	reader := strings.NewReader(abiJSON)
	multisendABI, _ := abi.JSON(reader)
	bytecode, _ := multisendABI.Pack(abiFunc, recipients, amounts, payload)
	fmt.Println("\nbytecode: " + hex.EncodeToString(bytecode))
}

func init() {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.WarnLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Panic("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)
}

func process(buckets []Bucket) (bps map[string](map[string]string)) {
	bps = make(map[string](map[string]string))
	for _, bucket := range buckets {
		vs, ok := bps[bucket.bpname]
		if ok {
			// Already have this BP
			_, ook := vs[bucket.ethAddr]
			if ook {
				// Already have this eth addr, need to combine the stakes
				vs[bucket.ethAddr] = addStrs(vs[bucket.ethAddr], bucket.stakes)
			} else {
				vs[bucket.ethAddr] = bucket.stakes
			}
		} else {
			vs := make(map[string]string)
			vs[bucket.ethAddr] = bucket.stakes
			name := "UNVOTED"
			if len(bucket.bpname) > 0 {
				name = bucket.bpname
			}
			bps[name] = vs
		}
	}

	return bps
}

func addStrs(a, b string) string {
	aa := new(big.Int)
	aaa, ok := aa.SetString(a, 10)
	if !ok {
		panic("SetString: error")
	}
	bb := new(big.Int)
	bbb, ok := bb.SetString(b, 10)
	if !ok {
		panic("SetString: error")
	}
	c := new(big.Int)
	c.Add(aaa, bbb)
	return c.String()
}

func toIoAddr(addr string) string {
	ethAddr := common.HexToAddress(addr)
	pkHash := ethAddr.Bytes()
	ioAddr, _ := address.FromBytes(pkHash)
	return ioAddr.String()
}
