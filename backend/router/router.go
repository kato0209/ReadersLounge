package router

import (
	"backend/controller"
	"backend/controller/openapi"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server *controller.Server) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8082", "http://localhost:8081", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:   "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		//CookieDomain: "localhost",
		//CookieSameSite: http.SameSiteNoneMode,
		CookieHTTPOnly: true,
		//CookieSecure:   false,
		CookieSameSite: http.SameSiteDefaultMode,
		CookieMaxAge:   60,
	}))

	openapi.RegisterHandlers(e, server)
	p := e.Group("/posts")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))
	p.GET("", server.GetPosts)
	p.POST("", server.CreatePost)

	u := e.Group("/user")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))
	u.GET("", server.User)

	return e
}
