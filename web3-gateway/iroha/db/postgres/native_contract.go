package postgres

import (
	"fmt"
	"math/big"
	"strings"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/consts"
)

func hasQueryPermission(
	creatorAccountID, targetAccountID string,
	indivPermissionID, allPermissionID, domainPermissionID pb.RolePermission,
	creatorDomain, targetDomain string,
) string {
	bits := consts.RolePermissionEnumLength

	return fmt.Sprintf(`
has_root_perm AS (%[1]v),
has_indiv_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[4]v') = '%[4]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = %[3]v
),
has_all_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[5]v') = '%[5]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = %[3]v
),
has_domain_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[6]v') = '%[6]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = %[3]v
),
has_query_perm AS (
  SELECT (SELECT * from has_root_perm)
	  OR (%[3]v = %[7]v AND (SELECT * FROM has_indiv_perm))
	  OR (SELECT * FROM has_all_perm)
	  OR (%[8]v = %[9]v AND (SELECT * FROM has_domain_perm)) AS perm
)
`,
		checkAccountRolePermission(pb.RolePermission_root, creatorAccountID),
		bits,
		creatorAccountID,
		roleToBitString(indivPermissionID),
		roleToBitString(allPermissionID),
		roleToBitString(domainPermissionID),
		targetAccountID,
		creatorDomain,
		targetDomain,
	)
}

func hasQueryPermissionTarget(
	creatorAccountID, targetAccountID string,
	indivPermissionID, allPermissionID, domainPermissionID pb.RolePermission,
) string {
	return fmt.Sprintf(`target AS (select '%s'::::text as t), %s`,
		targetAccountID,
		hasQueryPermissionInternal(
			creatorAccountID,
			indivPermissionID, allPermissionID, domainPermissionID,
		),
	)
}

func hasQueryPermissionInternal(
	creatorAccountID string,
	indivPermissionID, allPermissionID, domainPermissionID pb.RolePermission,
) string {
	bits := consts.RolePermissionEnumLength
	creatorAccountIDQuoted := fmt.Sprintf("'%s'", creatorAccountID)
	domain := getDomainFromName(creatorAccountID)

	return fmt.Sprintf(`
target_domain AS (select split_part(target.t, '@', 2) as td from target),
has_root_perm AS (%[1]v),
has_indiv_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[4]v') = '%[4]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = '%[3]v'
),
has_all_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[5]v') = '%[5]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = '%[3]v'
),
has_domain_perm AS (
  SELECT (COALESCE(bit_or(rp.permission), '0'::::bit(%[2]v))
  & '%[6]v') = '%[6]v' FROM role_has_permissions AS rp
	  JOIN account_has_roles AS ar on ar.role_id = rp.role_id
	  WHERE ar.account_id = '%[3]v'
),
has_perms as (
  SELECT (SELECT * from has_root_perm)
	  OR ('%[3]v' = (select t from target) AND (SELECT * FROM has_indiv_perm))
	  OR (SELECT * FROM has_all_perm)
	  OR ('%[7]v' = (select td from target_domain) AND (SELECT * FROM has_domain_perm)) AS perm
)
`,
		getAccountRolePermissionCheckSql(pb.RolePermission_root, creatorAccountIDQuoted),
		bits,
		creatorAccountID,
		roleToBitString(indivPermissionID),
		roleToBitString(allPermissionID),
		roleToBitString(domainPermissionID),
		domain,
	)
}

func getAccountRolePermissionCheckSql(permission pb.RolePermission, accountAlias string) string {
	bits := consts.RolePermissionEnumLength
	permStr := roleToBitString(permission)

	return fmt.Sprintf(`
SELECT
	(
	  COALESCE(bit_or(rp.permission), '0'::::bit(%[1]v))
	  & ('%[2]v'::::bit(%[1]v) | '%[3]v'::::bit(%[1]v))
	) != '0'::::bit(%[1]v)
	AS perm
FROM role_has_permissions AS rp
JOIN account_has_roles AS ar on ar.role_id = rp.role_id
WHERE ar.account_id = %[4]v
`,
		bits, permStr, roleToBitString(pb.RolePermission_root), accountAlias)
}

func checkAccountRolePermission(permission pb.RolePermission, accountAlias string) string {
	bits := consts.RolePermissionEnumLength
	permStr := roleToBitString(permission)

	return fmt.Sprintf(`
SELECT
	COALESCE(bit_or(rp.permission), '0'::::bit(%[1]v))
	& ('%[2]v'::::bit(%[1]v) | '%[3]v'::::bit(%[1]v))
	!= '0'::::bit(%[1]v) has_rp
FROM role_has_permissions AS rp
JOIN account_has_roles AS ar on ar.role_id = rp.role_id
WHERE ar.account_id = %[4]v
`,
		bits, permStr, roleToBitString(pb.RolePermission_root), accountAlias)
}

