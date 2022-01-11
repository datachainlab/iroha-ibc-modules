package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func Tx() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	commandClient := command.New(conn, time.Second*60)

	pubKey, privKey, err := crypto.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("privKey", privKey.Hex())
	fmt.Println("pubKey", pubKey.Hex())

	accountID := randStringRunes(6)
	fmt.Println("accountID:", accountID)
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.CreateAccount(accountID, DomainId, pubKey.Hex()),
			},
			command.CreatorAccountId(AdminAccountId),
		),
	)

	sigs, err := crypto.SignTransaction(tx, AdminPrivateKey)
	if err != nil {
		panic(err)
	}

	tx.Signatures = sigs

	txHash, err := commandClient.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}

	status, err := commandClient.TxStatusStream(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", status)
}
