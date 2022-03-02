package evm

import (
	"encoding/hex"
	"fmt"

	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/crypto"
	"github.com/hyperledger/burrow/execution/errors"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

var _ acmstate.ReaderWriter = (*storage)(nil)

type storage struct {
	dbExecer db.DBExecer
}

func newStorage(dbExecer db.DBExecer) *storage {
	return &storage{
		dbExecer: dbExecer,
	}
}

func (i storage) GetAccount(address crypto.Address) (*burrow.Account, error) {
	accData, err := i.dbExecer.GetBurrowAccountDataByAddress(address.String())
	if err != nil {
		return nil, err
	}
	if accData == nil {
		return nil, fmt.Errorf(
			"account does not exists at account %s",
			address,
		)
	}

	bz, err := hex.DecodeString(accData.Data)
	if err != nil {
		return nil, err
	}

	account := &burrow.Account{}
	if err = account.Unmarshal(bz); err != nil {
		return nil, err
	}

	// Unmarshalling of account data replaces account.EVMCode == nil with an empty slice []byte{}
	// Hence this workaround to revert that and make native.InitCode work
	if account.EVMCode != nil && len(account.EVMCode) == 0 {
		account.EVMCode = nil
	}
	if account.WASMCode != nil && len(account.WASMCode) == 0 {
		account.WASMCode = nil
	}

	return account, nil
}

func (i storage) UpdateAccount(account *burrow.Account) error {
	if account == nil {
		return errors.Errorf(errors.Codes.IllegalWrite, "UpdateAccount passed nil account")
	}

	marshalledData, err := account.Marshal()
	if err != nil {
		return err
	}

	data := hex.EncodeToString(marshalledData)

	return i.dbExecer.UpsertBurrowAccountDataByAddress(account.Address.String(), data)
}

func (i storage) RemoveAccount(address crypto.Address) error {
	return i.dbExecer.DeleteBurrowAccountKeyValueByAddress(address.String())
}

func (i storage) GetStorage(address crypto.Address, key binary.Word256) (value []byte, err error) {
	kv, err := i.dbExecer.GetBurrowAccountKeyValueByAddressAndKey(address.String(), key.String())
	if err != nil {
		return nil, err
	}

	if kv != nil && len(kv.Value) > 0 {
		value, err = hex.DecodeString(kv.Value)
		if err != nil {
			return nil, err
		}
	}

	return
}

func (i storage) SetStorage(address crypto.Address, key binary.Word256, value []byte) error {
	return i.dbExecer.UpsertBurrowAccountKeyValue(
		address.String(),
		hex.EncodeToString(key.Bytes()),
		hex.EncodeToString(value),
	)
}
