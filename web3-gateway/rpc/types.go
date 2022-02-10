package rpc

import (
	"fmt"

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

type Filter struct {
	// The hex representation of the block's height
	FromBlock string `json:"fromBlock"`
	// The hex representation of the block's height
	ToBlock string `json:"toBlock"`
	// address is a string or an array of strings
	Address_ interface{} `json:"address"`
	// Array of 32 Bytes DATA topics. Topics are order-dependent. Each topic can also be an array of DATA with 'or' options
	Topics [][]string `json:"topics"`
}

type EthGetLogsParams struct {
	Filter
}

func (f *Filter) Address() ([]string, error) {
	switch addr := f.Address_.(type) {
	case string:
		return []string{addr}, nil
	case []interface{}:
		a := make([]string, len(addr))
		for i, v := range addr {
			if v, ok := v.(string); ok {
				a[i] = v
			} else {
				return nil, fmt.Errorf("unexpected element type of EthGetLogsParams.address: %T", v)
			}
		}
		return a, nil
	case nil:
		return []string{}, nil
	default:
		return nil, fmt.Errorf("unexpected type of EthGetLogsParams.address: %T", addr)
	}
}

type EthGetLogsResult struct {
	Logs []Logs `json:"logs"`
}

type Logs struct {
	web3.Logs
	Topics []string `json:"topics"`
}
