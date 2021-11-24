package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func Batch() {
	conn, err := conn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	commandClient := command.New(conn, time.Second*60)

	assetID := randStringRunes(6)
	fmt.Println("assetID:", assetID)
	batchTx1 := commandClient.BuildTransaction(
		commandClient.BuildPayload(
			[]*protocol.Command{
				commandClient.CreateAsset(assetID, "test", 2),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)

	batchTx2 := commandClient.BuildTransaction(
		commandClient.BuildPayload(
			[]*protocol.Command{
				commandClient.AddAssetQuantity(fmt.Sprintf("%s#%s", assetID, DomainId), "100"),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)
	txList, err := commandClient.BuildBatchTransactions(
		[]*protocol.Transaction{batchTx1, batchTx2},
		protocol.Transaction_Payload_BatchMeta_ATOMIC,
	)
	if err != nil {
		panic(err)
	}
	txHashList, err := commandClient.SendBatchTransaction(
		context.Background(),
		txList,
		AdminPrivateKey,
	)
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
