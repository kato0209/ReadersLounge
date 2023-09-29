package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Signup(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Signup successful")
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
