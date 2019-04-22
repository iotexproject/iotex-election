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
	"crypto/tls"
	"encoding/csv"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	. "github.com/logrusorgru/aurora"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-core/action/protocol/rewarding/rewardingpb"
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

// Hard code
var bpHexMap map[string]string
var abiJSON string
var abiFunc string

// Flags
var configPath string
var epochStart uint64
var epochEnd uint64
var bp string
var endpoint string
var distPercentage uint64
var rewardAddress string
var withFoundationBonus bool
var csvFile bool

func main() {
	totalReward := big.NewInt(0)
	distributions := make(map[string]*big.Int)
	for i := epochStart; i <= epochEnd; i++ {
		reward := getReward(i)
		if reward.Sign() == 0 {
			continue
		}
		totalReward.Add(totalReward, reward)
		buckets := dump(i)
		bps := process(buckets)
		bp0 := getBP(bps)
		calculate(distributions, bp0, reward)
	}
	if csvFile {
		fileCSV(distributions)
	} else {
		printResult(distributions, totalReward)
	}
}

func getReward(epoch uint64) *big.Int {
	lastBlock := epoch * 24 * 15 // numDelegate: 24, subEpoch: 15
	conn, err := util.ConnectToEndpoint(false)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	cli := iotexapi.NewAPIServiceClient(conn)
	blockRequest := &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: lastBlock,
				Count: 1,
			},
		},
	}
	ctx := context.Background()
	blockResponse, err := cli.GetBlockMetas(ctx, blockRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(blockResponse.BlkMetas) == 0 {
		log.Fatalln("failed to get last block in epoch", epoch)
	}
	actionsRequest := &iotexapi.GetActionsRequest{
		Lookup: &iotexapi.GetActionsRequest_ByBlk{
			ByBlk: &iotexapi.GetActionsByBlockRequest{
				BlkHash: blockResponse.BlkMetas[0].Hash,
				Start:   uint64(blockResponse.BlkMetas[0].NumActions) - 1,
				Count:   1,
			},
		},
	}
	actionsResponse, err := cli.GetActions(ctx, actionsRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(actionsResponse.ActionInfo) == 0 {
		log.Fatalln("failed to get last action in epoch", epoch)
	}
	if actionsResponse.ActionInfo[0].Action.Core.GetGrantReward() == nil {
		log.Fatalln("Not grantReward action")
	}
	receiptRequest := &iotexapi.GetReceiptByActionRequest{
		ActionHash: actionsResponse.ActionInfo[0].ActHash,
	}
	receiptResponse, err := cli.GetReceiptByAction(ctx, receiptRequest)
	if err != nil {
		log.Fatalln(err.Error())
	}
	eReward := big.NewInt(0)
	fReward := big.NewInt(0)
	for _, receiptLog := range receiptResponse.ReceiptInfo.Receipt.Logs {
		var rewardLog rewardingpb.RewardLog
		var ok bool
		if err := proto.Unmarshal(receiptLog.Data, &rewardLog); err != nil {
			log.Fatalln(err.Error())
		}
		if rewardLog.Addr == rewardAddress {
			if rewardLog.Type == rewardingpb.RewardLog_EPOCH_REWARD {
				eReward, ok = new(big.Int).SetString(rewardLog.Amount, 10)
				if !ok {
					log.Fatalln("SetString: error")
				}
			} else if rewardLog.Type == rewardingpb.RewardLog_FOUNDATION_BONUS && withFoundationBonus {
				fReward, ok = new(big.Int).SetString(rewardLog.Amount, 10)
				if !ok {
					log.Fatalln("SetString: error")
				}
			}
		}
	}
	fmt.Printf("epoch %-10d", epoch)
	fmt.Printf("epoch reward: %-30s", util.RauToString(eReward, util.IotxDecimalNum))
	if withFoundationBonus {
		fmt.Printf("foundation bonus: %-30s", util.RauToString(fReward, util.IotxDecimalNum))
	}
	fmt.Println()
	return new(big.Int).Add(eReward, fReward)
}

func dump(epoch uint64) (buckets []Bucket) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("failed to load config file")
	}
	var config committee.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln("failed to unmarshal config")
	}
	committee, err := committee.NewCommittee(nil, config)
	if err != nil {
		log.Fatalln("failed to create committee")
	}
	var height uint64
	if epoch != 0 {
		conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
		if err != nil {
			log.Fatalln("failed to connect endpoint")
		}
		defer conn.Close()
		cli := iotexapi.NewAPIServiceClient(conn)
		request := iotexapi.GetEpochMetaRequest{EpochNumber: epoch}
		ctx := context.Background()
		response, err := cli.GetEpochMeta(ctx, &request)
		if err != nil {
			log.Fatalln("failed to get epoch meta")
		}
		height = response.EpochData.GravityChainStartHeight
	}
	result, err := committee.FetchResultByHeight(height)
	if err != nil {
		log.Fatalln("failed to fetch result")
	}
	for _, delegate := range result.Delegates() {
		for _, vote := range result.VotesByDelegate(delegate.Name()) {
			buckets = append(buckets, Bucket{
				ethAddr: hex.EncodeToString(vote.Voter()),
				stakes:  vote.WeightedAmount().String(),
				bpname:  string(vote.Candidate()),
			})
		}
	}
	return
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

