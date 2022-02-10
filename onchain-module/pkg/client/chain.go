package client

import (
	"context"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"

	irohatypes "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/irohaeth/types"
)

type Client struct {
	endpoint   string
	clientType string

	conn *rpc.Client
	ETHClient
}

func (cl Client) ClientType() string {
	return cl.clientType
}

func (cl Client) WaitForReceiptAndGet(ctx context.Context, tx *irohatypes.Transaction) (Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := cl.TransactionReceipt(ctx, tx.ID)
		if err != nil {
			return nil, err
		} else if receipt != nil {
			return receipt, nil
		} else {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		}
	}
}

func (cl Client) SendTx(ctx context.Context, args *core.SendTxArgs) (*irohatypes.Transaction, error) {
	var txID common.Hash
	err := cl.conn.CallContext(ctx, &txID, "eth_sendTransaction", args)
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

	return &irohatypes.Transaction{
		Transaction: tx,
		ID:          txID,
	}, nil
}

type ETHClient interface {
	bind.ContractBackend
	BlockByNumber(ctx context.Context, bn *big.Int) (*gethtypes.Block, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (Receipt, error)
}

type Receipt interface {
	PostState() []byte
	Status() uint64
	CumulativeGasUsed() uint64
	Bloom() gethtypes.Bloom
	Logs() []*gethtypes.Log
	TxHash() common.Hash
	ContractAddress() common.Address
	GasUsed() uint64
	BlockHash() common.Hash
	BlockNumber() *big.Int
	TransactionIndex() uint
	RevertReason() string
}

type GenTxOpts func(ctx context.Context) *bind.TransactOpts

func MakeGenTxOpts(accountID string) GenTxOpts {
	addr := gethcrypto.Keccak256([]byte(accountID))
	return func(ctx context.Context) *bind.TransactOpts {
		return &bind.TransactOpts{
			From:     common.HexToAddress(hex.EncodeToString(addr[12:32])),
			GasLimit: 0,
		}
	}
}
