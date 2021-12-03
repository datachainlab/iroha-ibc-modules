package e2e

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

type AccountTestSuite struct {
	TestSuite
}

func (suite *AccountTestSuite) TestSetAccountDetail() {
	var (
		key   = suite.randStringRunes(10)
		value = suite.randStringRunes(10)
	)

	{
		tx := suite.BuildTransaction(
			command.SetAccountDetail(AdminAccountId, key, value),
			AdminAccountId,
		)

		suite.SendTransaction(tx, AdminPrivateKey)
	}

	q := query.GetAccountDetail(
		&pb.GetAccountDetail_AccountId{AccountId: AdminAccountId},
		&pb.GetAccountDetail_Key{Key: key},
		&pb.GetAccountDetail_Writer{Writer: AdminAccountId},
		&pb.AccountDetailPaginationMeta{PageSize: math.MaxUint32},
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	detail := res.GetAccountDetailResponse().GetDetail()
	suite.Equal(detail, value)
}

func (suite *AccountTestSuite) TestSetAccountQuorum() {
	{
		tx := suite.BuildTransaction(
			command.SetAccountQuorum(UserAccountId, 2),
			AdminAccountId,
		)

		suite.SendTransaction(tx, AdminPrivateKey)
	}

	q := query.GetAccount(
		UserAccountId,
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	acc := res.GetAccountResponse().GetAccount()
	suite.Require().NotNil(acc)
}

func (suite *AccountTestSuite) randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestAccountTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
