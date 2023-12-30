package chat

import (
	"backend/models"
)

type Room struct {
	RoomID      int `db:"chat_room_id"`
	ChatPartner models.User
	LastMessage Message
}
