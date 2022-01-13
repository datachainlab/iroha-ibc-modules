package e2e

import (
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	TestSuite
}

func (suite *DomainTestSuite) TestTransaction() {
	var accountName = suite.AddUnixSuffix("test_account", "_")
	pubKey, _ := suite.CreateKeyPair()

	hash := suite.CreateAccount(accountName, pubKey)
	{
		txs := suite.getTransaction(hash)
		suite.Require().Condition(func() bool {
			if len(txs) == 0 {
				return false
			}
			return true
		}, "transaction must be more than 0")
	}
}

func (suite *DomainTestSuite) getTransaction(hash string) []*pb.Transaction {
	q := query.GetTransactions(
		[]string{hash},
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetTransactionsResponse().Transactions
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}
