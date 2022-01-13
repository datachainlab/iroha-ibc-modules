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
	var domainId = suite.AddUnixSuffix("testdomain", "")
	{
		// create domain
		suite.createDomain("admin", domainId)
	}
	// Note: no remove domain commands
}

func (suite *DomainTestSuite) createDomain(defaultRole, domainId string) string {
	tx := suite.BuildTransaction(
		command.CreateDomain(defaultRole, domainId),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func TestDomainTestSuite(t *testing.T) {
	suite.Run(t, new(DomainTestSuite))
}
