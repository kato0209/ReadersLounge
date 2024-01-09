-- +goose Up
CREATE TABLE IF NOT EXISTS post_likes (
    post_like_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id)
);

CREATE TABLE IF NOT EXISTS comment_likes (
    comment_like_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    comment_id integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
    UNIQUE (user_id, comment_id)
);