package controller

import (
	"backend/usecase"
	"net/http"

	"backend/controller/openapi"

	"github.com/labstack/echo/v4"
)

type Server struct {
	uu usecase.IUserUsecase
	pu usecase.IPostUsecase
	bu usecase.IBookUsecase
}

func NewServer(uu usecase.IUserUsecase, pu usecase.IPostUsecase, bu usecase.IBookUsecase) *Server {
	return &Server{uu, pu, bu}
}

func (s *Server) Csrftoken(ctx echo.Context) error {
	token := ctx.Get("csrf").(string)
	csrfToken := openapi.ResCsrfToken{CsrfToken: &token}
	return ctx.JSON(http.StatusOK, csrfToken)
}
