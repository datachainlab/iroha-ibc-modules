package rpc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gogo/protobuf/proto"
	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/crypto"
	x "github.com/hyperledger/burrow/encoding/hex"
	"github.com/hyperledger/burrow/rpc/web3"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/acm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

const (
	chainID      = math.MaxInt32
	networkID    = math.MaxInt32
	maxGasLimit  = 2<<52 - 1
	hexZero      = "0x0"
	hexZeroNonce = "0x0000000000000000"
	hexOne       = "0x1"
	pending      = "null"
)

var _ web3.Service = (*EthService)(nil)

type EthService struct {
	accountState *acm.AccountState
	keyStore     acm.KeyStore
	irohaClient  *iroha.Client
}

func NewEthService(
	accountState *acm.AccountState,
	keyStore acm.KeyStore,
	irohaClient *iroha.Client,
) *EthService {
	return &EthService{
		accountState: accountState,
		keyStore:     keyStore,
		irohaClient:  irohaClient,
	}
}

func (e EthService) Web3ClientVersion() (*web3.Web3ClientVersionResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) Web3Sha3(params *web3.Web3Sha3Params) (*web3.Web3Sha3Result, error) {
	data, err := x.DecodeToBytes(params.Data)
	if err != nil {
		return nil, err
	}

	return &web3.Web3Sha3Result{
		HashedData: x.EncodeBytes(crypto.Keccak256(data)),
	}, nil
}

func (e EthService) NetListening() (*web3.NetListeningResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) NetPeerCount() (*web3.NetPeerCountResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) NetVersion() (*web3.NetVersionResult, error) {
	return &web3.NetVersionResult{
		ChainID: x.EncodeNumber(uint64(networkID)),
	}, nil
}

func (e EthService) EthBlockNumber() (*web3.EthBlockNumberResult, error) {
	height, err := e.irohaClient.GetLatestHeight()
	if err != nil {
		return nil, err
	}
	return &web3.EthBlockNumberResult{
		BlockNumber: util.ToEthereumHexString(fmt.Sprintf("%x", height)),
	}, nil
}