func checkAccountRolePermissionWithAdditional(permission, additional pb.RolePermission, accountAlias string) string {
	bits := consts.RolePermissionEnumLength
	permStr := roleToBitString(permission)
	additionalPermStr := roleToBitString(additional)

	return fmt.Sprintf(`
SELECT
	COALESCE(bit_or(rp.permission), '0'::::bit(%[1]v))
	& ('%[2]v'::::bit(%[1]v) | '%[5]v'::::bit(%[1]v) | '%[3]v'::::bit(%[1]v))
	!= '0'::::bit(%[1]v) has_rp
FROM role_has_permissions AS rp
JOIN account_has_roles AS ar on ar.role_id = rp.role_id
WHERE ar.account_id = %[4]v
`,
		bits, permStr, roleToBitString(pb.RolePermission_root), accountAlias, additionalPermStr)
}

func checkAccountGrantablePermission(permission pb.GrantablePermission, creatorAccountID, targetAccountID string) string {
	bits := consts.GrantableRolePermissionEnumLength
	permStr := grantableRoleToBitString(permission)

	return fmt.Sprintf(`
SELECT
	  COALESCE(bit_or(permission), '0'::::bit(%[1]v)) & '%[2]v' = '%[2]v'
	  or (%[3]v)
FROM account_has_grantable_permissions
WHERE account_id = %[4]v AND
permittee_account_id = %[5]v`,
		bits,
		permStr,
		checkAccountRolePermission(pb.RolePermission_root, creatorAccountID),
		targetAccountID,
		creatorAccountID,
	)
}

func checkAccountDomainRoleOrGlobalRolePermission(
	globalPermission, domainPermission pb.RolePermission,
	creatorAccountID string,
	idWithTargetDomain string,
) string {
	return fmt.Sprintf(`
WITH
	has_global_role_perm AS (%[1]v),
	has_domain_role_perm AS (%[2]v)
	SELECT CASE
					   WHEN (SELECT * FROM has_global_role_perm) THEN true
					   WHEN ((split_part(%[3]v, '@', 2) = split_part(%[4]v, '#', 2))) THEN
						   CASE
							   WHEN (SELECT * FROM has_domain_role_perm) THEN true
							   ELSE false
							END
					   ELSE false END
`,
		checkAccountRolePermission(globalPermission, creatorAccountID),
		checkAccountRolePermission(domainPermission, creatorAccountID),
		creatorAccountID,
		idWithTargetDomain,
	)
}

func checkAccountHasRoleOrGrantablePerm(
	role pb.RolePermission, grantable pb.GrantablePermission,
	creatorAccountID, targetAccountID string,
) string {
	return fmt.Sprintf(`
WITH
	has_role_perm AS (%s),
	has_root_perm AS (%s),
	has_grantable_perm AS (%s)
	SELECT CASE
		WHEN (SELECT * FROM has_root_perm) THEN true
		WHEN (SELECT * FROM has_grantable_perm) THEN true
		WHEN (%s = %s) THEN
			CASE
				WHEN (SELECT * FROM has_role_perm) THEN true
				ELSE false
			END
		ELSE false END
`,
		checkAccountRolePermission(role, creatorAccountID),
		checkAccountRolePermission(pb.RolePermission_root, creatorAccountID),
		checkAccountGrantablePermission(grantable, creatorAccountID, targetAccountID),
		creatorAccountID,
		targetAccountID,
	)
}

func getDomainFromName(accountID string) string {
	return strings.Split(accountID, "@")[1]
}

func roleToBitString(perm pb.RolePermission) string {
	return bitString(int(perm), consts.RolePermissionEnumLength)
}

func grantableRoleToBitString(perm pb.GrantablePermission) string {
	return bitString(int(perm), consts.GrantableRolePermissionEnumLength)
}

func bitString(perm int, strLen int) string {
	var m = new(big.Int)
	m.SetBit(m, perm, 1)

	return fmt.Sprintf(fmt.Sprintf(`%%0%ds`, strLen), m.Text(2))
}

func getPrecisionFromAmount(amount string) int {
	pos := strings.Index(amount, ".")
	if pos == -1 {
		return 0
	}
	return len(amount) - 1 - pos
}

func permissionFor(perm pb.GrantablePermission) pb.RolePermission {
	switch perm {
	case pb.GrantablePermission_can_add_my_signatory:
		return pb.RolePermission_can_add_signatory
	case pb.GrantablePermission_can_remove_my_signatory:
		return pb.RolePermission_can_remove_signatory
	case pb.GrantablePermission_can_set_my_quorum:
		return pb.RolePermission_can_set_quorum
	case pb.GrantablePermission_can_set_my_account_detail:
		return pb.RolePermission_can_grant_can_set_my_account_detail
	case pb.GrantablePermission_can_transfer_my_assets:
		return pb.RolePermission_can_grant_can_transfer_my_assets
	case pb.GrantablePermission_can_call_engine_on_my_behalf:
		return pb.RolePermission_can_grant_can_call_engine_on_my_behalf
	default:
		return consts.RolePermissionEnumLength
	}
}
