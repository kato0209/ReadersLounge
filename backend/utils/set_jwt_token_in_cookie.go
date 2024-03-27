package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func SetJwtTokenInCookie(ctx echo.Context, tokenString string) {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	//cookie.SameSite = http.SameSiteDefaultMode
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
}
