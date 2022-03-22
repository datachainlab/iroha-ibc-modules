package cmd

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func accountIDToAddress(accountID string) common.Address {
	hash := gethcrypto.Keccak256([]byte(accountID))
	return common.BytesToAddress(hash[12:32])
}

func transactOpts(ctx context.Context, accountID string) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:    accountIDToAddress(accountID),
		Context: ctx,
	}
}

func callOpts(ctx context.Context, accountID string) *bind.CallOpts {
	return &bind.CallOpts{
		From:    accountIDToAddress(accountID),
		Context: ctx,
	}
}

func waitForReceipt(ctx context.Context, rpcCli *rpc.Client, txHash common.Hash) (*types.Receipt, error) {
	ethCli := ethclient.NewClient(rpcCli)

	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	for {
		receipt, err := ethCli.TransactionReceipt(ctx, txHash)
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
