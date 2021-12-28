package postgres

import (
	"fmt"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/consts"
)

func (c *postgresExecer) TransferAsset(
	src string, dst string,
	assetID string, description string, amount string,
) error {
	query := fmt.Sprintf(`
WITH %s
	new_src_quantity AS
	(
		SELECT coalesce(sum(amount), 0) - cast(:quantity as decimal) as value
		FROM account_has_asset
		   WHERE asset_id = :asset_id AND
		   account_id = :source_account_id
	),
	new_dest_quantity AS
	(
		SELECT coalesce(sum(amount), 0) + cast(:quantity as decimal) as value
		FROM account_has_asset
		   WHERE asset_id = :asset_id AND
		   account_id = :dest_account_id
	),
	checks AS -- error code and check result
	(
		-- source account exists
		SELECT 3 code, count(1) = 1 result
		FROM account
		WHERE account_id = :source_account_id

		-- dest account exists
		UNION
		SELECT 4, count(1) = 1
		FROM account
		WHERE account_id = :dest_account_id

		-- asset exists
		UNION
		SELECT 5, count(1) = 1
		FROM asset
		WHERE asset_id = :asset_id
		   AND precision >= :precision

		-- enough source quantity
		UNION
		SELECT 6, value >= 0
		FROM new_src_quantity

		-- dest quantity overflow
		UNION
		SELECT
			7,
			value < (2::::decimal ^ 256) / (10::::decimal ^ precision)
		FROM new_dest_quantity, asset
		WHERE asset_id = :asset_id

		-- description length
		UNION
		SELECT 8, :description_length <= setting_value::::integer
		FROM setting
		WHERE setting_key = '%s'
	),
	insert_src AS
	(
		UPDATE account_has_asset
		SET amount = value
		FROM new_src_quantity
		WHERE
			account_id = :source_account_id
			AND asset_id = :asset_id
			AND (SELECT bool_and(checks.result) FROM checks) %s
	),
	insert_dest AS
	(
		INSERT INTO account_has_asset(account_id, asset_id, amount)
		(
			SELECT :dest_account_id, :asset_id, value
			FROM new_dest_quantity
			WHERE (SELECT bool_and(checks.result) FROM checks) %s
		)
		ON CONFLICT (account_id, asset_id)
		DO UPDATE SET amount = EXCLUDED.amount
		RETURNING (1)
	)
  SELECT CASE
	  WHEN EXISTS (SELECT * FROM insert_dest LIMIT 1) THEN 0
	  WHEN EXISTS (SELECT * FROM checks WHERE not result and code = 4) THEN 4
	  %s
	  ELSE (SELECT code FROM checks WHERE not result ORDER BY code ASC LIMIT 1)
  END AS result
`, fmt.Sprintf(`
has_role_perm AS (%s),
  has_grantable_perm AS (%s),
  dest_can_receive AS (%s),
  has_perm AS
  (
	  SELECT
		  CASE WHEN (SELECT * FROM dest_can_receive) THEN
			  CASE WHEN NOT (:creator = :source_account_id) THEN
				  CASE WHEN (SELECT * FROM has_grantable_perm)
					  THEN true
				  ELSE false END
			  ELSE
				  CASE WHEN (SELECT * FROM has_role_perm)
					  THEN true
				  ELSE false END
			  END
		  ELSE false END
  ),
`,
		checkAccountRolePermission(pb.RolePermission_can_transfer, ":creator"),
		checkAccountGrantablePermission(pb.GrantablePermission_can_transfer_my_assets, ":creator", ":source_account_id"),
		checkAccountRolePermission(pb.RolePermission_can_receive, ":dest_account_id"),
	),
		consts.MaxDescriptionSizeKey,
		" AND (SELECT * FROM has_perm) ",
		" AND (SELECT * FROM has_perm) ",
		" WHEN NOT (SELECT * FROM has_perm) THEN 2 ",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":            c.caller,
		"source_account_id":  src,
		"dest_account_id":    dst,
		"asset_id":           assetID,
		"quantity":           amount,
		"precision":          getPrecisionFromAmount(amount),
		"description_length": len(description),
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed TransferAsset")
	}

	return nil
}

