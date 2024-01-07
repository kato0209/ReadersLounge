package controller

import (
	"fmt"
	"net/http"

	"backend/controller/openapi"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreateConnection(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	CreateConnectionBody := openapi.CreateConnectionJSONBody{}
	if err := ctx.Bind(&CreateConnectionBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(CreateConnectionBody.TargetUserId)

	if err := s.cnu.CreateConnection(ctx, userID, CreateConnectionBody.TargetUserId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}

func (s *Server) DeleteConnection(ctx echo.Context, connectionId int) error {
	if err := s.cnu.DeleteConnection(ctx, connectionId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetFollowingList(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	followingList, err := s.cnu.GetFollowingList(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, followingList)
}

func (s *Server) GetFollowerList(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	followerList, err := s.cnu.GetFollowerList(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, followerList)
}
