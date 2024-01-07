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

	fmt.Println(11)
	if err := s.cnu.CreateConnection(ctx, userID, CreateConnectionBody.TargetUserId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(22)

	return ctx.NoContent(http.StatusCreated)
}

func (s *Server) DeleteConnection(ctx echo.Context, connectionId int) error {
	if err := s.cnu.DeleteConnection(ctx, connectionId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetFollowingConnections(ctx echo.Context, params openapi.GetFollowingConnectionsParams) error {
	userID := params.UserId

	followingConnections, err := s.cnu.GetFollowingConnections(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res := []openapi.Connection{}
	for _, connection := range followingConnections {
		res = append(res, openapi.Connection{
			ConnectionId:           connection.ConnectionID,
			TargetUserId:           connection.Following.UserID,
			TargetUserName:         connection.Following.Name,
			TargetUserProfileImage: connection.Following.ProfileImage.ClassifyPathType(),
		})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetFollowerConnections(ctx echo.Context, params openapi.GetFollowerConnectionsParams) error {
	userID := params.UserId

	followerConnections, err := s.cnu.GetFollowerConnections(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	res := []openapi.Connection{}
	for _, connection := range followerConnections {
		res = append(res, openapi.Connection{
			ConnectionId:           connection.ConnectionID,
			TargetUserId:           connection.Follower.UserID,
			TargetUserName:         connection.Follower.Name,
			TargetUserProfileImage: connection.Follower.ProfileImage.ClassifyPathType(),
		})
	}

	return ctx.JSON(http.StatusOK, res)
}
