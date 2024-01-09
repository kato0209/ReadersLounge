package usecase

import (
	"backend/repository"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ILikeUsecase interface {
	CreatePostLike(ctx echo.Context, userID, postID int) error
	DeletePostLike(ctx echo.Context, postID, userID int) error
}

type likeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &likeUsecase{lr}
}

func (lu *likeUsecase) CreatePostLike(ctx echo.Context, userID, postID int) error {
	if err := lu.lr.CreatePostLike(ctx, userID, postID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lu *likeUsecase) DeletePostLike(ctx echo.Context, postID, userID int) error {
	if err := lu.lr.DeletePostLike(ctx, postID, userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
