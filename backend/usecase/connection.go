package usecase

import (
	"backend/models"
	"backend/repository"

	"github.com/labstack/echo/v4"
)

type IConnectionUsecase interface {
	CreateConnection(ctx echo.Context, followerID, followingID int) error
	DeleteConnection(ctx echo.Context, connectionId int) error
	GetFollowingList(ctx echo.Context, userID int) ([]models.Connection, error)
	GetFollowerList(ctx echo.Context, userID int) ([]models.Connection, error)
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

func (cu *connectionUsecase) GetFollowingList(ctx echo.Context, userID int) ([]models.Connection, error) {
	followingList := []models.Connection{}
	err := cu.cr.GetFollowingList(ctx, userID, &followingList)
	if err != nil {
		return nil, err
	}

	return followingList, nil
}

func (cu *connectionUsecase) GetFollowerList(ctx echo.Context, userID int) ([]models.Connection, error) {
	followerList := []models.Connection{}
	err := cu.cr.GetFollowerList(ctx, userID, &followerList)
	if err != nil {
		return nil, err
	}

	return followerList, nil
}
