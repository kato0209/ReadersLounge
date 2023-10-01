package repository

import (
	"backend/db"
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IUserRepository interface {
	CreateUser(ctx echo.Context, user *models.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(ctx echo.Context, user *models.User) error {
	c := ctx.Request().Context()
	if err := db.Tx(c, ur.db, func(tx *sqlx.Tx) error {
		"""
		err := tx.QueryRowContext(c, `
		INSERT INTO users (name, book_id)
		VALUES ($1, $2)
		RETURNING comment_id;
		`,
			authorID,
			bookID,
		).Scan(&CommentID)
		if err != nil {
			return errors.WithStack(err)
		}
		"""
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
