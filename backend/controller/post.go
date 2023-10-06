package controller

import (
	"backend/controller/openapi"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (s *Server) Posts(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	fmt.Println(userID)

	posts := []models.Post{}
	if err := s.pu.GetAllPosts(ctx, &posts); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	resPosts := []openapi.Post{}

	return ctx.JSON(http.StatusCreated, resPosts)
}
