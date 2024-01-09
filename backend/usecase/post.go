package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IPostUsecase interface {
	GetAllPosts(ctx echo.Context, posts *[]models.Post) error
	CreatePost(ctx echo.Context, post *models.Post) error
	DeletePost(ctx echo.Context, postID int) error
	GetLikedPostList(ctx echo.Context, userID int, posts *[]models.Post) error
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
	for i := range *posts {
		if !utils.IsRemotePath((*posts)[i].User.ProfileImage.FileName) {
			profileImage, err := utils.LoadImage(ctx, (*posts)[i].User.ProfileImage.FileName)
			if err != nil {
				return errors.WithStack(err)
			}
			(*posts)[i].User.ProfileImage.EncodedImage = &profileImage
		}
		if (*posts)[i].Image != nil && (*posts)[i].Image.FileName != nil {
			postImage, err := utils.LoadImage(ctx, *(*posts)[i].Image.FileName)
			if err != nil {
				return errors.WithStack(err)
			}
			(*posts)[i].Image.EncodedImage = &postImage
		}
	}
	return nil
}

func (pu *postUsecase) CreatePost(ctx echo.Context, post *models.Post) error {
	if post.Image != nil {
		if err := pu.pr.SavePostImage(ctx, post.Image); err != nil {
			return errors.WithStack(err)
		}
	}

	if err := pu.pr.CreatePost(ctx, post); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (pu *postUsecase) DeletePost(ctx echo.Context, postID int) error {
	if err := pu.pr.DeletePost(ctx, postID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (pu *postUsecase) GetLikedPostList(ctx echo.Context, userID int, posts *[]models.Post) error {
	if err := pu.pr.GetLikedPostList(ctx, userID, posts); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
