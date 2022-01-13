package e2e

import (
	"testing"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"

	"github.com/stretchr/testify/suite"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type RoleTestSuite struct {
	TestSuite
}

func (suite *RoleTestSuite) TestRole() {
	var getRoleName = suite.AddUnixSuffix("get_roles", "_")
	var getRolePermission = []pb.RolePermission{pb.RolePermission_can_get_roles}

	{
		// check current all roles by admin
		roles := suite.GetRoles(AdminAccountId, AdminPrivateKey)
		suite.Require().NotContains(roles, getRoleName)
	}

	suite.CreateRole(getRoleName, getRolePermission)
	{
		// role permission
		permissions := suite.GetRolePermissions(getRoleName)
		suite.Require().Contains(permissions, pb.RolePermission_can_get_roles)
	}

	{
		// check current role by user account, but it must be failed
		err := suite.getRolesByUser()
		suite.Require().Error(err)
		// failed to execute query: user must have at least one of the permissions: can_get_roles, " error_code:2

		// then, append role to user
		suite.AppendRole(UserAccountId, getRoleName)
		// try checking current all roles by user account
		roles := suite.GetRoles(UserAccountId, UserPrivateKey)
		suite.Require().Contains(roles, getRoleName)

		// detach role from user account by admin
		suite.DetachRole(UserAccountId, getRoleName)

		// check current all roles by user account, but it must be failed
		err = suite.getRolesByUser()
		suite.Require().Error(err)
	}
	// Note: no remove roll commands
}

func (suite *RoleTestSuite) getRolesByUser() error {
	// check current role by user account, but it must be failed
	q := query.GetRoles(
		query.CreatorAccountId(UserAccountId),
	)
	_, err := suite.SendQueryWithError(q, UserPrivateKey)
	return err
}

func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}
