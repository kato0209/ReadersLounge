package repository

import (
	"backend/models"
	"backend/utils"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IPostRepository interface {
	GetAllPosts(ctx echo.Context, posts *[]models.Post) error
	CreatePost(ctx echo.Context, post *models.Post) error
	SavePostImage(ctx echo.Context, image *models.PostImage) error
	DeletePost(ctx echo.Context, postID int) error
	GetLikedPostList(ctx echo.Context, userID int, posts *[]models.Post) error
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
			b.item_url AS item_url,
			pl.post_like_id,
			pl.user_id AS like_user_id
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
		LEFT JOIN
            post_likes AS pl ON p.post_id = pl.post_id
		ORDER BY p.created_at DESC;
	`
	rows, err := pr.db.QueryContext(c, query)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	postMap := make(map[int]*models.Post)
	var orderedPostIDs []int

	for rows.Next() {
		post := models.Post{}

		var postID int
		var fileName sql.NullString
		var likeID sql.NullInt64
		var likeUserID sql.NullInt64

		err := rows.Scan(
			&postID,
			&post.Content,
			&post.Rating,
			&fileName,
			&post.CreatedAt,
			&post.User.UserID,
			&post.User.Name,
			&post.User.ProfileImage.FileName,
			&post.Book.BookID,
			&post.Book.ISBNcode,
			&post.Book.Title,
			&post.Book.Author,
			&post.Book.Price,
			&post.Book.Publisher,
			&post.Book.PublishedAt,
			&post.Book.Image,
			&post.Book.ItemURL,
			&likeID,
			&likeUserID,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		if fileName.Valid {
			post.Image = &models.PostImage{FileName: &fileName.String}
		}

		if existingPost, exists := postMap[postID]; exists {
			if likeID.Valid {
				existingPost.Like = append(existingPost.Like, models.PostLike{
					PostLikeID: int(likeID.Int64),
					User:       models.User{UserID: int(likeUserID.Int64)},
				})
			}
		} else {
			post.PostID = postID
			postMap[postID] = &post
			orderedPostIDs = append(orderedPostIDs, postID)

			if likeID.Valid {
				post.Like = append(post.Like, models.PostLike{
					PostLikeID: int(likeID.Int64),
					User:       models.User{UserID: int(likeUserID.Int64)},
				})
			}
		}
	}

	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	for _, id := range orderedPostIDs {
		*posts = append(*posts, *postMap[id])
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

	var postImageStr sql.NullString
	if post.Image != nil && post.Image.FileName != nil {
		postImageStr = sql.NullString{String: *post.Image.FileName, Valid: true}
	}
	_, err = tx.ExecContext(
		c,
		`
		INSERT INTO post_details ( post_id, content, rating, image ) VALUES ($1, $2, $3, $4);
	`,
		post.PostID,
		post.Content,
		post.Rating,
		postImageStr,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (pr *postRepository) SavePostImage(ctx echo.Context, image *models.PostImage) error {

	err := utils.SaveImage(ctx, *image.FileName, image.Source)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (pr *postRepository) DeletePost(ctx echo.Context, postID int) error {
	c := ctx.Request().Context()
	_, err := pr.db.ExecContext(
		c,
		`
        DELETE FROM posts WHERE post_id = $1;
        `,
		postID,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (pr *postRepository) GetLikedPostList(ctx echo.Context, userID int, posts *[]models.Post) error {
	query := `
			SELECT post_id
			FROM post_likes
			WHERE user_id = $1;
		`

	rows, err := pr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(
			&post.PostID,
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
