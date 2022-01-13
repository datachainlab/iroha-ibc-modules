package e2e

import (
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OthersTestSuite struct {
	TestSuite
}

func (suite *DomainTestSuite) TestGetEngineReceipts() {
	var accountName = suite.AddUnixSuffix("test_account", "_")
	pubKey, _ := suite.CreateKeyPair()

	hash := suite.CreateAccount(accountName, pubKey)
	{
		// GetEngineReceipts
		// TODO: response has no receipts
		q := query.GetEngineReceipts(
			hash,
			query.CreatorAccountId(AdminAccountId),
		)
		res := suite.SendQuery(q, AdminPrivateKey)
		engineReceipts := res.GetEngineReceiptsResponse().EngineReceipts
		suite.T().Log(engineReceipts)
	}
}

func TestOthersTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(OthersTestSuite))
}
