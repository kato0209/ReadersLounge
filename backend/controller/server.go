package controller

import (
	"backend/models/chat"
	"backend/usecase"
	"net/http"

	"backend/controller/openapi"

	"github.com/labstack/echo/v4"
)

type Server struct {
	uu  usecase.IUserUsecase
	pu  usecase.IPostUsecase
	bu  usecase.IBookUsecase
	cu  usecase.IChatUsecase
	cnu usecase.IConnectionUsecase
	lu  usecase.ILikeUsecase
	cmu usecase.ICommentUsecase
	hub *chat.Hub
}

func NewServer(uu usecase.IUserUsecase, pu usecase.IPostUsecase, bu usecase.IBookUsecase, cu usecase.IChatUsecase, cnu usecase.IConnectionUsecase, lu usecase.ILikeUsecase, cmu usecase.ICommentUsecase, hub chat.Hub) *Server {
	return &Server{uu, pu, bu, cu, cnu, lu, cmu, &hub}
}

func (s *Server) Csrftoken(ctx echo.Context) error {
	token := ctx.Get("csrf").(string)
	csrfToken := openapi.ResCsrfToken{CsrfToken: &token}
	return ctx.JSON(http.StatusOK, csrfToken)
}

func (s *Server) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
}
