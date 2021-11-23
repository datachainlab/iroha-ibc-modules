package example

import (
	"context"
	"fmt"
	"time"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

func Query() {
	conn, err := conn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queryClient := query.New(conn, time.Second*60)

	q := query.GetAccountAsset(
		AdminAccountId,
		&pb.AssetPaginationMeta{
			PageSize: 500,
		},
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

	for _, asset := range res.GetAccountAssetsResponse().GetAccountAssets() {
		fmt.Println(asset)
	}
}
