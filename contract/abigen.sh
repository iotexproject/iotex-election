#!/bin/bash

abigen --abi iotx.abi --pkg contract --type IOTX --out iotx.go.tmp
abigen --abi register.abi --pkg contract --type Register --out register.go.tmp
abigen --abi staking.abi --pkg contract --type Staking --out staking.go.tmp

sed 's/github.com\/ethereum/github.com\/iotexproject/g' iotx.go.tmp > iotx.go
sed 's/github.com\/ethereum/github.com\/iotexproject/g' register.go.tmp > register.go
sed 's/github.com\/ethereum/github.com\/iotexproject/g' staking.go.tmp > staking.go

rm staking.go.tmp iotx.go.tmp register.go.tmp
