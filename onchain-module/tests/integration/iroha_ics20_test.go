package integration

import (
	"context"
	"testing"
	"time"

	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/channel"
	clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/client"
	ibctesting "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/testing"
)

const (
	ToriiAddress = "localhost:50051"

	AssetId = "coin#test"

	RelayerPrivateKey = "e517af47112e4f501afb26e4f34eadc8b0ad8eadaf4962169fc04bc8ddbfe091"
	AlicePrivateKey   = "a6924c9781c46df18c05545735f127eaf788a60816a7aeb9d5e928460b51cb2f"
	BobPrivateKey     = "f66c1f19a52bf2955d00bf050793a80056ccfa6237b46f4d7d3a9e20af669c29"
	AdminPrivateKey   = "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70"

	RelayerAccountID = "relayer@test"
	AliceAccountID   = "alice@test"
	BobAccountID     = "bob@test"
	AdminAccountID   = "admin@test"

	relayer  = ibctesting.RelayerKeyIndex // the key-index of relayer
	deployer = ibctesting.RelayerKeyIndex // the key-index of contract
	alice    = 1                          // the key-index of alice
	bob      = 2                          // the key-index of bob
	bank     = 3                          // the key-index of bank (= admin)
)

type IrohaIcs20TestSuite struct {
	suite.Suite

	coordinator ibctesting.Coordinator
	chainA      *ibctesting.Chain
	chainB      *ibctesting.Chain
	irohadA     *client.IrohadClient
	irohadB     *client.IrohadClient
}

func (suite *IrohaIcs20TestSuite) SetupTest() {
	chainClient, err := client.NewETHClient("http://127.0.0.1:8545", clienttypes.MockClient)
	suite.Require().NoError(err)

	networkID := 1000
	accountIds := []string{RelayerAccountID, AliceAccountID, BobAccountID, AdminAccountID}
	accountIdsA := map[uint32]string{
		relayer: RelayerAccountID,
		alice:   AliceAccountID,
		bank:    AdminAccountID,
	}
	accountIdsB := map[uint32]string{
		relayer: RelayerAccountID,
		bob:     BobAccountID,
		bank:    AdminAccountID,
	}
	keysA := map[uint32]string{
		relayer: RelayerPrivateKey,
		alice:   AlicePrivateKey,
		bank:    AdminPrivateKey,
	}
	keysB := map[uint32]string{
		relayer: RelayerPrivateKey,
		bob:     BobPrivateKey,
		bank:    AdminPrivateKey,
	}
	contractConfig := ibctesting.NewTruffleContractConfig(networkID, "../../build/contracts")

	suite.chainA = ibctesting.NewChain(suite.T(), *chainClient, contractConfig, accountIds, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), *chainClient, contractConfig, accountIds, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
	suite.irohadA = client.NewIrohadClient(ToriiAddress, accountIdsA, keysA)
	suite.irohadB = client.NewIrohadClient(ToriiAddress, accountIdsB, keysB)
}

