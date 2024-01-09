package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ILikeRepository interface {
	CreatePostLike(ctx echo.Context, userID, postID int) error
	DeletePostLike(ctx echo.Context, postID, userID int) error
}

type likeRepository struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) ILikeRepository {
	return &likeRepository{db}
}

func (lr *likeRepository) CreatePostLike(ctx echo.Context, userID, postID int) error {
	query := "INSERT INTO post_likes (user_id, post_id) VALUES ($1, $2);"

	_, err := lr.db.ExecContext(ctx.Request().Context(), query, userID, postID)
	if err != nil {
		return err
	}

	return nil
}

func (lr *likeRepository) DeletePostLike(ctx echo.Context, postID, userID int) error {
	query := "DELETE FROM post_likes WHERE post_id = $1 AND user_id = $2;"

	_, err := lr.db.ExecContext(ctx.Request().Context(), query, postID, userID)
	if err != nil {
		return err
	}

	return nil
}
