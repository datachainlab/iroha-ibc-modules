package postgres

import (
	"bytes"
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}

func (c *postgresClient) GetBurrowAccountKeyValueByAddressAndKey(address, key string) (*entity.BurrowAccountKeyValue, error) {
	address = strings.ToLower(x.RemovePrefix(address))
	key = strings.ToLower(x.RemovePrefix(key))

	var kv entity.BurrowAccountKeyValue

	query := "SELECT address, key, value FROM burrow_account_key_value WHERE address=$1 AND key=$2"

	if err := c.db.Get(&kv, query, address, key); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &kv, nil
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &receipt, nil
}

func (c *postgresClient) GetEngineReceiptLogsByTxHash(txHash string) ([]*entity.EngineReceiptLog, error) {
	txHash = strings.ToLower(x.RemovePrefix(txHash))

	var logs []*entity.EngineReceiptLog

	queryLogs := `
SELECT
	btl.log_idx,
	btl.call_id,
	btl.address,
	btl.data,
	tp.creator_id,
	tp.height,
	tp.index,
	tp.hash as tx_hash
FROM burrow_tx_logs AS btl
INNER JOIN engine_calls AS ec
ON btl.call_id = ec.call_id
INNER JOIN tx_positions AS tp
ON tp.hash = ec.tx_hash
WHERE ec.tx_hash = $1
`

	if err := c.db.Select(&logs, queryLogs, txHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	queryTopics := `
SELECT
	btlt.topic
FROM burrow_tx_logs_topics AS btlt
WHERE btlt.log_idx = $1
`

	for _, log := range logs {
		var topics []entity.EngineReceiptLogTopic
		if err := c.db.Select(&topics, queryTopics, log.LogIdx); err != nil {
			return nil, err
		}
		log.Topics = topics
	}

	return logs, nil
}

func (c *postgresClient) GetEngineReceiptLogsByFilters(opts ...db.LogFilterOption) ([]*entity.EngineReceiptLog, error) {

	filter := &db.TxReceiptLogFilter{
		FromBlock: 0,
		ToBlock:   0,
		Address:   "",
		Topics:    nil,
	}

	for _, opt := range opts {
		opt(filter)
	}

	var conditions bytes.Buffer

	clause := func() {
		if conditions.Len() == 0 {
			conditions.WriteString(" WHERE ")
		} else {
			conditions.WriteString(" AND ")
		}
	}

	if filter.FromBlock > 0 {
		clause()
		conditions.WriteString(fmt.Sprintf("height>=%d", filter.FromBlock))
	}

	if filter.ToBlock > 0 {
		clause()
		conditions.WriteString(fmt.Sprintf("height<=%d", filter.ToBlock))
	}

	if len(filter.Address) > 0 {
		clause()
		address := x.RemovePrefix(filter.Address)
		conditions.WriteString(fmt.Sprintf("address IN (LOWER('%s'), UPPER('%s'))", address, address))
	}

	if len(filter.Topics) > 0 {
		clause()
		var topicsStr string
		for i, topic := range filter.Topics {
			sep := ""
			if i > 0 {
				sep = ","
			}
			topicsStr = fmt.Sprintf("%s%s'%s'", topicsStr, sep, topic)
		}
		conditions.WriteString(fmt.Sprintf(`log_idx IN (
SELECT
	DISTINCT btl.log_idx
FROM 
	burrow_tx_logs_topics AS btlt
INNER JOIN 
	burrow_tx_logs btl 
ON 
	btlt.log_idx = btl.log_idx
WHERE 
	btlt.topic IN (%s)
)`, topicsStr))
	}

	queryLogs := fmt.Sprintf(`
SELECT
	btl.log_idx as log_idx,
	btl.call_id as call_id,
	btl.address as address,
	btl.data as data,
	tp.creator_id,
	tp.height as height,
	tp.index as index,
	tp.hash as tx_hash
FROM 
	burrow_tx_logs AS btl
INNER JOIN 
	engine_calls AS ec
ON 
	btl.call_id = ec.call_id
INNER JOIN 
	tx_positions AS tp
ON 
	tp.hash = ec.tx_hash
%s
ORDER BY log_idx`, conditions.String())

	var logs []*entity.EngineReceiptLog
	if err := c.db.Select(&logs, queryLogs); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	queryTopics := `
SELECT
	btlt.topic
FROM burrow_tx_logs_topics AS btlt
WHERE btlt.log_idx = $1
`

	for _, log := range logs {
		var topics []entity.EngineReceiptLogTopic
		if err := c.db.Select(&topics, queryTopics, log.LogIdx); err != nil {
			return nil, err
		}
		log.Topics = topics
	}

	return logs, nil
}