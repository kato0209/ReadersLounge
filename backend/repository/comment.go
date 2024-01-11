package repository

import (
	"backend/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ICommentRepository interface {
	CreateComment(ctx echo.Context, comment *models.Comment) error
	GetCommentsByPostID(ctx echo.Context, postID int, comments *[]models.Comment) error
	GetLikedCommentList(ctx echo.Context, userID int, comments *[]models.Comment) error
	DeleteCommentByCommentID(ctx echo.Context, commentID int) error
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
			ud.profile_image,
			cl.comment_like_id,
			cl.user_id
		FROM
			comments c
		INNER JOIN
			comment_details cd ON c.comment_id = cd.comment_id
		INNER JOIN
			users u ON c.user_id = u.user_id
		INNER JOIN
			user_details ud ON u.user_id = ud.user_id
		LEFT JOIN
			comment_likes cl ON c.comment_id = cl.comment_id
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

	commentMap := make(map[int]*models.Comment)
	var orderedCommentIDs []int
	for rows.Next() {

		var commentID int
		var likeID sql.NullInt64
		var likeUserID sql.NullInt64
		comment := models.Comment{}
		if err := rows.Scan(
			&commentID,
			&comment.User.UserID,
			&comment.Post.PostID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.User.Name,
			&comment.User.ProfileImage.FileName,
			&likeID,
			&likeUserID,
		); err != nil {
			return errors.WithStack(err)
		}

		if existingComment, exists := commentMap[commentID]; exists {
			if likeID.Valid {
				existingComment.Likes = append(existingComment.Likes, models.CommentLike{
					CommentLikeID: int(likeID.Int64),
					User:          models.User{UserID: int(likeUserID.Int64)},
				})
			}
		} else {
			comment.CommentID = commentID
			commentMap[commentID] = &comment
			orderedCommentIDs = append(orderedCommentIDs, commentID)

			if likeID.Valid {
				comment.Likes = append(comment.Likes, models.CommentLike{
					CommentLikeID: int(likeID.Int64),
					User:          models.User{UserID: int(likeUserID.Int64)},
				})
			}
		}
	}

	for _, id := range orderedCommentIDs {
		*comments = append(*comments, *commentMap[id])
	}

	return nil
}

func (cr *commentRepository) GetLikedCommentList(ctx echo.Context, userID int, comments *[]models.Comment) error {
	query := `
		SELECT
			c.comment_id
		FROM
			comment_likes c
		WHERE
			c.user_id = $1;
	`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(
			&comment.CommentID,
		); err != nil {
			return errors.WithStack(err)
		}
		*comments = append(*comments, comment)
	}

	return nil
}

func (cr *commentRepository) DeleteCommentByCommentID(ctx echo.Context, commentID int) error {
	query := `
		DELETE FROM comments WHERE comment_id = $1;
	`
	_, err := cr.db.ExecContext(ctx.Request().Context(), query, commentID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
