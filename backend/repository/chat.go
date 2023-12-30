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
	GetAllChatRooms(ctx echo.Context, userID int) ([]chat.Room, error)
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

func (cr *chatRepository) GetAllChatRooms(ctx echo.Context, userID int) ([]chat.Room, error) {
	query := `SELECT 
				chat_rooms.chat_room_id AS room_id,
				other_users.user_id AS target_user_id,
				other_user_details.name AS target_user_name,
				other_user_details.profile_image AS target_user_profile_image,
				latest_messages.content AS last_message,
				latest_messages.created_at AS last_message_created_at
			FROM entries
			INNER JOIN chat_rooms ON entries.chat_room_id = chat_rooms.chat_room_id
			INNER JOIN entries as other_entries ON chat_rooms.chat_room_id = other_entries.chat_room_id AND other_entries.user_id <> entries.user_id
			INNER JOIN users AS other_users ON other_users.user_id = other_entries.user_id
			INNER JOIN user_details AS other_user_details ON other_users.user_id = other_user_details.user_id
			LEFT JOIN (
				SELECT 
					chat_room_id, 
					content, 
					created_at
				FROM chat_messages
				WHERE chat_room_id IN (
					SELECT chat_room_id FROM entries WHERE user_id = $1
				)
				ORDER BY created_at DESC
				LIMIT 1
			) AS latest_messages ON chat_rooms.chat_room_id = latest_messages.chat_room_id
			WHERE entries.user_id = $1;
			`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var rooms []chat.Room
	for rows.Next() {
		room := chat.Room{}
		err := rows.Scan(
			&room.RoomID,
			&room.ChatPartner.UserID,
			&room.ChatPartner.Name,
			&room.ChatPartner.ProfileImage.FileName,
			&room.LastMessage.Content,
			&room.LastMessage.CreatedAt,
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}
	return rooms, nil
}
