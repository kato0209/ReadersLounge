package controller

import (
	"backend/controller/openapi"
	"backend/models/chat"
	models "backend/models/chat"
	"backend/utils"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (s *Server) RunLoop(h *chat.Hub) {
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

	client := models.NewClient(ws, userID, roomID)
	go s.cu.ReadLoop(client, s.hub.BroadcastCh, s.hub.UnRegisterCh)
	go s.cu.WriteLoop(client)
	s.hub.RegisterCh <- client

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
			RoomId:                 room.RoomID,
			TargetUserId:           room.ChatPartner.UserID,
			TargetUserName:         room.ChatPartner.Name,
			TargetUserProfileImage: room.ChatPartner.ProfileImage.ClassifyPathType(),
			LastMessage:            room.LastMessage.Content,
			LastMessageSentAt:      room.LastMessage.CreatedAt.Format("2006-01-02 15:04"),
		}
		resRooms = append(resRooms, resRoom)
	}

	return ctx.JSON(http.StatusOK, resRooms)
}

func (s *Server) GetMessages(ctx echo.Context, params openapi.GetMessagesParams) error {
	roomID := params.RoomId
	messages, err := s.cu.GetMessages(ctx, roomID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	resMessages := []openapi.Message{}
	for _, message := range messages {
		resMessage := openapi.Message{
			MessageId: message.MessageID,
			UserId:    message.User.UserID,
			Content:   message.Content,
			SentAt:    message.CreatedAt.Format("2006-01-02 15:04"),
		}
		resMessages = append(resMessages, resMessage)
	}
	return ctx.JSON(http.StatusOK, resMessages)
}
