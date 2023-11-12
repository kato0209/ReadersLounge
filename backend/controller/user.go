package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"backend/controller/openapi"
	"backend/models"

	"github.com/labstack/echo/v4"
)

func (s *Server) Signup(ctx echo.Context) error {
	reqSignupBody := openapi.ReqSignupBody{}
	if err := ctx.Bind(&reqSignupBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := s.uu.Signup(ctx, models.User{
		Name:         *reqSignupBody.Username,
		ProfileText:  "",
		IdentityType: "EmailPassword",
		Identifier:   *reqSignupBody.Identifier,
		Credential:   *reqSignupBody.Credential,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resSignupBody := openapi.ResSignupBody{
		UserId: &userRes.UserID,
	}

	return ctx.JSON(http.StatusCreated, resSignupBody)
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
		UserId:       &user.UserID,
		Name:         &user.Name,
		ProfileImage: &user.ProfileImage,
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	//cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)

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

	// stateの検証

	tokenString, user, err := s.uu.GoogleOAuthCallback(ctx, params.Code)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       &user.UserID,
		Name:         &user.Name,
		ProfileImage: &user.ProfileImage,
	}
	fmt.Println(resUser)

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	//cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)

	return ctx.Redirect(http.StatusMovedPermanently, os.Getenv("FRONTEND_URL"))
}

func (s *Server) User(ctx echo.Context) error {
	user := models.User{}
	return ctx.JSON(http.StatusOK, user)
}
