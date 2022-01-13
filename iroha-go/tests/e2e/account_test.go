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
		grantSetMyQuorumRole = suite.AddUnixSuffix("grant_set_my_quorum", "_")
		signatoryRoleName    = suite.AddUnixSuffix("signatory", "_")
	)
	suite.CreateRole(grantSetMyQuorumRole, []pb.RolePermission{pb.RolePermission_can_grant_can_set_my_quorum})
	suite.CreateRole(signatoryRoleName, []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory})

	// Scenario1: user call `SetAccountQuorum` to user itself
	suite.AppendRole(UserAccountId, signatoryRoleName)
	{
		// add account
		var accountName = suite.AddUnixSuffix("test_account", "_")
		//var accountId = fmt.Sprintf("%s@%s", accountName, DomainId)
		pubKey, privKey := suite.CreateKeyPair()
		suite.CreateAccount(accountName, pubKey)

		suite.AddSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)

		keys := suite.GetSignatory(UserAccountId)
		quorum := uint32(len(keys))
		suite.setAccountQuorum(UserAccountId, quorum, UserAccountId, UserPrivateKey)

		// revert condition
		multiSig := MultiSigInfo{
			Quorum:          quorum,
			AccountId:       UserAccountId,
			AccountPrivKeys: []string{UserPrivateKey, privKey},
		}
		suite.setAccountQuorumWithMultiSig(UserAccountId, quorum-1, &multiSig)
		suite.RemoveSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)
	}

	// Scenario2: admin call `SetAccountQuorum` to user
	suite.AppendRole(UserAccountId, grantSetMyQuorumRole)
	{
		// FIXME: if same test runs twice, error happens
		// - tx_status:REJECTED  tx_hash:"d6996039f71366f446d870893804f877a2b88ad9bf4a240af89c57d680bbf09f"
		suite.GrantPermission(AdminAccountId, pb.GrantablePermission_can_set_my_quorum, UserAccountId, UserPrivateKey)

		pubKey, privKey := suite.CreateKeyPair()
		suite.AddSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)

		keys := suite.GetSignatory(UserAccountId)
		quorum := uint32(len(keys))
		suite.setAccountQuorum(UserAccountId, quorum, AdminAccountId, AdminPrivateKey)

		// revert condition
		multiSig := MultiSigInfo{
			Quorum:          quorum,
			AccountId:       UserAccountId,
			AccountPrivKeys: []string{UserPrivateKey, privKey},
		}
		suite.setAccountQuorumWithMultiSig(UserAccountId, quorum-1, &multiSig)
		suite.RemoveSignatory(UserAccountId, pubKey, UserAccountId, UserPrivateKey)
		suite.RevokePermission(AdminAccountId, pb.GrantablePermission_can_set_my_quorum, UserAccountId, UserPrivateKey)
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
	// FIXME below error
	// - message:"SignedData: [Child errors=[Signatures list: [Child errors=[Signature #1 (Signature: [publicKey=313a07e6384776ed95447710d15e59148473ccfc052a681317a72a69f2a49910, signedData=754639601154dca07ede7b3a30a53c61eecb4ad93f2a44b6a2739b6333ccd91affcfa875b52b9e1c51bc1ed8105ae9e40069e9717ddfd0977924b5b4d1cbcb05]): [Errors=[Bad signature.]]]]]]"
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

type MultiSigInfo struct {
	Quorum          uint32
	AccountId       string
	AccountPrivKeys []string
}

func (suite *AccountTestSuite) setAccountQuorumWithMultiSig(targetAccountId string, quorum uint32, multiSig *MultiSigInfo) string {
	tx := suite.BuildTransactionWithQuorum(
		command.SetAccountQuorum(targetAccountId, quorum),
		multiSig.AccountId,
		multiSig.Quorum,
	)
	return suite.SendTransactions(tx, multiSig.AccountPrivKeys...)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
