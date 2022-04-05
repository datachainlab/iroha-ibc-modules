package iroha

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchandler"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/contract/irohaics20bank"
	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/contract/irohaics20transferbank"
)

type BoundContract struct {
	address common.Address
	abi     abi.ABI
	conn    *rpc.Client
}

func NewBoundContract(address common.Address, abi abi.ABI, conn *rpc.Client) BoundContract {
	return BoundContract{
		address: address,
		abi:     abi,
		conn:    conn,
	}
}

func (c BoundContract) Abi() abi.ABI {
	return c.abi
}

func (c BoundContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*Transaction, error) {
	from := common.NewMixedcaseAddress(opts.From)
	to := common.NewMixedcaseAddress(c.address)
	gas := hexutil.Uint64(0)
	gasPrice := hexutil.Big(*big.NewInt(0))
	value := hexutil.Big(*big.NewInt(0))
	nonce := hexutil.Uint64(0)
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
	return c.sendTx(ctx, &args)
}

func (c BoundContract) sendTx(ctx context.Context, args *core.SendTxArgs) (*Transaction, error) {
	var txID common.Hash
	err := c.conn.CallContext(ctx, &txID, "eth_sendTransaction", args)
	if err != nil {
		return nil, err
	}

	var tx *gethtypes.Transaction
	var input []byte
	if args.Data != nil {
		input = *args.Data
	} else if args.Input != nil {
		input = *args.Input
	}
	if args.To == nil {
		tx = gethtypes.NewContractCreation(uint64(args.Nonce), (*big.Int)(&args.Value), uint64(args.Gas), (*big.Int)(&args.GasPrice), input)
	} else {
		tx = gethtypes.NewTransaction(uint64(args.Nonce), args.To.Address(), (*big.Int)(&args.Value), (uint64)(args.Gas), (*big.Int)(&args.GasPrice), input)
	}

	return &Transaction{
		Transaction: tx,
		ID:          txID,
	}, nil
}

type IbcHost struct {
	ibchost.Ibchost
	BoundContract
}

func NewIbcHost(address common.Address, conn *rpc.Client) (*IbcHost, error) {
	ibcHost, err := ibchost.NewIbchost(address, ethclient.NewClient(conn))
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ibchost.IbchostABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, conn)

	return &IbcHost{
		Ibchost:       *ibcHost,
		BoundContract: boundContract,
	}, nil
}

type IbcHandler struct {
	ibchandler.Ibchandler
	BoundContract
}

func NewIbcHandler(address common.Address, conn *rpc.Client) (*IbcHandler, error) {
	ibcHandler, err := ibchandler.NewIbchandler(address, ethclient.NewClient(conn))
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(ibchandler.IbchandlerABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, conn)

	return &IbcHandler{
		Ibchandler:    *ibcHandler,
		BoundContract: boundContract,
	}, nil
}

type IrohaIcs20Bank struct {
	irohaics20bank.Irohaics20bank
	BoundContract
}

func NewIrohaIcs20Bank(address common.Address, conn *rpc.Client) (*IrohaIcs20Bank, error) {
	irohaIcs20Bank, err := irohaics20bank.NewIrohaics20bank(address, ethclient.NewClient(conn))
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(irohaics20bank.Irohaics20bankABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, conn)

	return &IrohaIcs20Bank{
		Irohaics20bank: *irohaIcs20Bank,
		BoundContract:  boundContract,
	}, nil
}

type IrohaIcs20Transfer struct {
	irohaics20transferbank.Irohaics20transferbank
	BoundContract
}

func NewIrohaIcs20Transfer(address common.Address, conn *rpc.Client) (*IrohaIcs20Transfer, error) {
	irohaIcs20TransferBank, err := irohaics20transferbank.NewIrohaics20transferbank(address, ethclient.NewClient(conn))
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(irohaics20transferbank.Irohaics20transferbankABI))
	if err != nil {
		return nil, err
	}

	boundContract := NewBoundContract(address, parsedABI, conn)

	return &IrohaIcs20Transfer{
		Irohaics20transferbank: *irohaIcs20TransferBank,
		BoundContract:          boundContract,
	}, nil
}
