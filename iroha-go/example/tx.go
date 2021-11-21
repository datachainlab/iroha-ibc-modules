package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func Tx() {
	conn, err := conn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	commandClient := command.New(conn, time.Second*60)

	assetID := randStringRunes(6)
	fmt.Println("assetID:", assetID)
	tx := commandClient.BuildTransaction(
		commandClient.BuildPayload(
			[]*protocol.Command{
				commandClient.CreateAsset(assetID, "test", 2),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)
	txHash, err := commandClient.SendTransaction(
		context.Background(),
		tx,
		AdminPrivateKey,
	)
	if err != nil {
		panic(err)
	}

	status, err := commandClient.TxStatusStream(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", status)
}
