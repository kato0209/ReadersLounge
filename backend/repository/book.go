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
	GetAllBooksGenres(ctx echo.Context, bookGenres *[]models.BooksGenreNode) error
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

func buildTree(genres []models.BooksGenre) []models.BooksGenreNode {
	childrenMap := make(map[string][]models.BooksGenreNode)
	for _, genre := range genres {
		node := models.BooksGenreNode{CurrentGenre: genre}
		parentID := genre.ParentGenreID
		childrenMap[parentID] = append(childrenMap[parentID], node)
	}

	var buildTree func(parentID string) []models.BooksGenreNode
	buildTree = func(parentID string) []models.BooksGenreNode {
		children := childrenMap[parentID]
		for i, child := range children {
			children[i].Children = buildTree(child.CurrentGenre.BooksGenreID)
		}
		return children
	}

	return buildTree("001")
}

func (br *bookRepository) GetAllBooksGenres(ctx echo.Context, booksGenreNode *[]models.BooksGenreNode) error {
	c := ctx.Request().Context()
	query := `
		SELECT
			id,
			books_genre_id,
			books_genre_name,
			genre_level,
			parent_genre_id
		FROM
			books_genres;
	`

	rows, err := br.db.QueryContext(c, query)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	bookGenres := []models.BooksGenre{}
	for rows.Next() {
		bookGenre := models.BooksGenre{}
		err := rows.Scan(
			&bookGenre.ID,
			&bookGenre.BooksGenreID,
			&bookGenre.BooksGenreName,
			&bookGenre.GenreLevel,
			&bookGenre.ParentGenreID,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		bookGenres = append(bookGenres, bookGenre)
	}
	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	*booksGenreNode = buildTree(bookGenres)

	return nil
}
