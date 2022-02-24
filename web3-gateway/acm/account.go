package acm

import (
	"errors"
	"fmt"

	"github.com/hyperledger/burrow/crypto"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

var (
	ErrNotFound = errors.New("account not found")
)

type AccountState struct {
	db DB
}

func NewAccountState(db DB) *AccountState {
	return &AccountState{
		db: db,
	}
}

func (s *AccountState) Add(irohaAccountID string, idx uint64) error {
	acc, err := NewAccount(irohaAccountID, idx)
	if err != nil {
		return err
	}

	return s.db.Add(acc)
}

func (s *AccountState) GetAll() (addresses []*Account, err error) {
	accounts, err := s.db.All()
	if len(accounts) == 0 {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *AccountState) GetByIrohaAccountID(accountID string) (*Account, error) {
	account, err := s.db.GetByIrohaAccountID(accountID)
	if err != nil {
		return nil, fmt.Errorf("%w(%s)", err, accountID)
	}

	return account, nil
}

func (s *AccountState) GetByIrohaAddress(address string) (*Account, error) {
	account, err := s.db.GetByIrohaAddress(util.ToIrohaHexString(address))
	if err != nil {
		return nil, fmt.Errorf("%w(%s)", err, address)
	}

	return account, nil
}

type Account struct {
	Id             uint64
	IrohaAccountID string
	IrohaAddress   string
}

func NewAccount(irohaAccountID string, index uint64) (*Account, error) {
	irohaAddressHex := util.IrohaAccountIDToAddressHex(irohaAccountID)

	irohaAddress, err := crypto.AddressFromHexString(irohaAddressHex)
	if err != nil {
		return nil, err
	}

	return &Account{
		Id:             index,
		IrohaAccountID: irohaAccountID,
		IrohaAddress:   util.ToIrohaHexString(irohaAddress.String()),
	}, nil
}

func (a *Account) GetIrohaAccountID() string {
	return a.IrohaAccountID
}

func (a *Account) GetIrohaAddress() string {
	return a.IrohaAddress
}
