package models

import "time"

type User struct {
	UserID       int       `db:"user_id"`
	Name         string    `db:"name"`
	ProfileText  *string   `db:"profile_text"`
	ProfileImage string    `db:"profile_image"`
	IdentityType string    `db:"identity_type"`
	Identifier   string    `db:"identifier"`
	Credential   string    `db:"credential"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

/*
type User struct {
	UserID       int     `db:"user_id"`
	Name         string  `db:"name"`
	ProfileText  *string `db:"profile_text"`
	ProfileImage *ProfileImage
	IdentityType string    `db:"identity_type"`
	Identifier   string    `db:"identifier"`
	Credential   string    `db:"credential"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
*/

type ProfileImage struct {
	Source       []byte
	FileName     *string `db:"image"`
	EncodedImage *string
}
