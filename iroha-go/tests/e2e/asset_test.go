package e2e

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

type AssetTestSuite struct {
	TestSuite
}

func (suite *AssetTestSuite) TestAsset() {
	var (
		err       error
		amount    float64 = 100.50
		balance   float64 = 0
		precision uint32  = 2
	)

	assets := suite.getAccountAsset()

	// CreateAsset if not exists
	if len(assets) == 0 {
		tx := suite.BuildTransaction(
			command.CreateAsset(strings.Split(AssetId, "#")[0], DomainId, precision),
			AdminAccountId,
		)

		suite.SendTransaction(tx, AdminPrivateKey)
	} else {
		balance, err = strconv.ParseFloat(assets[0].Balance, 64)
		suite.NoError(err)
	}

	// AddAssetQuantity
	tx := suite.BuildTransaction(
		command.AddAssetQuantity(AssetId, strconv.FormatFloat(amount, 'f', int(precision), 64)),
		AdminAccountId,
	)

	suite.SendTransaction(tx, AdminPrivateKey)

	q := query.GetAccountAsset(
		AdminAccountId,
		&pb.AssetPaginationMeta{
			PageSize:        math.MaxUint32,
			OptFirstAssetId: &pb.AssetPaginationMeta_FirstAssetId{FirstAssetId: AssetId},
		},
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	assets = res.GetAccountAssetsResponse().AccountAssets
	suite.Equal(assets[0].AssetId, AssetId)
	suite.Equal(assets[0].Balance, strconv.FormatFloat(balance+amount, 'f', int(precision), 64))

}

func (suite *AssetTestSuite) getAccountAsset() []*pb.AccountAsset {
	q := query.GetAccountAsset(
		AdminAccountId,
		&pb.AssetPaginationMeta{
			PageSize:        math.MaxUint32,
			OptFirstAssetId: &pb.AssetPaginationMeta_FirstAssetId{FirstAssetId: AssetId},
		},
		query.CreatorAccountId(AdminAccountId),
	)

	res, err := suite.SendQueryWithError(q, AdminPrivateKey)
	if err != nil {
		return nil
	}

	return res.GetAccountAssetsResponse().AccountAssets
}

func TestAssetTestSuite(t *testing.T) {
	suite.Run(t, new(AssetTestSuite))
}