func (c *postgresExecer) CreateAccount(name string, domain string, key string) error {
	query := fmt.Sprintf(`
WITH get_domain_default_role AS (SELECT default_role FROM domain
								 WHERE domain_id = :domain),
%s
insert_signatory AS
(
	INSERT INTO signatory(public_key)
	(
		SELECT lower(:pubkey)
		WHERE EXISTS (SELECT * FROM get_domain_default_role)
		  %s
	)
	ON CONFLICT (public_key)
	  DO UPDATE SET public_key = excluded.public_key
	RETURNING (1)
),
insert_account AS
(
	INSERT INTO account(account_id, domain_id, quorum, data)
	(
		SELECT :account_id, :domain, 1, '{}'
		WHERE EXISTS (SELECT * FROM insert_signatory)
		  AND EXISTS (SELECT * FROM get_domain_default_role)
	) RETURNING (1)
),
insert_account_signatory AS
(
	INSERT INTO account_has_signatory(account_id, public_key)
	(
		SELECT :account_id, lower(:pubkey) WHERE
		   EXISTS (SELECT * FROM insert_account)
	)
	RETURNING (1)
),
insert_account_role AS
(
	INSERT INTO account_has_roles(account_id, role_id)
	(
		SELECT :account_id, default_role FROM get_domain_default_role
		WHERE EXISTS (SELECT * FROM get_domain_default_role)
		  AND EXISTS (SELECT * FROM insert_account_signatory)
	) RETURNING (1)
)
SELECT CASE
WHEN EXISTS (SELECT * FROM insert_account_role) THEN 0
WHEN NOT EXISTS (SELECT * FROM get_domain_default_role) THEN 3
%s
ELSE 1
END AS result
`,
		fmt.Sprintf(`
domain_role_permissions_bits AS (
                 SELECT COALESCE(bit_or(rhp.permission), '0'::::bit(%[1]v)) AS bits
                 FROM role_has_permissions AS rhp
                 WHERE rhp.role_id = (SELECT * FROM get_domain_default_role)),
           account_permissions AS (
                 SELECT COALESCE(bit_or(rhp.permission), '0'::::bit(%[1]v)) AS perm
                 FROM role_has_permissions AS rhp
                 JOIN account_has_roles AS ar ON ar.role_id = rhp.role_id
                 WHERE ar.account_id = :creator
           ),
           creator_has_enough_permissions AS (
                SELECT ap.perm & dpb.bits = dpb.bits OR has_root_perm.has_rp
                FROM
                    account_permissions AS ap
                  , domain_role_permissions_bits AS dpb
                  , (%[3]v) as has_root_perm

           ),
           has_perm AS (%[2]v),
`,
			consts.RolePermissionEnumLength,
			checkAccountRolePermission(pb.RolePermission_can_create_account, ":creator"),
			checkAccountRolePermission(pb.RolePermission_root, ":creator"),
		),
		"AND (SELECT * FROM has_perm) AND (SELECT * FROM creator_has_enough_permissions)",
		"WHEN NOT (SELECT * FROM has_perm) THEN 2 WHEN NOT (SELECT * FROM creator_has_enough_permissions) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":    c.caller,
		"account_id": name,
		"domain":     domain,
		"pubkey":     key,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed CreateAccount")
	}

	return nil
}

