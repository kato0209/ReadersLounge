package controller

import (
	"net/http"

	"backend/controller/openapi"
	"backend/models"

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
	return nil
}

func (s *Server) Logout(ctx echo.Context) error {
	return nil
}
