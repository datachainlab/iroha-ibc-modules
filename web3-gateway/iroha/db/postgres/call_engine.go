package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	x "github.com/hyperledger/burrow/encoding/hex"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

func (c *postgresExecer) GetEngineTransaction(txHash string) (*entity.EngineTransaction, error) {
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

	if err := c.execer.Get(&tx, query, txHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &tx, nil
}

func (c *postgresExecer) GetEngineReceipt(txHash string) (*entity.EngineReceipt, error) {
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

	if err := c.execer.Get(&receipt, query, txHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &receipt, nil
}

func (c *postgresExecer) GetEngineReceiptLogsByTxHash(txHash string) ([]*entity.EngineReceiptLog, error) {
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

	if err := c.execer.Select(&logs, queryLogs, txHash); err != nil {
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
		if err := c.execer.Select(&topics, queryTopics, log.LogIdx); err != nil {
			return nil, err
		}
		log.Topics = topics
	}

	return logs, nil
}

func (c *postgresExecer) GetEngineReceiptLogsByFilters(opts ...db.LogFilterOption) ([]*entity.EngineReceiptLog, error) {

	filter := &db.TxReceiptLogFilter{
		FromBlock: 0,
		ToBlock:   0,
		Addresses: nil,
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

	if len(filter.Addresses) > 0 {
		clause()
		var addressStr string
		for i, addr := range filter.Addresses {
			sep := ""
			if i > 0 {
				sep = ", "
			}
			addr = x.RemovePrefix(addr)
			addressStr = fmt.Sprintf("%s%sLOWER('%s'), UPPER('%s')", addressStr, sep, addr, addr)
		}
		conditions.WriteString(fmt.Sprintf("address IN (%s)", addressStr))
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
	if err := c.execer.Select(&logs, queryLogs); err != nil {
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
		if err := c.execer.Select(&topics, queryTopics, log.LogIdx); err != nil {
			return nil, err
		}
		log.Topics = topics
	}

	return logs, nil
}
