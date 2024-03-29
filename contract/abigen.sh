#!/bin/bash

abigen --abi iotx.abi --pkg contract --type IOTX --out iotx.go
abigen --abi rotatablevps.abi --pkg contract --type RotatableVPS --out rotatablevps.go
abigen --abi broker.abi --pkg contract --type Broker --out broker.go
abigen --abi clerk.abi --pkg contract --type Clerk --out clerk.go
abigen --abi vita.abi --pkg contract --type Vita --out vita.go
abigen --abi register.abi --pkg contract --type Register --out register.go
abigen --abi staking.abi --pkg contract --type Staking --out staking.go
abigen --abi pyggstaking.abi --pkg contract --type PyggStaking --out pyggstaking.go
abigen --abi agent.abi --pkg contract --type Agent --out agent.go