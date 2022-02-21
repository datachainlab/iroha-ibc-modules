package integration

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ethmultisigtypes "github.com/datachainlab/ibc-ethmultisig-client/modules/light-clients/xx-ethmultisig/types"
	"github.com/datachainlab/ibc-ethmultisig-client/modules/relay/ethmultisig"
	"github.com/datachainlab/ibc-ethmultisig-client/pkg/contract/multisigclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/contract/ibchost"
	ibcsolidityclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
	"github.com/hyperledger-labs/yui-ibc-solidity/pkg/wallet"
	"github.com/stretchr/testify/suite"

	"github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/client"
	ibctesting "github.com/datachainlab/iroha-ibc-modules/onchain-module/pkg/testing"
)

const testMnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ETHMultisigTestSuite struct {
	suite.Suite

	chain *ibctesting.Chain
	cdc   codec.ProtoCodecMarshaler

	sigKeys []*ecdsa.PrivateKey
}

func (suite *ETHMultisigTestSuite) SetupTest() {
	chainClient, err := client.NewETHClient("http://127.0.0.1:8545", ethmultisigtypes.ClientType)
	suite.Require().NoError(err)

	accountIds := []string{"relayer@test"}
	contractConfig := ibctesting.NewTruffleContractConfig(math.MaxInt32, "../../build/contracts")

	suite.chain = ibctesting.NewChain(suite.T(), math.MaxInt32, *chainClient, contractConfig, accountIds, uint64(time.Now().UnixNano()))
	registry := codectypes.NewInterfaceRegistry()
	ethmultisigtypes.RegisterInterfaces(registry)
	suite.cdc = codec.NewProtoCodec(registry)

	suite.createSigKeys(1)
}

func (suite *ETHMultisigTestSuite) createSigKeys(length int) {
	for i := 0; i < length; i++ {
		key, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(testMnemonicPhrase, fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		suite.Require().NoError(err)
		suite.sigKeys = append(suite.sigKeys, key)
	}
}

func (suite *ETHMultisigTestSuite) TestMultisig() {
	ctx := context.TODO()

	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, suite.sigKeys, prefix)

	consensusState := makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		uint64(time.Now().UnixNano()),
	)
	anyConsensusStateBytes, err := suite.cdc.MarshalInterface(consensusState)
	suite.Require().NoError(err)
	err = suite.chain.WaitIfNoError(ctx)(
		suite.chain.IBCHost.Transact(
			suite.chain.TxOpts(ctx, 0),
			"setConsensusState",
			clientID,
			ibchost.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			anyConsensusStateBytes,
		))
	suite.Require().NoError(err)

	{
		timestamp, ok, err := suite.chain.MultisigClient.GetTimestampAtHeight(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{RevisionNumber: 0, RevisionHeight: 1},
		)
		suite.T().Log(timestamp)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	{
		targetClientState := makeMultisigClientState(1)
		proofClient, _, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
		suite.Require().NoError(err)
		anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
		addr := crypto.PubkeyToAddress(suite.sigKeys[0].PublicKey)
		ok, err := suite.chain.MultisigClient.VerifySignature(
			suite.chain.CallOpts(ctx, 0),
			multisigclient.ConsensusStateData{
				Addresses:   [][]byte{addr.Bytes()},
				Diversifier: diversifier,
				Timestamp:   proofClient.Timestamp,
			},
			multisigclient.MultiSignatureData{
				Signatures: proofClient.Signatures,
				Timestamp:  proofClient.Timestamp,
			},
			anyClientStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientState
	{
		targetClientState := makeMultisigClientState(1)
		proofClient, _, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
		suite.Require().NoError(err)
		anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofClient)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyClientState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientConsensusState consensusHeight.revisionNumber=0
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(0, 100)
		proofConsensus, _, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			counterpartyClientID,
			multisigclient.HeightData{
				RevisionNumber: consensusHeight.RevisionNumber,
				RevisionHeight: consensusHeight.RevisionHeight,
			},
			prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientConsensusState consensusHeight.revisionNumber=1
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(1, 100)
		proofConsensus, _, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			counterpartyClientID,
			multisigclient.HeightData{
				RevisionNumber: consensusHeight.RevisionNumber,
				RevisionHeight: consensusHeight.RevisionHeight,
			},
			prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyConnectionState
	{
		const connectionID = "connection-0"
		targetConnection := conntypes.NewConnectionEnd(conntypes.INIT, counterpartyClientID, conntypes.NewCounterparty(clientID, connectionID, types.NewMerklePrefix([]byte("ibc"))), []*conntypes.Version{}, 0)
		proof, _, err := prover.SignConnectionState(proofHeight, connectionID, targetConnection)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		connectionBytes, err := suite.cdc.Marshal(&targetConnection)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyConnectionState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			prefix, proofBytes, connectionID, connectionBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	const portID, channelID, cpPortID, cpChannelID = "port-0", "channel-0", "port-1", "channel-1"

	// VerifyChannelstate
	{
		targetChannel := chantypes.NewChannel(chantypes.INIT, chantypes.UNORDERED, chantypes.NewCounterparty(cpPortID, cpChannelID), []string{"connection-0"}, "1")
		proof, _, err := prover.SignChannelState(proofHeight, portID, channelID, targetChannel)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		channelBytes, err := suite.cdc.Marshal(&targetChannel)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyChannelState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			prefix, proofBytes, portID, channelID, channelBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketCommitment
	{
		commitment := sha256.Sum256([]byte("test"))
		proof, _, err := prover.SignPacketState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyPacketCommitment(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			0, 0,
			prefix, proofBytes, portID, channelID, 1, commitment,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketAcknowledgement
	{
		acknowledgement := []byte("ack")
		commitment := sha256.Sum256(acknowledgement)
		proof, _, err := prover.SignPacketAcknowledgementState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.MultisigClient.VerifyPacketAcknowledgement(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID,
			multisigclient.HeightData{
				RevisionNumber: 0,
				RevisionHeight: 1,
			},
			0, 0,
			prefix, proofBytes, portID, channelID, 1, acknowledgement,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}
}

func (suite *ETHMultisigTestSuite) TestMultisigSign() {
	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, suite.sigKeys, prefix)

	targetClientState := makeMultisigClientState(1)
	proofClient, signBytes, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
	suite.Require().NoError(err)

	err = ethmultisigtypes.VerifySignature(prover.Addresses(), proofClient, signBytes)
	suite.Require().NoError(err)
}

func makeMultisigClientState(latestHeight uint64) *ethmultisigtypes.ClientState {
	return &ethmultisigtypes.ClientState{
		LatestHeight: ibcsolidityclient.Height{
			RevisionNumber: 0,
			RevisionHeight: latestHeight,
		},
	}
}

func makeMultisigConsensusState(addresses []common.Address, diversifier string, timestamp uint64) *ethmultisigtypes.ConsensusState {
	var addrs [][]byte
	for _, addr := range addresses {
		addrs = append(addrs, addr[:])
	}
	return &ethmultisigtypes.ConsensusState{
		Addresses:   addrs,
		Diversifier: diversifier,
		Timestamp:   timestamp,
	}
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ETHMultisigTestSuite))
}
