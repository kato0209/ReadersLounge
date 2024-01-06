-- +goose Up
CREATE TABLE IF NOT EXISTS chat_rooms (
    chat_room_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS entries (
    entry_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    chat_room_id integer NOT NULL,
    user_id integer NOT NULL,
    joined_at timestamptz NOT NULL DEFAULT current_timestamp,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(chat_room_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS chat_messages (
    chat_message_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    chat_room_id integer NOT NULL,
    content text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(chat_room_id) ON DELETE CASCADE
);