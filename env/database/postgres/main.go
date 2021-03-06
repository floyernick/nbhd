package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"nbhd/env/config"
	"nbhd/env/database"
)

type Database struct {
	db *sql.DB
}

func NewPostgresDatabase(config config.DatabaseConfig) (database.Database, error) {

	pool, err := sql.Open("postgres", config.Url)

	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(config.ConnLifetime)
	pool.SetMaxOpenConns(config.OpenConns)
	pool.SetMaxIdleConns(config.IdleConns)

	if err := pool.Ping(); err != nil {
		return nil, err
	}

	db := Database{pool}

	return db, nil

}
