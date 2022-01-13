package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type PermissionTestSuite struct {
	TestSuite
}

func (suite *PermissionTestSuite) TestPermission() {
	var roleName = suite.AddUnixSuffix("grant_detail", "_")
	suite.CreateRole(roleName, []pb.RolePermission{pb.RolePermission_can_grant_can_set_my_account_detail})
	suite.AppendRole(AdminAccountId, roleName)

	suite.GrantPermission(AdminAccountId, pb.GrantablePermission_can_set_my_account_detail, UserAccountId, UserPrivateKey)
	suite.RevokePermission(AdminAccountId, pb.GrantablePermission_can_set_my_account_detail, UserAccountId, UserPrivateKey)
}

func TestPermissionTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(PermissionTestSuite))
}
