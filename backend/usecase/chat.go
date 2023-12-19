package usecase

import (
	"backend/models"
	"backend/models/chat"
	"backend/repository"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IChatUsecase interface {
	RunLoop(hub *chat.Hub)
	ReadLoop(client *chat.Client, broadCast chan<- []byte, unregister chan<- *chat.Client)
	WriteLoop(client *chat.Client)
	CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error)
	GetChatRooms(ctx echo.Context, userID int) ([]chat.Room, error)
}

type chatUsecase struct {
	cr repository.IChatRepository
}

func NewChatUsecase(cr repository.IChatRepository) IChatUsecase {
	return &chatUsecase{cr}
}

func (cu *chatUsecase) RunLoop(h *chat.Hub) {
	for {
		select {
		case client := <-h.RegisterCh:
			h.Register(client)

		case client := <-h.UnRegisterCh:
			h.Unregister(client)

		case msg := <-h.BroadcastCh:
			h.BroadCastToAllClient(msg)
		}
	}
}

func (cu *chatUsecase) ReadLoop(client *chat.Client, broadCast chan<- []byte, unregister chan<- *chat.Client) {
	defer func() {
		client.Disconnect(unregister)
	}()

	for {
		_, jsonMsg, err := client.WsConnect.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		broadCast <- jsonMsg

		message := chat.Message{
			User:    models.User{UserID: client.ClientID},
			Content: string(jsonMsg),
		}
		if err := cu.cr.SaveMessage(&message); err != nil {
			log.Printf("failed to save message: %v", err)
			break
		}

	}
}

func (cu *chatUsecase) WriteLoop(client *chat.Client) {
	defer func() {
		client.WsConnect.Close()
	}()

	for {
		message := <-client.SendCh
		w, err := client.WsConnect.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		w.Write(message)

		if err := w.Close(); err != nil {
			return
		}
	}
}

func (cu *chatUsecase) CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error) {
	hasPermission, err := cu.cr.CheckRoomAccessPermission(ctx, userID, roomID)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return hasPermission, nil
}

func (cu *chatUsecase) GetChatRooms(ctx echo.Context, userID int) ([]chat.Room, error) {
	rooms, err := cu.cr.GetAllChatRooms(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rooms, nil
}
