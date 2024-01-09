package controller

import (
	"backend/controller/openapi"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreatePostLike(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	createPostLikeBody := openapi.CreatePostLikeReqBody{}
	if err := ctx.Bind(&createPostLikeBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := s.lu.CreatePostLike(ctx, userID, createPostLikeBody.PostId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (s *Server) DeletePostLike(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	deletePostLikeBody := openapi.DeletePostLikeReqBody{}
	if err := ctx.Bind(&deletePostLikeBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := s.lu.DeletePostLike(ctx, deletePostLikeBody.PostId, userID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
