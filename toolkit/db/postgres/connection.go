package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgresConnection creates new DB connection
func NewPostgresConnection(dsn string) (*sqlx.DB, error) {
	return sqlx.Open("postgres", dsn)
}
