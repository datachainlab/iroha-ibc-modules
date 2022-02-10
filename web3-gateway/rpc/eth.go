package rpc

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
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
	chainID       = math.MaxInt32 // TODO configurable?
	networkID     = math.MaxInt32 // TODO configurable?
	hexZero       = "0x0"
	hexOne        = "0x1"
	zeroHash      = "0x0000000000000000000000000000000000000000000000000000000000000000"
	zeroAddress   = "0x0000000000000000000000000000000000000000"
	zeroLogBlooms = "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	zeroNonce     = "0x0000000000000000"
	zeroBytes     = ""
)

var _ EthService = (*ethService)(nil)

type EthService interface {
	Web3ClientVersion() (*web3.Web3ClientVersionResult, error)
	Web3Sha3(params *web3.Web3Sha3Params) (*web3.Web3Sha3Result, error)
	NetListening() (*web3.NetListeningResult, error)
	NetPeerCount() (*web3.NetPeerCountResult, error)
	NetVersion() (*web3.NetVersionResult, error)
	EthBlockNumber() (*web3.EthBlockNumberResult, error)
	EthCall(params *web3.EthCallParams) (*web3.EthCallResult, error)
	EthChainId() (*web3.EthChainIdResult, error)
	EthCoinbase() (*web3.EthCoinbaseResult, error)
	EthEstimateGas(*web3.EthEstimateGasParams) (*web3.EthEstimateGasResult, error)
	EthGasPrice() (*web3.EthGasPriceResult, error)
	EthGetBalance(*web3.EthGetBalanceParams) (*web3.EthGetBalanceResult, error)
	EthGetBlockByHash(*web3.EthGetBlockByHashParams) (*web3.EthGetBlockByHashResult, error)
	EthGetBlockByNumber(params *web3.EthGetBlockByNumberParams) (*web3.EthGetBlockByNumberResult, error)
	EthGetBlockTransactionCountByHash(*web3.EthGetBlockTransactionCountByHashParams) (*web3.EthGetBlockTransactionCountByHashResult, error)
	EthGetBlockTransactionCountByNumber(*web3.EthGetBlockTransactionCountByNumberParams) (*web3.EthGetBlockTransactionCountByNumberResult, error)
	EthGetCode(params *web3.EthGetCodeParams) (*web3.EthGetCodeResult, error)
	EthGetFilterChanges(*web3.EthGetFilterChangesParams) (*web3.EthGetFilterChangesResult, error)
	EthGetFilterLogs(*web3.EthGetFilterLogsParams) (*web3.EthGetFilterLogsResult, error)
	EthGetRawTransactionByHash(*web3.EthGetRawTransactionByHashParams) (*web3.EthGetRawTransactionByHashResult, error)
	EthGetRawTransactionByBlockHashAndIndex(*web3.EthGetRawTransactionByBlockHashAndIndexParams) (*web3.EthGetRawTransactionByBlockHashAndIndexResult, error)
	EthGetRawTransactionByBlockNumberAndIndex(*web3.EthGetRawTransactionByBlockNumberAndIndexParams) (*web3.EthGetRawTransactionByBlockNumberAndIndexResult, error)
	EthGetLogs(params *EthGetLogsParams) (*EthGetLogsResult, error)
	EthGetStorageAt(*web3.EthGetStorageAtParams) (*web3.EthGetStorageAtResult, error)
	EthGetTransactionByBlockHashAndIndex(*web3.EthGetTransactionByBlockHashAndIndexParams) (*web3.EthGetTransactionByBlockHashAndIndexResult, error)
	EthGetTransactionByBlockNumberAndIndex(*web3.EthGetTransactionByBlockNumberAndIndexParams) (*web3.EthGetTransactionByBlockNumberAndIndexResult, error)
	EthGetTransactionByHash(params *web3.EthGetTransactionByHashParams) (*web3.EthGetTransactionByHashResult, error)
	EthGetTransactionCount(*web3.EthGetTransactionCountParams) (*web3.EthGetTransactionCountResult, error)
	EthGetTransactionReceipt(params *web3.EthGetTransactionReceiptParams) (*EthGetTransactionReceiptResult, error)
	EthGetUncleByBlockHashAndIndex(*web3.EthGetUncleByBlockHashAndIndexParams) (*web3.EthGetUncleByBlockHashAndIndexResult, error)
	EthGetUncleByBlockNumberAndIndex(*web3.EthGetUncleByBlockNumberAndIndexParams) (*web3.EthGetUncleByBlockNumberAndIndexResult, error)
	EthGetUncleCountByBlockHash(*web3.EthGetUncleCountByBlockHashParams) (*web3.EthGetUncleCountByBlockHashResult, error)
	EthGetUncleCountByBlockNumber(*web3.EthGetUncleCountByBlockNumberParams) (*web3.EthGetUncleCountByBlockNumberResult, error)
	EthGetProof(*web3.EthGetProofParams) (*web3.EthGetProofResult, error)
	EthGetWork() (*web3.EthGetWorkResult, error)
	EthHashrate() (*web3.EthHashrateResult, error)
	EthMining() (*web3.EthMiningResult, error)
	EthNewBlockFilter() (*web3.EthNewBlockFilterResult, error)
	EthNewFilter(*web3.EthNewFilterParams) (*web3.EthNewFilterResult, error)
	EthNewPendingTransactionFilter() (*web3.EthNewPendingTransactionFilterResult, error)
	EthPendingTransactions() (*web3.EthPendingTransactionsResult, error)
	EthProtocolVersion() (*web3.EthProtocolVersionResult, error)
	EthSign(*web3.EthSignParams) (*web3.EthSignResult, error)
	EthAccounts() (*web3.EthAccountsResult, error)
	EthSendTransaction(params *web3.EthSendTransactionParams) (*web3.EthSendTransactionResult, error)
	EthSendRawTransaction(params *web3.EthSendRawTransactionParams) (*web3.EthSendRawTransactionResult, error)
	EthSubmitHashrate(*web3.EthSubmitHashrateParams) (*web3.EthSubmitHashrateResult, error)
	EthSubmitWork(*web3.EthSubmitWorkParams) (*web3.EthSubmitWorkResult, error)
	EthSyncing() (*web3.EthSyncingResult, error)
	EthUninstallFilter(*web3.EthUninstallFilterParams) (*web3.EthUninstallFilterResult, error)
}

