package db

import (
	"context"

	x "github.com/hyperledger/burrow/encoding/hex"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

type DBTransactor interface {
	Close() error
	Exec(ctx context.Context, caller string, f func(execer DBExecer) error) error
	ExecWithTxBoundary(ctx context.Context, caller string, f func(execer DBExecer) error) error
}

type DBExecer interface {
	EVMStorageExecer
	CallEngineExecer
	NativeContractExecer

	GetLatestHeight() (uint64, error)
}

type CallEngineExecer interface {
	GetEngineTransaction(txHash string) (*entity.EngineTransaction, error)
	GetEngineReceipt(txHash string) (*entity.EngineReceipt, error)
	GetEngineReceiptLogsByTxHash(txHash string) ([]*entity.EngineReceiptLog, error)
	GetEngineReceiptLogsByFilters(opts ...LogFilterOption) ([]*entity.EngineReceiptLog, error)
}

type EVMStorageExecer interface {
	GetBurrowAccountDataByAddress(address string) (*entity.BurrowAccountData, error)
	UpsertBurrowAccountDataByAddress(address string, data string) error
	GetBurrowAccountKeyValueByAddressAndKey(address, key string) (*entity.BurrowAccountKeyValue, error)
	DeleteBurrowAccountKeyValueByAddress(address string) error
}

type NativeContractExecer interface {
	GetAccountAssets(accountID string) ([]entity.AccountAsset, error)
	TransferAsset(src, dst string, assetID string, description string, amount string) error
	CreateAccount(name string, domain string, key string) error
	AddAssetQuantity(asset string, amount string) error
	SubtractAssetQuantity(asset string, amount string) error
	SetAccountDetail(accountID string, key string, value string) error
	GetAccountDetail() (string, error)
	SetAccountQuorum(accountID string, quorum string) error
	AddSignatory(accountID string, key string) error
	RemoveSignatory(accountID string, key string) error
	CreateDomain(domain string, role string) error
	GetAccount(accountID string) (*entity.Account, error)
	CreateAsset(name string, domain string, precision string) error
	GetSignatories(accountID string) ([]string, error)
	GetAssetInfo(asset string) (*entity.AssetInfo, error)
	AppendRole(accountID string, role string) error
	DetachRole(accountID string, role string) error
	AddPeer(address string, key string) error
	RemovePeer(key string) error
	GrantPermission(id string, permission string) error
	RevokePermission(id string, permission string) error
	CompareAndSetAccountDetail(accountID string, key string, value string, oldValue string, checkEmpty string) error
	CreateRole(name string, permissions string) error
	GetPeers() ([]entity.Peer, error)
	GetRoles() ([]string, error)
	GetRolePermissions(role string) (string, error)
	// ### not implemented ###
	//GetAccountTransactions(accountID string, meta *TxPaginationMeta) (interface{}, error)
	//GetPendingTransactions(meta *TxPaginationMeta) (interface{}, error)
	//GetAccountAssetTransactions(accountID string, assetID string, t *TxPaginationMeta) (interface{}, error)
	//GetTransactions(hashes []string) ([]entity.Transaction, error)
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
