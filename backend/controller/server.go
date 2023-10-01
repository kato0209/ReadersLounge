package controller

import (
	"backend/usecase"
	"net/http"
)

type Server struct {
	uu usecase.IUserUsecase
}

func NewServer(uu usecase.IUserUsecase) *Server {
	return &Server{uu}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
