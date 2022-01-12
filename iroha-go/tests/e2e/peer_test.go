package e2e

import (
	"fmt"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PeerTestSuite struct {
	TestSuite
}

func (suite *PeerTestSuite) TestPeer() {
	var port = suite.RandInt(20000, 65535)
	pubKey, _, err := suite.CreateKeyPair()
	suite.Require().NoError(err)
	{
		// add peer
		tx := suite.BuildTransaction(
			command.AddPeer(fmt.Sprintf("127.0.0.1:%d", port), pubKey, nil),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// remove peer
		// FIXME: rpc error: code = DeadlineExceeded desc = Deadline Exceeded
		tx := suite.BuildTransaction(
			command.RemovePeer(pubKey),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}
}

func TestPeerTestSuiteTestSuite(t *testing.T) {
	t.SkipNow()
	suite.Run(t, new(PeerTestSuite))
}
