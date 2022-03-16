package client

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

type IrohadClient struct {
	conn          *grpc.ClientConn
	commandClient command.CommandClient
	queryClient   query.QueryClient

	accountIds map[uint32]string
	keys       map[uint32]string
}

func NewIrohadClient(endpoint string, accountIds map[uint32]string, keys map[uint32]string) *IrohadClient {
	conn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}
	commandClient := command.New(conn, time.Minute)
	queryClient := query.New(conn, time.Minute)

	return &IrohadClient{
		conn:          conn,
		commandClient: commandClient,
		queryClient:   queryClient,
		accountIds:    accountIds,
		keys:          keys,
	}
}

func (cli *IrohadClient) AccountIdOf(index uint32) string {
	if accountId, ok := cli.accountIds[index]; !ok {
		panic("account id not found")
	} else {
		return accountId
	}
}

func (cli *IrohadClient) keyOf(index uint32) string {
	if key, ok := cli.keys[index]; !ok {
		panic("key not found")
	} else {
		return key
	}
}

func (cli *IrohadClient) signTx(tx *protocol.Transaction, signers ...uint32) error {
	var keys []string
	for _, signer := range signers {
		keys = append(keys, cli.keyOf(signer))
	}
	if sigs, err := crypto.SignTransaction(tx, keys...); err != nil {
		return err
	} else {
		tx.Signatures = sigs
		return nil
	}
}

func (cli *IrohadClient) sendTx(ctx context.Context, tx *protocol.Transaction) error {
	if txHash, err := cli.commandClient.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("SendTransaction failed: %v", err)
	} else if _, err := cli.commandClient.TxStatusStream(ctx, txHash); err != nil {
		return fmt.Errorf("TxStatusStream failed: %v", err)
	}
	return nil
}

func (cli *IrohadClient) signQuery(q *protocol.Query, signer uint32) error {
	// MEMO: query is always signed by user 0.
	if sig, err := crypto.SignQuery(q, cli.keyOf(signer)); err != nil {
		return err
	} else {
		q.Signature = sig
		return nil
	}
}

func (cli *IrohadClient) sendQuery(ctx context.Context, q *protocol.Query) (*protocol.QueryResponse, error) {
	return cli.queryClient.SendQuery(ctx, q)
}

func (cli *IrohadClient) AddAssetQuantity(ctx context.Context, signer uint32, assetId, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.AddAssetQuantity(assetId, amount),
			},
			command.CreatorAccountId(cli.AccountIdOf(signer)),
		),
	)
	if err := cli.signTx(tx, signer); err != nil {
		return err
	} else if err := cli.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (cli *IrohadClient) SubtractAssetQuantity(ctx context.Context, signer uint32, assetId, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.SubtractAssetQuantity(assetId, amount),
			},
			command.CreatorAccountId(cli.AccountIdOf(signer)),
		),
	)
	if err := cli.signTx(tx, signer); err != nil {
		return err
	} else if err := cli.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (cli *IrohadClient) TransferAsset(ctx context.Context, signer uint32, srcAccountId, destAccountId, assetID, description, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.TransferAsset(srcAccountId, destAccountId, assetID, description, amount),
			},
			command.CreatorAccountId(cli.AccountIdOf(signer)),
		),
	)
	if err := cli.signTx(tx, signer); err != nil {
		return err
	} else if err := cli.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (cli *IrohadClient) GetAccountAsset(ctx context.Context, signer uint32, accountId, assetId string) (int, error) {
	q := query.GetAccountAsset(accountId, nil, query.CreatorAccountId(accountId))
	if err := cli.signQuery(q, signer); err != nil {
		return 0, err
	} else if res, err := cli.sendQuery(ctx, q); err != nil {
		return 0, err
	} else {
		assets := res.GetAccountAssetsResponse().AccountAssets
		for _, a := range assets {
			if a.AssetId == assetId {
				return strconv.Atoi(a.Balance)
			}
		}
		return 0, nil
	}
}
