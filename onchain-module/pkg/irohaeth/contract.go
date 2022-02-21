package irohaeth

import (
	"context"
	"math/big"
	"strings"

	"github.com/datachainlab/ibc-ethmultisig-client/pkg/contract/multisigclient"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibcidentifier"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20bank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ics20transferbank"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/simpletoken"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/client"
	irohatypes "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/irohaeth/types"
)

type BoundContract struct {
	address common.Address
	abi     abi.ABI
	client  client.Client
}

func NewBoundContract(address common.Address, abi abi.ABI, client client.Client) BoundContract {
	return BoundContract{
		address: address,
		abi:     abi,
		client:  client,
	}
}

func (c BoundContract) Abi() abi.ABI {
	return c.abi
}

func (c BoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*irohatypes.Transaction, error) {
	from := common.NewMixedcaseAddress(opts.From)
	to := common.NewMixedcaseAddress(c.address)
	gas := hexutil.Uint64(0)
	gasPrice := hexutil.Big(*big.NewInt(0))
	value := hexutil.Big(*big.NewInt(0))
	nonce := (hexutil.Uint64)(0)
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	data := hexutil.Bytes(input)

	// make a sendTxArgs
	args := core.SendTxArgs{
		From:     from,
		To:       &to,
		Gas:      gas,
		GasPrice: gasPrice,
		Value:    value,
		Nonce:    nonce,
		Data:     &data,
		Input:    &data,
	}

	// send the transaction
	ctx := opts.Context
	if ctx == nil {
		ctx = context.TODO()
	}
	return c.client.SendTx(ctx, &args)
}

type Ibchost struct {
	ibchost.Ibchost
	BoundContract
}

func NewIbchost(address common.Address, client client.Client) (*Ibchost, error) {
	ibcHost, err := ibchost.NewIbchost(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ibchost.IbchostABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Ibchost{
		Ibchost:       *ibcHost,
		BoundContract: boundContract,
	}, nil
}

type IBCHandler struct {
	ibchandler.Ibchandler
	BoundContract
}

func NewIBCHandler(address common.Address, client client.Client) (*IBCHandler, error) {
	ibcHandler, err := ibchandler.NewIbchandler(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &IBCHandler{
		Ibchandler:    *ibcHandler,
		BoundContract: boundContract,
	}, nil
}

type Ibcidentifier struct {
	ibcidentifier.Ibcidentifier
	BoundContract
}

func NewIbcidentifier(address common.Address, client client.Client) (*Ibcidentifier, error) {
	ibcIdentifier, err := ibcidentifier.NewIbcidentifier(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ibcidentifier.IbcidentifierABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Ibcidentifier{
		Ibcidentifier: *ibcIdentifier,
		BoundContract: boundContract,
	}, nil
}

type Simpletoken struct {
	simpletoken.Simpletoken
	BoundContract
}

func NewSimpletoken(address common.Address, client client.Client) (*Simpletoken, error) {
	simpleToken, err := simpletoken.NewSimpletoken(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(simpletoken.SimpletokenABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Simpletoken{
		Simpletoken:   *simpleToken,
		BoundContract: boundContract,
	}, nil
}

type Ics20transferbank struct {
	ics20transferbank.Ics20transferbank
	BoundContract
}

func NewIcs20transferbank(address common.Address, client client.Client) (*Ics20transferbank, error) {
	ics20TransferBank, err := ics20transferbank.NewIcs20transferbank(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ics20transferbank.Ics20transferbankABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Ics20transferbank{
		Ics20transferbank: *ics20TransferBank,
		BoundContract:     boundContract,
	}, nil
}

type Ics20bank struct {
	ics20bank.Ics20bank
	BoundContract
}

func NewIcs20bank(address common.Address, client client.Client) (*Ics20bank, error) {
	ics20Bank, err := ics20bank.NewIcs20bank(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ics20bank.Ics20bankABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Ics20bank{
		Ics20bank:     *ics20Bank,
		BoundContract: boundContract,
	}, nil
}

type Multisigclient struct {
	multisigclient.Multisigclient
	BoundContract
}

func NewMultisigclient(address common.Address, client client.Client) (*Multisigclient, error) {
	multisigClient, err := multisigclient.NewMultisigclient(address, client)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(multisigclient.MultisigclientABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, client)

	return &Multisigclient{
		Multisigclient: *multisigClient,
		BoundContract:  boundContract,
	}, nil
}
