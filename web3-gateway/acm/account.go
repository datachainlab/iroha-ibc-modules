package acm

import (
	"bytes"
	"encoding/hex"
	"errors"

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

func (s *AccountState) Add(irohaAccountID string, privKey string) error {
	acc, err := NewAccount(irohaAccountID, privKey)
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

func (s *AccountState) GetDefaultAccount() (*Account, error) {
	account, err := s.db.First()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountState) GetByIrohaAccountID(accountID string) (*Account, error) {
	account, err := s.db.GetByIrohaAccountID(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountState) GetByIrohaAddress(address string) (*Account, error) {
	account, err := s.db.GetByIrohaAddress(util.ToIrohaHexString(address))
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountState) GetByEthereumAddress(address string) (*Account, error) {
	account, err := s.db.GetByEthereumAddress(address)
	if err != nil {
		return nil, err
	}

	return account, nil
}

type Account struct {
	EthereumAddress string
	IrohaAccountID  string
	IrohaAddress    string
}

func NewAccount(irohaAccountID string, privKey string) (*Account, error) {
	bz, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}

	irohaAddressHex := util.IrohaAccountIDToAddressHex(irohaAccountID)

	irohaAddress, err := crypto.AddressFromHexString(irohaAddressHex)
	if err != nil {
		return nil, err
	}

	ethPrivKey, err := crypto.GeneratePrivateKey(bytes.NewBuffer(bz), crypto.CurveTypeSecp256k1)
	if err != nil {
		return nil, err
	}

	return &Account{
		EthereumAddress: util.ToEthereumHexString(ethPrivKey.GetPublicKey().GetAddress().String()),
		IrohaAccountID:  irohaAccountID,
		IrohaAddress:    util.ToIrohaHexString(irohaAddress.String()),
	}, nil
}

func (a *Account) GetIrohaAccountID() string {
	return a.IrohaAccountID
}

func (a *Account) GetIrohaAddress() string {
	return a.IrohaAddress
}

func (a *Account) GetEthereumAddress() string {
	return a.EthereumAddress
}
