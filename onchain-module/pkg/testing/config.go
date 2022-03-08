package testing

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/truffle"
)

var _ ContractConfig = (*TruffleContractConfig)(nil)

type TruffleContractConfig struct {
	ibcHostAddress                common.Address
	ibcHandlerAddress             common.Address
	ibcIdentifierAddress          common.Address
	mockClientAddress             common.Address
	simpleTokenAddress            common.Address
	ics20BankAddress              common.Address
	ics20TransferBankAddress      common.Address
	irohaIcs20BankAddress         common.Address
	irohaIcs20TransferBankAddress common.Address
}

func NewTruffleContractConfig(networkID int, buildPath string) TruffleContractConfig {
	c := TruffleContractConfig{}
	networkIDStr := fmt.Sprint(networkID)

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "IBCHost.json")
		c.ibcHostAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "IBCHandler.json")
		c.ibcHandlerAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "IBCIdentifier.json")
		c.ibcIdentifierAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "MockClient.json")
		c.mockClientAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "SimpleToken.json")
		c.simpleTokenAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "ICS20Bank.json")
		c.ics20BankAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "ICS20TransferBank.json")
		c.ics20TransferBankAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "IrohaICS20Bank.json")
		c.irohaIcs20BankAddress = common.HexToAddress(string(n.Address))
	}

	{
		n := truffle.UnmarshallConfig(networkIDStr, buildPath, "IrohaICS20TransferBank.json")
		c.irohaIcs20TransferBankAddress = common.HexToAddress(string(n.Address))
	}

	return c
}

func (c TruffleContractConfig) GetIBCHostAddress() common.Address {
	return c.ibcHostAddress
}

func (c TruffleContractConfig) GetIBCHandlerAddress() common.Address {
	return c.ibcHandlerAddress
}

func (c TruffleContractConfig) GetIBCIdentifierAddress() common.Address {
	return c.ibcIdentifierAddress
}

func (c TruffleContractConfig) GetMockClientAddress() common.Address {
	return c.mockClientAddress
}

func (c TruffleContractConfig) GetSimpleTokenAddress() common.Address {
	return c.simpleTokenAddress
}

func (c TruffleContractConfig) GetICS20BankAddress() common.Address {
	return c.ics20BankAddress
}

func (c TruffleContractConfig) GetICS20TransferBankAddress() common.Address {
	return c.ics20TransferBankAddress
}

func (c TruffleContractConfig) GetIrohaICS20BankAddress() common.Address {
	return c.irohaIcs20BankAddress
}

func (c TruffleContractConfig) GetIrohaICS20TransferBankAddress() common.Address {
	return c.irohaIcs20TransferBankAddress
}
