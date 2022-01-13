package e2e

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PeerTestSuite struct {
	TestSuite
}

func (suite *PeerTestSuite) TestPeer() {
	var port = suite.RandInt(20000, 65535)
	address := fmt.Sprintf("127.0.0.1:%d", port)
	pubKey, _, err := suite.CreateKeyPair()
	suite.Require().NoError(err)
	{
		// add peer
		// FIXME: after running this command, subsequent commands fail
		// rpc error: code = DeadlineExceeded desc = Deadline Exceeded
		suite.AddPeer(address, pubKey)

		peers := suite.GetPeers()
		suite.Require().Equal(2, len(peers))

		suite.RemovePeer(pubKey)
	}
}

func TestPeerTestSuiteTestSuite(t *testing.T) {
	t.SkipNow()
	suite.Run(t, new(PeerTestSuite))
}
