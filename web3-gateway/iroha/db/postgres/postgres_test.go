package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	iroha "github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

const source = "postgres://postgres:mysecretpassword@localhost:5432/iroha_data"

func TestClient(t *testing.T) {
	t.SkipNow()
	db, err := sql.Open("pgx", source)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	ddb := sqlx.NewDb(db, "postgres")

	filter := &iroha.TxReceiptLogFilter{
		FromBlock: 0,
		ToBlock:   0,
		Address:   "",
		Topics:    nil,
	}

	logFilterOpts := []iroha.LogFilterOption{
		//iroha.FromBlockOption(0),
		//iroha.ToBlockOption(0),
		//iroha.AddressOption("708b5cd2603fa4f8fc3f3071598bf2080be16093"),
		//iroha.TopicsOption(
		//	//"981503e0fe5bde28ee3eefdc476eb74e602bbb29f68150d095bfae68a6875178",
		//),
	}

	for _, opt := range logFilterOpts {
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
		conditions.WriteString(fmt.Sprintf("address IN (LOWER('%s'), UPPER('%s'))", filter.Address, filter.Address))
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

	query := fmt.Sprintf(`
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

	log.Println(query)

	var logs []*entity.EngineReceiptLog
	if err = ddb.Select(&logs, query); err != nil {
		t.Error(err)
	}

	queryTopics := `
	SELECT
		btlt.topic
	FROM burrow_tx_logs_topics AS btlt
	WHERE btlt.log_idx = $1
		`

	for _, l := range logs {
		var topics []entity.EngineReceiptLogTopic
		if err = ddb.Select(&topics, queryTopics, l.LogIdx); err != nil {
			t.Error(err)
		}
		l.Topics = topics
		log.Printf("%+v", l)
	}

}
