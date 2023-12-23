package controller

import (
	"backend/controller/openapi"
	models "backend/models/chat"
	"backend/utils"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (s *Server) ChatSocket(ctx echo.Context, params openapi.ChatSocketParams) error {

	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	roomID := params.RoomId
	hasPermisson, err := s.cu.CheckRoomAccessPermission(ctx, userID, roomID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	} else if !hasPermisson {
		return ctx.JSON(http.StatusUnauthorized, "You don't have permission to access this room")
	}

	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	hub := models.NewHub()
	go s.cu.RunLoop(hub)

	client := models.NewClient(ws, userID, roomID)
	go s.cu.ReadLoop(client, hub.BroadcastCh, hub.UnRegisterCh)
	go s.cu.WriteLoop(client)
	hub.RegisterCh <- client

	return nil
}

func (s *Server) GetChatRooms(ctx echo.Context) error {
	userID, err := utils.ExtractUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	rooms, err := s.cu.GetChatRooms(ctx, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resRooms := []openapi.ChatRoom{}
	for _, room := range rooms {
		resRoom := openapi.ChatRoom{
			RoomId: room.RoomID,
		}
		resRooms = append(resRooms, resRoom)
	}

	return ctx.JSON(http.StatusOK, resRooms)
}
