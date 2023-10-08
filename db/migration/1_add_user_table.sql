-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS user_details (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    name char(20) NOT NULL,
    profile_text char(255),
    profile_image char(255) NOT NULL DEFAULT 'https://res.cloudinary.com/dvh5ehszr/image/upload/v1689442197/media/default_img.png',
    updated_at date,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS user_auths (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer NOT NULL,
    identity_type char(100) NOT NULL,
    identifier char(100) NOT NULL,
    credential char(100) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    UNIQUE (identifier)
);