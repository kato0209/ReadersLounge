package repository

import (
	"backend/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IPostRepository interface {
	GetAllPosts(ctx echo.Context, posts *[]models.Post) error
	CreatePost(ctx echo.Context, post *models.Post) error
}

type postRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) IPostRepository {
	return &postRepository{db}
}

func (pr *postRepository) GetAllPosts(ctx echo.Context, posts *[]models.Post) error {
	c := ctx.Request().Context()
	query := `
		SELECT
			p.post_id AS post_id,
			pd.content AS content,
			pd.rating AS rating,
			pd.image AS post_image,
			p.created_at AS created_at,
			u.user_id AS user_id,
			ud.name AS name,
			ud.profile_image AS profile_image,
			b.book_id AS book_id,
			b.ISBNcode AS ISBNcode,
			b.title AS title,
			b.author AS author,
			b.price AS price,
			b.publisher AS publisher,
			b.published_at AS published_at,
			b.image AS book_image,
			b.item_url AS item_url
		FROM 
			posts AS p
		INNER JOIN
			post_details AS pd ON p.post_id = pd.post_id
		INNER JOIN
			users AS u ON u.user_id = p.user_id
		INNER JOIN 
			user_details AS ud ON u.user_id = ud.user_id
		INNER JOIN
			books AS b ON p.book_id = b.book_id
	`
	rows, err := pr.db.QueryContext(c, query)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		book := models.Book{}
		user := models.User{}
		post.Book = book
		post.User = user

		err := rows.Scan(
			&post.PostID,
			&post.Content,
			&post.Rating,
			&post.Image,
			&post.CreatedAt,
			&post.User.UserID,
			&post.User.Name,
			&post.User.ProfileImage,
			&post.Book.BookID,
			&post.Book.ISBNcode,
			&post.Book.Title,
			&post.Book.Author,
			&post.Book.Price,
			&post.Book.Publisher,
			&post.Book.PublishedAt,
			&post.Book.Image,
			&post.Book.ItemURL,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		*posts = append(*posts, post)
	}
	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (pr *postRepository) CreatePost(ctx echo.Context, post *models.Post) error {
	c := ctx.Request().Context()
	tx, err := pr.db.BeginTxx(c, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(
		c,
		`
		INSERT INTO posts ( user_id, book_id ) VALUES ($1, $2) RETURNING post_id;
	`,
		post.User.UserID,
		post.Book.BookID,
	).Scan(&post.PostID)
	if err != nil {
		return errors.WithStack(err)
	}

	var postImage sql.NullString
	if post.Image != nil {
		postImage = sql.NullString{String: *post.Image, Valid: true}
	}
	_, err = tx.ExecContext(
		c,
		`
		INSERT INTO post_details ( post_id, content, rating, image ) VALUES ($1, $2, $3, $4);
	`,
		post.PostID,
		post.Content,
		post.Rating,
		postImage,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
