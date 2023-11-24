-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
    comment_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS comment_details (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    comment_id integer NOT NULL,
    content varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (comment_id) REFERENCES comments(comment_id)
);