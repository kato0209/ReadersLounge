package usecase

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IBookUsecase interface {
	RegisterBook(ctx echo.Context, ISBNcode string) (models.Book, error)
}

type bookUsecase struct {
	br repository.IBookRepository
}

func NewBookUsecase(br repository.IBookRepository) IBookUsecase {
	return &bookUsecase{br}
}

func fetchBookInfo(ISBNcode string) (models.Book, error) {
	url := fmt.Sprintf("%s?format=%s&applicationId=%s&isbn=%s",
		os.Getenv("RAKUTEN_BOOKS_API_URL"),
		"json",
		os.Getenv("RAKUTEN_APPLICATION_ID"),
		ISBNcode,
	)
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}

	res := models.RakutenBooksApiResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
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

func (bu *bookUsecase) RegisterBook(ctx echo.Context, ISBNcode string) (models.Book, error) {
	book, err := fetchBookInfo(ISBNcode)
	if err != nil {
		return models.Book{}, errors.WithStack(err)
	}

	bookExists, err := bu.br.CheckExistsBookData(ctx, ISBNcode)
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
