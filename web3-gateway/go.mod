module github.com/datachainlab/iroha-ibc-modules/web3-gateway

go 1.16

require (
	github.com/hyperledger/burrow v0.29.7
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace (
	github.com/datachainlab/iroha-ibc-modules/iroha-go => ../iroha-go
	github.com/perlin-network/life => github.com/silasdavis/life v0.0.0-20191009191257-e9c2a5fdbc96
)
