# iotex-election
Collect and process election information from the governance chain (which is Ethereum for now)

# Run as a service
0. dep ensure --vendor-only
1. rm election.db
2. go build -o ./bin/server -v ./server
3. ./bin/server
