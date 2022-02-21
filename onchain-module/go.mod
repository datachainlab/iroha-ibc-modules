module github.com/datachainlab/iroha-ibc-modules/onchain-module

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-beta1
	github.com/datachainlab/ibc-ethmultisig-client v0.1.1-0.20220216060713-5f9cc55814c0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gogo/protobuf v1.3.3
	github.com/hyperledger-labs/yui-ibc-solidity v0.0.0-20220214080515-0f917e10509b
	github.com/stretchr/testify v1.7.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
