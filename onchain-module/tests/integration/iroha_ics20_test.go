package integration

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"

	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/channel"
	clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/client"
	ibctesting "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/testing"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/crypto"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

const (
	ToriiAddress = "localhost:50051"

	AssetId = "coin#test"

	AdminPrivateKey   = "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70"
	AlicePrivateKey   = "a6924c9781c46df18c05545735f127eaf788a60816a7aeb9d5e928460b51cb2f"
	BobPrivateKey     = "f66c1f19a52bf2955d00bf050793a80056ccfa6237b46f4d7d3a9e20af669c29"
	RelayerPrivateKey = "e517af47112e4f501afb26e4f34eadc8b0ad8eadaf4962169fc04bc8ddbfe091"
	TestPrivateKey    = "7e00405ece477bb6dd9b03a78eee4e708afc2f5bcdce399573a5958942f4a390"

	AdminAccountID   = "admin@test"
	AliceAccountID   = "alice@test"
	BobAccountID     = "bob@test"
	RelayerAccountID = "relayer@test"
	TestAccountID    = "test@test"

	relayer         = ibctesting.RelayerKeyIndex // the key-index of relayer
	deployer        = ibctesting.RelayerKeyIndex // the key-index of contract
	bankA           = ibctesting.RelayerKeyIndex // the key-index of bankA
	bankB           = ibctesting.RelayerKeyIndex // the key-index of bankB
	alice    uint32 = 1                          // the key-index of alice
	bob      uint32 = 2                          // the key-index of bob
)

type IrohaIcs20TestSuite struct {
	suite.Suite

	irohadConn    *grpc.ClientConn
	commandClient command.CommandClient
	queryClient   query.QueryClient
	accountIdOf   map[uint32]string
	privateKeyOf  map[uint32]string
	coordinator   ibctesting.Coordinator
	chainA        *ibctesting.Chain
	chainB        *ibctesting.Chain
}

func (suite *IrohaIcs20TestSuite) SetupTest() {
	var irohadConn *grpc.ClientConn
	var commandClient command.CommandClient
	var queryClient query.QueryClient
	{
		var err error
		irohadConn, err = grpc.Dial(
			"localhost:50051",
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			panic(err)
		}
		commandClient = command.New(irohadConn, time.Minute)
		queryClient = query.New(irohadConn, time.Minute)
	}

	accountIdOf := map[uint32]string{
		relayer: RelayerAccountID,
		alice:   AliceAccountID,
		bob:     BobAccountID,
	}

	privateKeyOf := map[uint32]string{
		relayer: RelayerPrivateKey,
		alice:   AlicePrivateKey,
		bob:     BobPrivateKey,
	}

	chainClient, err := client.NewETHClient("http://127.0.0.1:8545", clienttypes.MockClient)
	suite.Require().NoError(err)

	networkID := 1000
	accountIds := []string{"relayer@test", "alice@test", "bob@test"}
	contractConfig := ibctesting.NewTruffleContractConfig(networkID, "../../build/contracts")

	suite.irohadConn = irohadConn
	suite.commandClient = commandClient
	suite.queryClient = queryClient
	suite.accountIdOf = accountIdOf
	suite.privateKeyOf = privateKeyOf
	suite.chainA = ibctesting.NewChain(suite.T(), *chainClient, contractConfig, accountIds, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), *chainClient, contractConfig, accountIds, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
}

