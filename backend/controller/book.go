package controller

import (
	"backend/models"
	"net/http"

	"backend/controller/openapi"

	"github.com/labstack/echo/v4"
)

func (s *Server) FetchBookData(ctx echo.Context, params openapi.FetchBookDataParams) error {
	books := []models.Book{}
	var keyword, booksGenreId string
	if params.Keyword != nil {
		keyword = *params.Keyword
	}
	if params.BooksGenreId != nil {
		booksGenreId = *params.BooksGenreId
	}

	if err := s.bu.FetchBookData(ctx, &books, keyword, booksGenreId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resBooks := []openapi.Book{}
	for _, book := range books {
		book := openapi.Book{
			ISBNcode:    book.ISBNcode,
			Author:      book.Author,
			Image:       book.Image,
			ItemUrl:     book.ItemURL,
			PublishedAt: book.PublishedAt,
			Publisher:   book.Publisher,
			Price:       book.Price,
			Title:       book.Title,
		}

		resBooks = append(resBooks, book)
	}

	return ctx.JSON(http.StatusOK, resBooks)
}

func convertToNewStructure(node models.BooksGenreNode) openapi.BookGenreNode {
	newNode := openapi.BookGenreNode{
		BooksGenreId:   node.CurrentGenre.BooksGenreID,
		BooksGenreName: node.CurrentGenre.BooksGenreName,
		GenreLevel:     node.CurrentGenre.GenreLevel,
		Id:             node.CurrentGenre.ID,
		ParentGenreId:  node.CurrentGenre.ParentGenreID,
	}

	for _, child := range node.Children {
		newNode.Children = append(newNode.Children, convertToNewStructure(child))
	}

	return newNode
}

func (s *Server) GetBooksGenres(ctx echo.Context) error {

	bookGenresNodes, err := s.bu.GetBooksGenres(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resBookGenreNodes := []openapi.BookGenreNode{}
	for _, bookGenreNode := range bookGenresNodes {
		resBookGenreNodes = append(resBookGenreNodes, convertToNewStructure(bookGenreNode))
	}

	return ctx.JSON(http.StatusOK, resBookGenreNodes)
}
