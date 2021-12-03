package rpc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"

	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/crypto"
	x "github.com/hyperledger/burrow/encoding/hex"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/rpc/web3"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/acm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/evm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/api"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/keyring"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

const (
	chainID   = math.MaxInt32
	networkID = math.MaxInt32
	hexZero   = "0x0"
	hexOne    = "0x1"
)

var _ web3.Service = (*EthService)(nil)

type EthService struct {
	accountState      *acm.AccountState
	keyStore          keyring.KeyStore
	irohaAPIClient    api.ApiClient
	irohaDBTransactor db.DBTransactor
	logger            *logging.Logger
	querier           string
}

func NewEthService(
	accountState *acm.AccountState,
	keyStore keyring.KeyStore,
	irohaAPIClient api.ApiClient,
	irohaDBTransactor db.DBTransactor,
	logger *logging.Logger,
	querier string,
) *EthService {
	return &EthService{
		accountState:      accountState,
		keyStore:          keyStore,
		irohaAPIClient:    irohaAPIClient,
		irohaDBTransactor: irohaDBTransactor,
		logger:            logger,
		querier:           querier,
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
	var height uint64

	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		height, err = querier.GetLatestHeight()
		return
	}); err != nil {
		return nil, err
	}
	return &web3.EthBlockNumberResult{
		BlockNumber: util.ToEthereumHexString(fmt.Sprintf("%x", height)),
	}, nil
}

func (e EthService) EthCall(params *web3.EthCallParams) (*web3.EthCallResult, error) {
	input, err := x.DecodeToBytes(params.Transaction.Data)
	if err != nil {
		return nil, err
	}

	callerAccount, err := e.accountState.GetByIrohaAddress(params.From)
	if err != nil {
		return nil, err
	}

	res, err := evm.CallSim(
		e.irohaDBTransactor, e.logger,
		callerAccount.IrohaAccountID,
		params.From, params.To,
		input,
	)
	if err != nil {
		return nil, err
	}

	return &web3.EthCallResult{
		ReturnValue: x.EncodeBytes(res),
	}, nil
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

func (e EthService) EthGetBlockByHash(*web3.EthGetBlockByHashParams) (*web3.EthGetBlockByHashResult, error) {
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
		err = e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
			height, err = querier.GetLatestHeight()
			return
		})
	} else {
		height, err = strconv.ParseUint(x.RemovePrefix(params.BlockNumber), 16, 64)
	}
	if err != nil {
		return nil, err
	}

	q := query.GetBlock(height, query.CreatorAccountId(e.querier))

	_, err = e.keyStore.SignQuery(q, e.querier)
	if err != nil {
		return nil, err
	}

	res, err := e.irohaAPIClient.SendQuery(context.Background(), q)
	if err != nil {
		return nil, err
	}

	block := res.GetBlockResponse().GetBlock().GetBlockV1()

	return &web3.EthGetBlockByNumberResult{
		GetBlockByNumberResult: web3.Block{
			Number:    util.ToEthereumHexString(fmt.Sprintf("%x", block.Payload.GetHeight())),
			Hash:      hexZero,
			Timestamp: util.ToEthereumHexString(fmt.Sprintf("%x", block.Payload.GetCreatedTime())),
		},
	}, nil
}

func (e EthService) EthGetBlockTransactionCountByHash(*web3.EthGetBlockTransactionCountByHashParams) (*web3.EthGetBlockTransactionCountByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetBlockTransactionCountByNumber(*web3.EthGetBlockTransactionCountByNumberParams) (*web3.EthGetBlockTransactionCountByNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetCode(params *web3.EthGetCodeParams) (*web3.EthGetCodeResult, error) {
	var data *entity.BurrowAccountData
	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		data, err = querier.GetBurrowAccountDataByAddress(params.Address)
		return
	}); err != nil {
		return nil, err
	} else if data == nil {
		return nil, nil
	}

	bz, err := hex.DecodeString(data.Data)
	if err != nil {
		return nil, err
	}

	irohaAccount := &burrow.Account{}
	if err = irohaAccount.Unmarshal(bz); err != nil {
		return nil, err
	}

	bytes := util.ToEthereumHexString(irohaAccount.EVMCode.String())

	return &web3.EthGetCodeResult{
		Bytes: bytes,
	}, nil
}

