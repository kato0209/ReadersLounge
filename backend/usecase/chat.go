package usecase

import (
	"backend/models/chat"
	"backend/repository"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type IChatUsecase interface {
	SaveMessage(*chat.Message) error
	RunLoop(*chat.Hub) error
	ReadLoop(client *chat.Client, broadCast chan<- []byte, unregister chan<- *chat.Client) error
	WriteLoop(client *chat.Client) error
}

type chatUsecase struct {
	cr repository.IChatRepository
}

func NewChatUsecase(cr repository.IChatRepository) IChatUsecase {
	return &chatUsecase{cr}
}

func (cu *chatUsecase) SaveMessage(*chat.Message) error {
	return nil
}

func (cu *chatUsecase) RunLoop(h *chat.Hub) error {
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

func (cu *chatUsecase) ReadLoop(client *chat.Client, broadCast chan<- []byte, unregister chan<- *chat.Client) error {
	defer func() {
		client.Disconnect(unregister)
	}()

	for {
		_, jsonMsg, err := client.WsConnect.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			return errors.WithStack(err)
		}

		broadCast <- jsonMsg
	}
}

func (cu *chatUsecase) WriteLoop(client *chat.Client) error {
	defer func() {
		client.WsConnect.Close()
	}()

	for {
		message := <-client.SendCh
		w, err := client.WsConnect.NextWriter(websocket.TextMessage)
		if err != nil {
			return errors.WithStack(err)
		}
		w.Write(message)

		if err := w.Close(); err != nil {
			return errors.WithStack(err)
		}
	}
}
