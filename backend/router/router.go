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

	openapi.RegisterHandlers(e, server)

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:jwt_token",
		Skipper: func(c echo.Context) bool {
			skipPaths := []string{"/health", "/csrftoken", "/set-state", "/signup", "/login", "/logout", "/oauth/google", "/oauth/google/callback"}

			path := c.Path()

			for _, skipPath := range skipPaths {
				if path == skipPath {
					return true
				}
			}

			return false
		},
	}))

	return e
}
