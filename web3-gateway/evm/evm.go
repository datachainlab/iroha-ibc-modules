package evm

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/execution/engine"
	"github.com/hyperledger/burrow/execution/errors"
	"github.com/hyperledger/burrow/execution/evm"
	"github.com/hyperledger/burrow/execution/native"
	"github.com/hyperledger/burrow/logging"
	"github.com/hyperledger/burrow/logging/structure"
	"github.com/hyperledger/burrow/permission"

	evmCtx "github.com/datachainlab/iroha-ibc-modules/web3-gateway/evm/context"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

var reqLock sync.Mutex

var (
	defaultValue = big.NewInt(0)
	defaultGas   = big.NewInt(10000000)
)

func CallSim(
	dbTransactor db.DBTransactor,
	logger *logging.Logger,
	callerAccountID string,
	caller, callee string,
	input []byte,
) ([]byte, error) {
	caller = util.RemoveHexPrefix(caller)
	callee = util.RemoveHexPrefix(callee)

	if isNative(callee) {
		return nil, fmt.Errorf(
			"the callee address %s is reserved for a native contract and cannot be called directly",
			callee,
		)
	}

	callerAddress, err := crypto.AddressFromHexString(caller)
	if err != nil {
		return nil, err
	}

	var calleeAccount *burrow.Account

	if err = dbTransactor.Exec(context.Background(), callerAccountID, func(execer db.DBExecer) (err error) {
		calleeAccount, err = getBurrowAccount(execer, callee)
		return
	}); err != nil {
		return nil, err
	}

	calleeAddress := calleeAccount.Address

	callParams := engine.CallParams{
		Caller: callerAddress,
		Callee: calleeAddress,
		Input:  input,
		Value:  *defaultValue,
		Gas:    defaultGas,
	}

	reqLock.Lock()
	defer reqLock.Unlock()

	var vmResult []byte
	var vmErr error

	if err = dbTransactor.ExecWithTxBoundary(context.Background(), callerAccountID, func(execer db.DBExecer) error {
		evmCtx.StoreCallContext(callParams, &evmCtx.CallContext{Caller: callerAccountID, Execer: execer})
		defer evmCtx.DeleteCallContext(callParams)

		vmResult, vmErr = execute(execer, logger, callParams, calleeAccount.EVMCode)

		_ = logger.TraceMsg("Successful execution")
		_ = logger.TraceMsg("VM Call complete",
			"caller", callerAddress,
			"callee", calleeAddress,
			"return", vmResult,
			structure.ErrorKey, vmErr)

		return nil
	}); err != nil {
		return nil, err
	}

	return vmResult, vmErr

}

func createNatives() (*native.Natives, error) {
	ns, err := native.Merge(serviceContract, native.Permissions, native.Precompiles)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func getBurrowAccount(execer db.DBExecer, address string) (*burrow.Account, error) {
	calleeAccData, err := execer.GetBurrowAccountDataByAddress(address)
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

func execute(
	dbExecer db.DBExecer,
	logger *logging.Logger,
	callParams engine.CallParams,
	evmCode []byte,
) ([]byte, error) {
	natives, err := createNatives()
	if err != nil {
		return nil, err
	}

	logger = logger.WithScope("EVM")
	vm := evm.New(engine.Options{
		Natives:      natives,
		Logger:       logger,
		DebugOpcodes: false,
	})

	state := newStorage(dbExecer)
	bc := newBlockchain()
	es := newEventSink(bc.blockHeight)

	if caller, err := state.GetAccount(callParams.Caller); err != nil {
		return nil, err
	} else {
		if err = state.UpdateAccount(&burrow.Account{
			Address:     caller.Address,
			Permissions: permission.DefaultAccountPermissions,
		}); err != nil {
			return nil, err
		}
	}

	ret, err := vm.Execute(state, bc, es, callParams, evmCode)
	if err != nil {
		logger.InfoMsg("Error on EVM execution", structure.ErrorKey, err)

		err = errors.AsException(
			errors.Wrapf(err, "call error: %v\nEVM call trace: %s",
				err, es.CallTrace(),
			),
		)
	}

	return ret, err
}
