package repository

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ICommentRepository interface {
	GetCommentsByPostID(ctx echo.Context, postID int, comments *[]models.Comment) error
	CreateComment(ctx echo.Context, comment *models.Comment) error
}

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(ctx echo.Context, comment *models.Comment) error {
	c := ctx.Request().Context()
	tx, err := cr.db.BeginTxx(c, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO comments ( user_id, post_id ) VALUES ($1, $2) RETURNING comment_id, created_at;
	`
	err = tx.QueryRowContext(
		c,
		query,
		comment.User.UserID,
		comment.Post.PostID,
	).Scan(&comment.CommentID, &comment.CreatedAt)
	if err != nil {
		return errors.WithStack(err)
	}

	query = `
		INSERT INTO comment_details ( comment_id, content ) VALUES ($1, $2);
	`
	_, err = tx.ExecContext(
		c,
		query,
		comment.CommentID,
		comment.Content,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (cr *commentRepository) GetCommentsByPostID(ctx echo.Context, postID int, comments *[]models.Comment) error {
	query := `
		SELECT
			c.comment_id,
			c.user_id,
			c.post_id,
			cd.content,
			c.created_at,
			ud.name,
			ud.profile_image
		FROM
			comments c
		INNER JOIN
			comment_details cd ON c.comment_id = cd.comment_id
		INNER JOIN
			users u ON c.user_id = u.user_id
		INNER JOIN
			user_details ud ON u.user_id = ud.user_id
		WHERE
			c.post_id = $1
		ORDER BY
			c.created_at ASC;
	`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, postID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(
			&comment.CommentID,
			&comment.User.UserID,
			&comment.Post.PostID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.User.Name,
			&comment.User.ProfileImage.FileName,
		); err != nil {
			return errors.WithStack(err)
		}
		*comments = append(*comments, comment)
	}

	return nil
}
