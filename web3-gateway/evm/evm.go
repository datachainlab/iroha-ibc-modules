package evm

import (
	"encoding/hex"
	"fmt"

	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/crypto"
	x "github.com/hyperledger/burrow/encoding/hex"
	"github.com/hyperledger/burrow/execution/engine"
	"github.com/hyperledger/burrow/execution/errors"
	"github.com/hyperledger/burrow/execution/evm"
	"github.com/hyperledger/burrow/execution/native"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/logging/structure"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

func CallSim(
	dbClient db.DBClient,
	logger *logging.Logger,
	caller, callee string,
	txHash string,
	input []byte,
) ([]byte, error) {
	caller = x.RemovePrefix(caller)
	callee = x.RemovePrefix(callee)

	if isNative(callee) {
		return nil, fmt.Errorf(
			"the callee address %s is reserved for a native contract and cannot be called directly",
			callee,
		)
	}

	natives, err := createNatives()
	if err != nil {
		return nil, err
	}

	logger = logger.WithScope("EVM")
	vm := evm.New(evm.Options{
		Natives: natives,
		Logger:  logger.With(structure.TxHashKey, txHash),
	})

	state := acmstate.NewCache(
		newStorage(dbClient),
		acmstate.Named("TxCache"),
		acmstate.ReadOnly,
	)

	blockchain := newBlockchain()

	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}
	eventSink := newEventSink(txHashBytes, blockchain.blockHeight)

	callerAddress, err := crypto.AddressFromHexString(caller)
	if err != nil {
		return nil, err
	}

	calleeAccount, err := getBurrowAccount(dbClient, callee)
	if err != nil {
		return nil, err
	}

	calleeAddress := calleeAccount.Address

	var gasLimit uint64 = 10000000
	params := engine.CallParams{
		Caller: callerAddress,
		Callee: calleeAddress,
		Input:  input,
		Value:  0,
		Gas:    &gasLimit,
	}

	ret, err := vm.Execute(state, blockchain, eventSink, params, calleeAccount.EVMCode)
	if err != nil {
		logger.InfoMsg("Error on EVM execution",
			structure.ErrorKey, err)
		err = errors.AsException(
			errors.Wrapf(err, "call error: %v\nEVM call trace: %s",
				err, eventSink.CallTrace(),
			),
		)
	}

	logger.TraceMsg("Successful execution")
	logger.TraceMsg("VM Call complete",
		"caller", callerAddress,
		"callee", calleeAddress,
		"return", ret,
		structure.ErrorKey, err)

	return ret, err

}

func createNatives() (*native.Natives, error) {
	ns, err := native.Merge(serviceContract, native.Permissions, native.Precompiles)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func getBurrowAccount(dbClient db.DBClient, address string) (*burrow.Account, error) {
	calleeAccData, err := dbClient.GetBurrowAccountDataByAddress(address)
	if err != nil {
		return nil, fmt.Errorf(
			"error getting account at address %s: %s",
			address, err.Error(),
		)
	}
	if calleeAccData == nil {
		return nil, fmt.Errorf(
			"contract account does not exists at account %s",
			address,
		)
	}

	bz, err := hex.DecodeString(calleeAccData.Data)
	if err != nil {
		return nil, err
	}

	calleeAccount := &burrow.Account{}
	if err = calleeAccount.Unmarshal(bz); err != nil {
		return nil, err
	}

	return calleeAccount, nil
}
