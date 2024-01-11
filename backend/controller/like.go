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

func (s *Server) CreateCommentLike(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	createCommentLikeBody := openapi.CreateCommentLikeReqBody{}
	if err := ctx.Bind(&createCommentLikeBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	commentLike, err := s.lu.CreateCommentLike(ctx, userID, createCommentLikeBody.CommentId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resCommentLike := openapi.CommentLike{
		CommentLikeId: commentLike.CommentLikeID,
	}

	return ctx.JSON(http.StatusCreated, resCommentLike)
}

func (s *Server) DeleteCommentLike(ctx echo.Context, commentId int) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := s.lu.DeleteCommentLike(ctx, commentId, userID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
