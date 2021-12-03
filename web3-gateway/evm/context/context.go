package context

import (
	"fmt"
	"sync"

	"github.com/hyperledger/burrow/execution/engine"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

type CallContext struct {
	Execer db.DBExecer
	Caller string
}

var contextMap sync.Map

func contextKey(params engine.CallParams) string {
	return fmt.Sprintf("%s:%s", params.Caller, params.Callee)
}

func StoreCallContext(params engine.CallParams, ctx *CallContext) {
	contextMap.Store(contextKey(params), ctx)
}

func LoadCallContext(params engine.CallParams) (*CallContext, error) {
	key := contextKey(params)
	v, ok := contextMap.Load(key)
	if !ok {
		return nil, fmt.Errorf("not found CallContext key:%s", key)
	}
	return v.(*CallContext), nil
}

func DeleteCallContext(params engine.CallParams) {
	contextMap.Delete(contextKey(params))
}
