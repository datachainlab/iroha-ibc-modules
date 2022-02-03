package db

import (
	"context"
	"fmt"
	"os"

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
	Addresses []string
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

func AddressesOption(as []string) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		filter.Addresses = make([]string, len(as))
		for i, a := range as {
			filter.Addresses[i] = util.ToIrohaHexString(a)
		}
	}
}

func TopicsOption(topics [][]string) LogFilterOption {
	return func(filter *TxReceiptLogFilter) {
		filter.Topics = make([]string, len(topics))
		for i, topic := range topics {
			switch len(topic) {
			case 0:
				// skip
			case 1:
				filter.Topics[i] = x.RemovePrefix(topic[0])
			default:
				fmt.Fprintf(os.Stderr, "up to 1 topic is acceptable for each position, but %d topics are specified", len(topic))
			}
		}
	}
}
