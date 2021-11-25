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

type EngineTransaction struct {
	CallID               int            `db:"call_id"`
	TxHash               string         `db:"tx_hash"`
	Callee               sql.NullString `db:"callee"`
	CreatorID            string         `db:"creator_id"`
	Height               int            `db:"height"`
	TransactionTimeStamp uint64         `db:"ts"`
	Index                int            `db:"index"`
	Data                 sql.NullString `db:"data"`
}

type EngineReceipt struct {
	CallID               int            `db:"call_id"`
	TxHash               string         `db:"tx_hash"`
	Callee               sql.NullString `db:"callee"`
	CreatedAddress       sql.NullString `db:"created_address"`
	CreatorID            string         `db:"creator_id"`
	Height               uint64         `db:"height"`
	TransactionTimeStamp uint64         `db:"ts"`
	Index                int            `db:"index"`
	Status               bool           `db:"status"`
}
