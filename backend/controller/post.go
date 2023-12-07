package controller

import (
	"backend/controller/openapi"
	"backend/models"
	"backend/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetPosts(ctx echo.Context) error {
	posts := []models.Post{}
	if err := s.pu.GetAllPosts(ctx, &posts); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	resPosts := []openapi.Post{}
	for _, post := range posts {

		fmt.Println(*post.User.ProfileImage.EncodedImage)
		profileImage := post.User.ProfileImage.ClassifyPathType()

		resUser := openapi.User{
			UserId:       post.User.UserID,
			Name:         post.User.Name,
			ProfileImage: profileImage,
		}
		resBook := openapi.Book{
			BookId:      post.Book.BookID,
			ISBNcode:    post.Book.ISBNcode,
			Author:      post.Book.Author,
			Image:       post.Book.Image,
			ItemUrl:     post.Book.ItemURL,
			PublishedAt: post.Book.PublishedAt,
			Publisher:   post.Book.Publisher,
			Price:       post.Book.Price,
			Title:       post.Book.Title,
		}
		var encodedImage *string
		if post.Image != nil {
			encodedImage = post.Image.EncodedImage
		}
		formattedTime := post.CreatedAt.Format("2006-01-02 15:04")
		resPosts = append(resPosts, openapi.Post{
			PostId:    post.PostID,
			Content:   post.Content,
			Rating:    post.Rating,
			Image:     encodedImage,
			CreatedAt: formattedTime,
			User:      resUser,
			Book:      resBook,
		})
	}

	return ctx.JSON(http.StatusOK, resPosts)
}

func (s *Server) CreatePost(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	rating, err := strconv.Atoi(form.Value["rating"][0])
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	reqCreatePostBody := openapi.ReqCreatePostBody{
		ISBNcode: form.Value["ISBNcode"][0],
		Content:  form.Value["content"][0],
		Rating:   rating,
	}

	book, err := s.bu.RegisterBook(ctx, reqCreatePostBody.ISBNcode)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	post := models.Post{
		Content: reqCreatePostBody.Content,
		Rating:  reqCreatePostBody.Rating,
		User: models.User{
			UserID: userID,
		},
		Book: book,
	}

	file, err := ctx.FormFile("image")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		fileModel := strings.Split(file.Filename, ".")
		fileName := fileModel[0]
		extension := fileModel[1]

		if extension == "jpeg" || extension == "png" {

			data, err := io.ReadAll(src)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}

			generatedFileName := fmt.Sprintf("%s_%s.%s", fileName, uuid.New().String(), extension)

			image := models.PostImage{
				Source:   &data,
				FileName: &generatedFileName,
			}
			post.Image = &image
		} else {
			return ctx.JSON(http.StatusBadRequest, "Unsupported file type")
		}

	} else if err != http.ErrMissingFile {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := s.pu.CreatePost(ctx, &post); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (s *Server) DeletePost(ctx echo.Context, postID int) error {
	if err := s.pu.DeletePost(ctx, postID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
