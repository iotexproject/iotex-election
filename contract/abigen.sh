#!/bin/bash

abigen --abi staking.abi --pkg contract --type Staking --out staking.go.tmp
abigen --abi iotx.abi --pkg contract --type IOTX --out iotx.go.tmp

sed 's/github.com\/ethereum/github.com\/iotexproject/g' staking.go.tmp > staking.go
sed 's/github.com\/ethereum/github.com\/iotexproject/g' iotx.go.tmp > iotx.go

rm staking.go.tmp iotx.go.tmp
