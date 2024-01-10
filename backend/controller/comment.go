package controller

import (
	"backend/controller/openapi"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetCommentsByPostID(ctx echo.Context, postId int) error {
	comments, err := s.cmu.GetCommentsByPostID(ctx, postId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	resComments := []openapi.Comment{}
	for _, comment := range comments {
		resUser := openapi.User{
			UserId:       comment.User.UserID,
			Name:         comment.User.Name,
			ProfileImage: comment.User.ProfileImage.ClassifyPathType(),
		}

		resComment := openapi.Comment{
			CommentId: comment.CommentID,
			User:      resUser,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04"),
		}
		resComments = append(resComments, resComment)
	}
	return ctx.JSON(http.StatusOK, resComments)
}

func (s *Server) CreateComment(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	reqCreateCommentBody := openapi.ReqCreateCommentBody{}
	if err := ctx.Bind(&reqCreateCommentBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	comment := models.Comment{
		Post:    models.Post{PostID: reqCreateCommentBody.PostId},
		Content: reqCreateCommentBody.Content,
		User:    models.User{UserID: userID},
	}
	if err := s.cmu.CreateComment(ctx, &comment); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       comment.User.UserID,
		Name:         comment.User.Name,
		ProfileImage: comment.User.ProfileImage.ClassifyPathType(),
	}

	resComment := openapi.Comment{
		CommentId: comment.CommentID,
		User:      resUser,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04"),
	}
	return ctx.JSON(http.StatusCreated, resComment)
}
