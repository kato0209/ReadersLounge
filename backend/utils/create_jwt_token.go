package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func CreateJwtTokenByUserID(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return tokenString, nil

}
