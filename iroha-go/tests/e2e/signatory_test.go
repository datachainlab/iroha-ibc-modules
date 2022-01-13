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

	var signatoryRoleName = suite.AddUnixSuffix("signatory", "_")
	var signatoryPermissions = []pb.RolePermission{pb.RolePermission_can_add_signatory, pb.RolePermission_can_remove_signatory}
	suite.CreateRole(signatoryRoleName, signatoryPermissions)
	suite.AppendRole(AdminAccountId, signatoryRoleName)

	pubKey, _ := suite.CreateKeyPair()

	// test
	{
		// add signatory to admin
		suite.AddSignatory(AdminAccountId, pubKey, AdminAccountId, AdminPrivateKey)

		// get signatory
		keys := suite.GetSignatory(AdminAccountId)
		suite.Require().Contains(keys, pubKey)

		// remove signatory from admin account
		suite.RemoveSignatory(AdminAccountId, pubKey, AdminAccountId, AdminPrivateKey)

		// get signatory again
		keys = suite.GetSignatory(AdminAccountId)
		suite.Require().NotContains(keys, pubKey)
	}
}

func TestSignatoryTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(SignatoryTestSuite))
}
