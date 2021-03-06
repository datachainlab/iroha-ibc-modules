package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	ibcclient "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
)

type ContractState interface {
	Header() *gethtypes.Header
	ETHProof() *ETHProof
}

func (cl Client) GetContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (ContractState, error) {
	switch cl.clientType {
	case ibcclient.MockClient:
		return cl.GetMockContractState(ctx, address, storageKeys, bn)
	default:
		panic(fmt.Sprintf("unknown client type '%v'", cl.clientType))
	}
}

func (cl Client) GetMockContractState(ctx context.Context, address common.Address, storageKeys [][]byte, bn *big.Int) (ContractState, error) {
	block, err := cl.BlockByNumber(ctx, bn)
	if err != nil {
		return nil, err
	}
	// this is dummy
	proof := &ETHProof{
		StorageProofRLP: make([][]byte, len(storageKeys)),
	}
	return ETHContractState{header: block.Header(), ethProof: proof}, nil
}

type ETHContractState struct {
	header   *gethtypes.Header
	ethProof *ETHProof
}

func (cs ETHContractState) Header() *gethtypes.Header {
	return cs.header
}

func (cs ETHContractState) ETHProof() *ETHProof {
	return cs.ethProof
}
