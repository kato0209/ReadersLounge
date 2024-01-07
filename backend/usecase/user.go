package usecase

import (
	"backend/models"
	"backend/repository"
	"backend/validator"
	"fmt"
	"os"

	"backend/utils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthApi "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type IUserUsecase interface {
	Signup(ctx echo.Context, user models.User) error
	Login(ctx echo.Context, user models.User) (string, models.User, error)
	GoogleOAuthCallback(ctx echo.Context, code string) (string, error)
	GetUserByUserID(ctx echo.Context, userID int) (models.User, error)
	UpdateUser(ctx echo.Context, user *models.User) error
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) Signup(ctx echo.Context, user models.User) error {
	if err := uu.uv.SignupValidator(user); err != nil {
		return errors.WithStack(err)
	}

	userExists, err := uu.ur.CheckExistsUserByIdentifier(ctx, user.Identifier)
	if err != nil {
		return errors.WithStack(err)
	}

	if userExists {
		return errors.New("EMAIL_ALREADY_USED")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Credential), 10)
	if err != nil {
		return errors.WithStack(err)
	}
	user.Credential = string(hash)

	if err := uu.ur.CreateUser(ctx, &user); err != nil {
		return errors.WithStack(err)
	}

	if !utils.IsRemotePath(user.ProfileImage.FileName) {
		profileImage, err := uu.ur.LoadProfileImage(ctx, user.ProfileImage.FileName)
		if err != nil {
			return errors.WithStack(err)
		}
		user.ProfileImage.EncodedImage = &profileImage
	}

	return nil
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

	if !utils.IsRemotePath(storedUser.ProfileImage.FileName) {
		profileImage, err := uu.ur.LoadProfileImage(ctx, storedUser.ProfileImage.FileName)
		if err != nil {
			return "", models.User{}, errors.WithStack(err)
		}
		storedUser.ProfileImage.EncodedImage = &profileImage
	}

	resUser := models.User{
		UserID:       storedUser.UserID,
		Name:         storedUser.Name,
		ProfileImage: storedUser.ProfileImage,
	}
	return tokenString, resUser, nil
}

func (uu *userUsecase) GoogleOAuthCallback(ctx echo.Context, code string) (string, error) {
	c := ctx.Request().Context()

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL: fmt.Sprintf(
			"%s://%s:%s/%s",
			os.Getenv("API_PROTOCOL"),
			os.Getenv("API_DOMAIN"),
			os.Getenv("API_PORT"),
			os.Getenv("GOOGLE_OAUTH_PATH"),
		),
		Scopes:   []string{os.Getenv("GOOGLE_OAUTH_USER_INFO_EMAIL_URL"), os.Getenv("GOOGLE_OAUTH_USER_INFO_PROFILE_URL")},
		Endpoint: google.Endpoint,
	}

	token, err := config.Exchange(c, code)
	if err != nil {
		return "", errors.WithStack(err)
	}

	service, err := oauthApi.NewService(c, option.WithTokenSource(config.TokenSource(c, token)))
	if err != nil {
		return "", errors.WithStack(err)
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return "", errors.WithStack(err)
	}

	profileImage := models.ProfileImage{FileName: userInfo.Picture}
	user := models.User{
		Name:         userInfo.Name,
		ProfileImage: profileImage,
		IdentityType: "GoogleOAuth",
		Identifier:   userInfo.Email,
		Credential:   "",
	}

	storedUser := models.User{}
	err = uu.ur.GetUserByIdentifier(ctx, &storedUser, user.Identifier)
	if err == nil {
		tokenString, err := utils.CreateJwtTokenByUserID(storedUser.UserID)
		if err != nil {
			return "", errors.WithStack(err)
		}
		return tokenString, nil
	} else if err.Error() == "user is not found" {
		if err := uu.ur.CreateUser(ctx, &user); err != nil {
			return "", errors.WithStack(err)
		}

		tokenString, err := utils.CreateJwtTokenByUserID(user.UserID)
		if err != nil {
			return "", errors.WithStack(err)
		}
		return tokenString, nil

	} else {
		return "", errors.WithStack(err)
	}
}

func (uu *userUsecase) GetUserByUserID(ctx echo.Context, userID int) (models.User, error) {
	resUser := models.User{}
	if err := uu.ur.GetUserByUserID(ctx, &resUser, userID); err != nil {
		return models.User{}, errors.WithStack(err)
	}

	if !utils.IsRemotePath(resUser.ProfileImage.FileName) {
		profileImage, err := uu.ur.LoadProfileImage(ctx, resUser.ProfileImage.FileName)
		if err != nil {
			return models.User{}, errors.WithStack(err)
		}
		resUser.ProfileImage.EncodedImage = &profileImage
	}

	return resUser, nil
}

func (uu *userUsecase) UpdateUser(ctx echo.Context, user *models.User) error {
	if user.ProfileImage != (models.ProfileImage{}) {
		if err := uu.ur.SaveProfileImage(ctx, &user.ProfileImage); err != nil {
			return errors.WithStack(err)
		}
	}
	if err := uu.ur.UpdateUserByUserID(ctx, user); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