func (c *postgresExecer) AddAssetQuantity(asset string, amount string) error {
	query := fmt.Sprintf(`
WITH %s
	 new_quantity AS
	 (
		 SELECT CAST(:quantity AS decimal) + coalesce(sum(amount), 0) as value
		 FROM account_has_asset
		 WHERE asset_id = :asset_id
			 AND account_id = :creator
	 ),
	 checks AS -- error code and check result
	 (
		 -- account exists
		 SELECT 1 code, count(1) = 1 result
		 FROM account
		 WHERE account_id = :creator
	
		 -- asset exists
		 UNION
		 SELECT 3, count(1) = 1
		 FROM asset
		 WHERE asset_id = :asset_id
			AND precision >= :precision
	
		 -- quantity overflow
		 UNION
		 SELECT
			4,
			value < (2::::decimal ^ 256) / (10::::decimal ^ precision)
		 FROM new_quantity, asset
		 WHERE asset_id = :asset_id
	 ),
	 inserted AS
	 (
		INSERT INTO account_has_asset(account_id, asset_id, amount)
		(
			SELECT :creator, :asset_id, value FROM new_quantity
			WHERE (SELECT bool_and(checks.result) FROM checks) %s
		)
		ON CONFLICT (account_id, asset_id) DO UPDATE
		SET amount = EXCLUDED.amount
		RETURNING (1)
	 )
	SELECT CASE
	  %s
	  WHEN EXISTS (SELECT * FROM inserted LIMIT 1) THEN 0
	  ELSE (SELECT code FROM checks WHERE not result ORDER BY code ASC LIMIT 1)
	END AS result;
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountDomainRoleOrGlobalRolePermission(
				pb.RolePermission_can_add_asset_qty,
				pb.RolePermission_can_add_domain_asset_qty,
				":creator",
				":asset_id",
			),
		),
		"AND (SELECT * from has_perm)",
		"WHEN NOT (SELECT * from has_perm) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":   c.caller,
		"asset_id":  asset,
		"precision": getPrecisionFromAmount(amount),
		"quantity":  amount,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed AddAssetQuantity")
	}

	return nil
}

func (c *postgresExecer) SubtractAssetQuantity(asset string, amount string) error {
	query := fmt.Sprintf(`
WITH %s
	has_account AS (SELECT account_id FROM account
					WHERE account_id = :creator LIMIT 1),
	has_asset AS (SELECT asset_id FROM asset
				  WHERE asset_id = :asset_id
				  AND precision >= :precision LIMIT 1),
	amount AS (SELECT amount FROM account_has_asset
			   WHERE asset_id = :asset_id
			   AND account_id = :creator LIMIT 1),
	new_value AS (SELECT
				   (SELECT
					   CASE WHEN EXISTS
						   (SELECT amount FROM amount LIMIT 1)
						   THEN (SELECT amount FROM amount LIMIT 1)
					   ELSE 0::::decimal
				   END) - CAST(:quantity AS decimal) AS value
			   ),
	inserted AS
	(
	   INSERT INTO account_has_asset(account_id, asset_id, amount)
	   (
		   SELECT :creator, :asset_id, value FROM new_value
		   WHERE EXISTS (SELECT * FROM has_account LIMIT 1) AND
			 EXISTS (SELECT * FROM has_asset LIMIT 1) AND
			 EXISTS (SELECT value FROM new_value WHERE value >= 0 LIMIT 1)
			 %s
	   )
	   ON CONFLICT (account_id, asset_id)
	   DO UPDATE SET amount = EXCLUDED.amount
	   RETURNING (1)
	)
  SELECT CASE
	  WHEN EXISTS (SELECT * FROM inserted LIMIT 1) THEN 0
	  %s
	  WHEN NOT EXISTS (SELECT * FROM has_asset LIMIT 1) THEN 3
	  WHEN NOT EXISTS
		  (SELECT value FROM new_value WHERE value >= 0 LIMIT 1) THEN 4
	  ELSE 1
  END AS result
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountDomainRoleOrGlobalRolePermission(
				pb.RolePermission_can_subtract_asset_qty,
				pb.RolePermission_can_subtract_domain_asset_qty,
				":creator",
				":asset_id",
			),
		),
		"AND (SELECT * from has_perm)",
		"WHEN NOT (SELECT * from has_perm) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":   c.caller,
		"asset_id":  asset,
		"precision": getPrecisionFromAmount(amount),
		"quantity":  amount,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed SubtractAssetQuantity")
	}

	return nil
}

