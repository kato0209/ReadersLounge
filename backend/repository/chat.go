package repository

import (
	"backend/models/chat"

	"github.com/jmoiron/sqlx"
)

type IChatRepository interface {
}

type chatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) IChatRepository {
	return &chatRepository{db}
}

func (cr *chatRepository) SaveMessage(*chat.Message) error {
	return nil
}
