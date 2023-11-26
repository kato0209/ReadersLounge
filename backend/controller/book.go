package controller

import (
	"backend/models"
	"net/http"

	"backend/controller/openapi"

	"github.com/labstack/echo/v4"
)

func (s *Server) FetchBookData(ctx echo.Context, params openapi.FetchBookDataParams) error {
	books := []models.Book{}
	keyword := *params.Keyword
	booksGenreId := *params.BooksGenreId
	if err := s.bu.FetchBookData(ctx, &books, keyword, booksGenreId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, nil)
}