func (c *postgresExecer) SetAccountDetail(accountID string, key string, value string) error {
	query := fmt.Sprintf(`
WITH %s
	inserted AS
	(
		UPDATE account SET data = jsonb_set(
		CASE WHEN data ? :creator THEN data ELSE
		jsonb_set(data, array[:creator], '{}') END,
		array[:creator, :key], CAST(:value AS jsonb)) WHERE account_id=:target %s
		RETURNING (1)
	)
  SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
	WHEN NOT EXISTS
			(SELECT * FROM account WHERE account_id=:target) THEN 3
	%s
	ELSE 1
  END AS result
`,
		fmt.Sprintf(`
has_role_perm AS (%s),
has_grantable_perm AS (%s),
has_perm AS (SELECT CASE
			   WHEN (SELECT * FROM has_grantable_perm) THEN true
			   WHEN (:creator = :target) THEN true
			   WHEN (SELECT * FROM has_role_perm) THEN true
			   ELSE false END
),
`,
			checkAccountRolePermission(pb.RolePermission_can_set_detail, ":creator"),
			checkAccountGrantablePermission(pb.GrantablePermission_can_set_my_account_detail, ":creator", ":target"),
		),
		"AND (SELECT * from has_perm)",
		"WHEN NOT (SELECT * from has_perm) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"key":     key,
		"value":   fmt.Sprintf(`"%s"`, value),
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed SetAccountDetail")
	}

	return nil
}

func (c *postgresExecer) SetAccountQuorum(accountID string, quorum string) error {
	query := fmt.Sprintf(`
WITH %s
	updated AS (
		UPDATE account SET quorum=:quorum
		WHERE account_id=:target
		%s
		RETURNING (1)
	)
SELECT CASE
	WHEN EXISTS (SELECT * FROM updated) THEN 0
	%s
	ELSE 1
END AS result
`,
		fmt.Sprintf(`
get_signatories AS (
	SELECT public_key FROM account_has_signatory
	WHERE account_id = :target
),
check_account_signatories AS (
	SELECT 1 FROM account
	WHERE :quorum <= (SELECT COUNT(*) FROM get_signatories)
	AND account_id = :target
),
has_perm AS (%s),
`,
			checkAccountHasRoleOrGrantablePerm(
				pb.RolePermission_can_set_quorum,
				pb.GrantablePermission_can_set_my_quorum,
				":creator", ":target"),
		),
		`
AND EXISTS 
	(SELECT * FROM get_signatories) 
	AND EXISTS (SELECT * FROM check_account_signatories) 
	AND (SELECT * FROM has_perm)
`,
		`
WHEN NOT (SELECT * FROM has_perm) THEN 2
WHEN NOT EXISTS (SELECT * FROM get_signatories) THEN 4
WHEN NOT EXISTS (SELECT * FROM check_account_signatories) THEN 5
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"quorum":  quorum,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed SetAccountQuorum")
	}

	return nil
}

func (c *postgresExecer) AddSignatory(accountID string, pubkey string) error {
	query := fmt.Sprintf(`
WITH %s
	insert_signatory AS
	(
		INSERT INTO signatory(public_key)
		(SELECT lower(:pubkey) %s)
		ON CONFLICT (public_key)
		  DO UPDATE SET public_key = excluded.public_key
		RETURNING (1)
	),
	insert_account_signatory AS
	(
		INSERT INTO account_has_signatory(account_id, public_key)
		(
			SELECT :target, lower(:pubkey)
			WHERE EXISTS (SELECT * FROM insert_signatory)
		)
		RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM insert_account_signatory) THEN 0
	%s
	ELSE 1
	END AS RESULT;
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountHasRoleOrGrantablePerm(
				pb.RolePermission_can_add_signatory,
				pb.GrantablePermission_can_add_my_signatory,
				":creator", ":target"),
		),
		"WHERE (SELECT * FROM has_perm)",
		"WHEN NOT (SELECT * from has_perm) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"pubkey":  pubkey,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed AddSignatory")
	}

	return nil
}

