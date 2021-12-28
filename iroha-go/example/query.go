package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

func Query() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queryClient := query.New(conn, time.Second*60)
	q := query.GetAccountAsset(
		UserAccountId,
		nil,
		query.CreatorAccountId(AdminAccountId),
	)

	sig, err := crypto.SignQuery(q, AdminPrivateKey)
	if err != nil {
		panic(err)
	}

	q.Signature = sig

	res, err := queryClient.SendQuery(context.Background(), q)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.GetAccountAssetsResponse().GetAccountAssets())
	//for _, asset := range res.GetAccountDetailResponse().GetDetail() {
	//	fmt.Println(asset)
	//}
}
