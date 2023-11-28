package repository

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IBookRepository interface {
	InsertBookData(ctx echo.Context, book *models.Book) error
	UpdateBookData(ctx echo.Context, book *models.Book) error
	CheckExistsBookDataByISBNcode(ctx echo.Context, ISBNcode string) (bool, error)
	FetchBookInfo(ctx echo.Context, ISBNcode string) (models.Book, error)
	FetchBookData(ctx echo.Context, books *[]models.Book, keyword, booksGenreID string) error
	FetchAllBooksGenre(ctx echo.Context) ([]models.BooksGenre, error)
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

func (br *bookRepository) FetchBookInfo(ctx echo.Context, ISBNcode string) (models.Book, error) {
	url := fmt.Sprintf("%s?format=json&applicationId=%s&isbn=%s&outOfStockFlag=1",
		os.Getenv("RAKUTEN_BOOKS_API_URL"),
		os.Getenv("RAKUTEN_APPLICATION_ID"),
		ISBNcode,
	)

	res := models.RakutenApiBooksResponse{}
	if err := utils.FetchApi(url, &res); err != nil {
		return models.Book{}, errors.WithStack(err)
	}

	book := models.Book{
		ISBNcode:    res.Items[0].Item.ISBNcode,
		Title:       res.Items[0].Item.Title,
		Author:      res.Items[0].Item.Author,
		Price:       res.Items[0].Item.Price,
		Publisher:   res.Items[0].Item.Publisher,
		PublishedAt: res.Items[0].Item.PublishedAt,
		ItemURL:     res.Items[0].Item.ItemURL,
		Image:       res.Items[0].Item.Image,
	}

	return book, nil
}

func (br *bookRepository) FetchBookData(ctx echo.Context, books *[]models.Book, keyword, booksGenreID string) error {
	baseUrl := fmt.Sprintf("%s?format=json&applicationId=%s&outOfStockFlag=1",
		os.Getenv("RAKUTEN_BOOKS_API_URL"),
		os.Getenv("RAKUTEN_APPLICATION_ID"),
	)

	var url string
	if keyword != "" && booksGenreID != "" {
		url = fmt.Sprintf("%s&title=%s&booksGenreId=%s", baseUrl, keyword, booksGenreID)
	} else if keyword != "" {
		url = fmt.Sprintf("%s&title=%s", baseUrl, keyword)
	} else if booksGenreID != "" {
		url = fmt.Sprintf("%s&booksGenreId=%s", baseUrl, booksGenreID)
	} else {
		url = baseUrl
	}

	res := models.RakutenApiBooksResponse{}
	if err := utils.FetchApi(url, &res); err != nil {
		return errors.WithStack(err)
	}

	for _, item := range res.Items {
		book := models.Book{
			ISBNcode:    item.Item.ISBNcode,
			Title:       item.Item.Title,
			Author:      item.Item.Author,
			Price:       item.Item.Price,
			Publisher:   item.Item.Publisher,
			PublishedAt: item.Item.PublishedAt,
			ItemURL:     item.Item.ItemURL,
			Image:       item.Item.Image,
		}
		*books = append(*books, book)
	}

	return nil
}

func (br *bookRepository) FetchAllBooksGenre(ctx echo.Context) ([]models.BooksGenre, error) {
	genreRootUrl := fmt.Sprintf("%s?applicationId=%s&booksGenreId=001",
		os.Getenv("RAKUTEN_BOOKS_GENRE_API_URL"),
		os.Getenv("RAKUTEN_APPLICATION_ID"),
	)

	res := models.RakutenApiBooksGenreResponse{}
	if err := utils.FetchApi(genreRootUrl, &res); err != nil {
		return []models.BooksGenre{}, errors.WithStack(err)
	}

	bookGenreStack := []models.BooksGenre{}
	bookGenreResults := []models.BooksGenre{}
	for _, child := range res.Children {
		rootBooksGenre := models.BooksGenre{
			BooksGenreID:   child.Child.BooksGenreID,
			BooksGenreName: child.Child.BooksGenreName,
			GenreLevel:     child.Child.GenreLevel,
			ParentGenreID:  "001",
		}

		bookGenreStack = append(bookGenreStack, rootBooksGenre)
		bookGenreResults = append(bookGenreResults, rootBooksGenre)
		for len(bookGenreStack) > 0 {
			genreRootUrl := fmt.Sprintf("%s?applicationId=%s&booksGenreId=%s",
				os.Getenv("RAKUTEN_BOOKS_GENRE_API_URL"),
				os.Getenv("RAKUTEN_APPLICATION_ID"),
				bookGenreStack[0].BooksGenreID,
			)
			res := models.RakutenApiBooksGenreResponse{}
			if err := utils.FetchApi(genreRootUrl, &res); err != nil {
				return []models.BooksGenre{}, errors.WithStack(err)
			}
			for _, child := range res.Children {
				bookGenre := models.BooksGenre{
					BooksGenreID:   child.Child.BooksGenreID,
					BooksGenreName: child.Child.BooksGenreName,
					GenreLevel:     child.Child.GenreLevel,
					ParentGenreID:  bookGenreStack[0].BooksGenreID,
				}
				bookGenreStack = append(bookGenreStack, bookGenre)
				bookGenreResults = append(bookGenreResults, bookGenre)
			}
			bookGenreStack = bookGenreStack[1:]
		}

	}

	return bookGenreResults, nil
}
