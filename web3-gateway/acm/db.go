package acm

import (
	"github.com/hashicorp/go-memdb"
)

type DB interface {
	Add(account *Account) error
	All() ([]*Account, error)
	GetByIrohaAccountID(accountID string) (*Account, error)
	GetByIrohaAddress(address string) (*Account, error)
}

var _ DB = (*MemDB)(nil)

const (
	MemDBAccountTable               = "account"
	MemDBAccountIndexId             = "id"
	MemDBAccountIndexIrohaAccountId = "iroha_account_id"
	MemDBAccountIndexIrohaAddress   = "iroha_address"
)

type MemDB struct {
	db *memdb.MemDB
}

func (m MemDB) Add(account *Account) error {
	txn := m.db.Txn(true)
	if err := txn.Insert(MemDBAccountTable, account); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (m MemDB) All() ([]*Account, error) {
	// Create read-only transaction
	txn := m.db.Txn(false)
	it, err := txn.Get(MemDBAccountTable, MemDBAccountIndexId)
	if err != nil {
		return nil, err
	}

	var accounts []*Account

	for obj := it.Next(); obj != nil; obj = it.Next() {
		acc := obj.(*Account)
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (m MemDB) GetByIrohaAccountID(accountID string) (*Account, error) {
	txn := m.db.Txn(false)
	raw, err := txn.First(MemDBAccountTable, MemDBAccountIndexIrohaAccountId, accountID)
	if err != nil {
		return nil, err
	} else if raw == nil {
		return nil, ErrNotFound
	}

	return raw.(*Account), nil
}

func (m MemDB) GetByIrohaAddress(address string) (*Account, error) {
	txn := m.db.Txn(false)
	raw, err := txn.First(MemDBAccountTable, MemDBAccountIndexIrohaAddress, address)
	if err != nil {
		return nil, err
	} else if raw == nil {
		return nil, ErrNotFound
	}

	return raw.(*Account), nil
}

func NewMemDB() (DB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			MemDBAccountTable: {
				Name: MemDBAccountTable,
				Indexes: map[string]*memdb.IndexSchema{
					MemDBAccountIndexId: {
						Name:    MemDBAccountIndexId,
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"},
					},
					MemDBAccountIndexIrohaAccountId: {
						Name:    MemDBAccountIndexIrohaAccountId,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "IrohaAccountID"},
					},
					MemDBAccountIndexIrohaAddress: {
						Name:    MemDBAccountIndexIrohaAddress,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "IrohaAddress"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &MemDB{
		db: db,
	}, nil
}
