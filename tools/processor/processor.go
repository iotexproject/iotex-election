// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// Usage:
// go run dumper/dumper.go to generate the csv
// cp the csv to processor/
// go run processor/processor.go to generate the staking info.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"sort"

	. "github.com/logrusorgru/aurora"
)

type Bucket struct {
	ethAddr string
	stakes  string
	bpname  string
}

// voter,startTime,duration,decay,tokens,votes,votee
//e16d79bc05455fb4e3a56dec88d1a338d6d6c354,2019-03-05 00:59:00 -0800 PST,336h0m0s,true,10000000000000000000000,10380178401692392586853,iosg

func main() {
	buckets := load()
	bps := process(buckets)

	// Sort for prettier printing
	keys := make([]string, 0, len(bps))
	for k := range bps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(Bold(Cyan(">>>" + k + "<<<")))
		for kk, vv := range bps[k] {
			fmt.Println(kk, ":", vv)
		}
	}
}

func load() []Bucket {
	csvFile, _ := os.Open("stats.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var buckets []Bucket
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		buckets = append(buckets, Bucket{
			ethAddr: line[0],
			stakes:  line[5],
			bpname:  line[6],
		})
	}
	return buckets
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
			name := "UNVOITED"
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
