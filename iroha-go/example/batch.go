package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func Batch() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	commandClient := command.New(conn, time.Second*10)

	assetID := randStringRunes(6)
	fmt.Println("assetID:", assetID)
	batchTx1 := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.CreateAsset(assetID, DomainId, 2),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)

	batchTx2 := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.AddAssetQuantity(fmt.Sprintf("%s#%s", assetID, DomainId), "100"),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)

	batchTx3 := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.TransferAsset(
					AdminAccountId,
					UserAccountId,
					fmt.Sprintf("%s#%s", assetID, DomainId),
					"Transfer",
					"100"),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)
	txList, err := command.BuildBatchTransactions(
		[]*pb.Transaction{batchTx1, batchTx2, batchTx3},
		pb.Transaction_Payload_BatchMeta_ATOMIC,
	)
	if err != nil {
		panic(err)
	}

	for _, tx := range txList.Transactions {
		sigs, err := crypto.SignTransaction(tx, AdminPrivateKey)
		if err != nil {
			panic(err)
		}

		tx.Signatures = sigs
	}

	txHashList, err := commandClient.SendBatchTransaction(context.Background(), txList)
	if err != nil {
		panic(err)
	}

	for _, txHash := range txHashList {
		status, err := commandClient.TxStatusStream(context.Background(), txHash)
		if err != nil {
			panic(err)
		}

		fmt.Println("status:", status)
	}
}
