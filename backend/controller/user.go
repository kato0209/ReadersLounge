package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"backend/controller/openapi"
	"backend/models"
	"backend/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) Signup(ctx echo.Context) error {
	reqSignupBody := openapi.ReqSignupBody{}
	if err := ctx.Bind(&reqSignupBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := s.uu.Signup(ctx, models.User{
		Name:         *reqSignupBody.Username,
		ProfileText:  nil,
		IdentityType: "EmailPassword",
		Identifier:   *reqSignupBody.Identifier,
		Credential:   *reqSignupBody.Credential,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	tokenString, user, err := s.uu.Login(ctx, models.User{
		Identifier: *reqSignupBody.Identifier,
		Credential: *reqSignupBody.Credential,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       user.UserID,
		Name:         user.Name,
		ProfileImage: user.ProfileImage,
	}

	utils.SetJwtTokenInCookie(ctx, tokenString)

	return ctx.JSON(http.StatusCreated, resUser)
}

func (s *Server) Login(ctx echo.Context) error {
	reqLoginBody := openapi.ReqLoginBody{}
	if err := ctx.Bind(&reqLoginBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, user, err := s.uu.Login(ctx, models.User{
		Identifier: *reqLoginBody.Identifier,
		Credential: *reqLoginBody.Credential,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       user.UserID,
		Name:         user.Name,
		ProfileImage: user.ProfileImage,
	}

	utils.SetJwtTokenInCookie(ctx, tokenString)

	return ctx.JSON(http.StatusOK, resUser)
}

func (s *Server) Logout(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	//cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}

func (s *Server) GoogleOauthCallback(ctx echo.Context, params openapi.GoogleOauthCallbackParams) error {

	cookieState, err := ctx.Cookie("state")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	queryState := ctx.QueryParam("state")

	if queryState != cookieState.Value {
		return ctx.JSON(http.StatusInternalServerError, "invalid state")
	}

	tokenString, err := s.uu.GoogleOAuthCallback(ctx, params.Code)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	utils.SetJwtTokenInCookie(ctx, tokenString)

	return ctx.Redirect(http.StatusMovedPermanently, os.Getenv("FE_URL"))
}

func (s *Server) GetUser(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	loginUser, err := s.uu.GetUserByUserID(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       loginUser.UserID,
		Name:         loginUser.Name,
		ProfileImage: loginUser.ProfileImage,
	}

	return ctx.JSON(http.StatusOK, resUser)
}

func (s *Server) UpdateUser(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	updateUser := models.User{
		Name:        form.Value["name"][0],
		ProfileText: &form.Value["profile_text"][0],
	}

	file, err := ctx.FormFile("image")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		fileModel := strings.Split(file.Filename, ".")
		fileName := fileModel[0]
		extension := fileModel[1]

		if extension == "jpeg" || extension == "png" {

			data, err := io.ReadAll(src)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, err.Error())
			}

			generatedFileName := fmt.Sprintf("%s_%s.%s", fileName, uuid.New().String(), extension)

			image := models.ProfileImage{
				Source:   data,
				FileName: &generatedFileName,
			}
			fmt.Println(image)

		} else {
			return ctx.JSON(http.StatusBadRequest, "Unsupported file type")
		}

	} else if err != http.ErrMissingFile {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.uu.UpdateUser(ctx, &updateUser, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       updateUser.UserID,
		Name:         updateUser.Name,
		ProfileImage: updateUser.ProfileImage,
	}
	return ctx.JSON(http.StatusOK, resUser)
}
