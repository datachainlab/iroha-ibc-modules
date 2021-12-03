package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

type PermissionTestSuite struct {
	TestSuite
}

func (suite *PermissionTestSuite) TestGrantPermission() {
	{
		tx := suite.BuildTransaction(
			command.GrantPermission(AdminAccountId, pb.GrantablePermission_can_set_my_account_detail),
			UserAccountId,
		)
		suite.SendTransaction(tx, UserPrivateKey)
	}

	{
		tx := suite.BuildTransaction(
			command.RevokePermission(AdminAccountId, pb.GrantablePermission_can_set_my_account_detail),
			UserAccountId,
		)
		suite.SendTransaction(tx, UserPrivateKey)
	}
}

func TestPermissionTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(PermissionTestSuite))
}
