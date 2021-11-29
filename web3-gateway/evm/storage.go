package evm

import (
	"encoding/hex"
	"fmt"

	burrow "github.com/hyperledger/burrow/acm"
	"github.com/hyperledger/burrow/acm/acmstate"
	"github.com/hyperledger/burrow/binary"
	"github.com/hyperledger/burrow/crypto"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
)

var _ acmstate.ReaderWriter = (*storage)(nil)

type storage struct {
	dbClient db.DBClient
}

func newStorage(dbClient db.DBClient) *storage {
	return &storage{dbClient: dbClient}
}

func (i storage) GetAccount(address crypto.Address) (*burrow.Account, error) {
	accData, err := i.dbClient.GetBurrowAccountDataByAddress(address.String())
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

func (i storage) UpdateAccount(*burrow.Account) error {
	return nil
}

func (i storage) GetStorage(address crypto.Address, key binary.Word256) (value []byte, err error) {
	kv, err := i.dbClient.GetBurrowAccountKeyValueByAddressAndKey(address.String(), key.String())
	if err != nil {
		return nil, err
	}

	if len(kv.Value) > 0 {
		value, err = hex.DecodeString(kv.Value)
		if err != nil {
			return nil, err
		}
	}

	return
}

func (i storage) RemoveAccount(crypto.Address) error {
	return nil
}

func (i storage) SetStorage(crypto.Address, binary.Word256, []byte) error {
	return nil
}
