# iotex-election
Collect and process election information from the governance chain (which is Ethereum for now)

# Run as a service
1. rm election.db
2. go build -o ./bin/server -v ./server
3. ./bin/server
