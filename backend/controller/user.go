package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"backend/controller/openapi"
	"backend/models"

	"backend/auth/oidc"

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
		IdentityType: "email",
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

func (s *Server) SignupWithGoogle(ctx echo.Context) error {
	client := oidc.NewGoogleOidcClient()

	state, err := oidc.RandomState()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "state"
	cookie.Value = state
	ctx.SetCookie(cookie)

	redirectUrl := client.AuthUrl(
		"code",
		[]string{"openid", "email", "profile"},
		fmt.Sprintf(
			"%s://%s:%s/auth/google/callback",
			os.Getenv("API_PROTOCOL"),
			os.Getenv("API_DOMAIN"),
			os.Getenv("API_PORT"),
		),
		state,
	)
	fmt.Println(redirectUrl)
	for _, cookie := range ctx.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func (s *Server) GoogleSignupCallback(ctx echo.Context, params openapi.GoogleSignupCallbackParams) error {
	cookieState, _ := ctx.Cookie("state")
	fmt.Println(cookieState)

	params = openapi.GoogleSignupCallbackParams{}
	if err := ctx.Bind(&params); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	/*
		if params.State != cookieState.Value {
			err := fmt.Errorf("state parameter does not match for query: %s, cookie: %s", params.State, cookieState)
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}*/

	// 認可コードを取り出しトークンエンドポイントに投げることでid_tokenを取得できる
	client := oidc.NewGoogleOidcClient()
	tokenResp, err := client.PostTokenEndpoint(
		params.Code,
		fmt.Sprintf(
			"%s://%s:%s/auth/google/callback",
			os.Getenv("API_PROTOCOL"),
			os.Getenv("API_DOMAIN"),
			os.Getenv("API_PORT"),
		),
		"authorization_code",
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// JWKsエンドポイントから公開鍵を取得しid_token(JWT)の署名を検証。改竄されていないことを確認する
	idToken, err := oidc.NewIdToken(tokenResp.IdToken, oidc.Google)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err = idToken.Validate(client.JwksEndpoint, client.ClientId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	//cookieの設定

	return ctx.NoContent(http.StatusOK)
}