func (suite *IrohaIcs20TestSuite) TestChannel() {
	ctx := context.Background()

	chainA := suite.chainA
	chainB := suite.chainB
	irohadA := suite.irohadA
	irohadB := suite.irohadB

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
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, deployer), "setBank", irohadA.AccountIdOf(bank)),
	))

	// check alice's initial balance
	aliceBalance0, err = irohadA.GetAccountAsset(ctx, alice, irohadA.AccountIdOf(alice), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("alice's initial balance:", aliceBalance0)

	// check bob's initial balance
	bobBalance0, err = irohadB.GetAccountAsset(ctx, bob, irohadB.AccountIdOf(bob), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bob's initial balance:", bobBalance0)

	// check bankA's initial balance
	bankABalance0, err = irohadA.GetAccountAsset(ctx, bank, irohadA.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bankA's initial balance:", bankABalance0)

	// check bankB's initial balance
	bankBBalance0, err = irohadB.GetAccountAsset(ctx, bank, irohadB.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("bankB's initial balance:", bankBBalance0)

	// bankA mints tokens to alice
	suite.T().Log("bankA mints tokens to alice")
	suite.Require().NoError(irohadA.AddAssetQuantity(ctx, bank, AssetId, "1000"))
	suite.Require().NoError(irohadA.TransferAsset(ctx, bank, irohadA.AccountIdOf(bank), irohadA.AccountIdOf(alice), AssetId, "initial mint", "1000"))

	// check alice's balance after initial mint
	aliceBalance, err = irohadA.GetAccountAsset(ctx, alice, irohadA.AccountIdOf(alice), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceBalance, aliceBalance0+1000)
	suite.T().Log("alice's balance after initial mint:", aliceBalance)

	// alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens
	suite.T().Log("alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Transfer.Transact(chainA.TxOpts(ctx, alice), "sendTransfer",
			irohadA.AccountIdOf(alice),
			irohadB.AccountIdOf(bob),
			AssetId,
			"hoge",
			"1000",
			chanA.PortID,
			chanA.ID,
			uint64(chainA.LastHeader().Number.Int64())+1000,
		),
	))

	// check alice's balance after sendTransfer
	aliceBalance, err = irohadA.GetAccountAsset(ctx, alice, irohadA.AccountIdOf(alice), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(aliceBalance, aliceBalance0)
	suite.T().Log("alice's balance after sendTransfer:", aliceBalance)

	// check bankA's balance after sendTransfer
	bankABalance, err = irohadA.GetAccountAsset(ctx, bank, irohadA.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankABalance, bankABalance0+1000)
	suite.T().Log("bankA's balance after sendTransfer:", bankABalance)

	// bankA burns the tokens
	suite.T().Log("bankA burns the tokens")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bank), "burn"),
	))

	// check bankA's balance after burn
	bankABalance, err = irohadA.GetAccountAsset(ctx, bank, irohadA.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankABalance, bankABalance0)
	suite.T().Log("bankA's balance after burn:", bankABalance)

	// relayer executes recvPacket, while requesting for bankB to mint the tokens to bob
	suite.T().Log("relayer executes recvPacket, while requesting for bankB to mint the tokens to bob")
	packet, err := chainA.GetLastSentPacket(ctx, chanA.PortID, chanA.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainB, chainA, chanB, chanA, *packet))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainA, chainB, chanA, chanB, *packet, []byte{1}))

	// bankB mints the tokens to bob
	suite.T().Log("bankB mints the tokens to bob")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bank), "mint"),
	))

	// check bob's balance after mint
	bobBalance, err = irohadB.GetAccountAsset(ctx, bob, irohadB.AccountIdOf(bob), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bobBalance, bobBalance0+1000)
	suite.T().Log("bob's balance after mint:", bobBalance)

	// bob executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens
	suite.T().Log("bob executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Transfer.Transact(chainB.TxOpts(ctx, bob), "sendTransfer",
			irohadB.AccountIdOf(bob),
			irohadA.AccountIdOf(alice),
			AssetId,
			"hoge",
			"1000",
			chanB.PortID,
			chanB.ID,
			uint64(chainB.LastHeader().Number.Int64())+1000,
		),
	))

	// check bob's balance after sendTransfer
	bobBalance, err = irohadB.GetAccountAsset(ctx, bob, irohadB.AccountIdOf(bob), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bobBalance, bobBalance0)
	suite.T().Log("bob's balance after sendTransfer:", bobBalance)

	// check bankB's balance after sendTransfer
	bankBBalance, err = irohadB.GetAccountAsset(ctx, bank, irohadB.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankBBalance, bankBBalance0+1000)
	suite.T().Log("bankB's balance after sendTransfer:", bankBBalance)

	// bankB burns the tokens
	suite.T().Log("bankB burns the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bank), "burn"),
	))

	// check bankB's balance after burn
	bankBBalance, err = irohadB.GetAccountAsset(ctx, bank, irohadB.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankBBalance, bankBBalance0)
	suite.T().Log("bankB's balance after burn:", bankBBalance)

	// relayer executes recvPacket, while requesting for bankA to mint the tokens to alice
	suite.T().Log("relayer executes recvPacket, while requesting for bankA to mint the tokens to alice")
	packet, err = chainB.GetLastSentPacket(ctx, chanB.PortID, chanB.ID)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.coordinator.HandlePacketRecv(ctx, chainA, chainB, chanA, chanB, *packet))
	suite.Require().NoError(suite.coordinator.HandlePacketAcknowledgement(ctx, chainB, chainA, chanB, chanA, *packet, []byte{1}))

	// bankA mints the token to alice
	suite.T().Log("bankA mints the tokens to alice")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bank), "mint"),
	))

	// check alice's balance after recvPacket
	aliceBalance, err = irohadA.GetAccountAsset(ctx, alice, irohadA.AccountIdOf(alice), AssetId)
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
}

func TestIrohaIcs20TestSuite(t *testing.T) {
	suite.Run(t, new(IrohaIcs20TestSuite))
}
