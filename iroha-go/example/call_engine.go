package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

const contractInput = "60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f9418a77c4cc64e8bdf3d840ca5c4cbec4b1f48f6004caa36249359db1348e7064736f6c634300080a0033"

func CallEngine() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	commandClient := command.New(conn, time.Second*60)

	var tx *pb.Transaction

	tx = command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.CallEngine(
					AdminAccountId,
					"",
					contractInput,
				),
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
