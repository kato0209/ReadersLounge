package initialize

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

func Init(db *sqlx.DB) {

	booksGenreExists, err := checkExistsBooksGenreData(db)
	if err != nil {
		slog.Error("check exists books genre data error", "error", err)
	}

	if !booksGenreExists {

		res, err := fetchAllBooksGenre()
		if err != nil {
			slog.Error("fetch all books genre error", "error", err)
		}

		saveBooksGenreData(db, res)
	}
}

func checkExistsBooksGenreData(db *sqlx.DB) (bool, error) {
	query := `
		SELECT EXISTS (SELECT * FROM books_genres);
	`
	var exists bool
	err := db.QueryRowx(query).Scan(&exists)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return exists, nil
}

func fetchAllBooksGenre() ([]models.BooksGenre, error) {
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

func saveBooksGenreData(db *sqlx.DB, booksGenres []models.BooksGenre) error {

	tx, err := db.Beginx()
	if err != nil {
		slog.Error("tx begin error", "error", err)
	}
	for _, bookGenre := range booksGenres {
		query := `
			INSERT INTO books_genres (
				books_genre_id,
				books_genre_name,
				genre_level,
				parent_genre_id
			) VALUES (
				:books_genre_id,
				:books_genre_name,
				:genre_level,
				:parent_genre_id
			)
		`
		_, err := tx.NamedExec(query, bookGenre)
		if err != nil {
			_ = tx.Rollback()
			slog.Error("save books genre data error", "error", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("tx commit error", "error", err)
	}

	return nil
}
