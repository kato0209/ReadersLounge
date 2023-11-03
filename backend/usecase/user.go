package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(ctx echo.Context, user models.User) (models.UserResponse, error)
	Login(ctx echo.Context, user models.User) (string, models.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) Signup(ctx echo.Context, user models.User) (models.UserResponse, error) {
	if err := uu.uv.SignupValidator(user); err != nil {
		return models.UserResponse{}, errors.WithStack(err)
	}
	existingUser := models.User{}
	err := uu.ur.GetUserByIdentifier(ctx, &existingUser, user.Identifier)
	if err == nil {
		return models.UserResponse{}, errors.WithStack(errors.New("email already exists"))
	} else if err.Error() != "user is not found" {
		return models.UserResponse{}, errors.WithStack(err)
	}

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

func (uu *userUsecase) Login(ctx echo.Context, user models.User) (string, models.User, error) {
	if err := uu.uv.LoginValidator(user); err != nil {
		return "", models.User{}, errors.WithStack(err)
	}

	storedUser := models.User{}
	if err := uu.ur.GetUserByIdentifier(ctx, &storedUser, user.Identifier); err != nil {
		return "", models.User{}, errors.WithStack(err)
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Credential), []byte(user.Credential))
	if err != nil {
		return "", models.User{}, errors.WithStack(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.UserID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", models.User{}, errors.WithStack(err)
	}
	resUser := models.User{
		UserID:       storedUser.UserID,
		Name:         storedUser.Name,
		ProfileImage: storedUser.ProfileImage,
	}
	return tokenString, resUser, nil
}
