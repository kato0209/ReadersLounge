package usecase

import (
	"backend/models"
	"backend/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(ctx echo.Context, user models.User) (models.UserResponse, error)
	Login(ctx echo.Context, user models.User) (string, error)
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
		return models.UserResponse{}, errors.WithStack(err)
	}
	newUser := models.User{Name: user.Name, IdentityType: user.IdentityType, Identifier: user.Identifier, Credential: string(hash)}
	if err := uu.ur.CreateUser(ctx, &newUser); err != nil {
		return models.UserResponse{}, errors.WithStack(err)
	}
	resUser := models.UserResponse{
		UserID: newUser.UserID,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(ctx echo.Context, user models.User) (string, error) {

	storedUser := models.User{}
	if err := uu.ur.GetUserByIdentifier(ctx, &storedUser, user.Identifier); err != nil {
		return "", errors.WithStack(err)
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Credential), []byte(user.Credential))
	if err != nil {
		return "", errors.WithStack(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.UserID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return tokenString, nil
}
