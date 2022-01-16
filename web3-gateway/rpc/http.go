package rpc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/gorilla/handlers"
	"github.com/hyperledger/burrow/rpc/web3"
)

type HTTPServer struct {
	handler http.Handler
}

func NewHTTPServer(ethService EthService) *HTTPServer {
	srv := &HTTPServer{}

	cdcMap := srv.endpointCodecMap(ethService)

	srv.handler = handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}), // TODO configurable
	)(jsonrpc.NewServer(cdcMap))

	return srv
}

func (srv *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.handler.ServeHTTP(w, r)
}

func (srv *HTTPServer) endpointCodecMap(ethService EthService) jsonrpc.EndpointCodecMap {
	return jsonrpc.EndpointCodecMap{
		"web3_clientVersion": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.Web3ClientVersion()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"web3_sha3": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.Web3Sha3Params)
				res, err = ethService.Web3Sha3(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (request interface{}, err error) {
				req := new(web3.Web3Sha3Params)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, err
			},
			Encode: jsonEncoder,
		},
		"net_listening": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.NetListening()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"net_peerCount": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.NetPeerCount()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"net_version": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.NetVersion()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_blockNumber": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthBlockNumber()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_call": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthCallParams)
				res, err = ethService.EthCall(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthCallParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_chainId": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthChainId()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_coinbase": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthCoinbase()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_estimateGas": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthEstimateGasParams)
				res, err = ethService.EthEstimateGas(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthEstimateGasParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_gasPrice": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthGasPrice()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_getBalance": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetBalanceParams)
				res, err = ethService.EthGetBalance(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetBalanceParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getBlockByHash": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetBlockByHashParams)
				res, err = ethService.EthGetBlockByHash(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetBlockByHashParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getBlockByNumber": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetBlockByNumberParams)
				res, err = ethService.EthGetBlockByNumber(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetBlockByNumberParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getBlockTransactionCountByHash": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetBlockTransactionCountByHashParams)
				res, err = ethService.EthGetBlockTransactionCountByHash(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetBlockTransactionCountByHashParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getBlockTransactionCountByNumber": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetBlockTransactionCountByNumberParams)
				res, err = ethService.EthGetBlockTransactionCountByNumber(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetBlockTransactionCountByNumberParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getCode": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetCodeParams)
				res, err = ethService.EthGetCode(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetCodeParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getFilterChanges": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetFilterChangesParams)
				res, err = ethService.EthGetFilterChanges(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetFilterChangesParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getFilterLogs": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetFilterLogsParams)
				res, err = ethService.EthGetFilterLogs(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetFilterLogsParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getRawTransactionByHash": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetRawTransactionByHashParams)
				res, err = ethService.EthGetRawTransactionByHash(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetRawTransactionByHashParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getRawTransactionByBlockHashAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetRawTransactionByBlockHashAndIndexParams)
				res, err = ethService.EthGetRawTransactionByBlockHashAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetRawTransactionByBlockHashAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getRawTransactionByBlockNumberAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetRawTransactionByBlockNumberAndIndexParams)
				res, err = ethService.EthGetRawTransactionByBlockNumberAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetRawTransactionByBlockNumberAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getLogs": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetLogsParams)
				res, err = ethService.EthGetLogs(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetLogsParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getStorageAt": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetStorageAtParams)
				res, err = ethService.EthGetStorageAt(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetStorageAtParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getTransactionByBlockHashAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetTransactionByBlockHashAndIndexParams)
				res, err = ethService.EthGetTransactionByBlockHashAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetTransactionByBlockHashAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getTransactionByBlockNumberAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetTransactionByBlockNumberAndIndexParams)
				res, err = ethService.EthGetTransactionByBlockNumberAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetTransactionByBlockNumberAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getTransactionByHash": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetTransactionByHashParams)
				res, err = ethService.EthGetTransactionByHash(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetTransactionByHashParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getTransactionCount": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetTransactionCountParams)
				res, err = ethService.EthGetTransactionCount(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetTransactionCountParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getTransactionReceipt": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetTransactionReceiptParams)
				res, err = ethService.EthGetTransactionReceipt(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetTransactionReceiptParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getUncleByBlockHashAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetUncleByBlockHashAndIndexParams)
				res, err = ethService.EthGetUncleByBlockHashAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetUncleByBlockHashAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getUncleByBlockNumberAndIndex": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetUncleByBlockNumberAndIndexParams)
				res, err = ethService.EthGetUncleByBlockNumberAndIndex(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetUncleByBlockNumberAndIndexParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getUncleCountByBlockHash": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetUncleCountByBlockHashParams)
				res, err = ethService.EthGetUncleCountByBlockHash(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetUncleCountByBlockHashParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getUncleCountByBlockNumber": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetUncleCountByBlockNumberParams)
				res, err = ethService.EthGetUncleCountByBlockNumber(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetUncleCountByBlockNumberParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getProof": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthGetProofParams)
				res, err = ethService.EthGetProof(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthGetProofParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_getWork": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthGetWork()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_hashrate": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthHashrate()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_mining": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthMining()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_newBlockFilter": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthNewBlockFilter()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_newFilter": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthNewFilterParams)
				res, err = ethService.EthNewFilter(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthNewFilterParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_newPendingTransactionFilter": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthNewPendingTransactionFilter()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_pendingTransactions": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthPendingTransactions()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_protocolVersion": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthProtocolVersion()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_sign": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthSignParams)
				res, err = ethService.EthSign(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthSignParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_accounts": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthAccounts()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_sendTransaction": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthSendTransactionParams)
				res, err = ethService.EthSendTransaction(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthSendTransactionParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_sendRawTransaction": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthSendRawTransactionParams)
				res, err = ethService.EthSendRawTransaction(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthSendRawTransactionParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_submitHashrate": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthSubmitHashrateParams)
				res, err = ethService.EthSubmitHashrate(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthSubmitHashrateParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_submitWork": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthSubmitWorkParams)
				res, err = ethService.EthSubmitWork(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthSubmitWorkParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
		"eth_syncing": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				res, err = ethService.EthSyncing()
				return
			},
			Decode: nopDecoder,
			Encode: jsonEncoder,
		},
		"eth_uninstallFilter": jsonrpc.EndpointCodec{
			Endpoint: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				params := req.(*web3.EthUninstallFilterParams)
				res, err = ethService.EthUninstallFilter(params)
				return
			},
			Decode: func(ctx context.Context, msg json.RawMessage) (interface{}, error) {
				req := new(web3.EthUninstallFilterParams)
				if err := web3.ParamsToStruct(msg, req); err != nil {
					return nil, err
				}
				return req, nil
			},
			Encode: jsonEncoder,
		},
	}
}

func nopDecoder(context.Context, json.RawMessage) (interface{}, error) { return struct{}{}, nil }
func jsonEncoder(ctx context.Context, result interface{}) (json.RawMessage, error) {
	b, err := json.Marshal(web3.StructToResult(result))
	if err != nil {
		return nil, err
	}

	return b, nil
}