func (e EthService) EthGetFilterChanges(*web3.EthGetFilterChangesParams) (*web3.EthGetFilterChangesResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetFilterLogs(*web3.EthGetFilterLogsParams) (*web3.EthGetFilterLogsResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByHash(*web3.EthGetRawTransactionByHashParams) (*web3.EthGetRawTransactionByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByBlockHashAndIndex(*web3.EthGetRawTransactionByBlockHashAndIndexParams) (*web3.EthGetRawTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetRawTransactionByBlockNumberAndIndex(*web3.EthGetRawTransactionByBlockNumberAndIndexParams) (*web3.EthGetRawTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetLogs(params *web3.EthGetLogsParams) (*web3.EthGetLogsResult, error) {
	var filterOpts []db.LogFilterOption

	switch params.FromBlock {
	case "":
		fallthrough
	case "pending":
		fallthrough
	case "latest":
		if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) error {
			height, err := querier.GetLatestHeight()
			if err != nil {
				return err
			}
			filterOpts = append(filterOpts, db.FromBlockOption(height))

			return nil
		}); err != nil {
			return nil, err
		}
	case "earliest":
		filterOpts = append(filterOpts, db.FromBlockOption(1))
	default:
		height, err := strconv.ParseUint(x.RemovePrefix(params.FromBlock), 10, 64)
		if err != nil {
			return nil, err
		}
		filterOpts = append(filterOpts, db.FromBlockOption(height))
	}

	switch params.ToBlock {
	case "":
		fallthrough
	case "pending":
		fallthrough
	case "latest":
		if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) error {
			height, err := querier.GetLatestHeight()
			if err != nil {
				return err
			}

			filterOpts = append(filterOpts, db.ToBlockOption(height))

			return nil
		}); err != nil {
			return nil, err
		}

	case "earliest":
		filterOpts = append(filterOpts, db.ToBlockOption(1))
	default:
		height, err := strconv.ParseUint(x.RemovePrefix(params.ToBlock), 10, 64)
		if err != nil {
			return nil, err
		}
		filterOpts = append(filterOpts, db.ToBlockOption(height))
	}

	if len(params.Address) > 0 {
		filterOpts = append(filterOpts, db.AddressOption(params.Address))
	}

	if len(params.Topics) > 0 {
		filterOpts = append(filterOpts, db.TopicsOption(params.Topics...))
	}

	var eLogs []*entity.EngineReceiptLog

	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		eLogs, err = querier.GetEngineReceiptLogsByFilters(filterOpts...)
		return
	}); err != nil {
		return nil, err
	}

	return &web3.EthGetLogsResult{
		Logs: irohaToEthereumTxReceiptLogs(eLogs),
	}, nil
}

func (e EthService) EthGetStorageAt(*web3.EthGetStorageAtParams) (*web3.EthGetStorageAtResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByBlockHashAndIndex(*web3.EthGetTransactionByBlockHashAndIndexParams) (*web3.EthGetTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByBlockNumberAndIndex(*web3.EthGetTransactionByBlockNumberAndIndexParams) (*web3.EthGetTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionByHash(params *web3.EthGetTransactionByHashParams) (*web3.EthGetTransactionByHashResult, error) {
	var eTx *entity.EngineTransaction
	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		eTx, err = querier.GetEngineTransaction(params.TransactionHash)
		return
	}); err != nil {
		return nil, err
	} else if eTx == nil {
		return nil, nil
	}

	acc, err := e.accountState.GetByIrohaAccountID(eTx.CreatorID)
	if err != nil {
		return nil, err
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

func (e EthService) EthGetTransactionCount(*web3.EthGetTransactionCountParams) (*web3.EthGetTransactionCountResult, error) {
	return &web3.EthGetTransactionCountResult{
		NonceOrNull: hexZero,
	}, nil
}

func (e EthService) EthGetTransactionReceipt(params *web3.EthGetTransactionReceiptParams) (*web3.EthGetTransactionReceiptResult, error) {
	var eReceipt *entity.EngineReceipt

	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		eReceipt, err = querier.GetEngineReceipt(params.TransactionHash)
		return
	}); err != nil {
		return nil, err
	} else if eReceipt == nil {
		return nil, nil
	}

	var eLogs []*entity.EngineReceiptLog

	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		eLogs, err = querier.GetEngineReceiptLogsByTxHash(eReceipt.TxHash)
		return
	}); err != nil {
		return nil, err
	}

	from := util.IrohaAccountIDToAddressHex(eReceipt.CreatorID)
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

func (e EthService) EthGetUncleByBlockHashAndIndex(*web3.EthGetUncleByBlockHashAndIndexParams) (*web3.EthGetUncleByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleByBlockNumberAndIndex(*web3.EthGetUncleByBlockNumberAndIndexParams) (*web3.EthGetUncleByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleCountByBlockHash(*web3.EthGetUncleCountByBlockHashParams) (*web3.EthGetUncleCountByBlockHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetUncleCountByBlockNumber(*web3.EthGetUncleCountByBlockNumberParams) (*web3.EthGetUncleCountByBlockNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetProof(*web3.EthGetProofParams) (*web3.EthGetProofResult, error) {
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

func (e EthService) EthNewFilter(*web3.EthNewFilterParams) (*web3.EthNewFilterResult, error) {
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

func (e EthService) EthSign(*web3.EthSignParams) (*web3.EthSignResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthAccounts() (*web3.EthAccountsResult, error) {
	accounts, err := e.accountState.GetAll()
	if err != nil {
		if err == acm.ErrNotFound {
			return nil, nil
		}
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	txHash, err := e.irohaAPIClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return nil, err
	}

	_, err = e.irohaAPIClient.TxStatusStream(context.Background(), txHash)
	if err != nil {
		return nil, err
	}

	return &web3.EthSendTransactionResult{
		TransactionHash: util.ToEthereumHexString(txHash),
	}, nil
}

func (e EthService) EthSendRawTransaction(*web3.EthSendRawTransactionParams) (*web3.EthSendRawTransactionResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSubmitHashrate(*web3.EthSubmitHashrateParams) (*web3.EthSubmitHashrateResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSubmitWork(*web3.EthSubmitWorkParams) (*web3.EthSubmitWorkResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSyncing() (*web3.EthSyncingResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthUninstallFilter(*web3.EthUninstallFilterParams) (*web3.EthUninstallFilterResult, error) {
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
