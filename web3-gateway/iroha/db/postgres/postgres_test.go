package postgres

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

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

	query := `
	SELECT
		ec.call_id,
		tp.index,
		ec.tx_hash,
		ec.callee,
		tp.creator_id,
		tp.height,
		btl.data
	FROM engine_calls AS ec
	INNER JOIN tx_positions AS tp
	ON ec.tx_hash = tp.hash
	LEFT OUTER JOIN burrow_tx_logs AS btl
	ON ec.call_id = btl.call_id
	WHERE ec.tx_hash = $1
`

	var tx entity.EngineTransaction

	if err = ddb.Get(&tx, query, "8f47d3d6de50b9d78690c6589d833c3e0eee1701011b0f25ab8dcfb50b3d7004"); err != nil {
		t.Error(err)
	}

	log.Printf("%+v", tx)
}