func (e EthService) EthCall(params *web3.EthCallParams) (*web3.EthCallResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthChainId() (*web3.EthChainIdResult, error) {
	return &web3.EthChainIdResult{
		ChainId: x.EncodeNumber(uint64(chainID)),
	}, nil
}

func (e EthService) EthCoinbase() (*web3.EthCoinbaseResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthEstimateGas(*web3.EthEstimateGasParams) (*web3.EthEstimateGasResult, error) {
	return &web3.EthEstimateGasResult{
		GasUsed: hexZero,
	}, nil
}

func (e EthService) EthGasPrice() (*web3.EthGasPriceResult, error) {
	return &web3.EthGasPriceResult{
		GasPrice: hexZero,
	}, nil
}

func (e EthService) EthGetBalance(*web3.EthGetBalanceParams) (*web3.EthGetBalanceResult, error) {
	return &web3.EthGetBalanceResult{
		GetBalanceResult: hexZero,
	}, nil
}

func (e EthService) EthGetBlockByHash(params *web3.EthGetBlockByHashParams) (*web3.EthGetBlockByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetBlockByNumber(params *web3.EthGetBlockByNumberParams) (*web3.EthGetBlockByNumberResult, error) {
	num := hexOne
	if params.BlockNumber == "earliest" {
		return &web3.EthGetBlockByNumberResult{
			GetBlockByNumberResult: web3.Block{
				Number: num,
			},
		}, nil
	}

	var height uint64
	var err error
	if params.BlockNumber == "latest" || params.BlockNumber == "pending" {
		height, err = e.irohaClient.GetLatestHeight()
		if err != nil {
			return nil, web3.ErrServer
		}
	} else {
		height, err = strconv.ParseUint(x.RemovePrefix(params.BlockNumber), 16, 64)
		if err != nil {
			return nil, web3.ErrServer
		}
	}

	acc, err := e.accountState.GetDefaultAccount()
	if err != nil {
		return nil, web3.ErrServer
	}

	q := query.GetBlock(
		height,
		query.CreatorAccountId(acc.IrohaAccountID),
	)

	_, err = e.keyStore.SignQuery(q, acc.IrohaAccountID)
	if err != nil {
		return nil, web3.ErrServer
	}

	res, err := e.irohaClient.SendQuery(context.Background(), q)
	if err != nil {
		return nil, web3.ErrServer
	}

	block := res.GetBlockResponse().GetBlock().GetBlockV1()

	return &web3.EthGetBlockByNumberResult{
		GetBlockByNumberResult: web3.Block{
			Number: util.ToEthereumHexString(fmt.Sprintf("%x", block.Payload.Height)),
			Hash:   hexZero,
		},
	}, nil
}

func (e EthService) EthGetBlockTransactionCountByHash(params *web3.EthGetBlockTransactionCountByHashParams) (*web3.EthGetBlockTransactionCountByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetBlockTransactionCountByNumber(params *web3.EthGetBlockTransactionCountByNumberParams) (*web3.EthGetBlockTransactionCountByNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetCode(params *web3.EthGetCodeParams) (*web3.EthGetCodeResult, error) {
	data, err := e.irohaClient.GetBurrowAccountDataByAddress(params.Address)
	if err != nil {
		return nil, web3.ErrServer
	} else if data == nil {
		return nil, nil
	}

	bz, err := hex.DecodeString(data.Data)
	if err != nil {
		return nil, web3.ErrServer
	}

	var irohaAccount burrow.Account
	if err = proto.Unmarshal(bz, &irohaAccount); err != nil {
		return nil, web3.ErrServer
	}

	bytes := util.ToEthereumHexString(irohaAccount.EVMCode.String())

	return &web3.EthGetCodeResult{
		Bytes: bytes,
	}, nil
}

func (e EthService) EthGetFilterChanges(params *web3.EthGetFilterChangesParams) (*web3.EthGetFilterChangesResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetFilterLogs(params *web3.EthGetFilterLogsParams) (*web3.EthGetFilterLogsResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByHash(params *web3.EthGetRawTransactionByHashParams) (*web3.EthGetRawTransactionByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByBlockHashAndIndex(params *web3.EthGetRawTransactionByBlockHashAndIndexParams) (*web3.EthGetRawTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByBlockNumberAndIndex(params *web3.EthGetRawTransactionByBlockNumberAndIndexParams) (*web3.EthGetRawTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetLogs(params *web3.EthGetLogsParams) (*web3.EthGetLogsResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetStorageAt(params *web3.EthGetStorageAtParams) (*web3.EthGetStorageAtResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByBlockHashAndIndex(params *web3.EthGetTransactionByBlockHashAndIndexParams) (*web3.EthGetTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByBlockNumberAndIndex(params *web3.EthGetTransactionByBlockNumberAndIndexParams) (*web3.EthGetTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByHash(params *web3.EthGetTransactionByHashParams) (*web3.EthGetTransactionByHashResult, error) {
	eTx, err := e.irohaClient.GetEngineTransaction(params.TransactionHash)
	if err != nil {
		return nil, web3.ErrServer
	} else if eTx == nil {
		return nil, nil
	}

	acc, err := e.accountState.GetByIrohaAccountID(eTx.CreatorID)
	if err != nil {
		return nil, web3.ErrServer
	}

	tx := web3.Transaction{
		BlockNumber:      util.ToEthereumHexString(fmt.Sprintf("%x", eTx.Height)),
		BlockHash:        hexZero,
		TransactionIndex: util.ToEthereumHexString(fmt.Sprintf("%x", eTx.Index)),
		Hash:             util.ToEthereumHexString(eTx.TxHash),
		From:             util.ToEthereumHexString(acc.IrohaAddress),
		Nonce:            hexZero,
		Gas:              hexZero,
		Value:            hexZero,
		GasPrice:         hexZero,
		S:                hexZero,
		R:                hexZero,
		V:                hexZero,
	}

	if eTx.Callee.Valid {
		tx.To = util.ToEthereumHexString(eTx.Callee.String)
	}

	if eTx.Data.Valid {
		tx.Data = util.ToEthereumHexString(eTx.Data.String)
	}

	return &web3.EthGetTransactionByHashResult{
		Transaction: tx,
	}, nil
}

func (e EthService) EthGetTransactionCount(params *web3.EthGetTransactionCountParams) (*web3.EthGetTransactionCountResult, error) {
	return &web3.EthGetTransactionCountResult{
		NonceOrNull: hexZero,
	}, nil
}

func (e EthService) EthGetTransactionReceipt(params *web3.EthGetTransactionReceiptParams) (*web3.EthGetTransactionReceiptResult, error) {
	eReceipt, err := e.irohaClient.GetEngineReceipt(params.TransactionHash)
	if err != nil {
		return nil, web3.ErrServer
	} else if eReceipt == nil {
		return nil, nil
	}

	eLogs, err := e.irohaClient.GeEngineReceiptLogsByTxHash(eReceipt.TxHash)
	if err != nil {
		return nil, web3.ErrServer
	}

	from := util.IrohaAccountIDToAddressHex(eReceipt.CreatorID)
	if err != nil {
		return nil, web3.ErrServer
	}

	receipt := web3.Receipt{
		From:              util.ToEthereumHexString(from),
		TransactionHash:   util.ToEthereumHexString(eReceipt.TxHash),
		TransactionIndex:  util.ToEthereumHexString(fmt.Sprintf("%x", eReceipt.Index)),
		BlockNumber:       util.ToEthereumHexString(fmt.Sprintf("%x", eReceipt.Height)),
		BlockHash:         hexZero,
		GasUsed:           hexZero,
		CumulativeGasUsed: hexZero,
		Logs:              irohaToEthereumTxReceiptLogs(eLogs),
		LogsBloom:         "",
	}

	if eReceipt.Status {
		receipt.Status = hexOne
	} else {
		receipt.Status = hexZero
	}

	if eReceipt.Callee.Valid {
		receipt.To = util.ToEthereumHexString(eReceipt.Callee.String)
	}

	if eReceipt.CreatedAddress.Valid {
		receipt.ContractAddress = util.ToEthereumHexString(eReceipt.CreatedAddress.String)
	}

	return &web3.EthGetTransactionReceiptResult{
		Receipt: receipt,
	}, nil
}

func (e EthService) EthGetUncleByBlockHashAndIndex(params *web3.EthGetUncleByBlockHashAndIndexParams) (*web3.EthGetUncleByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleByBlockNumberAndIndex(params *web3.EthGetUncleByBlockNumberAndIndexParams) (*web3.EthGetUncleByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleCountByBlockHash(params *web3.EthGetUncleCountByBlockHashParams) (*web3.EthGetUncleCountByBlockHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleCountByBlockNumber(params *web3.EthGetUncleCountByBlockNumberParams) (*web3.EthGetUncleCountByBlockNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetProof(params *web3.EthGetProofParams) (*web3.EthGetProofResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetWork() (*web3.EthGetWorkResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthHashrate() (*web3.EthHashrateResult, error) {
	return &web3.EthHashrateResult{HashesPerSecond: hexZero}, nil
}

func (e EthService) EthMining() (*web3.EthMiningResult, error) {
	return &web3.EthMiningResult{Mining: false}, nil
}

func (e EthService) EthNewBlockFilter() (*web3.EthNewBlockFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthNewFilter(params *web3.EthNewFilterParams) (*web3.EthNewFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthNewPendingTransactionFilter() (*web3.EthNewPendingTransactionFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthPendingTransactions() (*web3.EthPendingTransactionsResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthProtocolVersion() (*web3.EthProtocolVersionResult, error) {
	return &web3.EthProtocolVersionResult{ProtocolVersion: hexZero}, nil
}

func (e EthService) EthSign(params *web3.EthSignParams) (*web3.EthSignResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthAccounts() (*web3.EthAccountsResult, error) {
	accounts, err := e.accountState.GetAll()
	if err != nil {
		if err == acm.ErrNotFound {
			return nil, nil
		}
		return nil, web3.ErrServer
	}

	addresses := make([]string, 0, len(accounts))

	for _, acc := range accounts {
		addresses = append(addresses, util.ToEthereumHexString(acc.GetIrohaAddress()))
	}

	return &web3.EthAccountsResult{
		Addresses: addresses,
	}, nil
}

func (e EthService) EthSendTransaction(params *web3.EthSendTransactionParams) (*web3.EthSendTransactionResult, error) {
	acc, err := e.accountState.GetByIrohaAddress(params.From)
	if err != nil {
		if err == acm.ErrNotFound {
			return nil, nil
		}
		return nil, web3.ErrServer
	}

	accountID := acc.GetIrohaAccountID()
	contractAddress := util.ToIrohaHexString(params.To)
	input := x.RemovePrefix(params.Data)

	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.CallEngine(accountID, contractAddress, input),
			},
			command.CreatorAccountId(accountID),
		),
	)

	_, err = e.keyStore.SignTransaction(tx, accountID)
	if err == acm.ErrNotFound {
		return nil, web3.ErrServer
	}

	txHash, err := e.irohaClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, web3.ErrServer
	}

	_, err = e.irohaClient.TxStatusStream(context.Background(), txHash)
	if err != nil {
		return nil, web3.ErrServer
	}

	return &web3.EthSendTransactionResult{
		TransactionHash: util.ToEthereumHexString(txHash),
	}, nil
}

func (e EthService) EthSendRawTransaction(params *web3.EthSendRawTransactionParams) (*web3.EthSendRawTransactionResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSubmitHashrate(params *web3.EthSubmitHashrateParams) (*web3.EthSubmitHashrateResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSubmitWork(params *web3.EthSubmitWorkParams) (*web3.EthSubmitWorkResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSyncing() (*web3.EthSyncingResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthUninstallFilter(params *web3.EthUninstallFilterParams) (*web3.EthUninstallFilterResult, error) {
	return nil, errors.New("implement me")
}

func irohaToEthereumTxReceiptLogs(logs []*entity.EngineReceiptLog) []web3.Logs {
	ethLogs := make([]web3.Logs, 0, len(logs))

	for i, log := range logs {
		ethLog := web3.Logs{
			LogIndex:         util.ToEthereumHexString(fmt.Sprintf("%x", i)),
			TransactionIndex: util.ToEthereumHexString(fmt.Sprintf("%x", log.Index)),
			TransactionHash:  util.ToEthereumHexString(log.TxHash),
			Address:          util.ToEthereumHexString(log.Address),
			BlockHash:        hexZero,
			BlockNumber:      util.ToEthereumHexString(fmt.Sprintf("%x", log.Height)),
			Data:             util.ToEthereumHexString(log.Data),
			Topics:           make([]web3.Topics, 0, len(log.Topics)),
		}

		for _, topic := range log.Topics {
			ethTopic := web3.Topics{
				DataWord: util.ToEthereumHexString(topic.Topic),
			}
			ethLog.Topics = append(ethLog.Topics, ethTopic)
		}

		ethLogs = append(ethLogs, ethLog)
	}

	return ethLogs
}
