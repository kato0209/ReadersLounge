-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    post_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    created_at date NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

CREATE TABLE IF NOT EXISTS post_details (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    post_id integer NOT NULL,
    content char(255) NOT NULL,
    rating int NOT NULL,
    image char(255),
    updated_at date,
    FOREIGN KEY (post_id) REFERENCES posts(post_id)
);