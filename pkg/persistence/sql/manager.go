package sql

import (
	gosql "database/sql"
	"log"

	"github.com/mcustiel/go-blog/pkg/persistence"
)

// type ConnectionData struct {
// 	Driver string
// 	DSN    string
// }

type DefaultConnectionManager struct {
	connection persistence.Connection[gosql.Result]
}

func NewDefaultConnectionManager() *DefaultConnectionManager {
	// TODO: bring the connection data from a configuration or pass it as parameter
	conn := NewSqlConnection(
		"sqlite3",
		"file:_resources/blog.db?cache=shared",
		func(pool *gosql.DB) error {
			pool.SetConnMaxLifetime(0)
			pool.SetMaxIdleConns(1)
			pool.SetMaxOpenConns(1)
			return nil
		})
	return &DefaultConnectionManager{conn}
}

func (dcm *DefaultConnectionManager) Open() error {
	log.Println("[DEBUG] - Opening connection in DB Manager")
	return dcm.connection.Open()
}

func (dcm *DefaultConnectionManager) Close() error {
	return dcm.connection.Close()
}

func (dcm *DefaultConnectionManager) GetConnection() persistence.Connection[gosql.Result] {
	return dcm.connection
}
