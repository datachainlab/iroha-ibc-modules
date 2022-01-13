package e2e

import (
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
		suite.CreateDomain("admin", domainId)
	}
	// Note: no remove domain commands
}

func TestDomainTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(DomainTestSuite))
}
