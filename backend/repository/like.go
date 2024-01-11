package repository

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ILikeRepository interface {
	CreatePostLike(ctx echo.Context, userID, postID int, postLike *models.PostLike) error
	DeletePostLike(ctx echo.Context, postID, userID int) error
	CreateCommentLike(ctx echo.Context, userID, commentID int, commentLike *models.CommentLike) error
	DeleteCommentLike(ctx echo.Context, commentID, userID int) error
}

type likeRepository struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) ILikeRepository {
	return &likeRepository{db}
}

func (lr *likeRepository) CreatePostLike(ctx echo.Context, userID, postID int, postLike *models.PostLike) error {
	query := "INSERT INTO post_likes (user_id, post_id) VALUES ($1, $2) RETURNING post_like_id;"

	err := lr.db.QueryRowContext(ctx.Request().Context(), query, userID, postID).Scan(&postLike.PostLikeID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lr *likeRepository) DeletePostLike(ctx echo.Context, postID, userID int) error {
	query := "DELETE FROM post_likes WHERE post_id = $1 AND user_id = $2;"

	_, err := lr.db.ExecContext(ctx.Request().Context(), query, postID, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lr *likeRepository) CreateCommentLike(ctx echo.Context, userID, commentID int, commentLike *models.CommentLike) error {
	query := "INSERT INTO comment_likes (user_id, comment_id) VALUES ($1, $2) RETURNING comment_like_id;"

	err := lr.db.QueryRowContext(ctx.Request().Context(), query, userID, commentID).Scan(&commentLike.CommentLikeID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lr *likeRepository) DeleteCommentLike(ctx echo.Context, commentID, userID int) error {
	query := "DELETE FROM comment_likes WHERE comment_id = $1 AND user_id = $2;"

	_, err := lr.db.ExecContext(ctx.Request().Context(), query, commentID, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
