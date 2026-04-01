package database

import "github.com/jmoiron/sqlx"

type Database struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (db Database) Initialize() error {
	return nil
}
