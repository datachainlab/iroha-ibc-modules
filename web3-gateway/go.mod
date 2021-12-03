module github.com/datachainlab/iroha-ibc-modules/web3-gateway

go 1.16

require (
	github.com/avast/retry-go/v4 v4.0.1
	github.com/datachainlab/iroha-ibc-modules/iroha-go v0.0.0
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-memdb v1.3.2
	github.com/hyperledger/burrow v0.29.7
	github.com/jackc/pgx/v4 v4.14.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.42.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace (
	github.com/datachainlab/iroha-ibc-modules/iroha-go => ../iroha-go
	github.com/perlin-network/life => github.com/silasdavis/life v0.0.0-20191009191257-e9c2a5fdbc96
)
