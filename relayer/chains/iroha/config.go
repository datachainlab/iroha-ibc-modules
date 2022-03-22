package iroha

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.ChainConfigI = (*ChainConfig)(nil)

func (c ChainConfig) Build() (core.ChainI, error) {
	return NewChain(c)
}

func (c ChainConfig) IBCHostAddress() common.Address {
	return common.HexToAddress(c.IbcHostAddress)
}

func (c ChainConfig) IBCHandlerAddress() common.Address {
	return common.HexToAddress(c.IbcHandlerAddress)
}

func (c ChainConfig) IrohaICS20BankAddress() common.Address {
	return common.HexToAddress(c.IrohaIcs20BankAddress)
}

func (c ChainConfig) IrohaICS20TransferAddress() common.Address {
	return common.HexToAddress(c.IrohaIcs20TransferAddress)
}
