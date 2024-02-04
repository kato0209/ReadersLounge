package db

import (
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

func Open() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
		os.Getenv("PGSSLMODE"),
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
