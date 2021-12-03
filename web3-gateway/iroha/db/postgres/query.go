package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

func (c *postgresExecer) GetAccountAssets(accountID string) ([]entity.AccountAsset, error) {
	query := fmt.Sprintf(`
with %s,
all_data as (
  select row_number() over () rn, *
  from (
	  select *
	  from account_has_asset
	  where account_id = :account_id
	  order by asset_id
  ) t
),
total_number as (
  select rn total_number
  from all_data
  order by rn desc
  limit 1
),
page_start as (
  select rn
  from all_data
  where coalesce(asset_id = :first_asset_id, true)
  limit 1
),
page_data as (
  select * from all_data, page_start, total_number
  where
	  all_data.rn >= page_start.rn and
	  coalesce( -- TODO remove after pagination is mandatory IR-516
		  all_data.rn < page_start.rn + :page_size,
		  true
	  )
)
select account_id, asset_id, amount, total_number, perm
  from
	  page_data
	  right join has_perms on true
`, hasQueryPermissionTarget(
		c.caller,
		accountID,
		pb.RolePermission_can_get_my_acc_ast,
		pb.RolePermission_can_get_all_acc_ast,
		pb.RolePermission_can_get_domain_acc_ast,
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"account_id":     accountID,
		"first_asset_id": nil,
		"page_size":      nil,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var assets []entity.AccountAsset
	if err := sqlx.StructScan(rows, &assets); err != nil {
		return nil, err
	}

	return assets, nil
}

func (c *postgresExecer) GetAccountDetail() (string, error) {
	query := fmt.Sprintf(`
with %s,
detail AS (
	with filtered_plain_data as (
		select row_number() over () rn, *
		from (
		  select
			  data_by_writer.key writer,
			  plain_data.key as key,
			  plain_data.value as value
		  from
			  jsonb_each((
				  select data
				  from account
				  where account_id = :account_id
			  )) data_by_writer,
		  jsonb_each(data_by_writer.value) plain_data
		  where
			  coalesce(data_by_writer.key = :writer, true) and
			  coalesce(plain_data.key = :key, true)
		  order by data_by_writer.key asc, plain_data.key asc
		) t
	),
	page_limits as (
		select start.rn as start, start.rn + :page_size as end
		  from (
			  select rn
			  from filtered_plain_data
			  where
				  coalesce(writer = :first_record_writer, true) and
				  coalesce(key = :first_record_key, true)
			  limit 1
		  ) start
	),
	total_number as (select count(1) total_number from filtered_plain_data),
	next_record as (
		select writer, key
		from
		  filtered_plain_data,
		  page_limits
		where rn = page_limits.end
	),
	page as (
		select json_object_agg(writer, data_by_writer) json
		from (
		  select writer, json_object_agg(key, value) data_by_writer
		  from
			  filtered_plain_data,
			  page_limits
		  where
			  rn >= page_limits.start and
			  coalesce(rn < page_limits.end, true)
		  group by writer
		) t
	),
	target_account_exists as (
		select count(1) val
		from account
		where account_id = :account_id
	)
	select
		page.json json,
		total_number,
		next_record.writer next_writer,
		next_record.key next_key,
		target_account_exists.val target_account_exists
	from
		page
		left join total_number on true
		left join next_record on true
		right join target_account_exists on true
)
select detail.*, perm from detail
right join has_perms on true
`, hasQueryPermissionTarget(
		c.caller,
		c.caller,
		pb.RolePermission_can_get_my_acc_detail,
		pb.RolePermission_can_get_all_acc_detail,
		pb.RolePermission_can_get_domain_acc_detail,
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"account_id":          c.caller,
		"writer":              nil,
		"key":                 nil,
		"first_record_writer": nil,
		"first_record_key":    nil,
		"page_size":           nil,
	})
	if !rows.Next() {
		return "", sql.ErrNoRows
	} else if err != nil {
		return "", err
	}

	var detail entity.AccountDetail
	if err := rows.StructScan(&detail); err != nil {
		return "", err
	}

	return detail.Json, nil
}

func (c *postgresExecer) GetAccount(accountID string) (*entity.Account, error) {
	query := fmt.Sprintf(`
WITH %s,
t AS (
	SELECT a.account_id, a.domain_id, a.quorum, a.data, ARRAY_AGG(ar.role_id) AS roles
	FROM account AS a, account_has_roles AS ar
	WHERE a.account_id = :target_account_id
	AND ar.account_id = a.account_id
	GROUP BY a.account_id
)
SELECT account_id, domain_id, quorum, data, roles, perm
FROM t RIGHT OUTER JOIN has_perms AS p ON TRUE
`, hasQueryPermissionTarget(
		c.caller,
		accountID,
		pb.RolePermission_can_get_my_acc_ast,
		pb.RolePermission_can_get_all_acc_ast,
		pb.RolePermission_can_get_domain_acc_ast,
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"target_account_id": accountID,
	})
	if !rows.Next() {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}

	var account entity.Account
	if err := rows.StructScan(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *postgresExecer) GetSignatories(accountID string) ([]string, error) {
	query := fmt.Sprintf(`
with %s,
t AS (
	  SELECT public_key FROM account_has_signatory
	  WHERE account_id = :account_id
)
SELECT public_key, perm FROM t
RIGHT OUTER JOIN has_perms ON TRUE
`, hasQueryPermissionTarget(
		c.caller,
		accountID,
		pb.RolePermission_can_get_my_signatories,
		pb.RolePermission_can_get_all_signatories,
		pb.RolePermission_can_get_domain_signatories,
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"account_id": accountID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var signatories []entity.Signatory
	if err := sqlx.StructScan(rows, &signatories); err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(signatories))
	for _, v := range signatories {
		ret = append(ret, v.PublicKey)
	}

	return ret, nil
}

func (c *postgresExecer) GetAssetInfo(asset string) (*entity.AssetInfo, error) {
	query := fmt.Sprintf(`
WITH has_perms AS (%s),
perms AS (SELECT domain_id, precision FROM asset
		WHERE asset_id = :asset_id)
SELECT domain_id, precision, perm FROM perms
RIGHT OUTER JOIN has_perms ON TRUE
`, getAccountRolePermissionCheckSql(
		pb.RolePermission_can_read_assets,
		":role_account_id",
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"role_account_id": c.caller,
		"asset_id":        asset,
	})
	if !rows.Next() {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}

	var assetInfo entity.AssetInfo
	if err := rows.StructScan(&assetInfo); err != nil {
		return nil, err
	}

	return &assetInfo, nil
}

func (c *postgresExecer) GetPeers() ([]entity.Peer, error) {
	query := fmt.Sprintf(`
WITH has_perms AS (%s)
SELECT public_key, address, tls_certificate, perm FROM peer
RIGHT OUTER JOIN has_perms ON TRUE
`, getAccountRolePermissionCheckSql(
		pb.RolePermission_can_get_peers,
		":role_account_id",
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"role_account_id": c.caller,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var peers []entity.Peer
	if err := sqlx.StructScan(rows, &peers); err != nil {
		return nil, err
	}

	return peers, nil
}

func (c *postgresExecer) GetRoles() ([]string, error) {
	query := fmt.Sprintf(`
WITH has_perms AS (%s)
SELECT role_id, perm FROM role
RIGHT OUTER JOIN has_perms ON TRUE
`, getAccountRolePermissionCheckSql(
		pb.RolePermission_can_get_roles,
		":role_account_id",
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"role_account_id": c.caller,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var roles []entity.Role
	if err := sqlx.StructScan(rows, &roles); err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(roles))
	for _, v := range roles {
		ret = append(ret, v.RoleID)
	}

	return ret, nil
}

func (c *postgresExecer) GetRolePermissions(role string) (string, error) {
	query := fmt.Sprintf(`
WITH has_perms AS (%s),
perms AS (SELECT permission FROM role_has_permissions
		WHERE role_id = :role_name)
SELECT permission, perm FROM perms
RIGHT OUTER JOIN has_perms ON TRUE
`, getAccountRolePermissionCheckSql(
		pb.RolePermission_can_get_roles,
		":role_account_id",
	))

	rows, err := c.execer.NamedQuery(query, map[string]interface{}{
		"role_account_id": c.caller,
		"role_name":       role,
	})
	if !rows.Next() {
		return "", sql.ErrNoRows
	} else if err != nil {
		return "", err
	}

	var perm entity.RolePermission
	if err := rows.StructScan(&perm); err != nil {
		return "", err
	}

	return perm.Permission, nil
}
