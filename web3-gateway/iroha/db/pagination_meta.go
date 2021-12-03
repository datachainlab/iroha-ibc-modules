package db

type TxPaginationMeta struct {
	PageSize      *string
	FirstTxHash   *string
	Ordering      *string
	FirstTxTime   *string
	LastTxTime    *string
	FirstTxHeight *string
	LastTxHeight  *string
}

type OrderingField struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}
