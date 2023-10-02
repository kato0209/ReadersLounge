package usecase

import (
	"backend/models"
	"backend/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(ctx echo.Context, user models.User) (models.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(ctx echo.Context, user models.User) (models.UserResponse, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Credential), 10)
	if err != nil {
		return models.UserResponse{}, err
	}
	newUser := models.User{Name: user.Name, IdentityType: user.IdentityType, Identifier: user.Identifier, Credential: string(hash)}
	if err := uu.ur.CreateUser(ctx, &newUser); err != nil {
		return models.UserResponse{}, err
	}
	resUser := models.UserResponse{
		UserID: newUser.UserID,
	}
	return resUser, nil
}
