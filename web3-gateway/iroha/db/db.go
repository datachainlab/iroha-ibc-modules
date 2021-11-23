package db

import (
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

type DBClient interface {
	GetLatestHeight() (uint64, error)
	GetBurrowAccountDataByAddress(address string) (*entity.BurrowAccountData, error)
	GetEngineTransaction(txHash string) (*entity.EngineTransaction, error)
	GetEngineReceipt(txHash string) (*entity.EngineReceipt, error)
}
