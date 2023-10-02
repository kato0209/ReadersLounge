package controller

import (
	"backend/usecase"
	"net/http"

	"backend/controller/openapi"

	"github.com/labstack/echo/v4"
)

type Server struct {
	uu usecase.IUserUsecase
}

func NewServer(uu usecase.IUserUsecase) *Server {
	return &Server{uu}
}

func (s *Server) Csrftoken(ctx echo.Context) error {
	token := ctx.Get("csrf").(string)
	csrfToken := openapi.ResCsrfToken{CsrfToken: &token}
	return ctx.JSON(http.StatusOK, csrfToken)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
