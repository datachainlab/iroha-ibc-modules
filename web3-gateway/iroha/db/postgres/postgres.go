package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/entity"
)

var _ db.DBTransactor = (*postgresTransactor)(nil)

type postgresTransactor struct {
	db       *sqlx.DB
	txOpts   *sql.TxOptions
	readOnly bool
}

type TransactorOption func(*postgresTransactor)

func TxOpts(txOpts *sql.TxOptions) TransactorOption {
	return func(execer *postgresTransactor) {
		execer.txOpts = txOpts
	}
}

func ReadOnly(readOnly bool) TransactorOption {
	return func(execer *postgresTransactor) {
		execer.readOnly = readOnly
	}
}

func NewTransactor(
	user string, password string, host string, port int, database string,
	opts ...TransactorOption,
) (db.DBTransactor, error) {
	conn, err := sql.Open(
		"pgx",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%v/%s",
			user, password, host, port, database,
		),
	)
	if err != nil {
		return nil, err
	}

	transactor := &postgresTransactor{
		db:       sqlx.NewDb(conn, "postgres"),
		txOpts:   nil,
		readOnly: true,
	}

	for _, opt := range opts {
		opt(transactor)
	}

	return transactor, nil
}

func (c *postgresTransactor) Close() error {
	return c.db.Close()
}

func (c *postgresTransactor) Exec(ctx context.Context, caller string, f func(execer db.DBExecer) error) error {
	return f(&postgresExecer{execer: c.db, caller: caller})
}

func (c *postgresTransactor) ExecWithTxBoundary(ctx context.Context, caller string, f func(execer db.DBExecer) error) (err error) {
	tx, err := c.db.BeginTxx(ctx, c.txOpts)
	if err != nil {
		return err
	}
	defer func() {
		if c.readOnly {
			err = tx.Rollback()
		} else if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				err = fmt.Errorf("%w:%s", err, txErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	if err = f(&postgresExecer{execer: tx, caller: caller}); err != nil {
		return err
	}

	return nil
}

var _ db.DBExecer = (*postgresExecer)(nil)

type execer interface {
	sqlx.Queryer
	sqlx.Execer

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

var _ execer = (*sqlx.DB)(nil)
var _ execer = (*sqlx.Tx)(nil)

type postgresExecer struct {
	execer execer
	caller string
}

func (c *postgresExecer) GetLatestHeight() (uint64, error) {
	var block entity.TopBlockInfo

	query := "SELECT lock, height, hash FROM top_block_info"

	if err := c.execer.Get(&block, query); err != nil {
		return 0, err
	}

	return block.Height, nil
}
