package entity

import (
	"database/sql"
)

type TopBlockInfo struct {
	Lock   string `db:"lock"`
	Height uint64 `db:"height"`
	Hash   string `db:"hash"`
}

type BurrowAccountData struct {
	Address string `db:"address"`
	Data    string `db:"data"`
}

type BurrowAccountKeyValue struct {
	Address string `db:"address"`
	Key     string `db:"key"`
	Value   string `db:"value"`
}

type EngineTransaction struct {
	CallID               int64          `db:"call_id"`
	TxHash               string         `db:"tx_hash"`
	Callee               sql.NullString `db:"callee"`
	CreatorID            string         `db:"creator_id"`
	Height               uint64         `db:"height"`
	TransactionTimeStamp uint64         `db:"ts"`
	Index                uint64         `db:"index"`
	Data                 sql.NullString `db:"data"`
}

type EngineReceipt struct {
	CallID               int64              `db:"call_id"`
	TxHash               string             `db:"tx_hash"`
	Callee               sql.NullString     `db:"callee"`
	CreatedAddress       sql.NullString     `db:"created_address"`
	CreatorID            string             `db:"creator_id"`
	Height               uint64             `db:"height"`
	TransactionTimeStamp uint64             `db:"ts"`
	Index                uint64             `db:"index"`
	Status               bool               `db:"status"`
	Logs                 []EngineReceiptLog `db:"-"`
}

type EngineReceiptLog struct {
	LogIdx    int64                   `db:"log_idx"`
	CallID    int64                   `db:"call_id"`
	Address   string                  `db:"address"`
	Data      string                  `db:"data"`
	CreatorID string                  `db:"creator_id"`
	Height    uint64                  `db:"height"`
	Index     uint64                  `db:"index"`
	TxHash    string                  `db:"tx_hash"`
	Topics    []EngineReceiptLogTopic `db:"-"`
}

type EngineReceiptLogTopic struct {
	Topic string `db:"topic"`
}