func getBP(bps map[string](map[string]string)) map[string]string {
	var bpHex string
	var bpByte []byte
	bpHex, ok := bpHexMap[bp]
	if !ok {
		zeroByte := []byte{}
		for i := 0; i < 12-len(bp); i++ {
			zeroByte = append(zeroByte, byte(0))
		}
		bpByte = append(zeroByte, []byte(bp)...)
		bpHex = string(bpByte)
	}
	bpByte, err := hex.DecodeString(bpHex)
	if err != nil {
		log.Fatalln(err.Error())
	}
	bp0, ok := bps[string(bpByte)]
	if !ok {
		log.Fatalln("invalid bp name: " + bp)
	}
	return bp0
}

func calculate(distributions map[string]*big.Int, bp0 map[string]string, rewardAmount *big.Int) {
	payoutAmount := new(big.Int).Div(new(big.Int).Mul(rewardAmount,
		big.NewInt(int64(distPercentage))), big.NewInt(100))
	totalVotes := big.NewInt(0)
	var keys []string
	for k, v := range bp0 {
		votes, ok := new(big.Int).SetString(v, 10)
		if !ok {
			log.Panic("SetString: error")
		}
		totalVotes.Add(totalVotes, votes)
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		votes, ok := new(big.Int).SetString(bp0[k], 10)
		if !ok {
			log.Panic("SetString: error")
		}
		amount := new(big.Int).Div(new(big.Int).Mul(votes, payoutAmount), totalVotes)
		if _, ok := distributions[k]; !ok {
			distributions[k] = amount
		} else {
			distributions[k].Add(distributions[k], amount)
		}
	}
}

func printResult(distributions map[string]*big.Int, totalReward *big.Int) {
	totalPayout := big.NewInt(0)
	recipients := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	payload := bp
	var list []string
	list = append(list, fmt.Sprintf("%-41s\t%-40s\t%s", "IOAddr", "ETHAddr", "Distribution(IOTX)"))
	for k, v := range distributions {
		ioAddr := toIoAddr(k)
		recipient, err := util.IoAddrToEvmAddr(ioAddr)
		if err != nil {
			log.Fatalln(err.Error())
		}
		recipients = append(recipients, recipient)
		amounts = append(amounts, v)
		totalPayout.Add(totalPayout, v)
		list = append(list, fmt.Sprintf("%s\t%s\t%s", toIoAddr(k), recipient.String(),
			util.RauToString(v, util.IotxDecimalNum)))
	}

	// generate bytecode
	reader := strings.NewReader(abiJSON)
	multisendABI, _ := abi.JSON(reader)
	bytecode, _ := multisendABI.Pack(abiFunc, recipients, amounts, payload)

	// print
	fmt.Printf("\nBytecode for invoking multisend contract\n")
	fmt.Println(hex.EncodeToString(bytecode))
	fmt.Println()
	fmt.Println(strings.Join(list, "\n"))
	fmt.Printf("\n%-15s%-30s%-30s%s", "Epoches", "Total Reward(IOTX)", "Percentage %", "Total Distribution(IOTX)")
	fmt.Printf("\n%-15d%-30s%-30d%s\n", epochEnd-epochStart+1, util.RauToString(totalReward, util.IotxDecimalNum),
		distPercentage, util.RauToString(totalPayout, util.IotxDecimalNum))
}

