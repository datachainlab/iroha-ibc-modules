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
	var TestAccountName = suite.AddUnixSuffix("test_account", "_")
	var TestAccountId = fmt.Sprintf("%s@%s", TestAccountName, DomainId)

	{
		// create key pair for new account
		pubKey, _, err := suite.CreateKeyPair()
		suite.Require().NoError(err)

		// create account
		tx := suite.BuildTransaction(
			command.CreateAccount(TestAccountName, DomainId, pubKey),
			AdminAccountId,
		)
		hash := suite.SendTransaction(tx, AdminPrivateKey)

		// check transaction by hash
		q := query.GetTransactions(
			[]string{hash},
			query.CreatorAccountId(AdminAccountId),
		)
		res := suite.SendQuery(q, AdminPrivateKey)
		txs := res.GetTransactionsResponse().Transactions
		suite.Require().Condition(func() bool {
			if len(txs) == 0 {
				return false
			}
			return true
		}, "transaction must be more than 0")

		// GetEngineReceipts
		// TODO: response has no receipts
		q = query.GetEngineReceipts(
			hash,
			query.CreatorAccountId(AdminAccountId),
		)
		res = suite.SendQuery(q, AdminPrivateKey)
		engineReceipts := res.GetEngineReceiptsResponse().EngineReceipts
		suite.T().Log(engineReceipts)
	}

	{
		// check new account
		q := query.GetAccount(
			TestAccountId,
			query.CreatorAccountId(AdminAccountId),
		)

		res := suite.SendQuery(q, AdminPrivateKey)
		acc := res.GetAccountResponse().GetAccount()
		suite.Require().NotNil(acc)
	}

	{
		q := query.GetAccountTransactions(
			AdminAccountId,
			&pb.TxPaginationMeta{PageSize: math.MaxUint32},
			query.CreatorAccountId(AdminAccountId),
		)

		res := suite.SendQuery(q, AdminPrivateKey)
		txs := res.GetTransactionsPageResponse().Transactions
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

	// below error
	// - Transaction deserialization failed: hash Hash: [c7aeb854bcca94c70dca7f1182a99287f6248a548a16c57063f8e496150cdc91], SignedData: [Child errors=[Transaction: [Child errors=[Command #1: [Child errors=[AppendRole: [Child errors=[AccountId: [Errors=[passed value: 'f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70' does not match regex '[a-z_0-9]{1,32}\@([a-zA-Z]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?'.]]]]]]]]]]
	// - paginationMeta is replaced by nil to avoid the error for now. but it must be fixed
	q := query.GetAccountDetail(
		&pb.GetAccountDetail_AccountId{AccountId: AdminAccountId},
		&pb.GetAccountDetail_Key{Key: key},
		&pb.GetAccountDetail_Writer{Writer: AdminAccountId},
		//&pb.AccountDetailPaginationMeta{PageSize: math.MaxUint32},
		nil,
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	// it returns only value of key
	detail := res.GetAccountDetailResponse().GetDetail()
	expected := fmt.Sprintf("{ \"%s\" : { \"%s\" : \"%s\" } }", AdminAccountId, key, value)
	suite.Equal(expected, detail)
}

func (suite *AccountTestSuite) TestSetAccountQuorum() {
	// Scenario1: user call `SetAccountQuorum` to user itself
	var (
		setQuorumRole        = suite.AddUnixSuffix("set_quorum", "_")
		grantSetMyQuorumRole = suite.AddUnixSuffix("grant_set_my_quorum", "_")
	)
	{
		// create role RolePermission_can_set_quorum
		tx := suite.BuildTransaction(
			command.CreateRole(setQuorumRole, []pb.RolePermission{pb.RolePermission_can_set_quorum}),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)

		// create role RolePermission_can_grant_can_set_my_quorum
		tx = suite.BuildTransaction(
			command.CreateRole(grantSetMyQuorumRole, []pb.RolePermission{pb.RolePermission_can_grant_can_set_my_quorum}),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}
	{
		// append setQuorumRole to admin
		//tx := suite.BuildTransaction(
		//	command.AppendRole(AdminAccountId, setQuorumRole),
		//	AdminAccountId,
		//)
		//suite.SendTransaction(tx, AdminPrivateKey)

		// append grantSetMyQuorumRole to user
		//tx := suite.BuildTransaction(
		//	command.AppendRole(UserAccountId, grantSetMyQuorumRole),
		//	AdminAccountId,
		//)
		//suite.SendTransaction(tx, AdminPrivateKey)
	}

	var signatoryRoles = suite.AddUnixSuffix("signatory", "_")
	{
		// create role `RolePermission_can_add_signatory`
		tx := suite.BuildTransaction(
			command.CreateRole(signatoryRoles, []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory}),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)

		// append roles to user
		tx = suite.BuildTransaction(
			command.AppendRole(UserAccountId, signatoryRoles),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}
	{
		pubKey, _, err := suite.CreateKeyPair()
		suite.Require().NoError(err)
		// add signatory to user account by user account
		tx := suite.BuildTransaction(
			command.AddSignatory(UserAccountId, pubKey),
			UserAccountId,
		)
		suite.SendTransaction(tx, UserPrivateKey)
	}

	{
		tx := suite.BuildTransaction(
			command.SetAccountQuorum(UserAccountId, 2),
			UserAccountId,
		)
		suite.SendTransaction(tx, UserPrivateKey)

		// FIXME:
		// command 'SetQuorum' with index '0' did not pass verification with code '2', query arguments: SetQuorum
		//tx := suite.BuildTransaction(
		//	command.SetAccountQuorum(UserAccountId, 2),
		//	AdminAccountId,
		//)
		//suite.SendTransaction(tx, AdminPrivateKey)
	}

	//q := query.GetAccount(
	//	UserAccountId,
	//	query.CreatorAccountId(AdminAccountId),
	//)
	//
	//res := suite.SendQuery(q, AdminPrivateKey)
	//acc := res.GetAccountResponse().GetAccount()
	//suite.Require().NotNil(acc)
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
