package postgres

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const (
	adminAccount     = "admin@test"
	adminPublicKey   = "313a07e6384776ed95447710d15e59148473ccfc052a681317a72a69f2a49910"
	userAccount      = "test@test"
	userPublicKey    = "716fe505f69f18511a1b083915aa9ff73ef36e6688199f3959750db38b8f4bfc"
	userPublicKeyTwo = "716fe505f69f18511a1b083915aa9ff73ef36e6688199f3959750db38b8f4bfb"
	assetID          = "testcoin#test"
	source           = "postgres://postgres:mysecretpassword@localhost:5432/iroha_data"

	testDomain  = "test2"
	defaultRole = "user"
)

func TestGetAccountDetail(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	detail, err := e.GetAccountDetail()
	assert.NoError(t, err)
	assert.NotEmpty(t, detail)
}

func TestGetAccount(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	account, err := e.GetAccount(adminAccount)
	assert.NoError(t, err)
	assert.Equal(t, adminAccount, account.AccountID)
}

func TestGetAssetInfo(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	assetInfo, err := e.GetAssetInfo(assetID)
	assert.NoError(t, err)
	assert.Equal(t, strings.Split(assetID, "#")[1], assetInfo.DomainID)
}

func TestGetSignatories(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	sigs, err := e.GetSignatories(adminAccount)
	assert.NoError(t, err)
	assert.Contains(t, sigs, adminPublicKey)
}

func TestGetPeers(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	peers, err := e.GetPeers()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, peers, 1)
}

func TestGetRoles(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	roles, err := e.GetRoles()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(roles), 1)
}

func TestGetRolePermissions(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	perms, err := e.GetRolePermissions("gateway_querier")
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(perms), 1)
}

func testDB(t *testing.T) *sqlx.DB {
	conn, err := sql.Open("pgx", source)
	assert.NoError(t, err)

	return sqlx.NewDb(conn, "postgres")
}
