package e2e

import (
	"fmt"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
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
		// FIXME: after running this command, subsequent commands fail
		// rpc error: code = DeadlineExceeded desc = Deadline Exceeded
		tx := suite.BuildTransaction(
			command.AddPeer(fmt.Sprintf("127.0.0.1:%d", port), pubKey, nil),
			AdminAccountId,
		)
		suite.SendTransaction(tx, AdminPrivateKey)
	}

	{
		// check peer
		q := query.GetPeers(
			query.CreatorAccountId(UserAccountId),
		)
		res := suite.SendQuery(q, UserPrivateKey)
		peers := res.GetPeersResponse().Peers
		suite.Require().Equal(2, len(peers))
	}

	{
		// remove peer
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
