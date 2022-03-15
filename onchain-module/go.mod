module github.com/datachainlab/iroha-ibc-modules/onchain-module

go 1.16

require (
	github.com/datachainlab/iroha-ibc-modules/iroha-go v0.0.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20220117004623-482544f9ca21
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.45.0
)

replace (
	github.com/datachainlab/iroha-ibc-modules/iroha-go => ../iroha-go
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)
