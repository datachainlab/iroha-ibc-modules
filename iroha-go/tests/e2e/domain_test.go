package e2e

import (
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DomainTestSuite struct {
	TestSuite
}

func (suite *DomainTestSuite) TestDomain() {
	var TestDomainId = suite.AddUnixSuffix("testdomain", "")
	{
		// create domain
		tx := suite.BuildTransaction(
			command.CreateDomain("admin", TestDomainId),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	// Note: no remove domain commands
}

func TestDomainTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(DomainTestSuite))
}
