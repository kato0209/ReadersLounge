package router

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8082", "http://localhost:8081", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		//CookieDomain: os.Getenv("API_DOMAIN"),
		CookieDomain: "localhost",
		//CookieHTTPOnly: true,
		//CookieSameSite: http.SameSiteNoneMode,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	p := e.Group("/posts")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))

	return e
}
