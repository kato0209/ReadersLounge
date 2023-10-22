-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
    comment_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at date NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS comment_details (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    comment_id integer NOT NULL,
    content char(255) NOT NULL,
    updated_at date,
    FOREIGN KEY (comment_id) REFERENCES comments(comment_id)
);