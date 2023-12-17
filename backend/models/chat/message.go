package models

import (
	"backend/models"
	"time"
)

type Message struct {
	MessageID int `db:"message_id"`
	User      models.User
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
