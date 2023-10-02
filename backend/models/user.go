package models

import "time"

type User struct {
	UserID       int
	CreatedAt    time.Time
	Name         string
	ProfileText  string
	ProfileImage string
	UpdatedAt    time.Time
	IdentityType string
	Identifier   string
	Credential   string
}

type UserResponse struct {
	UserID int
}
