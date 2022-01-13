package e2e

import (
	"testing"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/stretchr/testify/suite"
)

type SetValueTestSuite struct {
	TestSuite
}

func (suite *SetValueTestSuite) TestSetValue() {
	var (
		key   = suite.AddUnixSuffix("key", "_")
		value = suite.AddUnixSuffix("value", "_")
	)

	{
		// FIXME: Currently SetSettingValue is only allowed in genesis block
		tx := suite.BuildTransaction(
			command.SetSettingValue(key, value),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}
}

func TestSetValueTestSuite(t *testing.T) {
	t.SkipNow()
	suite.Run(t, new(SetValueTestSuite))
}
