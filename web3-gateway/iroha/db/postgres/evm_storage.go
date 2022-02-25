package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/util"
)

func (c *postgresExecer) GetBurrowAccountDataByAddress(address string) (*entity.BurrowAccountData, error) {
	address = strings.ToLower(util.RemoveHexPrefix(address))

	var account entity.BurrowAccountData

	query := "SELECT address, data FROM burrow_account_data WHERE address=$1"

	if err := c.execer.Get(&account, query, address); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (c *postgresExecer) UpsertBurrowAccountDataByAddress(address string, data string) error {
	query := `
insert into burrow_account_data (address, data) 
values (lower(:address), :data) 
on conflict (address) do update set data = excluded.data 
returning 1
`

	if _, err := c.execer.NamedExec(query, map[string]interface{}{
		"address": address,
		"data":    data,
	}); err != nil {
		return err
	}

	return nil
}

func (c *postgresExecer) GetBurrowAccountKeyValueByAddressAndKey(address, key string) (*entity.BurrowAccountKeyValue, error) {
	address = strings.ToLower(util.RemoveHexPrefix(address))
	key = strings.ToLower(util.RemoveHexPrefix(key))

	var kv entity.BurrowAccountKeyValue

	query := "SELECT address, key, value FROM burrow_account_key_value WHERE address=$1 AND key=$2"

	if err := c.execer.Get(&kv, query, address, key); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &kv, nil
}

func (c *postgresExecer) DeleteBurrowAccountKeyValueByAddress(address string) error {
	query := `
delete from burrow_account_key_value
where address = lower(:address);

delete from burrow_account_data
where address = lower(:address)
returning 1
`
	if res, err := c.execer.NamedExec(query, map[string]interface{}{"address": address}); err != nil {
		return err
	} else if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected > 0 {
		return fmt.Errorf("account deletion failed")
	}

	return nil
}

func (c *postgresExecer) UpsertBurrowAccountKeyValue(address string, key string, value string) error {
	query := `
insert into burrow_account_key_value (address, key, value)
values (lower(:address), lower(:key), :value)
on conflict (address, key) do update set value = excluded.value
returning 1
`
	if res, err := c.execer.NamedExec(query, map[string]interface{}{"address": address, "key": key, "value": value}); err != nil {
		return err
	} else if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected > 0 {
		return fmt.Errorf("account deletion failed")
	}

	return nil
}
