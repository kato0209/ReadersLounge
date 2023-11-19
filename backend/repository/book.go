package repository

import (
	"backend/models"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IBookRepository interface {
	InsertBookData(ctx echo.Context, book *models.Book) error
	UpdateBookData(ctx echo.Context, book *models.Book) error
	CheckExistsBookDataByISBNcode(ctx echo.Context, ISBNcode string) (bool, error)
	//GetBookByISBNcode(ctx echo.Context, ISBNcode string) (models.Book, error)
}

type bookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) InsertBookData(ctx echo.Context, book *models.Book) error {
	c := ctx.Request().Context()
	query := `
		INSERT INTO books (
			ISBNcode,
			title,
			author,
			price,
			publisher,
			published_at,
			image,
			item_url
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING book_id;
	`
	err := br.db.QueryRowxContext(
		c,
		query,
		book.ISBNcode,
		book.Title,
		book.Author,
		book.Price,
		book.Publisher,
		book.PublishedAt,
		book.Image,
		book.ItemURL,
	).Scan(&book.BookID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (br *bookRepository) UpdateBookData(ctx echo.Context, book *models.Book) error {
	c := ctx.Request().Context()
	query := `
		UPDATE
			books
		SET
			title = $1,
			author = $2,
			price = $3,
			publisher = $4,
			published_at = $5,
			image = $6,
			item_url = $7,
			updated_at = $8
		WHERE
			ISBNcode = $9
		RETURNING book_id;
	`
	err := br.db.QueryRowxContext(
		c,
		query,
		book.Title,
		book.Author,
		book.Price,
		book.Publisher,
		book.PublishedAt,
		book.Image,
		book.ItemURL,
		time.Now(),
		book.ISBNcode,
	).Scan(&book.BookID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (br *bookRepository) CheckExistsBookDataByISBNcode(ctx echo.Context, ISBNcode string) (bool, error) {
	c := ctx.Request().Context()
	query := `
		SELECT
			EXISTS (
				SELECT
					*
				FROM
					books
				WHERE
					ISBNcode = $1
			)
	`
	var exists bool
	err := br.db.QueryRowxContext(c, query, ISBNcode).Scan(&exists)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return exists, nil
}

/*
func (br *bookRepository) GetBookByISBNcode(ctx echo.Context, ISBNcode string) (models.Book, error) {
	c := ctx.Request().Context()
	query := `
		SELECT
			book_id,
			ISBNcode,
			title,
			author,
			price,
			publisher,
			published_at,
			image,
			item_url
		FROM
			books
		WHERE
			ISBNcode = $1
	`
	book := models.Book{}
	err := br.db.QueryRowxContext(c, query, ISBNcode).StructScan(&book)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}
	return book, nil
}
*/
