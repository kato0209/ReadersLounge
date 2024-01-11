package usecase

import (
	"backend/models"
	"backend/repository"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ILikeUsecase interface {
	CreatePostLike(ctx echo.Context, userID, postID int) (models.PostLike, error)
	DeletePostLike(ctx echo.Context, postID, userID int) error
	CreateCommentLike(ctx echo.Context, userID, commentID int) (models.CommentLike, error)
	DeleteCommentLike(ctx echo.Context, commentID, userID int) error
}

type likeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &likeUsecase{lr}
}

func (lu *likeUsecase) CreatePostLike(ctx echo.Context, userID, postID int) (models.PostLike, error) {
	postLike := models.PostLike{}
	if err := lu.lr.CreatePostLike(ctx, userID, postID, &postLike); err != nil {
		fmt.Println(err)
		return models.PostLike{}, errors.WithStack(err)
	}

	return postLike, nil
}

func (lu *likeUsecase) DeletePostLike(ctx echo.Context, postID, userID int) error {
	if err := lu.lr.DeletePostLike(ctx, postID, userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (lu *likeUsecase) CreateCommentLike(ctx echo.Context, userID, commentID int) (models.CommentLike, error) {
	commentLike := models.CommentLike{}
	if err := lu.lr.CreateCommentLike(ctx, userID, commentID, &commentLike); err != nil {
		return models.CommentLike{}, errors.WithStack(err)
	}

	return commentLike, nil
}

func (lu *likeUsecase) DeleteCommentLike(ctx echo.Context, commentID, userID int) error {
	if err := lu.lr.DeleteCommentLike(ctx, commentID, userID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
