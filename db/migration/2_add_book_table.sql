-- +goose Up
CREATE TABLE IF NOT EXISTS books (
    book_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ISBNcode char(100) NOT NULL,
    title char(100) NOT NULL,
    author char(100) NOT NULL,
    price char(100) NOT NULL,
    publisher char(100) NOT NULL,
    published_at date NOT NULL,
    item_url char(255) NOT NULL,
    image char(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS bookmarks (
    bookmark_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    created_at date NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

CREATE TABLE IF NOT EXISTS book_archives (
    book_archive_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    created_at date NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);