func fileCSV(distributions map[string]*big.Int) {
	file, err := os.Create("distributions.csv")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for k, v := range distributions {
		ioAddr := toIoAddr(k)
		recipient, err := util.IoAddrToEvmAddr(ioAddr)
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = writer.Write([]string{recipient.String(), v.String()})
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func addStrs(a, b string) string {
	aa := new(big.Int)
	aaa, ok := aa.SetString(a, 10)
	if !ok {
		log.Panic("SetString: error")
	}
	bb := new(big.Int)
	bbb, ok := bb.SetString(b, 10)
	if !ok {
		log.Panic("SetString: error")
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

func init() {
	// print disclaim
	disclaim := Red("This Bookkeeper is a REFERENCE IMPLEMENTATION of reward distribution tool provided by IOTEX FOUNDATION. IOTEX FOUNDATION disclaims all responsibility for any damages or losses (including, without limitation, financial loss, damages for loss in business projects, loss of profits or other consequential losses) arising in contract, tort or otherwise from the use of or inability to use the Bookkeeper, or from any action or decision taken as a result of using this Bookkeeper.")
	fmt.Printf("\n%s\n%s\n", Bold(Red("Attention")), disclaim)

	// flags
	flag.StringVar(&configPath, "config", "committee.yaml", "path of server config file")
	flag.Uint64Var(&epochStart, "start", 0, "iotex epoch start")
	flag.Uint64Var(&epochEnd, "end", 0, "iotex epoch end (included)")
	flag.StringVar(&bp, "bp", "", "bp name")
	flag.StringVar(&endpoint, "endpoint", "api.iotex.one:443", "set endpoint")
	flag.Uint64Var(&distPercentage, "dist-percentage", 0, "distribution percentage of epoch reward")
	flag.StringVar(&rewardAddress, "reward-address", "", "choose reward address in certain epoch")
	flag.BoolVar(&withFoundationBonus, "with-foundation-bonus", false, "add foundation bonus in distribution")
	flag.BoolVar(&csvFile, "csv", false, "write to a csv file")
	flag.Parse()

	// check
	if rewardAddress == "" {
		log.Fatalln("empty reward address")
	}
	if epochStart > epochEnd {
		log.Fatalln("start epoch is larger than end epoch")
	}

	// warning
	if distPercentage > 100 {
		fmt.Println(Brown("\nWarning: percentage " + strconv.Itoa(int(distPercentage)) + `% is larger than 100%`))
	}
	if epochEnd-epochStart >= 24 {
		fmt.Println(Brown("\nWarning: fetch more than 24 epoches' voters may cost much time"))
	}
	fmt.Println()

	// init zap
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapCfg.Level.SetLevel(zap.WarnLevel)
	l, err := zapCfg.Build()
	if err != nil {
		log.Fatalln("Failed to init zap global logger, no zap log will be shown till zap is properly initialized: ", err)
	}
	zap.ReplaceGlobals(l)

	// hard code
	abiJSON = `[{"constant":false,"inputs":[{"name":"recipients","type":"address[]"},
	{"name":"amounts","type":"uint256[]"},{"name":"payload","type":"string"}],
	"name":"multiSend","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},
	{"anonymous":false,"inputs":[{"indexed":false,"name":"recipient","type":"address"},
	{"indexed":false,"name":"amount","type":"uint256"}],"name":"Transfer","type":"event"},
	{"anonymous":false,"inputs":[{"indexed":false,"name":"refund","type":"uint256"}],
	"name":"Refund","type":"event"},{"anonymous":false,
	"inputs":[{"indexed":false,"name":"payload","type":"string"}],"name":"Payload","type":"event"}]`
	abiFunc = "multiSend"
	bpHexMap = map[string]string{
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
}
