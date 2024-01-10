package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetCommentsByPostID(ctx echo.Context, postId int) error {
	return ctx.JSON(http.StatusOK, nil)
}
