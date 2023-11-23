-- +goose Up
CREATE TABLE IF NOT EXISTS connections (
    connection_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    following_id integer NOT NULL,
    follower_id integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (following_id) REFERENCES users(user_id),
    FOREIGN KEY (follower_id) REFERENCES users(user_id)
);