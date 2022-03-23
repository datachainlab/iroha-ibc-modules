module github.com/datachainlab/iroha-ibc-modules/relayer

go 1.16

replace (
	github.com/datachainlab/iroha-ibc-modules/iroha-go => ../iroha-go
	github.com/datachainlab/iroha-ibc-modules/onchain-module => ../onchain-module
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/ibc-ethmultisig-client v0.1.1-0.20220302030309-1191ec233811
	github.com/datachainlab/iroha-ibc-modules/iroha-go v0.0.0
	github.com/datachainlab/iroha-ibc-modules/onchain-module v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20220214080515-0f917e10509b
	github.com/hyperledger-labs/yui-relayer v0.1.2-0.20220124061305-6b081dc42621
	github.com/spf13/cobra v1.1.3
	google.golang.org/grpc v1.45.0
)
