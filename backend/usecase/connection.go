package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IConnectionUsecase interface {
	CreateConnection(ctx echo.Context, followerID, followingID int) error
	DeleteConnection(ctx echo.Context, connectionId int) error
	GetFollowingConnections(ctx echo.Context, userID int) ([]models.Connection, error)
	GetFollowerConnections(ctx echo.Context, userID int) ([]models.Connection, error)
}

type connectionUsecase struct {
	cr repository.IConnectionRepository
}

func NewConnectionUsecase(br repository.IConnectionRepository) IConnectionUsecase {
	return &connectionUsecase{br}
}

func (cu *connectionUsecase) CreateConnection(ctx echo.Context, followerID, followingID int) error {
	if err := cu.cr.CreateConnection(ctx, followerID, followingID); err != nil {
		return err
	}

	return nil
}

func (cu *connectionUsecase) DeleteConnection(ctx echo.Context, connectionId int) error {
	if err := cu.cr.DeleteConnection(ctx, connectionId); err != nil {
		return err
	}

	return nil
}

func (cu *connectionUsecase) GetFollowingConnections(ctx echo.Context, userID int) ([]models.Connection, error) {
	followingConnections := []models.Connection{}
	err := cu.cr.GetFollowingConnections(ctx, userID, &followingConnections)
	if err != nil {
		return nil, err
	}
	for i := range followingConnections {
		if !utils.IsRemotePath((followingConnections)[i].Following.ProfileImage.FileName) {
			profileImage, err := utils.LoadImage(ctx, (followingConnections)[i].Following.ProfileImage.FileName)
			if err != nil {
				return []models.Connection{}, errors.WithStack(err)
			}
			(followingConnections)[i].Following.ProfileImage.EncodedImage = &profileImage
		}
	}

	return followingConnections, nil
}

func (cu *connectionUsecase) GetFollowerConnections(ctx echo.Context, userID int) ([]models.Connection, error) {
	followerConnections := []models.Connection{}
	err := cu.cr.GetFollowerConnections(ctx, userID, &followerConnections)
	if err != nil {
		return nil, err
	}
	for i := range followerConnections {
		if !utils.IsRemotePath((followerConnections)[i].Follower.ProfileImage.FileName) {
			profileImage, err := utils.LoadImage(ctx, (followerConnections)[i].Follower.ProfileImage.FileName)
			if err != nil {
				return []models.Connection{}, errors.WithStack(err)
			}
			(followerConnections)[i].Follower.ProfileImage.EncodedImage = &profileImage
		}
	}

	return followerConnections, nil
}
