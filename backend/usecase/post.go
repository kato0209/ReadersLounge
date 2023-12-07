package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IPostUsecase interface {
	GetAllPosts(ctx echo.Context, posts *[]models.Post) error
	CreatePost(ctx echo.Context, post *models.Post) error
	DeletePost(ctx echo.Context, postID int) error
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
	for _, post := range *posts {
		if !utils.IsRemotePath(post.User.ProfileImage.FileName) {
			profileImage, err := pu.pr.LoadPostImage(ctx, post.User.ProfileImage.FileName)
			if err != nil {
				return errors.WithStack(err)
			}
			post.User.ProfileImage.EncodedImage = &profileImage
			fmt.Println(444444)
			fmt.Println(post.User.ProfileImage.EncodedImage)
		}
		if post.Image != nil && post.Image.FileName != nil {
			postImage, err := pu.pr.LoadPostImage(ctx, *post.Image.FileName)
			if err != nil {
				return errors.WithStack(err)
			}
			post.Image.EncodedImage = &postImage
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
