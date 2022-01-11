package e2e

import (
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type RoleTestSuite struct {
	TestSuite
}

func (suite *RoleTestSuite) TestRole() {
	var TestGetRoles = suite.AddUnixSuffix("test_get_roles", "_")
	{
		// check current role first
		q := query.GetRoles(
			query.CreatorAccountId(AdminAccountId),
		)
		res := suite.SendQuery(q, AdminPrivateKey)
		roles := res.GetRolesResponse().Roles
		suite.T().Logf("admin roles: %v", roles)
		// admin, user, money_creator, evm_admin, gateway_querier
		suite.Require().NotContains(roles, TestGetRoles)
	}

	{
		// check current role by user account, but it must be failed
		q := query.GetRoles(
			query.CreatorAccountId(UserAccountId),
		)
		_, err := suite.SendQueryWithError(q, UserPrivateKey)
		suite.Require().Error(err)
		// failed to execute query: user must have at least one of the permissions: can_get_roles, " error_code:2
	}

	{
		// create role `can_get_roles` by admin
		tx := suite.BuildTransaction(
			command.CreateRole(TestGetRoles, []pb.RolePermission{pb.RolePermission_can_get_roles}),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// check current role
		q := query.GetRoles(
			query.CreatorAccountId(AdminAccountId),
		)
		res := suite.SendQuery(q, AdminPrivateKey)
		roles := res.GetRolesResponse().Roles
		suite.Require().Contains(roles, TestGetRoles)
	}

	{
		// append TestGetRoles role to user account by admin
		tx := suite.BuildTransaction(
			command.AppendRole(UserAccountId, TestGetRoles),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// try checking current role by user account
		q := query.GetRoles(
			query.CreatorAccountId(UserAccountId),
		)
		res := suite.SendQuery(q, UserPrivateKey)
		userRoles := res.GetRolesResponse().Roles
		suite.T().Logf("user roles: %v", userRoles)
	}

	{
		// detach role TestGetRoles role from user account by admin
		tx := suite.BuildTransaction(
			command.DetachRole(UserAccountId, TestGetRoles),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// check current role by user account, but it must be failed
		q := query.GetRoles(
			query.CreatorAccountId(UserAccountId),
		)
		_, err := suite.SendQueryWithError(q, UserPrivateKey)
		suite.Require().Error(err)
	}

	// Note: no remove roll commands
}

func TestRoleTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}