type ethService struct {
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
) EthService {
	return &ethService{
		accountState:      accountState,
		keyStore:          keyStore,
		irohaAPIClient:    irohaAPIClient,
		irohaDBTransactor: irohaDBTransactor,
		logger:            logger,
		querier:           querier,
	}
}

func (e ethService) Web3ClientVersion() (*web3.Web3ClientVersionResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) Web3Sha3(params *web3.Web3Sha3Params) (*web3.Web3Sha3Result, error) {
	data, err := x.DecodeToBytes(params.Data)
	if err != nil {
		return nil, err
	}

	return &web3.Web3Sha3Result{
		HashedData: x.EncodeBytes(crypto.Keccak256(data)),
	}, nil
}

func (e ethService) NetListening() (*web3.NetListeningResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) NetPeerCount() (*web3.NetPeerCountResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) NetVersion() (*web3.NetVersionResult, error) {
	return &web3.NetVersionResult{
		ChainID: x.EncodeNumber(uint64(networkID)),
	}, nil
}

func (e ethService) EthBlockNumber() (*web3.EthBlockNumberResult, error) {
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

func (e ethService) EthCall(params *web3.EthCallParams) (*web3.EthCallResult, error) {
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

func (e ethService) EthChainId() (*web3.EthChainIdResult, error) {
	return &web3.EthChainIdResult{
		ChainId: x.EncodeNumber(uint64(chainID)),
	}, nil
}

func (e ethService) EthCoinbase() (*web3.EthCoinbaseResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthEstimateGas(*web3.EthEstimateGasParams) (*web3.EthEstimateGasResult, error) {
	return &web3.EthEstimateGasResult{
		GasUsed: hexZero,
	}, nil
}

func (e ethService) EthGasPrice() (*web3.EthGasPriceResult, error) {
	return &web3.EthGasPriceResult{
		GasPrice: hexZero,
	}, nil
}

func (e ethService) EthGetBalance(*web3.EthGetBalanceParams) (*web3.EthGetBalanceResult, error) {
	return &web3.EthGetBalanceResult{
		GetBalanceResult: hexZero,
	}, nil
}

func (e ethService) EthGetBlockByHash(*web3.EthGetBlockByHashParams) (*web3.EthGetBlockByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetBlockByNumber(params *web3.EthGetBlockByNumberParams) (*web3.EthGetBlockByNumberResult, error) {
	var height uint64
	var err error

	switch params.BlockNumber {
	case "earliest":
		height = 1
	case "latest":
		fallthrough
	case "pending":
		err = e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
			height, err = querier.GetLatestHeight()
			return
		})
	default:
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
			ParentHash:       zeroHash,
			Sha3Uncles:       types.EmptyUncleHash.Hex(),
			Miner:            zeroAddress,
			StateRoot:        zeroHash,
			TransactionsRoot: types.EmptyRootHash.Hex(),
			ReceiptsRoot:     zeroHash,
			LogsBloom:        zeroLogBlooms,
			Difficulty:       hexZero,
			Number:           util.ToEthereumHexString(fmt.Sprintf("%x", block.Payload.GetHeight())),
			GasLimit:         hexZero,
			GasUsed:          hexZero,
			Timestamp:        util.ToEthereumHexString(fmt.Sprintf("%x", block.Payload.GetCreatedTime())),
			ExtraData:        zeroBytes,
			Hash:             zeroHash,
			TotalDifficulty:  hexZero,
			Size:             hexZero,
			Nonce:            zeroNonce,
			Transactions:     nil,
			Uncles:           nil,
		},
	}, nil
}

func (e ethService) EthGetBlockTransactionCountByHash(*web3.EthGetBlockTransactionCountByHashParams) (*web3.EthGetBlockTransactionCountByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetBlockTransactionCountByNumber(*web3.EthGetBlockTransactionCountByNumberParams) (*web3.EthGetBlockTransactionCountByNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetCode(params *web3.EthGetCodeParams) (*web3.EthGetCodeResult, error) {
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

func (e ethService) EthGetFilterChanges(*web3.EthGetFilterChangesParams) (*web3.EthGetFilterChangesResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetFilterLogs(*web3.EthGetFilterLogsParams) (*web3.EthGetFilterLogsResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetRawTransactionByHash(*web3.EthGetRawTransactionByHashParams) (*web3.EthGetRawTransactionByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetRawTransactionByBlockHashAndIndex(*web3.EthGetRawTransactionByBlockHashAndIndexParams) (*web3.EthGetRawTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetRawTransactionByBlockNumberAndIndex(*web3.EthGetRawTransactionByBlockNumberAndIndexParams) (*web3.EthGetRawTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetLogs(params *EthGetLogsParams) (*EthGetLogsResult, error) {
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

	if addresses, err := params.Address(); err != nil {
		return nil, err
	} else if len(addresses) > 0 {
		filterOpts = append(filterOpts, db.AddressesOption(addresses))
	}

	if len(params.Topics) > 0 {
		filterOpts = append(filterOpts, db.TopicsOption(params.Topics))
	}

	var eLogs []*entity.EngineReceiptLog

	if err := e.irohaDBTransactor.Exec(context.Background(), e.querier, func(querier db.DBExecer) (err error) {
		eLogs, err = querier.GetEngineReceiptLogsByFilters(filterOpts...)
		return
	}); err != nil {
		return nil, err
	}

	return &EthGetLogsResult{
		Logs: irohaToEthereumTxReceiptLogs(eLogs),
	}, nil
}

func (e ethService) EthGetStorageAt(*web3.EthGetStorageAtParams) (*web3.EthGetStorageAtResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetTransactionByBlockHashAndIndex(*web3.EthGetTransactionByBlockHashAndIndexParams) (*web3.EthGetTransactionByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetTransactionByBlockNumberAndIndex(*web3.EthGetTransactionByBlockNumberAndIndexParams) (*web3.EthGetTransactionByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetTransactionByHash(params *web3.EthGetTransactionByHashParams) (*web3.EthGetTransactionByHashResult, error) {
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
		BlockHash:        zeroHash,
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

func (e ethService) EthGetTransactionCount(*web3.EthGetTransactionCountParams) (*web3.EthGetTransactionCountResult, error) {
	return &web3.EthGetTransactionCountResult{
		NonceOrNull: hexZero,
	}, nil
}

func (e ethService) EthGetTransactionReceipt(params *web3.EthGetTransactionReceiptParams) (*EthGetTransactionReceiptResult, error) {
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
	receipt := Receipt{
		Receipt: web3.Receipt{
			From:              util.ToEthereumHexString(from),
			To:                zeroAddress,
			ContractAddress:   zeroAddress,
			TransactionHash:   util.ToEthereumHexString(eReceipt.TxHash),
			TransactionIndex:  util.ToEthereumHexString(fmt.Sprintf("%x", eReceipt.Index)),
			BlockNumber:       util.ToEthereumHexString(fmt.Sprintf("%x", eReceipt.Height)),
			BlockHash:         zeroHash,
			GasUsed:           hexZero,
			CumulativeGasUsed: hexZero,
			LogsBloom:         zeroLogBlooms,
		},
		Logs: irohaToEthereumTxReceiptLogs(eLogs),
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

	return &EthGetTransactionReceiptResult{
		Receipt: receipt,
	}, nil
}

func (e ethService) EthGetUncleByBlockHashAndIndex(*web3.EthGetUncleByBlockHashAndIndexParams) (*web3.EthGetUncleByBlockHashAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetUncleByBlockNumberAndIndex(*web3.EthGetUncleByBlockNumberAndIndexParams) (*web3.EthGetUncleByBlockNumberAndIndexResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetUncleCountByBlockHash(*web3.EthGetUncleCountByBlockHashParams) (*web3.EthGetUncleCountByBlockHashResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetUncleCountByBlockNumber(*web3.EthGetUncleCountByBlockNumberParams) (*web3.EthGetUncleCountByBlockNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetProof(*web3.EthGetProofParams) (*web3.EthGetProofResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthGetWork() (*web3.EthGetWorkResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthHashrate() (*web3.EthHashrateResult, error) {
	return &web3.EthHashrateResult{HashesPerSecond: hexZero}, nil
}

func (e ethService) EthMining() (*web3.EthMiningResult, error) {
	return &web3.EthMiningResult{Mining: false}, nil
}

func (e ethService) EthNewBlockFilter() (*web3.EthNewBlockFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthNewFilter(*web3.EthNewFilterParams) (*web3.EthNewFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthNewPendingTransactionFilter() (*web3.EthNewPendingTransactionFilterResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthPendingTransactions() (*web3.EthPendingTransactionsResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthProtocolVersion() (*web3.EthProtocolVersionResult, error) {
	return &web3.EthProtocolVersionResult{ProtocolVersion: hexZero}, nil
}

func (e ethService) EthSign(*web3.EthSignParams) (*web3.EthSignResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthAccounts() (*web3.EthAccountsResult, error) {
	accounts, err := e.accountState.GetAll()
	if err != nil {
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

func (e ethService) EthSendTransaction(params *web3.EthSendTransactionParams) (*web3.EthSendTransactionResult, error) {
	acc, err := e.accountState.GetByIrohaAddress(params.From)
	if err != nil {
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

func (e ethService) EthSendRawTransaction(
	params *web3.EthSendRawTransactionParams,
) (*web3.EthSendRawTransactionResult, error) {
	data, err := x.DecodeToBytes(params.SignedTransactionData)
	if err != nil {
		return nil, err
	}

	rawTx := new(types.Transaction)
	if err := rawTx.DecodeRLP(rlp.NewStream(bytes.NewReader(data), uint64(len(data)))); err != nil {
		return nil, err
	}

	var signer types.Signer = types.FrontierSigner{}
	if rawTx.Protected() {
		signer = types.NewEIP155Signer(rawTx.ChainId())
	}
	// Signature to Ethereum Address
	ethAddress, _ := types.Sender(signer, rawTx)
	// Get Iroha Account by Ethereum Address
	acc, err := e.accountState.GetByEthereumAddress(ethAddress.Hex())
	if err != nil {
		return nil, err
	}
	accountID := acc.GetIrohaAccountID()

	//Conversion to Iroha representation
	contractAddress := util.ToIrohaHexString(rawTx.To().Hex())

	input := hex.EncodeToString(rawTx.Data())

	tx := command.BuildTransaction(
		command.BuildPayload(
			[]*pb.Command{
				command.CallEngine(accountID, contractAddress, input),
			},
			command.CreatorAccountId(accountID),
		),
	)

	_, err = e.keyStore.SignTransaction(tx, accountID)
	if err != nil {
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

	return &web3.EthSendRawTransactionResult{
		TransactionHash: util.ToEthereumHexString(txHash),
	}, nil
}

func (e ethService) EthSubmitHashrate(*web3.EthSubmitHashrateParams) (*web3.EthSubmitHashrateResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthSubmitWork(*web3.EthSubmitWorkParams) (*web3.EthSubmitWorkResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthSyncing() (*web3.EthSyncingResult, error) {
	return nil, errors.New("implement me")
}

func (e ethService) EthUninstallFilter(*web3.EthUninstallFilterParams) (*web3.EthUninstallFilterResult, error) {
	return nil, errors.New("implement me")
}

func irohaToEthereumTxReceiptLogs(logs []*entity.EngineReceiptLog) []Logs {
	ethLogs := make([]Logs, 0, len(logs))

	for i, log := range logs {
		ethLog := Logs{
			Logs: web3.Logs{
				LogIndex:         util.ToEthereumHexString(fmt.Sprintf("%x", i)),
				TransactionIndex: util.ToEthereumHexString(fmt.Sprintf("%x", log.Index)),
				TransactionHash:  util.ToEthereumHexString(log.TxHash),
				Address:          util.ToEthereumHexString(log.Address),
				BlockHash:        zeroHash,
				BlockNumber:      util.ToEthereumHexString(fmt.Sprintf("%x", log.Height)),
				Data:             util.ToEthereumHexString(log.Data),
			},
			Topics: make([]string, 0, len(log.Topics)),
		}

		for _, topic := range log.Topics {
			ethTopic := util.ToEthereumHexString(topic.Topic)
			ethLog.Topics = append(ethLog.Topics, ethTopic)
		}

		ethLogs = append(ethLogs, ethLog)
	}

	return ethLogs
}
