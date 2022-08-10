package sql

import (
	"database/sql"

	"github.com/mcustiel/go-blog/pkg/persistence"
)

type SqlRows struct {
	rows   *sql.Rows
	mapper persistence.RowMapper
}

func NewSqlRows(rows *sql.Rows, mapper persistence.RowMapper) *SqlRows {
	return &SqlRows{rows, mapper}
}

func (rows *SqlRows) Next() bool {
	return rows.rows.Next()
}

func (rows *SqlRows) Get() (interface{}, error) {
	return rows.mapper(rows.rows)
}

func (rows *SqlRows) Error() error {
	return rows.rows.Err()
}

func (rows *SqlRows) Close() error {
	return rows.rows.Close()
}
