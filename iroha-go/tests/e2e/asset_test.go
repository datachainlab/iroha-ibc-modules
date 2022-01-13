package e2e

import (
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

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
		assetName         = strings.Split(AssetId, "#")[0]
	)

	assets := suite.getAccountAssetWithError(AssetId, AdminAccountId)

	// createAsset if not exists
	if len(assets) == 0 {
		suite.createAsset(assetName, DomainId, precision)
	} else {
		balance, err = strconv.ParseFloat(assets[0].Balance, 64)
		suite.NoError(err)
	}

	suite.addAssetQuantity(AssetId, strconv.FormatFloat(amount, 'f', int(precision), 64))

	{
		assets := suite.getAccountAsset(AssetId, AdminAccountId)
		suite.Equal(assets[0].AssetId, AssetId)
		suite.Equal(assets[0].Balance, strconv.FormatFloat(balance+amount, 'f', int(precision), 64))
	}

	{
		// FIXME: no transaction for now
		// - how can account asset transactions be retrieved?
		suite.getAccountAssetTransactions(AssetId, AdminAccountId)
	}

	{
		asset := suite.getAssetInfo(AssetId)
		suite.T().Logf("asset: %v", asset)
		suite.Equal(asset.AssetId, AssetId)
		suite.Equal(asset.DomainId, DomainId)
		suite.Equal(asset.Precision, precision)
	}
}

func (suite *AssetTestSuite) getAccountAssetWithError(assetId string, targetAccountId string) []*pb.AccountAsset {
	q := query.GetAccountAsset(
		targetAccountId,
		&pb.AssetPaginationMeta{
			PageSize:        math.MaxUint32,
			OptFirstAssetId: &pb.AssetPaginationMeta_FirstAssetId{FirstAssetId: assetId},
		},
		query.CreatorAccountId(AdminAccountId),
	)

	res, err := suite.SendQueryWithError(q, AdminPrivateKey)
	if err != nil {
		if strings.Contains(err.Error(), "error_code:4") {
			return nil
		}
		suite.NoError(err)
		return nil
	}
	return res.GetAccountAssetsResponse().AccountAssets
}

func (suite *AssetTestSuite) getAccountAsset(assetId string, targetAccountId string) []*pb.AccountAsset {
	q := query.GetAccountAsset(
		targetAccountId,
		&pb.AssetPaginationMeta{
			PageSize:        math.MaxUint32,
			OptFirstAssetId: &pb.AssetPaginationMeta_FirstAssetId{FirstAssetId: assetId},
		},
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetAccountAssetsResponse().AccountAssets
}

func (suite *AssetTestSuite) getAccountAssetTransactions(assetId string, targetAccountId string) []*pb.Transaction {
	q := query.GetAccountAssetTransactions(
		targetAccountId,
		assetId,
		&pb.TxPaginationMeta{PageSize: math.MaxUint32},
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetTransactionsPageResponse().Transactions
}

func (suite *AssetTestSuite) getAssetInfo(assetId string) *pb.Asset {
	q := query.GetAssetInfo(
		assetId,
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetAssetResponse().Asset
}

func (suite *AssetTestSuite) createAsset(assetName, domainID string, precision uint32) string {
	tx := suite.BuildTransaction(
		command.CreateAsset(assetName, domainID, precision),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *AssetTestSuite) addAssetQuantity(assetID, amount string) string {
	tx := suite.BuildTransaction(
		command.AddAssetQuantity(assetID, amount),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func TestAssetTestSuite(t *testing.T) {
	suite.Run(t, new(AssetTestSuite))
}
