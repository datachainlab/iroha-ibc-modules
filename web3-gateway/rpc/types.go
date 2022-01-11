package rpc

import (
	"github.com/hyperledger/burrow/rpc/web3"
)

/*
 * The definition of Log Topics in burrow is different from the spec in Ethereum JSON-RPC.
 * https://github.com/hyperledger/burrow/blob/v0.29.7/rpc/web3/types.go#L446
 * https://eth.wiki/json-rpc/API#returns-43
 * so create a Log struct with Topics set in string array.
 */
type EthGetTransactionReceiptResult struct {
	// returns either a receipt or null
	Receipt
}

type Receipt struct {
	web3.Receipt
	Logs []Logs `json:"logs"`
}

type EthGetLogsResult struct {
	Logs []Logs `json:"logs"`
}

type Logs struct {
	web3.Logs
	Topics []string `json:"topics"`
}
