package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ICommentUsecase interface {
	GetCommentsByPostID(ctx echo.Context, postID int) ([]models.Comment, error)
	GetLikedCommentList(ctx echo.Context, userID int) ([]models.Comment, error)
	CreateComment(ctx echo.Context, comment *models.Comment) error
	DeleteComment(ctx echo.Context, commentID int) error
}

type commentUsecase struct {
	cmr repository.ICommentRepository
	ur  repository.IUserRepository
}

func NewCommentUsecase(cmr repository.ICommentRepository, ur repository.IUserRepository) ICommentUsecase {
	return &commentUsecase{cmr, ur}
}

func (cu *commentUsecase) GetCommentsByPostID(ctx echo.Context, postID int) ([]models.Comment, error) {
	comments := []models.Comment{}
	err := cu.cmr.GetCommentsByPostID(ctx, postID, &comments)
	if err != nil {
		return nil, err
	}

	for i := range comments {
		if !utils.IsRemotePath(comments[i].User.ProfileImage.FileName) {
			profileImage, err := utils.LoadImage(ctx, comments[i].User.ProfileImage.FileName)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			comments[i].User.ProfileImage.EncodedImage = &profileImage
		}
		comments[i].User.ProfileImage = comments[i].User.ProfileImage
	}

	return comments, nil
}

func (cu *commentUsecase) GetLikedCommentList(ctx echo.Context, userID int) ([]models.Comment, error) {
	comments := []models.Comment{}
	err := cu.cmr.GetLikedCommentList(ctx, userID, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cu *commentUsecase) CreateComment(ctx echo.Context, comment *models.Comment) error {
	if err := cu.cmr.CreateComment(ctx, comment); err != nil {
		return errors.WithStack(err)
	}

	user := models.User{}
	if err := cu.ur.GetUserByUserID(ctx, &user, comment.User.UserID); err != nil {
		return errors.WithStack(err)
	}

	comment.User.Name = user.Name

	if !utils.IsRemotePath(user.ProfileImage.FileName) {
		profileImage, err := utils.LoadImage(ctx, user.ProfileImage.FileName)
		if err != nil {
			return errors.WithStack(err)
		}
		user.ProfileImage.EncodedImage = &profileImage
	}
	comment.User.ProfileImage = user.ProfileImage

	return nil
}

func (cu *commentUsecase) DeleteComment(ctx echo.Context, commentID int) error {
	if err := cu.cmr.DeleteCommentByCommentID(ctx, commentID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
