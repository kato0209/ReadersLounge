package usecase

import (
	"backend/models"
	"backend/repository"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IBookUsecase interface {
	RegisterBook(ctx echo.Context, ISBNcode string) (models.Book, error)
	FetchBookData(ctx echo.Context, books *[]models.Book, keyword, booksGenreID string) error
	FetchBookGenres(ctx echo.Context) ([]models.BooksGenre, error)
}

type bookUsecase struct {
	br repository.IBookRepository
}

func NewBookUsecase(br repository.IBookRepository) IBookUsecase {
	return &bookUsecase{br}
}

func (bu *bookUsecase) RegisterBook(ctx echo.Context, ISBNcode string) (models.Book, error) {
	book, err := bu.br.FetchBookInfo(ctx, ISBNcode)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}

	bookExists, err := bu.br.CheckExistsBookDataByISBNcode(ctx, ISBNcode)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}

	if bookExists {
		if err := bu.br.UpdateBookData(ctx, &book); err != nil {
			return models.Book{}, errors.WithStack(err)
		}
	} else {
		if err := bu.br.InsertBookData(ctx, &book); err != nil {
			return models.Book{}, errors.WithStack(err)
		}
	}

	return book, nil
}

func (bu *bookUsecase) FetchBookData(ctx echo.Context, books *[]models.Book, keyword, booksGenreID string) error {
	if err := bu.br.FetchBookData(ctx, books, keyword, booksGenreID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (bu *bookUsecase) FetchBookGenres(ctx echo.Context) ([]models.BooksGenre, error) {
	bookGenres, err := bu.br.FetchAllBooksGenre(ctx)
	if err != nil {
		return []models.BooksGenre{}, errors.WithStack(err)
	}

	return bookGenres, nil
}
