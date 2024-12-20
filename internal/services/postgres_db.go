package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(connectionString string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresDatabase{db: db}, nil
}
