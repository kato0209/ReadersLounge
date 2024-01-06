package repository

import (
	"backend/models/chat"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IChatRepository interface {
	SaveMessage(message *chat.Message) error
	CheckRoomAccessPermission(ctx echo.Context, userID, roomID int) (bool, error)
	GetAllChatRooms(ctx echo.Context, userID int) ([]chat.Room, error)
	GetMessagesByRoomID(ctx echo.Context, roomID int, messaegs *[]chat.Message) error
	CreateChatRoom(ctx echo.Context, userID, chatPartnerID int, room *chat.Room) error
	CheckRoomExists(ctx echo.Context, userID, chatPartnerID int) (bool, int, error)
	UpdateRoomAccessTime(ctx echo.Context, roomID, userID int) error
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
			) VALUES ($1, $2, $3) RETURNING chat_message_id, created_at;`

	err := cr.db.QueryRowx(
		query,
		message.User.UserID,
		message.RoomID,
		message.Content,
	).Scan(&message.MessageID, &message.CreatedAt)
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
	err := cr.db.QueryRowxContext(ctx.Request().Context(), query, roomID, userID).Scan(&exists)
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
					inner_chat_messages.chat_room_id, 
					inner_chat_messages.content, 
					inner_chat_messages.created_at
				FROM (
					SELECT 
						chat_room_id, 
						content, 
						created_at,
						ROW_NUMBER() OVER (PARTITION BY chat_room_id ORDER BY created_at DESC) AS rn
					FROM chat_messages
					WHERE chat_room_id IN (
						SELECT chat_room_id FROM entries WHERE user_id = $1
					)
				) AS inner_chat_messages
				WHERE inner_chat_messages.rn = 1
			) AS latest_messages ON chat_rooms.chat_room_id = latest_messages.chat_room_id
			WHERE entries.user_id = $1
			ORDER BY latest_messages.created_at DESC;
			`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var rooms []chat.Room
	for rows.Next() {
		var lastMessageContent sql.NullString
		var lastMessageCreatedAt sql.NullTime
		room := chat.Room{}
		err := rows.Scan(
			&room.RoomID,
			&room.ChatPartner.UserID,
			&room.ChatPartner.Name,
			&room.ChatPartner.ProfileImage.FileName,
			&lastMessageContent,
			&lastMessageCreatedAt,
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if lastMessageContent.Valid {
			room.LastMessage.Content = &lastMessageContent.String
		}
		if lastMessageCreatedAt.Valid {
			room.LastMessage.CreatedAt = &lastMessageCreatedAt.Time
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}
	return rooms, nil
}

func (cr *chatRepository) GetMessagesByRoomID(ctx echo.Context, roomID int, messages *[]chat.Message) error {
	query := `
		SELECT
			chat_message_id, user_id, content, created_at
		FROM chat_messages
		WHERE chat_room_id = $1
		ORDER BY created_at;
	`
	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, roomID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		message := chat.Message{}
		err := rows.Scan(
			&message.MessageID,
			&message.User.UserID,
			&message.Content,
			&message.CreatedAt,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		*messages = append(*messages, message)
	}
	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (cr *chatRepository) CreateChatRoom(ctx echo.Context, userID, chatPartnerID int, room *chat.Room) error {
	tx, err := cr.db.BeginTxx(ctx.Request().Context(), nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO chat_rooms VALUES (DEFAULT) RETURNING chat_room_id;
	`
	err = tx.QueryRowxContext(ctx.Request().Context(), query).Scan(&room.RoomID)
	if err != nil {
		return errors.WithStack(err)
	}

	query = `
		INSERT INTO entries (user_id, chat_room_id) VALUES ($1, $2);
	`
	_, err = tx.ExecContext(ctx.Request().Context(), query, userID, room.RoomID)
	if err != nil {
		return errors.WithStack(err)
	}

	query = `
		INSERT INTO entries (user_id, chat_room_id) VALUES ($1, $2);
	`
	_, err = tx.ExecContext(ctx.Request().Context(), query, chatPartnerID, room.RoomID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (cr *chatRepository) CheckRoomExists(ctx echo.Context, userID, chatPartnerID int) (bool, int, error) {
	var roomID int
	query := `
		SELECT e.chat_room_id FROM entries e
		INNER JOIN entries e2 ON e.chat_room_id = e2.chat_room_id
		WHERE e.user_id = $1 AND e2.user_id = $2
		LIMIT 1;
	`
	err := cr.db.GetContext(ctx.Request().Context(), &roomID, query, userID, chatPartnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0, nil
		}
		return false, 0, errors.WithStack(err)
	}
	return true, roomID, nil
}

func (cr *chatRepository) UpdateRoomAccessTime(ctx echo.Context, roomID, userID int) error {
	query := `
		UPDATE entries SET joined_at = CURRENT_TIMESTAMP WHERE chat_room_id = $1 AND user_id = $2;
	`
	_, err := cr.db.ExecContext(ctx.Request().Context(), query, roomID, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
