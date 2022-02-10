package types

import (
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type Transaction struct {
	*gethtypes.Transaction
	ID common.Hash
}
