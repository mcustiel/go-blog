package sql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/mcustiel/go-blog/pkg/persistence"
)

type SqlConnection struct {
	driver       string
	dsn          string
	pool         *sql.DB
	ctx          context.Context
	configurator func(*sql.DB) error
}

func NewSqlConnection(driver string, dsn string, configurator func(*sql.DB) error) *SqlConnection {
	ctx := context.Background()

	return &SqlConnection{
		driver,
		dsn,
		nil,
		ctx,
		configurator}
}

func (conn *SqlConnection) Open() error {
	var err error
	if conn.pool != nil {
		log.Println("[INFO] Trying to open an already open connection")
		return nil
	}
	log.Println("[DEBUG] - Opening connection in SQL Connection")
	conn.pool, err = sql.Open(conn.driver, conn.dsn)
	if err != nil {
		return err
	}

	return conn.configurator(conn.pool)
}

func (conn *SqlConnection) Close() error {
	if conn.pool == nil {
		log.Println("[INFO] Trying to close a closed connection")
		return nil
	}
	err := conn.pool.Close()
	if err != nil {
		conn.pool = nil
	}
	return err
}

func (conn *SqlConnection) Exec(query string, params []any) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(conn.ctx, 5*time.Second)
	defer cancel()

	return conn.pool.ExecContext(ctx, query, params...)
}

func (conn *SqlConnection) QueryOne(query string, params []any, mapper persistence.RowMapper) (interface{}, error) {
	ctx, cancel := context.WithTimeout(conn.ctx, 5*time.Second)
	defer cancel()

	row := conn.pool.QueryRowContext(ctx, query, params...)
	err := row.Err()
	if err != nil {

		return nil, err
	}

	return mapper(row)
}

func (conn *SqlConnection) Query(query string, params []any, mapper persistence.RowMapper) (persistence.RowsIterator, error) {
	rows, err := conn.pool.Query(query, params...)
	if err != nil {
		return nil, err
	}
	return NewSqlRows(rows, mapper), nil
}
