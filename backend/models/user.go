package models

import "time"

type User struct {
	UserID       int       `db:"user_id"`
	CreatedAt    time.Time `db:"created_at"`
	Name         string    `db:"name"`
	ProfileText  string    `db:"profile_text"`
	ProfileImage string    `db:"profile_image"`
	UpdatedAt    time.Time `db:"updated_at"`
	IdentityType string    `db:"identity_type"`
	Identifier   string    `db:"identifier"`
	Credential   string    `db:"credential"`
}

type UserResponse struct {
	UserID int
}
