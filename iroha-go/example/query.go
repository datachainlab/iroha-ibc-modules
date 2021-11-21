package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

func Query() {
	conn, err := conn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queryClient := query.New(conn, time.Second*60)
	res, err := queryClient.SendQuery(
		context.Background(),
		queryClient.GetAccountAsset(
			AdminAccountId,
			nil,
			query.CreatorAccountId(AdminAccountId),
		),
		AdminPrivateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
