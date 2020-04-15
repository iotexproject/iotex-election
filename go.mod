module github.com/iotexproject/iotex-election

go 1.12

require (
	github.com/bwmarrin/discordgo v0.19.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/ethereum/go-ethereum v1.8.27
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/hashicorp/golang-lru v0.5.1
	github.com/iotexproject/go-pkgs v0.1.2-0.20200212033110-8fa5cf96fc1b
	github.com/iotexproject/iotex-address v0.2.1
	github.com/iotexproject/iotex-antenna-go/v2 v2.3.2
	github.com/iotexproject/iotex-core v0.11.1
	github.com/iotexproject/iotex-proto v0.2.6-0.20200327040553-157f35632918
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	go.etcd.io/bbolt v1.3.2
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20191204025024-5ee1b9f4859a
	google.golang.org/grpc v1.21.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/ethereum/go-ethereum => github.com/iotexproject/go-ethereum v0.2.0
