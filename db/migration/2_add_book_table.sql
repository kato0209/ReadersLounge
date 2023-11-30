-- +goose Up
CREATE TABLE IF NOT EXISTS books (
    book_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ISBNcode varchar(100) NOT NULL UNIQUE,
    title varchar(100) NOT NULL,
    author varchar(100) NOT NULL,
    price integer NOT NULL,
    publisher varchar(100) NOT NULL,
    published_at varchar(100) NOT NULL,
    item_url varchar(255) NOT NULL,
    image varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS bookmarks (
    bookmark_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

CREATE TABLE IF NOT EXISTS book_archives (
    book_archive_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

CREATE TABLE IF NOT EXISTS books_genres (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    books_genre_id VARCHAR(255) NOT NULL UNIQUE,
    books_genre_name VARCHAR(255) NOT NULL,
    genre_level INT NOT NULL,
    parent_genre_id VARCHAR(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    updated_at timestamptz NOT NULL DEFAULT current_timestamp
);