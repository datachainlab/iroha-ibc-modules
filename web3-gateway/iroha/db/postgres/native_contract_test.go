package postgres

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func TestHasQueryPermissionTarget(t *testing.T) {
	query := hasQueryPermissionTarget(
		"admin@test", "test@test",
		pb.RolePermission_can_get_my_acc_ast,
		pb.RolePermission_can_get_all_acc_ast,
		pb.RolePermission_can_get_domain_acc_ast,
	)
	log.Printf("%s", query)
}

func TestCheckAccountRolePermission(t *testing.T) {
	query := checkAccountRolePermission(
		pb.RolePermission_can_transfer,
		"admin@test",
	)
	log.Printf("%s", query)
}

func TestCheckAccountGrantablePermission(t *testing.T) {
	query := checkAccountGrantablePermission(
		pb.GrantablePermission_can_transfer_my_assets,
		"admin@test",
		"test@test",
	)
	log.Printf("%s", query)
}

func TestRoleToBitString(t *testing.T) {
	tests := []struct {
		value    pb.RolePermission
		expected string
	}{
		{value: pb.RolePermission_can_get_all_engine_receipts, expected: "10000000000000000000000000000000000000000000000000000"},
		{value: pb.RolePermission_root, expected: "00000100000000000000000000000000000000000000000000000"},
		{value: pb.RolePermission_can_append_role, expected: "00000000000000000000000000000000000000000000000000001"},
	}
	for _, c := range tests {
		result := roleToBitString(c.value)
		assert.Equal(t, c.expected, result)
	}
}

func TestGrantableRoleToBitString(t *testing.T) {
	tests := []struct {
		value    pb.GrantablePermission
		expected string
	}{
		{value: pb.GrantablePermission_can_call_engine_on_my_behalf, expected: "100000"},
		{value: pb.GrantablePermission_can_transfer_my_assets, expected: "010000"},
		{value: pb.GrantablePermission_can_set_my_account_detail, expected: "001000"},
		{value: pb.GrantablePermission_can_set_my_quorum, expected: "000100"},
		{value: pb.GrantablePermission_can_remove_my_signatory, expected: "000010"},
		{value: pb.GrantablePermission_can_add_my_signatory, expected: "000001"},
	}
	for _, c := range tests {
		result := grantableRoleToBitString(c.value)
		assert.Equal(t, c.expected, result)
	}
}

func TestGetPrecisionFromAmount(t *testing.T) {
	tests := []struct {
		value    string
		expected int
	}{
		{value: "100", expected: 0},
		{value: "100.50", expected: 2},
		{value: "0.555", expected: 3},
	}

	for _, c := range tests {
		result := getPrecisionFromAmount(c.value)
		assert.Equal(t, c.expected, result)
	}
}
