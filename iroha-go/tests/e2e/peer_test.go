package e2e

import (
	"fmt"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"testing"

	"github.com/stretchr/testify/suite"
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
		suite.addPeer(address, pubKey)

		peers := suite.getPeers()
		suite.Require().Equal(2, len(peers))

		suite.removePeer(pubKey)
	}
}

func (suite *PeerTestSuite) addPeer(address, pubKey string) string {
	tx := suite.BuildTransaction(
		command.AddPeer(address, pubKey, nil),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *PeerTestSuite) getPeers() []*pb.Peer {
	q := query.GetPeers(
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetPeersResponse().Peers
}

func (suite *PeerTestSuite) removePeer(pubKey string) string {
	tx := suite.BuildTransaction(
		command.RemovePeer(pubKey),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func TestPeerTestSuiteTestSuite(t *testing.T) {
	t.SkipNow()
	suite.Run(t, new(PeerTestSuite))
}
