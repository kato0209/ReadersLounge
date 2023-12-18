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
	go s.cu.RunLoop(hub)

	client := models.NewClient(ws)
	go s.cu.ReadLoop(client, hub.BroadcastCh, hub.UnRegisterCh)
	go s.cu.WriteLoop(client)
	hub.RegisterCh <- client

	return nil
}
