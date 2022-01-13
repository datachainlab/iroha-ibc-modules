package e2e

import (
	"testing"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"

	"github.com/stretchr/testify/suite"
)

type SignatoryTestSuite struct {
	TestSuite
}

func (suite *SignatoryTestSuite) TestSignatory() {
	pubKey, _, err := suite.CreateKeyPair()
	suite.Require().NoError(err)

	var signatoryRoleName = suite.AddUnixSuffix("signatory", "_")
	var signatoryPermissions = []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory}

	// create role
	suite.CreateRole(signatoryRoleName, signatoryPermissions)
	// append role
	suite.AppendRole(AdminAccountId, signatoryRoleName)

	// test
	{
		// add signatory to admin
		suite.AddSignatory(AdminAccountId, pubKey)

		// get signatory
		keys := suite.GetSignatory(AdminAccountId)
		suite.Require().Contains(keys, pubKey)

		// remove signatory from admin account
		suite.RemoveSignatory(AdminAccountId, pubKey)

		// get signatory again
		keys = suite.GetSignatory(AdminAccountId)
		suite.Require().NotContains(keys, pubKey)
	}
}

func TestSignatoryTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(SignatoryTestSuite))
}
