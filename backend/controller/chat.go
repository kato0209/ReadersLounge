package controller

import (
	models "backend/models/chat"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (s *Server) ChatSocket(ctx echo.Context) error {
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
	go hub.RunLoop()

	client := models.NewClient(ws)
	go client.ReadLoop(hub.BroadcastCh, hub.UnRegisterCh)
	go client.WriteLoop()
	hub.RegisterCh <- client

	return nil
}
