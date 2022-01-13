package e2e

import (
	"fmt"
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

// TODO: test for CompareAndSetAccountDetail

func (suite *AccountTestSuite) TestCreateAccount() {
	//suite.T().SkipNow()
	var accountName = suite.AddUnixSuffix("test_account", "_")
	var accountId = fmt.Sprintf("%s@%s", accountName, DomainId)
	pubKey, _ := suite.CreateKeyPair()

	{
		suite.CreateAccount(accountName, pubKey)

		account := suite.getAccount(accountId)
		suite.Require().NotNil(account)

		txs := suite.getAccountTransactions(AdminAccountId)
		suite.Require().Condition(func() bool {
			if len(txs) == 0 {
				return false
			}
			return true
		}, "transaction must be more than 0")
	}
	// Note: no remove account commands
}

func (suite *AccountTestSuite) TestSetAccountDetail() {
	//suite.T().SkipNow()
	var (
		key   = suite.randStringRunes(10)
		value = suite.randStringRunes(10)
	)

	{
		suite.setAccountDetail(AdminAccountId, key, value)
		detail := suite.getAccountDetail(AdminAccountId, key)
		expected := fmt.Sprintf("{ \"%s\" : { \"%s\" : \"%s\" } }", AdminAccountId, key, value)
		suite.Equal(expected, detail)
	}
}

func (suite *AccountTestSuite) TestSetAccountQuorum() {
	var (
		setQuorumRole        = suite.AddUnixSuffix("set_quorum", "_")
		grantSetMyQuorumRole = suite.AddUnixSuffix("grant_set_my_quorum", "_")
		signatoryRoleName    = suite.AddUnixSuffix("signatory", "_")
	)
	suite.CreateRole(setQuorumRole, []pb.RolePermission{pb.RolePermission_can_set_quorum})
	suite.CreateRole(grantSetMyQuorumRole, []pb.RolePermission{pb.RolePermission_can_grant_can_set_my_quorum})
	suite.CreateRole(signatoryRoleName, []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory})

	// Scenario1: user call `SetAccountQuorum` to user itself
	suite.AppendRole(UserAccountId, signatoryRoleName)
	{
		pubKey, _ := suite.CreateKeyPair()
		suite.AddSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)

		keys := suite.GetSignatory(UserAccountId)
		suite.setAccountQuorum(UserAccountId, uint32(len(keys)), UserAccountId, UserPrivateKey)
	}

	// Scenario2: admin call `SetAccountQuorum` to admin itself
	{
		pubKey, _ := suite.CreateKeyPair()
		suite.AddSignatory(AdminAccountId, pubKey, AdminAccountId, AdminPrivateKey)

		keys := suite.GetSignatory(AdminAccountId)
		suite.setAccountQuorum(AdminAccountId, uint32(len(keys)), AdminAccountId, AdminPrivateKey)
	}

	// Scenario3: admin call `SetAccountQuorum` to user
	suite.AppendRole(UserAccountId, grantSetMyQuorumRole)
	suite.GrantPermission(AdminAccountId, pb.GrantablePermission_can_set_my_quorum, UserAccountId, UserPrivateKey)
	{
		pubKey, _ := suite.CreateKeyPair()
		suite.AddSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)

		keys := suite.GetSignatory(UserAccountId)
		suite.setAccountQuorum(UserAccountId, uint32(len(keys)), AdminAccountId, AdminPrivateKey)
	}
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

func (suite *AccountTestSuite) getAccount(targetAccountId string) *pb.Account {
	q := query.GetAccount(
		targetAccountId,
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetAccountResponse().GetAccount()
}

func (suite *AccountTestSuite) getAccountTransactions(targetAccountId string) []*pb.Transaction {
	q := query.GetAccountTransactions(
		targetAccountId,
		&pb.TxPaginationMeta{PageSize: math.MaxUint32},
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetTransactionsPageResponse().Transactions
}

func (suite *AccountTestSuite) setAccountDetail(targetAccountId, key, value string) string {
	tx := suite.BuildTransaction(
		command.SetAccountDetail(targetAccountId, key, value),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *AccountTestSuite) getAccountDetail(targetAccountId, key string) string {
	// below error
	// - Transaction deserialization failed: hash Hash: [c7aeb854bcca94c70dca7f1182a99287f6248a548a16c57063f8e496150cdc91], SignedData: [Child errors=[Transaction: [Child errors=[Command #1: [Child errors=[AppendRole: [Child errors=[AccountId: [Errors=[passed value: 'f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70' does not match regex '[a-z_0-9]{1,32}\@([a-zA-Z]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?'.]]]]]]]]]]
	// - paginationMeta is replaced by nil to avoid the error for now. but it must be fixed
	q := query.GetAccountDetail(
		&pb.GetAccountDetail_AccountId{AccountId: targetAccountId},
		&pb.GetAccountDetail_Key{Key: key},
		&pb.GetAccountDetail_Writer{Writer: targetAccountId},
		//&pb.AccountDetailPaginationMeta{PageSize: math.MaxUint32},
		nil,
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetAccountDetailResponse().GetDetail()
}

func (suite *AccountTestSuite) setAccountQuorum(targetAccountId string, quorum uint32, byAccountId, byAccountPrivKey string) string {
	tx := suite.BuildTransaction(
		command.SetAccountQuorum(targetAccountId, quorum),
		byAccountId,
	)
	return suite.SendTransaction(tx, byAccountPrivKey)
}

func TestAccountTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
