package e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	channeltypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/channel"
	clienttypes "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/client"
	ibctesting "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/testing"
)

const (
	ToriiAddressA = "localhost:50051"
	ToriiAddressB = "localhost:51051"

	GatewayAddressA = "http://127.0.0.1:8545"
	GatewayAddressB = "http://127.0.0.1:8645"

	AssetId = "coin#test"

	AdminAPrivateKey   = "f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70"
	AdminBPrivateKey   = "8d6a25122e3a61e0d76e5c6d2b590f97b254f9b1eaff0b450653e8f04f4d5911"
	AlicePrivateKey    = "a6924c9781c46df18c05545735f127eaf788a60816a7aeb9d5e928460b51cb2f"
	BobPrivateKey      = "f66c1f19a52bf2955d00bf050793a80056ccfa6237b46f4d7d3a9e20af669c29"
	CarolPrivateKey    = "2ec6a7b95aadbafadb7ee21f17e65dc8a3e223853af5c3d1974a7f32b6720295"
	DavePrivateKey     = "3a2e40aa0e008409282a58258dbca48857cb57bb170ff6bd5ef5de38a8f9ab0f"
	RelayerAPrivateKey = "e517af47112e4f501afb26e4f34eadc8b0ad8eadaf4962169fc04bc8ddbfe091"
	RelayerBPrivateKey = "a944b0d0a4548b69edf2fef8b14d4575decf40272b084dfad204e438192c0dd0"

	AdminAccountID   = "admin@test"
	AliceAccountID   = "alice@test"
	BobAccountID     = "bob@test"
	CarolAccountID   = "carol@test"
	DaveAccountID    = "dave@test"
	RelayerAccountID = "relayer@test"

	relayer  = ibctesting.RelayerKeyIndex // the key-index of relayer
	deployer = ibctesting.RelayerKeyIndex // the key-index of contract
	alice    = 1                          // the key-index of alice
	carol    = 1                          // the key-index of carol
	bob      = 2                          // the key-index of bob
	dave     = 2                          // the key-index of dave
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
	chainClientA, err := client.NewETHClient(GatewayAddressA, clienttypes.MockClient)
	suite.Require().NoError(err)
	chainClientB, err := client.NewETHClient(GatewayAddressB, clienttypes.MockClient)
	suite.Require().NoError(err)

	accountIdListA := []string{RelayerAccountID, AliceAccountID, BobAccountID, AdminAccountID}
	accountIdListB := []string{RelayerAccountID, CarolAccountID, DaveAccountID, AdminAccountID}
	accountIdsA := map[uint32]string{
		relayer: RelayerAccountID,
		alice:   AliceAccountID,
		bob:     BobAccountID,
		bank:    AdminAccountID,
	}
	accountIdsB := map[uint32]string{
		relayer: RelayerAccountID,
		carol:   CarolAccountID,
		dave:    DaveAccountID,
		bank:    AdminAccountID,
	}
	keysA := map[uint32]string{
		relayer: RelayerAPrivateKey,
		alice:   AlicePrivateKey,
		bob:     BobPrivateKey,
		bank:    AdminAPrivateKey,
	}
	keysB := map[uint32]string{
		relayer: RelayerBPrivateKey,
		carol:   CarolPrivateKey,
		dave:    DavePrivateKey,
		bank:    AdminBPrivateKey,
	}
	contractConfigA := ibctesting.NewTruffleContractConfig(1000, "../../build/contracts")
	contractConfigB := ibctesting.NewTruffleContractConfig(2000, "../../build/contracts")

	suite.chainA = ibctesting.NewChain(suite.T(), *chainClientA, contractConfigA, accountIdListA, uint64(time.Now().UnixNano()))
	suite.chainB = ibctesting.NewChain(suite.T(), *chainClientB, contractConfigB, accountIdListB, uint64(time.Now().UnixNano()))
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), suite.chainA, suite.chainB)
	suite.irohadA = client.NewIrohadClient(ToriiAddressA, accountIdsA, keysA)
	suite.irohadB = client.NewIrohadClient(ToriiAddressB, accountIdsB, keysB)
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
	 * 1. deployer sets bank account for chainA
	 * 2. deployer sets bank account for chainB
	 * 3. bankA mints tokens to alice
	 * 4. alice executes ICS-20 sendTransfer, while requesting for bankA to burn the tokens
	 * 5. bankA burns the tokens
	 * 6. relayer executes recvPacket, while requesting for bankB to mint the tokens to carol
	 * 7. bankB mints the tokens to carol
	 * 8. carol executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens
	 * 9. bankB burns the tokens
	 * 10. relayer executes recvPacket, while requesting for bankA to mint the tokens to alice
	 * 11. bankA mints the token to alice
	 */

	var err error
	var (
		aliceBalance0,
		aliceBalance,
		carolBalance0,
		carolBalance,
		bankABalance0,
		bankABalance,
		bankBBalance0,
		bankBBalance int
	)

	// deployer sets bank account for chainA
	suite.T().Log("deployer sets bank account for chainA")
	suite.Require().NoError(chainA.WaitIfNoError(ctx)(
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, deployer), "setBank", irohadA.AccountIdOf(bank)),
	))

	// deployer sets bank account for chainB
	suite.T().Log("deployer sets bank account for chainB")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, deployer), "setBank", irohadB.AccountIdOf(bank)),
	))

	// check alice's initial balance
	aliceBalance0, err = irohadA.GetAccountAsset(ctx, alice, irohadA.AccountIdOf(alice), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("alice's initial balance:", aliceBalance0)

	// check carol's initial balance
	carolBalance0, err = irohadB.GetAccountAsset(ctx, carol, irohadB.AccountIdOf(carol), AssetId)
	suite.Require().NoError(err)
	suite.T().Log("carol's initial balance:", carolBalance0)

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
			irohadB.AccountIdOf(carol),
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
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bank), "burn", burnRequestId),
	))

	// check bankA's balance after burn
	bankABalance, err = irohadA.GetAccountAsset(ctx, bank, irohadA.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankABalance, bankABalance0)
	suite.T().Log("bankA's balance after burn:", bankABalance)

	// relayer executes recvPacket, while requesting for bankB to mint the tokens to carol
	suite.T().Log("relayer executes recvPacket, while requesting for bankB to mint the tokens to carol")
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

	// bankB mints the tokens to carol
	suite.T().Log("bankB mints the tokens to carol")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bank), "mint", mintRequestId),
	))

	// check carol's balance after mint
	carolBalance, err = irohadB.GetAccountAsset(ctx, carol, irohadB.AccountIdOf(carol), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(carolBalance, carolBalance0+1000)
	suite.T().Log("carol's balance after mint:", carolBalance)

	// carol executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens
	suite.T().Log("carol executes ICS-20 sendTransfer, while requesting for bankB to burn the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Transfer.Transact(chainB.TxOpts(ctx, carol), "sendTransfer",
			irohadB.AccountIdOf(carol),
			irohadA.AccountIdOf(alice),
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

	// check carol's balance after sendTransfer
	carolBalance, err = irohadB.GetAccountAsset(ctx, carol, irohadB.AccountIdOf(carol), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(carolBalance, carolBalance0)
	suite.T().Log("carol's balance after sendTransfer:", carolBalance)

	// check bankB's balance after sendTransfer
	bankBBalance, err = irohadB.GetAccountAsset(ctx, bank, irohadB.AccountIdOf(bank), AssetId)
	suite.Require().NoError(err)
	suite.Require().Equal(bankBBalance, bankBBalance0+1000)
	suite.T().Log("bankB's balance after sendTransfer:", bankBBalance)

	// bankB burns the tokens
	suite.T().Log("bankB burns the tokens")
	suite.Require().NoError(chainB.WaitIfNoError(ctx)(
		chainB.IrohaICS20Bank.Transact(chainB.TxOpts(ctx, bank), "burn", burnRequestId),
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
		chainA.IrohaICS20Bank.Transact(chainA.TxOpts(ctx, bank), "mint", mintRequestId),
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
