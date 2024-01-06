package usecase

import (
	"backend/controller/openapi"
	"backend/models"
	"backend/models/chat"
	"backend/repository"
	"backend/utils"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IChatUsecase interface {
	ReadLoop(client *chat.Client, broadCast chan<- []byte, unregister chan<- *chat.Client)
	WriteLoop(client *chat.Client)
	CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error)
	GetChatRooms(ctx echo.Context, userID int) ([]chat.Room, error)
	GetMessages(ctx echo.Context, roomID int) ([]chat.Message, error)
	CreateChatRoom(ctx echo.Context, userID, chatPartnerID int, room *chat.Room) error
}

type chatUsecase struct {
	cr repository.IChatRepository
}

func NewChatUsecase(cr repository.IChatRepository) IChatUsecase {
	return &chatUsecase{cr}
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

		var decodedMessage chat.Message
		err = json.Unmarshal([]byte(jsonMsg), &decodedMessage)
		if err != nil {
			log.Printf("failed to decode message: %v", err)
			break
		}

		decodedMessage.User = models.User{UserID: client.ClientID}
		if err := cu.cr.SaveMessage(&decodedMessage); err != nil {
			log.Printf("failed to save message: %v", err)
			break
		}

		resMessage := openapi.Message{
			MessageId: decodedMessage.MessageID,
			UserId:    decodedMessage.User.UserID,
			Content:   decodedMessage.Content,
			SentAt:    decodedMessage.CreatedAt.Format("2006-01-02 15:04"),
		}

		encodedMessage, err := json.Marshal(resMessage)
		if err != nil {
			log.Printf("failed to encode message: %v", err)
			break
		}

		broadCast <- encodedMessage

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
	for i := range rooms {
		room := &rooms[i]
		if !utils.IsRemotePath(room.ChatPartner.ProfileImage.FileName) {
			profileImage, err := utils.LoadImage(ctx, room.ChatPartner.ProfileImage.FileName)
			if err != nil {
				return []chat.Room{}, errors.WithStack(err)
			}
			room.ChatPartner.ProfileImage.EncodedImage = &profileImage
		}
	}
	return rooms, nil
}

func (cu *chatUsecase) GetMessages(ctx echo.Context, roomID int) ([]chat.Message, error) {
	messages := []chat.Message{}
	err := cu.cr.GetMessagesByRoomID(ctx, roomID, &messages)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return messages, nil
}

func (cu *chatUsecase) CreateChatRoom(ctx echo.Context, userID, chatPartnerID int, room *chat.Room) error {

	roomExists, roomID, err := cu.cr.CheckRoomExists(ctx, userID, chatPartnerID)
	if err != nil {
		return errors.WithStack(err)
	}
	room.RoomID = roomID

	if roomExists {
		err = cu.cr.UpdateRoomAccessTime(ctx, roomID, userID)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	err = cu.cr.CreateChatRoom(ctx, userID, chatPartnerID, room)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
