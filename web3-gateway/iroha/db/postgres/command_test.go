package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/datachainlab/iroha-ibc-modules/iroha-go/iroha.generated/protocol"
)

func TestPostgresExecer_CreateAccount(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.CreateAccount("name", "test", "716fe505f69f18511a1b083915aa9ff73ef36e6688199f3959750db38b8f4bfb")
	assert.NoError(t, err)
}

func TestPostgresExecer_AddAssetQuantity(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.AddAssetQuantity(assetID, "10.50")
	assert.NoError(t, err)
}

func TestPostgresExecer_SubtractAssetQuantity(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.SubtractAssetQuantity(assetID, "1.30")
	assert.NoError(t, err)
}

func TestPostgresExecer_SetAccountDetail(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.SetAccountDetail(adminAccount, "key", "value")
	assert.NoError(t, err)
}

func TestPostgresExecer_SetAccountQuorum(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.SetAccountQuorum(userAccount, "2")
	assert.NoError(t, err)
}

func TestPostgresExecer_AddSignatory(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.AddSignatory(adminAccount, userPublicKeyTwo)
	assert.NoError(t, err)
}

func TestPostgresExecer_RemoveSignatory(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.RemoveSignatory(adminAccount, userPublicKeyTwo)
	assert.NoError(t, err)
}

func TestPostgresExecer_CreateDomain(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.CreateDomain(testDomain, defaultRole)
	assert.NoError(t, err)
}

func TestPostgresExecer_CreateAsset(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.CreateAsset("testcoin2", "test", "5")
	assert.NoError(t, err)
}

func TestPostgresExecer_AppendRole(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.AppendRole(userAccount, "money_creator")
	assert.NoError(t, err)
}

func TestPostgresExecer_DetachRole(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.DetachRole(userAccount, "money_creator")
	assert.NoError(t, err)
}

func TestPostgresExecer_AddPeer(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.AddPeer("127.0.0.1:10002", "bddd58404d1315e0eb27902c5d7c8eb0602c16238f005773df406bc191308928")
	assert.NoError(t, err)
}

func TestPostgresExecer_RemovePeer(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: adminAccount}

	err = e.RemovePeer("bddd58404d1315e0eb27902c5d7c8eb0602c16238f005773df406bc191308928")
	assert.NoError(t, err)
}

func TestPostgresExecer_GrantPermission(t *testing.T) {
	var err error

	db := testDB(t)
	defer db.Close()

	e := postgresExecer{execer: db, caller: userAccount}

	err = e.GrantPermission(adminAccount, pb.GrantablePermission_can_set_my_quorum.String())
	assert.NoError(t, err)
}
