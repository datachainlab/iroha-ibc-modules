package rpc

import (
	"errors"
	"math"

	"github.com/hyperledger/burrow/crypto"
	x "github.com/hyperledger/burrow/encoding/hex"
	"github.com/hyperledger/burrow/rpc/web3"
)

const (
	chainID      = math.MaxInt16
	networkID    = math.MaxInt16
	maxGasLimit  = 2<<52 - 1
	hexZero      = "0x0"
	hexZeroNonce = "0x0000000000000000"
	pending      = "null"
)

var _ web3.Service = (*EthService)(nil)

type EthService struct{}

func NewEthService() *EthService {
	return &EthService{}
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
	return nil, errors.New("implement me")
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
	return nil, errors.New("implement me")
}

func (e EthService) EthGetBlockTransactionCountByHash(params *web3.EthGetBlockTransactionCountByHashParams) (*web3.EthGetBlockTransactionCountByHashResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetBlockTransactionCountByNumber(params *web3.EthGetBlockTransactionCountByNumberParams) (*web3.EthGetBlockTransactionCountByNumberResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthGetCode(params *web3.EthGetCodeParams) (*web3.EthGetCodeResult, error) {
	return nil, errors.New("implement me")
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
	return nil, errors.New("implement me")
}

func (e EthService) EthGetTransactionCount(params *web3.EthGetTransactionCountParams) (*web3.EthGetTransactionCountResult, error) {
	return &web3.EthGetTransactionCountResult{
		NonceOrNull: hexZero,
	}, nil
}

func (e EthService) EthGetTransactionReceipt(params *web3.EthGetTransactionReceiptParams) (*web3.EthGetTransactionReceiptResult, error) {
	return nil, errors.New("implement me")
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
	return nil, errors.New("implement me")
}

func (e EthService) EthMining() (*web3.EthMiningResult, error) {
	return nil, errors.New("implement me")
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
	return nil, errors.New("implement me")
}

func (e EthService) EthSign(params *web3.EthSignParams) (*web3.EthSignResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthAccounts() (*web3.EthAccountsResult, error) {
	return nil, errors.New("implement me")
}

func (e EthService) EthSendTransaction(params *web3.EthSendTransactionParams) (*web3.EthSendTransactionResult, error) {
	return nil, errors.New("implement me")
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
	panic("implement me")
}

func (e EthService) EthUninstallFilter(params *web3.EthUninstallFilterParams) (*web3.EthUninstallFilterResult, error) {
	panic("implement me")
}