func (suite *IrohaIcs20TestSuite) TestChannel() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB

	clientA, clientB := suite.coordinator.SetupClients(ctx, chainA, chainB, clienttypes.MockClient)
	connA, connB := suite.coordinator.CreateConnection(ctx, chainA, chainB, clientA, clientB)
	chanA, chanB := suite.coordinator.CreateChannel(ctx, chainA, chainB, connA, connB, ibctesting.IrohaTransferPort, ibctesting.IrohaTransferPort, channeltypes.UNORDERED)

	/*
	 * # Scenario
	 * 1. deployer sets bank account
	 * 2. bankA mints tokens to alice
	 * 3. alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens
	 * 4. bankA burns the tokens
	 * 5. relayer executes recvPacket, while requesting for bankB to mint the tokens to bob
	 * 6. bankB mints the tokens to bob
	 * 7. bob executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens
	 * 8. bankB burns the tokens
	 * 9. relayer executes recvPacket, while requesting for bankA to mint the tokens to alice
	 * 10. bankA mints the token to alice
	 *
	 * # Roles
	 * deployer = "relayer@test"
	 * alice = "alice@test"
	 * bob = "bob@test"
	 * bank = "relayer@test" (= deployer)
	 * relayer = "relayer@test" (= deployer)
	 * token = "coin@test"
	 */

	var err error
	var (
		aliceBalance0,
		aliceBalance,
		bobBalance0,
		bobBalance,
		bankABalance0,
		bankABalance,
		bankBBalance0,
		bankBBalance int
	)

	// deployer sets bank account
	suite.T().Log("deployer sets bank account")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, deployer), "setBank", suite.accountIdOf[bankA]),
	))

	// check alice's initial balance
	aliceBalance0, err = suite.balanceOf(ctx, alice, AssetId)
	suite.Require().NoError(err)
	suite.T().Log("alice's initial balance:", aliceBalance0)

	// check bob's initial balance
	bobBalance0, err = suite.balanceOf(ctx, bob, AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bob's initial balance:", bobBalance0)

	// check bankA's initial balance
	bankABalance0, err = suite.balanceOf(ctx, bankA, AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bankA's initial balance:", bankABalance0)

	// check bankB's initial balance
	bankBBalance0, err = suite.balanceOf(ctx, bankB, AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bankB's initial balance:", bankBBalance0)

	// bankA mints tokens to alice
	suite.T().Log("bankA mints tokens to alice")
	suite.Require().NoError(suite.addAssetQuantity(ctx, bankA, AssetId, "1000"))
	suite.Require().NoError(suite.transferAsset(ctx, bankA, bankA, alice, AssetId, "initial mint", "1000"))

	// check alice's balance after initial mint
	aliceBalance, err = suite.balanceOf(ctx, alice, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceBalance, aliceBalance0+1000)
	suite.T().Log("alice's balance after initial mint:", aliceBalance)

	// alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens
	suite.T().Log("alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Transfer.Transact(chainA.TxOpts(ctx, alice), "sendTransfer",
			suite.accountIdOf[alice],
			suite.accountIdOf[bob],
			AssetId,
			"hoge",
			"1000",
			chanA.PortID,
			chanA.ID,
			uint64(chainA.LastHeader().Number.Int64())+1000,
		),
	))
	var burnRequestId *big.Int
	{
		events, err := chainA.FindBurnRequestedEvents(ctx)
		suite.Require().NoError(err)
		suite.Require().Len(events, 1)
		burnRequestId = events[0].Id
	}
	suite.T().Log("burn request id:", burnRequestId)

	// check alice's balance after sendTransfer
	aliceBalance, err = suite.balanceOf(ctx, alice, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceBalance, aliceBalance0)
	suite.T().Log("alice's balance after sendTransfer:", aliceBalance)

	// check bankA's balance after sendTransfer
	bankABalance, err = suite.balanceOf(ctx, bankA, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankABalance, bankABalance0+1000)
	suite.T().Log("bankA's balance after sendTransfer:", bankABalance)

	// bankA burns the tokens
	suite.T().Log("bankA burns the tokens")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bankA), "burn", burnRequestId),
	))

	// check bankA's balance after burn
	bankABalance, err = suite.balanceOf(ctx, bankA, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankABalance, bankABalance0)
	suite.T().Log("bankA's balance after burn:", bankABalance)

	// relayer executes recvPacket, while requesting for bankB to mint the tokens to bob
	suite.T().Log("relayer executes recvPacket, while requesting for bankB to mint the tokens to bob")
	packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)

	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB))
	suite.Require().NoError(chainB.HandlePacketRecv(ctx, chainA, chanB, chanA, *packet))

	var mintRequestId *big.Int
	{
		events, err := chainB.FindMintRequestedEvents(ctx)
		suite.Require().NoError(err)
		suite.Require().Len(events, 1)
		mintRequestId = events[0].Id
	}
	suite.T().Log("mint request id:", mintRequestId)

	chainB.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *packet, []byte{1}))

	// bankB mints the tokens to bob
	suite.T().Log("bankB mints the tokens to bob")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bankB), "mint", mintRequestId),
	))

	// check bob's balance after mint
	bobBalance, err = suite.balanceOf(ctx, bob, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bobBalance, bobBalance0+1000)
	suite.T().Log("bob's balance after mint:", bobBalance)

	// bob executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens
	suite.T().Log("bob executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Transfer.Transact(chainB.TxOpts(ctx, bob), "sendTransfer",
			suite.accountIdOf[bob],
			suite.accountIdOf[alice],
			AssetId,
			"hoge",
			"1000",
			chanB.PortID,
			chanB.ID,
			uint64(chainB.LastHeader().Number.Int64())+1000,
		),
	))
	{
		events, err := chainB.FindBurnRequestedEvents(ctx)
		suite.Require().NoError(err)
		suite.Require().Len(events, 1)
		burnRequestId = events[0].Id
	}
	suite.T().Log("burn request id:", burnRequestId)

	// check bob's balance after sendTransfer
	bobBalance, err = suite.balanceOf(ctx, bob, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bobBalance, bobBalance0)
	suite.T().Log("bob's balance after sendTransfer:", bobBalance)

	// check bankB's balance after sendTransfer
	bankBBalance, err = suite.balanceOf(ctx, bankB, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankBBalance, bankBBalance0+1000)
	suite.T().Log("bankB's balance after sendTransfer:", bankBBalance)

	// bankB burns the tokens
	suite.T().Log("bankB burns the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bankB), "burn", burnRequestId),
	))

	// check bankB's balance after burn
	bankBBalance, err = suite.balanceOf(ctx, bankB, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankBBalance, bankBBalance0)
	suite.T().Log("bankB's balance after burn:", bankBBalance)

	// relayer executes recvPacket, while requesting for bankA to mint the tokens to alice
	suite.T().Log("relayer executes recvPacket, while requesting for bankA to mint the tokens to alice")
	packet, err = chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
	suite.Require().NoError(err)

	chainB.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))
	suite.Require().NoError(chainA.HandlePacketRecv(ctx, chainB, chanA, chanB, *packet))

	{
		events, err := chainA.FindMintRequestedEvents(ctx)
		suite.Require().NoError(err)
		suite.Require().Len(events, 1)
		mintRequestId = events[0].Id
	}
	suite.T().Log("mint request id:", mintRequestId)

	chainA.UpdateHeader()
	suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *packet, []byte{1}))

	// bankA mints the token to alice
	suite.T().Log("bankA mints the tokens to alice")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bankA), "mint", mintRequestId),
	))

	// check alice's balance after recvPacket
	aliceBalance, err = suite.balanceOf(ctx, alice, AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceBalance, aliceBalance0+1000)
	suite.T().Log("alice's balance after recvPacket:", aliceBalance)

	// close channel
	suite.coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
	// confirm that the channel is CLOSED on chain A
	chanData, ok, err := chainA.IBCHost.GetChannel(chainA.CallOpts(ctx, relayer), chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	suite.Require().True(ok)
	suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)
	// confirm that the channel is CLOSED on chain B
	chanData, ok, err = chainB.IBCHost.GetChannel(chainB.CallOpts(ctx, relayer), chanB.PortID, chanB.ID)
	suite.Require().NoError(err)
	suite.Require().True(ok)
	suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)

	/*
		// deposit a simple token to the bank
		suite.T().Log("deposit a simple token to the bank")
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(chainA.ICS20Bank.Transact(
			chainA.TxOpts(ctx, deployer),
			"deposit",
			chainA.ContractConfig.GetSimpleTokenAddress(),
			big.NewInt(100),
			chainA.CallOpts(ctx, alice).From,
		)))

		//// ensure that the balance is reduced
		balance1, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
		suite.Require().NoError(err)
		suite.Require().Equal(balance0.Int64()-100, balance1.Int64())
		suite.T().Log("SimpleToken:balance:deployer", balance1)

		baseDenom := strings.ToLower(chainA.ContractConfig.GetSimpleTokenAddress().String())

		bankA, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, alice).From, baseDenom)
		suite.Require().NoError(err)
		suite.Require().GreaterOrEqual(bankA.Int64(), int64(100))
		suite.T().Log("ICS20Bank:balance:alice", bankA)

		// try to transfer the token to chainB
		suite.T().Log("transfer the token to chainB:amount", 100)
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(
			chainA.ICS20Transfer.Transact(
				chainA.TxOpts(ctx, alice),
				"sendTransfer",
				baseDenom,
				uint64(100),
				chainB.CallOpts(ctx, bob).From,
				chanA.PortID, chanA.ID,
				uint64(chainA.LastHeader().Number.Int64())+1000,
			),
		))
		chainA.UpdateHeader()
		suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainB, chainA, clientB))

		// ensure that escrow has correct balance
		escrowBalance, err := chainA.ICS20Bank.BalanceOf(chainA.CallOpts(ctx, alice), chainA.ContractConfig.GetICS20TransferBankAddress(), baseDenom)
		suite.Require().NoError(err)
		suite.Require().GreaterOrEqual(escrowBalance.Int64(), int64(100))
		suite.T().Log("ICS20Bank:balance:alice", escrowBalance)

		// relay the packet
		transferPacket, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, *transferPacket))
		suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *transferPacket, []byte{1}))

		// ensure that chainB has correct balance
		expectedDenom := fmt.Sprintf("%v/%v/%v", chanB.PortID, chanB.ID, baseDenom)
		balance, err := chainB.ICS20Bank.BalanceOf(chainB.CallOpts(ctx, relayer), chainB.CallOpts(ctx, bob).From, expectedDenom)
		suite.Require().NoError(err)
		suite.Require().Equal(int64(100), balance.Int64())
		suite.T().Log("ICS20Bank:balance:bob", balance)

		// try to transfer the token to chainA
		suite.T().Log("transfer the token to chainB:amount", 100)
		suite.Require().NoError(chainB.WaitIfNoError(ctx)(
			chainB.ICS20Transfer.Transact(
				chainB.TxOpts(ctx, bob),
				"sendTransfer",
				expectedDenom,
				uint64(100),
				chainA.CallOpts(ctx, alice).From,
				chanB.PortID,
				chanB.ID,
				uint64(chainB.LastHeader().Number.Int64())+1000,
			),
		))
		chainB.UpdateHeader()
		suite.Require().NoError(suite.coordinator.UpdateClient(ctx, chainA, chainB, clientA))

		// relay the packet
		transferPacket, err = chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
		suite.Require().NoError(err)
		suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainA, chainB, chanA, chanB, *transferPacket))
		suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *transferPacket, []byte{1}))

		// withdraw tokens from the bank
		suite.T().Log("withdraw tokens from the bank")
		suite.Require().NoError(chainA.WaitIfNoError(ctx)(
			chainA.ICS20Bank.Transact(
				chainA.TxOpts(ctx, alice),
				"withdraw",
				chainA.ContractConfig.GetSimpleTokenAddress(),
				big.NewInt(100),
				chainA.CallOpts(ctx, deployer).From,
			)))

		// ensure that token balance equals original value
		balanceA2, err := chainA.SimpleToken.BalanceOf(chainA.CallOpts(ctx, relayer), chainA.CallOpts(ctx, deployer).From)
		suite.Require().NoError(err)
		suite.Require().Equal(balance0.Int64(), balanceA2.Int64())
		suite.T().Log("ICS20Bank:balance:deployer", balanceA2)

		// close channel
		suite.coordinator.CloseChannel(ctx, chainA, chainB, chanA, chanB)
		// confirm that the channel is CLOSED on chain A
		chanData, ok, err := chainA.IBCHost.GetChannel(chainA.CallOpts(ctx, relayer), chanA.PortID, chanA.ID)
		suite.Require().NoError(err)
		suite.Require().True(ok)
		suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)
		// confirm that the channel is CLOSED on chain B
		chanData, ok, err = chainB.IBCHost.GetChannel(chainB.CallOpts(ctx, relayer), chanB.PortID, chanB.ID)
		suite.Require().NoError(err)
		suite.Require().True(ok)
		suite.Require().Equal(channeltypes.Channel_State(chanData.State), channeltypes.CLOSED)
	*/
}

