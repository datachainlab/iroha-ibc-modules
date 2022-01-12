package e2e

import (
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignatoryTestSuite struct {
	TestSuite
}

func (suite *SignatoryTestSuite) TestSignatory() {
	var TestAccountName = suite.AddUnixSuffix("test_signatory", "_")
	pubKey, privKey, err := suite.CreateKeyPair()
	suite.Require().NoError(err)
	{
		// create key pair for new account
		suite.Require().NoError(err)
		suite.T().Logf("pubKey: %s, privKey: %s", pubKey, privKey)

		// create account
		tx := suite.BuildTransaction(
			command.CreateAccount(TestAccountName, DomainId, pubKey),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	var signatoryRoles = suite.AddUnixSuffix("signatory", "_")
	{
		// add rolls `RolePermission_can_add_signatory` to admin
		tx := suite.BuildTransaction(
			command.CreateRole(signatoryRoles, []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory}),
			//command.CreateRole(TestSignatoryRoles, []pb.RolePermission{pb.RolePermission_can_grant_can_add_my_signatory, pb.RolePermission_can_grant_can_remove_my_signatory}),
			//command.CreateRole(TestSignatoryRoles, []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory, pb.RolePermission_can_grant_can_add_my_signatory, pb.RolePermission_can_grant_can_remove_my_signatory}),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)

		// append roles to admin
		tx = suite.BuildTransaction(
			command.AppendRole(AdminAccountId, signatoryRoles),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// add created account signatory to admin account
		tx := suite.BuildTransaction(
			command.AddSignatory(AdminAccountId, pubKey),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// get signatory
		q := query.GetSignatories(
			AdminAccountId,
			query.CreatorAccountId(AdminAccountId),
		)

		res := suite.SendQuery(q, AdminPrivateKey)
		keys := res.GetSignatoriesResponse().GetKeys()
		suite.Require().Contains(keys, pubKey)
	}

	{
		// remove created account signatory from admin account
		tx := suite.BuildTransaction(
			command.RemoveSignatory(AdminAccountId, pubKey),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// get signatory
		q := query.GetSignatories(
			AdminAccountId,
			query.CreatorAccountId(AdminAccountId),
		)

		res := suite.SendQuery(q, AdminPrivateKey)
		keys := res.GetSignatoriesResponse().GetKeys()
		suite.Require().NotContains(keys, pubKey)
	}
}

func TestSignatoryTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(SignatoryTestSuite))
}
