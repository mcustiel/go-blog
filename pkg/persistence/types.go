package persistence

type RowsIterator interface {
	Next() bool
	Get() (interface{}, error)
	Close() error
	Error() error
}

type Scanneable interface {
	Scan(dest ...any) error
}

type RowMapper func(s Scanneable) (interface{}, error)

type Connection[ER any] interface {
	Open() error
	Exec(query string, params []any) (ER, error)
	QueryOne(query string, params []any, mapper RowMapper) (interface{}, error)
	Query(query string, params []any, mapper RowMapper) (RowsIterator, error)
	Close() error
}

type ConnectionManager[ER any] interface {
	Open() error
	Close() error
	GetConnection() Connection[ER]
}
