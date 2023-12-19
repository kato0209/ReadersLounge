package repository

import (
	"backend/models/chat"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IChatRepository interface {
	SaveMessage(message *chat.Message) error
	CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error)
}

type chatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) IChatRepository {
	return &chatRepository{db}
}

func (cr *chatRepository) SaveMessage(message *chat.Message) error {
	query := `INSERT INTO chat_messages (
				user_id, chat_room_id, content
			) VALUES ($1, $2, $3) RETURNING chat_message_id;`

	chatRoomID := 1
	err := cr.db.QueryRowx(
		query,
		message.User.UserID,
		chatRoomID,
		message.Content,
	).Scan(&message.MessageID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (cr *chatRepository) CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error) {
	query := `SELECT EXISTS (
				SELECT 1 
				FROM entries 
				WHERE chat_room_id = $1 AND user_id = $2
			);`

	var exists bool
	err := cr.db.QueryRowxContext(ctx.Request().Context(), query, userID, roomID).Scan(&exists)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return exists, nil
}
