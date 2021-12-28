package entity

import (
	"database/sql"
)

type TopBlockInfo struct {
	Lock   string `db:"lock"`
	Height uint64 `db:"height"`
	Hash   string `db:"hash"`
}

type Account struct {
	Perm
	AccountID string `db:"account_id"`
	DomainID  string `db:"domain_id"`
	Quorum    int64  `db:"quorum"`
	Data      string `db:"data"`
	Roles     string `db:"roles"`
}

type AccountDetail struct {
	Perm
	Json                string         `db:"json"`
	TotalNumber         int64          `db:"total_number"`
	DomainID            sql.NullString `db:"next_writer"`
	Data                sql.NullString `db:"next_key"`
	TargetAccountExists int64          `db:"target_account_exists"`
}

type AssetInfo struct {
	Perm
	DomainID  string `db:"domain_id"`
	Precision int64  `db:"precision"`
}

type AccountAsset struct {
	Perm
	AccountID   string `db:"account_id"`
	AssetID     string `db:"asset_id"`
	Amount      int64  `db:"amount"`
	TotalNumber int64  `db:"total_number"`
}

type Signatory struct {
	Perm
	PublicKey string `db:"public_key"`
}

type Peer struct {
	Perm
	PublicKey      string         `db:"public_key"`
	Address        string         `db:"address"`
	TlsCertificate sql.NullString `db:"tls_certificate"`
}

type Role struct {
	Perm
	RoleID string `db:"role_id"`
}

type RolePermission struct {
	Perm
	Permission string `db:"permission"`
}

type Perm struct {
	Perm bool `db:"perm"`
}

type Transaction struct {
	Perm
	Height int64  `db:"height"`
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
