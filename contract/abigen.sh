#!/bin/bash

abigen --abi iotx.abi --pkg contract --type IOTX --out iotx.go
abigen --abi rotatablevps.abi --pkg contract --type RotatableVPS --out rotatablevps.go
abigen --abi broker.abi --bin broker.bin --pkg contract --type Broker --out broker.go
abigen --abi clerk.abi --bin clerk.bin --pkg contract --type Clerk --out clerk.go
abigen --abi vita.abi --bin vita.bin --pkg contract --type Vita --out vita.go
abigen --abi register.abi --pkg contract --type Register --out register.go
abigen --abi staking.abi --pkg contract --type Staking --out staking.go
