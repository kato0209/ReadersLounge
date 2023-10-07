package controller

import (
	"backend/controller/openapi"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Posts(ctx echo.Context) error {
	/*
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
	*/

	posts := []models.Post{}
	if err := s.pu.GetAllPosts(ctx, &posts); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	resPosts := []openapi.Post{}
	for _, post := range posts {
		resUser := openapi.User{
			UserId:       &post.User.UserID,
			Name:         &post.User.Name,
			ProfileImage: &post.User.ProfileImage,
		}
		fmt.Println(post.User.ProfileImage)
		resBook := openapi.Book{
			BookId:      &post.Book.BookID,
			ISBNcode:    &post.Book.ISBNcode,
			Author:      &post.Book.Author,
			Image:       &post.Book.Image,
			ItemUrl:     &post.Book.ItemURL,
			PublishedAt: &post.Book.PublishedAt,
			Publisher:   &post.Book.Publisher,
			Price:       &post.Book.Price,
			Title:       &post.Book.Title,
		}
		resPosts = append(resPosts, openapi.Post{
			PostId:    &post.PostID,
			Content:   &post.Content,
			Rating:    &post.Rating,
			Image:     &post.Image,
			CreatedAt: &post.CreatedAt,
			User:      &resUser,
			Book:      &resBook,
		})
	}
	fmt.Println(*resPosts[0].User.Name)
	fmt.Println(*resPosts[0].User.ProfileImage)

	return ctx.JSON(http.StatusCreated, resPosts)
}
