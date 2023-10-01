package controller

import (
	"net/http"

	generated_models "backend/controller/openapi/models"
	"backend/models"

	"github.com/labstack/echo/v4"
)

func (s *Server) Signup(ctx echo.Context) error {
	user := generated_models.ReqSignupBody{}
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := s.uu.Signup(ctx, models.User{
		Name:         *user.Username,
		ProfileText:  "",
		IdentityType: "email",
		Identifier:   *user.Identifier,
		Credential:   *user.Credential,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, userRes)
}
