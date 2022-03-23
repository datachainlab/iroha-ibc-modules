package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"github.com/datachainlab/iroha-ibc-modules/relayer/chains/iroha"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hyperledger-labs/yui-relayer/config"
	"google.golang.org/grpc"
)

func findIrohaChainConfig(ctx *config.Context, chainID string) (*iroha.ChainConfig, error) {
	for _, config := range ctx.Config.Chains {
		if err := config.Init(ctx.Codec); err != nil {
			return nil, err
		} else if chain, err := config.Build(); err != nil {
			return nil, err
		} else if chain.ChainID() == chainID {
			var cfg iroha.ChainConfig
			if err := json.Unmarshal(config.Chain, &cfg); err != nil {
				return nil, err
			}
			return &cfg, nil
		}
	}
	return nil, fmt.Errorf("chain config not found for chain_id:%s", chainID)
}

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

func dialToIrohad(endpoint string) (*grpc.ClientConn, error) {
	return grpc.Dial(
		endpoint,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func makeCommandClient(grpcConn *grpc.ClientConn) command.CommandClient {
	return command.New(grpcConn, time.Minute)
}

func makeQueryClient(grpcConn *grpc.ClientConn) query.QueryClient {
	return query.New(grpcConn, time.Minute)
}

func sendIrohaTx(ctx context.Context, grpcConn *grpc.ClientConn, tx *protocol.Transaction, keys ...string) (*protocol.ToriiResponse, error) {
	// sign tx
	if sigs, err := crypto.SignTransaction(tx, keys...); err != nil {
		return nil, fmt.Errorf("SignTransaction failed: %w", err)
	} else {
		tx.Signatures = sigs
	}

	// make client
	client := makeCommandClient(grpcConn)

	// send tx
	if txHash, err := client.SendTransaction(ctx, tx); err != nil {
		return nil, fmt.Errorf("SendTransaction failed: %w", err)
	} else if res, err := client.TxStatusStream(ctx, txHash); err != nil {
		return nil, fmt.Errorf("TxStatusStream failed: %w", err)
	} else {
		return res, nil
	}
}

func sendIrohaQuery(ctx context.Context, grpcConn *grpc.ClientConn, q *protocol.Query, key string) (*protocol.QueryResponse, error) {
	// sign query
	if sig, err := crypto.SignQuery(q, key); err != nil {
		return nil, fmt.Errorf("SignQuery failed: %w", err)
	} else {
		q.Signature = sig
	}

	// make client
	client := makeQueryClient(grpcConn)

	// send query
	if res, err := client.SendQuery(ctx, q); err != nil {
		return nil, fmt.Errorf("SendQuery failed: %w", err)
	} else {
		return res, nil
	}
}