func (c *postgresExecer) RemoveSignatory(accountID string, pubkey string) error {
	query := fmt.Sprintf(`
WITH %s
	delete_account_signatory AS (DELETE FROM account_has_signatory
		WHERE account_id = :target
		AND public_key = lower(:pubkey)
		%s
		RETURNING (1)),
	delete_signatory AS
	(
		DELETE FROM signatory WHERE public_key = lower(:pubkey) AND
			NOT EXISTS (SELECT 1 FROM account_has_signatory
						WHERE public_key = lower(:pubkey))
			AND NOT EXISTS (SELECT 1 FROM peer
							WHERE public_key = lower(:pubkey))
		RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM delete_account_signatory) THEN
	CASE
		WHEN EXISTS (SELECT * FROM delete_signatory) THEN 0
		WHEN EXISTS (SELECT 1 FROM account_has_signatory
					 WHERE public_key = lower(:pubkey)) THEN 0
		WHEN EXISTS (SELECT 1 FROM peer
					 WHERE public_key = lower(:pubkey)) THEN 0
		ELSE 1
	END
	%s
	ELSE 1
	END AS result
`,
		fmt.Sprintf(`
has_perm AS (%s),
get_account AS (
	SELECT quorum FROM account WHERE account_id = :target LIMIT 1
),
get_signatories AS (
	SELECT public_key FROM account_has_signatory
	WHERE account_id = :target
),
get_signatory AS (
	SELECT * FROM get_signatories
	WHERE public_key = lower(:pubkey)
),
check_account_signatories AS (
	SELECT quorum FROM get_account
	WHERE quorum < (SELECT COUNT(*) FROM get_signatories)
),
`,
			checkAccountHasRoleOrGrantablePerm(
				pb.RolePermission_can_remove_signatory,
				pb.GrantablePermission_can_remove_my_signatory,
				":creator", ":target"),
		),
		`
AND (SELECT * FROM has_perm)
AND EXISTS (SELECT * FROM get_account)
AND EXISTS (SELECT * FROM get_signatories)
AND EXISTS (SELECT * FROM check_account_signatories)
`,
		`
WHEN NOT EXISTS (SELECT * FROM get_account) THEN 3
WHEN NOT (SELECT * FROM has_perm) THEN 2
WHEN NOT EXISTS (SELECT * FROM get_signatory) THEN 4
WHEN NOT EXISTS (SELECT * FROM check_account_signatories) THEN 5
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"pubkey":  pubkey,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed RemoveSignatory")
	}

	return nil
}

func (c *postgresExecer) CreateDomain(domain string, role string) error {
	query := fmt.Sprintf(`
WITH %s
	inserted AS
	(
		INSERT INTO domain(domain_id, default_role)
		(
			SELECT :domain, :default_role
			%s
		) RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
	%s
	ELSE 1
	END AS result
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountRolePermission(pb.RolePermission_can_create_domain, ":creator"),
		),
		`WHERE (SELECT * FROM has_perm)`,
		`WHEN NOT (SELECT * FROM has_perm) THEN 2`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":      c.caller,
		"domain":       domain,
		"default_role": role,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed CreateDomain")
	}

	return nil
}

