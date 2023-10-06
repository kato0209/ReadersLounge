package usecase

import (
	"backend/models"
	"backend/repository"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IPostUsecase interface {
	GetAllPosts(ctx echo.Context, posts *[]models.Post) error
}

type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts(ctx echo.Context, posts *[]models.Post) error {

	if err := pu.pr.GetAllPosts(ctx, posts); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
