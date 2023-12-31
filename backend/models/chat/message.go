package chat

import (
	"backend/models"
	"time"
)

type Message struct {
	MessageID int `json:"message_id" db:"chat_message_id"`
	User      models.User
	RoomID    int       `json:"room_id" db:"chat_room_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"sent_at" db:"created_at"`
}
