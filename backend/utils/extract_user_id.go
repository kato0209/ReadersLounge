package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func ExtractUserID(ctx echo.Context) (int, error) {
	user, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("failed to extract user id")
	}
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	return userID, nil
}
