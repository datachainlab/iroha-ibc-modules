package e2e

import (
	"context"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

const (
	ToriiAddress    = "localhost:50051"
	DomainId        = "test"
	AdminAccountId  = "admin@test"
	UserAccountId   = "test@test"
	AdminPrivateKey = "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70"
	UserPrivateKey  = "7e00405ece477bb6dd9b03a78eee4e708afc2f5bcdce399573a5958942f4a390"
	AssetId         = "testcoin#test"
)

type TestSuite struct {
	suite.Suite

	CommandClient command.CommandClient
	QueryClient   query.QueryClient
}

func (suite *TestSuite) SetupTest() {
	conn, err := grpc.Dial(
		ToriiAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	suite.NoError(err)

	suite.CommandClient = command.New(conn, time.Second*5)
	suite.QueryClient = query.New(conn, time.Second*5)
}

func (suite *TestSuite) BuildTransaction(cmd *pb.Command, accountID string) *pb.Transaction {
	return command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{cmd},
			command.CreatorAccountId(accountID),
		),
	)
}

func (suite *TestSuite) SendTransaction(tx *pb.Transaction, privKey string) string {
	sig, err := crypto.SignTransaction(tx, privKey)
	suite.Require().NoError(err)

	tx.Signatures = sig

	txHash, err := suite.CommandClient.SendTransaction(context.Background(), tx)
	suite.Require().NoError(err)

	res, err := suite.CommandClient.TxStatusStream(context.Background(), txHash)
	suite.Require().NoError(err)

	return res.TxHash
}

func (suite *TestSuite) SendQuery(query *pb.Query, privKey string) *pb.QueryResponse {
	sig, err := crypto.SignQuery(query, privKey)
	suite.Require().NoError(err)

	query.Signature = sig

	res, err := suite.QueryClient.SendQuery(context.Background(), query)
	suite.Require().NoError(err)

	return res
}

func (suite *TestSuite) SendQueryWithError(query *pb.Query, privKey string) (*pb.QueryResponse, error) {
	sig, err := crypto.SignQuery(query, privKey)
	suite.NoError(err)

	query.Signature = sig

	res, err := suite.QueryClient.SendQuery(context.Background(), query)

	return res, err
}
