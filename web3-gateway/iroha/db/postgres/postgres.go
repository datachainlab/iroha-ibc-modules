package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	x "github.com/hyperledger/burrow/encoding/hex"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

var _ db.DBClient = (*postgresClient)(nil)

type postgresClient struct {
	db *sqlx.DB
}

func NewClient(user string, password string, host string, port int, database string) (db.DBClient, *sql.DB, error) {
	conn, err := sql.Open(
		"pgx",
		fmt.Sprintf(
			// postgres://user:password@host:port/database
			"postgres://%s:%s@%s:%v/%s",
			user, password, host, port, database,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	return &postgresClient{
		db: sqlx.NewDb(conn, "postgres"),
	}, conn, nil
}

func (c *postgresClient) GetLatestHeight() (uint64, error) {
	var block entity.TopBlockInfo

	query := "SELECT lock, height, hash FROM top_block_info"

	if err := c.db.Get(&block, query); err != nil {
		return 0, err
	}

	return block.Height, nil
}

func (c *postgresClient) GetBurrowAccountDataByAddress(address string) (*entity.BurrowAccountData, error) {
	address = strings.ToLower(x.RemovePrefix(address))

	var account entity.BurrowAccountData

	query := "SELECT address, data FROM burrow_account_data WHERE address=$1"

	if err := c.db.Get(&account, query, address); err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *postgresClient) GetEngineTransaction(txHash string) (*entity.EngineTransaction, error) {
	txHash = strings.ToLower(x.RemovePrefix(txHash))

	var tx entity.EngineTransaction

	query := `
SELECT
	ec.call_id,
	tp.index,
	ec.tx_hash,
	ec.callee,
	tp.creator_id,
	tp.height,
    tp.ts,
	btl.data
FROM engine_calls AS ec
INNER JOIN tx_positions AS tp
ON ec.tx_hash = tp.hash
LEFT OUTER JOIN burrow_tx_logs AS btl
ON ec.call_id = btl.call_id
WHERE ec.tx_hash = $1
`

	if err := c.db.Get(&tx, query, txHash); err != nil {
		return nil, err
	}

	return &tx, nil
}

func (c *postgresClient) GetEngineReceipt(txHash string) (*entity.EngineReceipt, error) {
	txHash = strings.ToLower(x.RemovePrefix(txHash))

	var receipt entity.EngineReceipt

	query := `
SELECT
	ec.call_id,
	ec.tx_hash,
	ec.callee,
	ec.created_address,
	tp.creator_id,
	tp.height,
    tp.ts,
	tp.index,
	tsbh.status
FROM engine_calls AS ec
INNER JOIN tx_positions AS tp
ON ec.tx_hash = tp.hash
INNER JOIN tx_status_by_hash AS tsbh
ON tp.hash = tsbh.hash
WHERE ec.tx_hash = $1
`

	if err := c.db.Get(&receipt, query, txHash); err != nil {
		return nil, err
	}

	return &receipt, nil
}
