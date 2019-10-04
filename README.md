# iotex-election
Collect and process election information from the governance chain (which is Ethereum for now)

# Run as a service
0. dep ensure --vendor-only
1. rm election.db
2. make run 


# Tools

## Dumper
### Build
```sh
[foo@bar iotex-election]$  go build -ldflags "-extldflags=-Wl,--allow-multiple-definition" -x -o ./bin/dumper -v ./tools/dumper
```
### Dump votes to csv
```sh
[foo@bar iotex-election]$ ./bin/dumper > stats.csv
```
### Run from source
```sh
[foo@bar iotex-election]$ go run -ldflags "-extldflags=-Wl,--allow-multiple-definition" tools/dumper/dumper.go > stats.csv
```

## Processor
### Build
```sh
[foo@bar iotex-election]$ go build -o ./bin/processor -v ./tools/processor
```
### Process votes
```sh
[foo@bar iotex-election]$ ./bin/processor
```
### Run from source
```sh
[foo@bar iotex-election]$ go run tools/processor/processor.go
```