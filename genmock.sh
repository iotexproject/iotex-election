#!/bin/bash

rm -rf ./test/mock
mkdir -p ./test/mock

mkdir -p ./test/mock/mock_committee
mockgen -destination=./test/mock/mock_committee/mock_committee.go  \
        -source=./committee/committee.go \
        -imports =github.com/iotexproject/iotex-election/committee \
        -package=mock_committee \
        Committee