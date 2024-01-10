package usecase

import (
	"backend/repository"
)

type ICommentUsecase interface {
}

type commentUsecase struct {
	cmr repository.ICommentRepository
}

func NewCommentUsecase(cmr repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{cmr}
}
