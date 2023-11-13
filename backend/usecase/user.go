package usecase

import (
	"backend/auth/oidc"
	"backend/models"
	"backend/repository"
	"backend/validator"
	"fmt"
	"os"

	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(ctx echo.Context, user models.User) (models.User, error)
	Login(ctx echo.Context, user models.User) (string, models.User, error)
	GoogleOAuthCallback(ctx echo.Context, code string) (string, models.User, error)
	GetUserByUserID(ctx echo.Context, userID int) (models.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) Signup(ctx echo.Context, user models.User) (models.User, error) {
	if err := uu.uv.SignupValidator(user); err != nil {
		return models.User{}, errors.WithStack(err)
	}
	existingUser := models.User{}
	err := uu.ur.GetUserByIdentifier(ctx, &existingUser, user.Identifier)
	if err == nil {
		return models.User{}, errors.WithStack(errors.New("email already exists"))
	} else if err.Error() != "user is not found" {
		return models.User{}, errors.WithStack(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Credential), 10)
	if err != nil {
		return models.User{}, errors.WithStack(err)
	}
	newUser := models.User{Name: user.Name, IdentityType: user.IdentityType, Identifier: user.Identifier, Credential: string(hash)}
	if err := uu.ur.CreateUser(ctx, &newUser); err != nil {
		return models.User{}, errors.WithStack(err)
	}
	resUser := models.User{
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

	tokenString, err := utils.CreateJwtTokenByUserID(storedUser.UserID)
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

func (uu *userUsecase) GoogleOAuthCallback(ctx echo.Context, code string) (string, models.User, error) {
	client := oidc.NewGoogleOidcClient()
	tokenResp, err := client.PostTokenEndpoint(
		code,
		fmt.Sprintf(
			"%s://%s:%s/oauth/google/callback",
			os.Getenv("API_PROTOCOL"),
			os.Getenv("API_DOMAIN"),
			os.Getenv("API_PORT"),
		),
		"authorization_code",
	)
	if err != nil {
		return "", models.User{}, errors.WithStack(err)
	}
	idToken, err := oidc.NewIdToken(tokenResp.IdToken, oidc.Google)
	if err != nil {
		return "", models.User{}, errors.WithStack(err)
	}
	if err = idToken.Validate(client.JwksEndpoint, client.ClientId); err != nil {
		return "", models.User{}, errors.WithStack(err)
	}

	email, err := idToken.Payload.GetEmail()
	if err != nil {
		return "", models.User{}, errors.WithStack(err)
	}

	user := models.User{
		Name:         idToken.Payload.GetName(),
		IdentityType: "GoogleOAuth",
		Identifier:   email,
		Credential:   "",
	}
	storedUser := models.User{}
	err = uu.ur.GetUserByIdentifier(ctx, &storedUser, user.Identifier)
	if err == nil {
		tokenString, err := utils.CreateJwtTokenByUserID(storedUser.UserID)
		if err != nil {
			return "", models.User{}, errors.WithStack(err)
		}
		return tokenString, storedUser, nil
	} else if err.Error() == "user is not found" {

		if err := uu.ur.CreateUser(ctx, &user); err != nil {
			return "", models.User{}, errors.WithStack(err)
		}

		tokenString, err := utils.CreateJwtTokenByUserID(user.UserID)
		if err != nil {
			return "", models.User{}, errors.WithStack(err)
		}
		return tokenString, user, nil

	} else {
		return "", models.User{}, errors.WithStack(err)
	}
}

func (uu *userUsecase) GetUserByUserID(ctx echo.Context, userID int) (models.User, error) {
	resUser := models.User{}
	if err := uu.ur.GetUserByUserID(ctx, &resUser, userID); err != nil {
		return models.User{}, errors.WithStack(err)
	}
	return resUser, nil
}