func (suite *IrohaIcs20TestSuite) signTx(tx *protocol.Transaction, signers ...uint32) error {
	var keys []string
	for _, signer := range signers {
		keys = append(keys, suite.privateKeyOf[signer])
	}
	if sigs, err := crypto.SignTransaction(tx, keys...); err != nil {
		return err
	} else {
		tx.Signatures = sigs
		return nil
	}
}

func (suite *IrohaIcs20TestSuite) sendTx(ctx context.Context, tx *protocol.Transaction) error {
	if txHash, err := suite.commandClient.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("SendTransaction failed: %v", err)
	} else if _, err := suite.commandClient.TxStatusStream(ctx, txHash); err != nil {
		return fmt.Errorf("TxStatusStream failed: %v", err)
	}
	return nil
}

func (suite *IrohaIcs20TestSuite) signQuery(q *protocol.Query, signer uint32) error {
	if sig, err := crypto.SignQuery(q, suite.privateKeyOf[signer]); err != nil {
		return err
	} else {
		q.Signature = sig
		return nil
	}
}

func (suite *IrohaIcs20TestSuite) sendQuery(ctx context.Context, q *protocol.Query) (*protocol.QueryResponse, error) {
	return suite.queryClient.SendQuery(ctx, q)
}

func (suite *IrohaIcs20TestSuite) addAssetQuantity(ctx context.Context, signer uint32, assetId, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.AddAssetQuantity(assetId, amount),
			},
			command.CreatorAccountId(suite.accountIdOf[signer]),
		),
	)
	if err := suite.signTx(tx, signer); err != nil {
		return err
	} else if err := suite.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (suite *IrohaIcs20TestSuite) subtractAssetQuantity(ctx context.Context, signer uint32, assetId, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.SubtractAssetQuantity(assetId, amount),
			},
			command.CreatorAccountId(suite.accountIdOf[signer]),
		),
	)
	if err := suite.signTx(tx, signer); err != nil {
		return err
	} else if err := suite.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (suite *IrohaIcs20TestSuite) transferAsset(ctx context.Context, signer uint32, src, dest uint32, assetID, description, amount string) error {
	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*protocol.Command{
				command.TransferAsset(suite.accountIdOf[src], suite.accountIdOf[dest], assetID, description, amount),
			},
			command.CreatorAccountId(suite.accountIdOf[signer]),
		),
	)
	if err := suite.signTx(tx, signer); err != nil {
		return err
	} else if err := suite.sendTx(ctx, tx); err != nil {
		return err
	}
	return nil
}

func (suite *IrohaIcs20TestSuite) balanceOf(ctx context.Context, account uint32, assetId string) (int, error) {
	accountId := suite.accountIdOf[account]
	q := query.GetAccountAsset(accountId, nil, query.CreatorAccountId(accountId))
	if err := suite.signQuery(q, account); err != nil {
		return 0, err
	} else if res, err := suite.sendQuery(ctx, q); err != nil {
		return 0, err
	} else {
		assets := res.GetAccountAssetsResponse().AccountAssets
		for _, a := range assets {
			if a.AssetId == assetId {
				return strconv.Atoi(a.Balance)
			}
		}
		return 0, nil
	}
}

func TestIrohaIcs20TestSuite(t *testing.T) {
	suite.Run(t, new(IrohaIcs20TestSuite))
}
