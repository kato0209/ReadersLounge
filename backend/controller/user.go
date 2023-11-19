package controller

import (
	"net/http"
	"os"
	"time"

	"backend/controller/openapi"
	"backend/models"
	"backend/utils"

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
		UserId:       &user.UserID,
		Name:         &user.Name,
		ProfileImage: &user.ProfileImage,
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
		UserId:       &user.UserID,
		Name:         &user.Name,
		ProfileImage: &user.ProfileImage,
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

func (s *Server) User(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	loginUser, err := s.uu.GetUserByUserID(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resUser := openapi.User{
		UserId:       &loginUser.UserID,
		Name:         &loginUser.Name,
		ProfileImage: &loginUser.ProfileImage,
	}

	return ctx.JSON(http.StatusOK, resUser)
}