func (c *postgresExecer) CreateAsset(name string, domain string, precision string) error {
	query := fmt.Sprintf(`
WITH %s
	inserted AS
	(
		INSERT INTO asset(asset_id, domain_id, precision)
		(
			SELECT :asset_id, :domain, :precision
			%s
		) RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
	%s
	ELSE 1
	END AS result
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountRolePermission(pb.RolePermission_can_create_asset, ":creator"),
		),
		`WHERE (SELECT * FROM has_perm)`,
		`WHEN NOT (SELECT * FROM has_perm) THEN 2`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":   c.caller,
		"asset_id":  fmt.Sprintf("%s#%s", name, domain),
		"domain":    domain,
		"precision": precision,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed CreateAsset")
	}

	return nil
}

func (c *postgresExecer) AppendRole(accountID string, role string) error {
	query := fmt.Sprintf(`
WITH %s
	role_exists AS (SELECT * FROM role WHERE role_id = :role),
	inserted AS (
		INSERT INTO account_has_roles(account_id, role_id)
		(
			SELECT :target, :role %s) RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
	WHEN NOT EXISTS (SELECT * FROM role_exists) THEN 4
	%s
	ELSE 1
	END AS result
`,
		fmt.Sprintf(`
has_perm AS (%[1]v),
has_root_perm AS (%[2]v),
role_permissions AS (
	SELECT permission FROM role_has_permissions
	WHERE role_id = :role
),
account_roles AS (
	SELECT role_id FROM account_has_roles WHERE account_id = :creator
),
account_has_role_permissions AS (
	SELECT COALESCE(bit_or(rp.permission), '0'::::bit(%[3]v)) &
		(SELECT * FROM role_permissions) =
		(SELECT * FROM role_permissions)
	FROM role_has_permissions AS rp
	JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	WHERE ar.account_id = :creator
),
`,
			checkAccountRolePermission(pb.RolePermission_can_append_role, ":creator"),
			checkAccountRolePermission(pb.RolePermission_root, ":creator"),
			consts.RolePermissionEnumLength,
		),
		`
WHERE
(SELECT * FROM has_root_perm)
OR (EXISTS (SELECT * FROM account_roles) AND
(SELECT * FROM account_has_role_permissions)
AND (SELECT * FROM has_perm))
`,
		`
WHEN NOT EXISTS (SELECT * FROM account_roles)
	AND NOT (SELECT * FROM has_root_perm) THEN 2
WHEN NOT (SELECT * FROM account_has_role_permissions)
	AND NOT (SELECT * FROM has_root_perm) THEN 2
WHEN NOT (SELECT * FROM has_perm) THEN 2
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"role":    role,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed AppendRole")
	}

	return nil
}

func (c *postgresExecer) DetachRole(accountID string, role string) error {
	query := fmt.Sprintf(`
WITH %s
	deleted AS
	(
	  DELETE FROM account_has_roles
	  WHERE account_id=:target
	  AND role_id=:role
	  %s
	  RETURNING (1)
	)
	SELECT CASE
	WHEN EXISTS (SELECT * FROM deleted) THEN 0
	WHEN NOT EXISTS (SELECT * FROM account
					 WHERE account_id = :target) THEN 3
	WHEN NOT EXISTS (SELECT * FROM role
					 WHERE role_id = :role) THEN 5
	WHEN NOT EXISTS (SELECT * FROM account_has_roles
					 WHERE account_id=:target AND role_id=:role) THEN 4
	%s
	ELSE 1
	END AS result
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountRolePermission(pb.RolePermission_can_detach_role, ":creator"),
		),
		`AND (SELECT * FROM has_perm)`,
		`WHEN NOT (SELECT * FROM has_perm) THEN 2`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"target":  accountID,
		"role":    role,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed DetachRole")
	}

	return nil
}

func (c *postgresExecer) AddPeer(address string, pubkey string) error {
	query := fmt.Sprintf(`
WITH %s
	inserted AS (
		INSERT INTO peer(public_key, address, tls_certificate)
		(
			SELECT lower(:pubkey), :address, :tls_certificate
			%s
		) RETURNING (1)
	)
	SELECT CASE WHEN EXISTS (SELECT * FROM inserted) THEN 0
	  %s
	  ELSE 1 END AS result
`,
		fmt.Sprintf(`has_perm AS (%s),`,
			checkAccountRolePermission(pb.RolePermission_can_add_peer, ":creator"),
		),
		`AND (SELECT * FROM has_perm)`,
		`WHEN NOT (SELECT * FROM has_perm) THEN 2`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":         c.caller,
		"address":         address,
		"pubkey":          pubkey,
		"tls_certificate": nil,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed AddPeer")
	}

	return nil
}

func (c *postgresExecer) RemovePeer(pubkey string) error {
	query := fmt.Sprintf(`
WITH %s
	removed AS (
		DELETE FROM peer WHERE public_key = lower(:pubkey)
		%s
		RETURNING (1)
	)
	SELECT CASE
		WHEN EXISTS (SELECT * FROM removed) THEN 0
		%s
		ELSE 1
	END AS result
`,
		fmt.Sprintf(`
has_perm AS (%s),
get_peer AS (
  SELECT * from peer WHERE public_key = lower(:pubkey) LIMIT 1
),
check_peers AS (
  SELECT 1 WHERE (SELECT COUNT(*) FROM peer) > 1
),
`,
			checkAccountRolePermissionWithAdditional(
				pb.RolePermission_can_add_peer, pb.RolePermission_can_remove_peer, ":creator",
			),
		),
		`
AND (SELECT * FROM has_perm)
AND EXISTS (SELECT * FROM get_peer)
AND EXISTS (SELECT * FROM check_peers)
`,
		`
WHEN NOT EXISTS (SELECT * from get_peer) THEN 3
WHEN NOT EXISTS (SELECT * from check_peers) THEN 4
WHEN NOT (SELECT * from has_perm) THEN 2
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"pubkey":  pubkey,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed RemovePeer")
	}

	return nil
}

func (c *postgresExecer) GrantPermission(accountID string, permission string) error {
	perm := pb.GrantablePermission(pb.GrantablePermission_value[permission])
	requiredPerm := permissionFor(perm)

	query := fmt.Sprintf(`
WITH %s
inserted AS (
	INSERT INTO account_has_grantable_permissions AS
	has_perm(permittee_account_id, account_id, permission)
	(SELECT :target, :creator, :granted_perm %s) ON CONFLICT
	(permittee_account_id, account_id)
	DO UPDATE SET permission=(SELECT has_perm.permission | :granted_perm
	WHERE (has_perm.permission & :granted_perm) <> :granted_perm)
	RETURNING (1)
)
SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
%s
ELSE 1
END AS result
`,
		fmt.Sprintf(`has_perm AS (%s)`,
			checkAccountRolePermission(requiredPerm, ":creator"),
		),
		`
AND (SELECT * FROM has_perm)
AND EXISTS (SELECT * FROM get_peer)
AND EXISTS (SELECT * FROM check_peers)
`,
		`
WHEN NOT EXISTS (SELECT * from get_peer) THEN 3
WHEN NOT EXISTS (SELECT * from check_peers) THEN 4
WHEN NOT (SELECT * from has_perm) THEN 2
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":      c.caller,
		"target":       accountID,
		"granted_perm": grantableRoleToBitString(perm),
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed GrantPermission")
	}

	return nil
}

func (c *postgresExecer) RevokePermission(accountID string, permission string) error {
	perm := pb.GrantablePermission(pb.GrantablePermission_value[permission])

	query := fmt.Sprintf(`
WITH %[2]v
inserted AS (
	UPDATE account_has_grantable_permissions as has_perm
	SET permission=(
	  SELECT has_perm.permission & (~ :revoked_perm::bit(%[1]v))
	  WHERE has_perm.permission & :revoked_perm::bit(%[1]v)
		  = :revoked_perm::bit(%[1]v) AND
	  has_perm.permittee_account_id=:target AND
	  has_perm.account_id=:creator
	)
	WHERE
	permittee_account_id=:target AND
	account_id=:creator %[3]v
  RETURNING (1)
)
SELECT CASE
	WHEN EXISTS (SELECT * FROM inserted) THEN 0
%[4]v
ELSE 1
END AS result
`,
		consts.GrantableRolePermissionEnumLength,
		fmt.Sprintf(`
has_perm AS (
	SELECT
	  (
		  COALESCE(bit_or(permission), '0'::bit(%[1]v))
		  & :revoked_perm::bit(%[1]v)
	  )
	  = :revoked_perm::bit(%[1]v)
	FROM account_has_grantable_permissions
	WHERE account_id = :creator AND
	permittee_account_id = :target),
`, consts.GrantableRolePermissionEnumLength),
		` AND (SELECT * FROM has_perm)`,
		` WHEN NOT (SELECT * FROM has_perm) THEN 2 `,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":      c.caller,
		"target":       accountID,
		"revoked_perm": grantableRoleToBitString(perm),
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed RevokePermission")
	}

	return nil
}

func (c *postgresExecer) CompareAndSetAccountDetail(
	accountID string, key string, value string, oldValue string, checkEmpty string,
) error {
	query := fmt.Sprintf(`
  WITH %s
	old_value AS
	(
		SELECT *
		FROM account
		WHERE
		  account_id = :target
		  AND CASE
			WHEN data ? :creator AND data->:creator ?:key
			  THEN CASE
				WHEN :have_expected_value::::boolean
					THEN data->:creator->:key = :expected_value::::jsonb
				ELSE FALSE
				END
			ELSE not (:check_empty::::boolean and :have_expected_value::::boolean)
		  END
	),
	inserted AS
	(
		UPDATE account
		SET data = jsonb_set(
		  CASE
			WHEN data ? :creator THEN data
			ELSE jsonb_set(data, array[:creator], '{}')
		  END,
		  array[:creator, :key], (CAST:new_value AS jsonb)
		)
		WHERE
		  EXISTS (SELECT * FROM old_value)
		  AND account_id = :target
		  %s
		RETURNING (1)
	)
  SELECT CASE
	  WHEN EXISTS (SELECT * FROM inserted) THEN 0
	  WHEN NOT EXISTS
		  (SELECT * FROM account WHERE account_id=:target) THEN 3
	  WHEN NOT EXISTS (SELECT * FROM old_value) THEN 4
	  %s
	  ELSE 1
  END AS result
`,
		fmt.Sprintf(`
has_role_perm AS (%s),
has_grantable_perm AS (%s),
%s,
has_perm AS
(
  SELECT CASE
	  WHEN (SELECT * FROM has_query_perm) THEN
		  CASE
			  WHEN (SELECT * FROM has_grantable_perm)
				  THEN true
			  WHEN (:creator = :target) THEN true
			  WHEN (SELECT * FROM has_role_perm)
				  THEN true
			  ELSE false END
	  ELSE false END
),
`,
			checkAccountRolePermission(pb.RolePermission_can_set_detail, ":creator"),
			checkAccountGrantablePermission(pb.GrantablePermission_can_set_my_account_detail, ":creator", ":target"),
			hasQueryPermission(
				":creator", ":target",
				pb.RolePermission_can_get_my_acc_detail,
				pb.RolePermission_can_get_all_acc_detail,
				pb.RolePermission_can_get_domain_acc_detail,
				":creator_domain", ":target_domain",
			),
		),
		"AND (SELECT * from has_perm)",
		"WHEN NOT (SELECT * from has_perm) THEN 2",
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator":             c.caller,
		"target":              accountID,
		"key":                 key,
		"new_value":           value,
		"check_empty":         checkEmpty,
		"have_expected_value": len(oldValue) > 0,
		"expected_value":      oldValue,
		"creator_domain":      getDomainFromName(c.caller),
		"target_domain":       getDomainFromName(accountID),
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("failed CompareAndSetAccountDetail")
	}

	return nil
}

func (c *postgresExecer) CreateRole(name string, permissions string) error {
	query := fmt.Sprintf(`
WITH %s
	insert_role AS (INSERT INTO role(role_id)
						(SELECT :role
						%s) RETURNING (1)),
	insert_role_permissions AS
	(
		INSERT INTO role_has_permissions(role_id, permission)
		(
			SELECT :role, :perms WHERE EXISTS
				(SELECT * FROM insert_role)
		) RETURNING (1)
	)
	SELECT CASE
		WHEN EXISTS (SELECT * FROM insert_role_permissions) THEN 0
		%s
		WHEN EXISTS (SELECT * FROM role WHERE role_id = :role) THEN 2
		ELSE 1
	END AS result
`,
		fmt.Sprintf(`
account_has_role_permissions AS (
	SELECT COALESCE(bit_or(rp.permission), '0'::::bit(%d)) &
		:perms = :perms
	FROM role_has_permissions AS rp
	JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	WHERE ar.account_id = :creator),
has_perm AS (%s),
has_root_perm AS (%s),
`,
			consts.RolePermissionEnumLength,
			checkAccountRolePermission(pb.RolePermission_can_create_role, ":creator"),
			checkAccountRolePermission(pb.RolePermission_root, ":creator"),
		),
		`
WHERE (SELECT * FROM has_root_perm) OR
	((SELECT * FROM account_has_role_permissions)
	 AND (SELECT * FROM has_perm))
`,
		`
WHEN NOT (SELECT * FROM account_has_role_permissions)
	AND NOT (SELECT * FROM has_root_perm) THEN 2
	WHEN NOT (SELECT * FROM has_perm) THEN 2
`,
	)

	result, err := c.execer.NamedExec(query, map[string]interface{}{
		"creator": c.caller,
		"role":    name,
		"perms":   permissions,
	})
	if err != nil {
		return err
	} else if affected, err := result.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return fmt.Errorf("CreateRole")
	}

	return nil
}
