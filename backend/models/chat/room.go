package chat

import (
	"backend/models"
	"time"
)

type Room struct {
	RoomID      int `db:"chat_room_id"`
	ChatPartner models.User
	LastMessage LastMessage
}

type LastMessage struct {
	Content   *string
	CreatedAt *time.Time
}
