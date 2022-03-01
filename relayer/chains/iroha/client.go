package iroha

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewRPCClient(endpoint string) (*rpc.Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}

	return conn, err
}

func (chain *Chain) CallOpts(ctx context.Context, height int64) *bind.CallOpts {
	txOpts := chain.TxOpts(ctx)
	opts := &bind.CallOpts{
		From:    txOpts.From,
		Context: txOpts.Context,
	}
	if height > 0 {
		opts.BlockNumber = big.NewInt(height)
	}
	return opts
}

func (chain *Chain) TxOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     chain.relayerAddress,
		GasLimit: 1000000,
		Context:  ctx,
	}
}
