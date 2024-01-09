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

	postLike, err := s.lu.CreatePostLike(ctx, userID, createPostLikeBody.PostId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resPostLike := openapi.PostLike{
		PostLikeId: postLike.PostLikeID,
	}

	return ctx.JSON(http.StatusCreated, resPostLike)
}

func (s *Server) DeletePostLike(ctx echo.Context, postId int) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := s.lu.DeletePostLike(ctx, postId, userID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
