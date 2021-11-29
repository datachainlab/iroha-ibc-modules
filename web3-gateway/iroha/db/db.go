package db

import (
	x "github.com/hyperledger/burrow/encoding/hex"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

type DBClient interface {
	GetLatestHeight() (uint64, error)
	GetBurrowAccountDataByAddress(address string) (*entity.BurrowAccountData, error)
	GetBurrowAccountKeyValueByAddressAndKey(address, key string) (*entity.BurrowAccountKeyValue, error)
	GetEngineTransaction(txHash string) (*entity.EngineTransaction, error)
	GetEngineReceipt(txHash string) (*entity.EngineReceipt, error)
	GeEngineReceiptLogsByTxHash(txHash string) ([]*entity.EngineReceiptLog, error)
}

type TxReceiptLogFilter struct {
	FromBlock uint64
	ToBlock   uint64
	Address   string
	Topics    []string
}

type LogFilterOption func(*TxReceiptLogFilter)

func FromBlockOption(n uint64) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		filter.FromBlock = n
	}
}

func ToBlockOption(n uint64) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		filter.ToBlock = n
	}
}

func AddressOption(addr string) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		filter.Address = util.ToIrohaHexString(addr)
	}
}

func TopicsOption(ts ...string) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		for _, topic := range ts {
			filter.Topics = append(filter.Topics, x.RemovePrefix(topic))
		}
	}
}
