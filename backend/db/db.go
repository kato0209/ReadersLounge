package db

import (
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

func Open() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
