package db

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

func Tx(ctx context.Context, db *sqlx.DB, f func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if p := recover(); p != nil {
			pErr := fmt.Errorf("panic error: %v", p)
			if rErr := tx.Rollback(); rErr != nil {
				err = errors.CombineErrors(pErr, rErr)
			} else {
				err = pErr
			}
		} else if err != nil {
			if rErr := tx.Rollback(); rErr != nil {
				err = errors.CombineErrors(err, rErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	return f(tx)
}
