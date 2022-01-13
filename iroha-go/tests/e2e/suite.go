package e2e

import (
	"context"
	"crypto/rand"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	mathrand "math/rand"
	"strconv"
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
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
	suite.Require().Condition(func() bool {
		if res.ErrorCode != 0 {
			return false
		}
		if res.ErrOrCmdName != "" {
			return false
		}
		if res.FailedCmdIndex != 0 {
			return false
		}
		return true
	}, "check *pb.ToriiResponse carefully")

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

func (suite *TestSuite) CreateKeyPair() (string, string, error) {
	pubKey, privKey, err := crypto.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	return pubKey.Hex(), privKey.Hex(), nil
}

func (suite *TestSuite) AddUnixSuffix(target, delimiter string) string {
	return fmt.Sprintf("%s%s%s", target, delimiter, strconv.FormatInt(time.Now().Unix(), 10))
}

func (suite *TestSuite) RandInt(min int, max int) int {
	mathrand.Seed(time.Now().UTC().UnixNano())
	return min + mathrand.Intn(max-min)
}

func (suite *TestSuite) CreateAccount(accountName, pubKey string) string {
	tx := suite.BuildTransaction(
		command.CreateAccount(accountName, DomainId, pubKey),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) CreateDomain(defaultRole, domainId string) string {
	tx := suite.BuildTransaction(
		command.CreateDomain(defaultRole, domainId),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) CreateRole(roleName string, permissions []pb.RolePermission) string {
	tx := suite.BuildTransaction(
		command.CreateRole(roleName, permissions),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) AppendRole(targetAccountId, roleName string) string {
	tx := suite.BuildTransaction(
		command.AppendRole(targetAccountId, roleName),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) DetachRole(targetAccountId, roleName string) string {
	tx := suite.BuildTransaction(
		command.DetachRole(targetAccountId, roleName),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) GetRoles(byAccountId, byAccountPrivKey string) []string {
	// check current role first
	q := query.GetRoles(
		query.CreatorAccountId(byAccountId),
	)
	res := suite.SendQuery(q, byAccountPrivKey)
	return res.GetRolesResponse().Roles
	// roles would like `admin, user, money_creator, evm_admin, gateway_querier`
}

func (suite *TestSuite) GetRolePermissions(roleName string) []pb.RolePermission {
	q := query.GetRolePermissions(
		roleName,
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetRolePermissionsResponse().Permissions
}

func (suite *TestSuite) GrantPermission(toUserAccountId string, permission pb.GrantablePermission, byAccountId, byAccountPrivKey string) string {
	tx := suite.BuildTransaction(
		command.GrantPermission(toUserAccountId, permission),
		byAccountId,
	)
	return suite.SendTransaction(tx, byAccountPrivKey)
}

func (suite *TestSuite) RevokePermission(fromUserAccountId string, permission pb.GrantablePermission, byAccountId, byAccountPrivKey string) string {
	tx := suite.BuildTransaction(
		command.RevokePermission(fromUserAccountId, permission),
		byAccountId,
	)
	return suite.SendTransaction(tx, byAccountPrivKey)
}

func (suite *TestSuite) AddSignatory(targetAccountId, pubKey string) string {
	tx := suite.BuildTransaction(
		command.AddSignatory(targetAccountId, pubKey),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) RemoveSignatory(targetAccountId, pubKey string) string {
	tx := suite.BuildTransaction(
		command.RemoveSignatory(targetAccountId, pubKey),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

// GetSignatory gets signatory, and returns keys
func (suite *TestSuite) GetSignatory(targetAccountId string) []string {
	q := query.GetSignatories(
		targetAccountId,
		query.CreatorAccountId(AdminAccountId),
	)

	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetSignatoriesResponse().GetKeys()
}

func (suite *TestSuite) AddPeer(address, pubKey string) string {
	tx := suite.BuildTransaction(
		command.AddPeer(address, pubKey, nil),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}

func (suite *TestSuite) GetPeers() []*pb.Peer {
	q := query.GetPeers(
		query.CreatorAccountId(AdminAccountId),
	)
	res := suite.SendQuery(q, AdminPrivateKey)
	return res.GetPeersResponse().Peers
}

func (suite *TestSuite) RemovePeer(pubKey string) string {
	tx := suite.BuildTransaction(
		command.RemovePeer(pubKey),
		AdminAccountId,
	)
	return suite.SendTransaction(tx, AdminPrivateKey)
}
