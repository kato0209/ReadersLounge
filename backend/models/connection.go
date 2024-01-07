package models

import "time"

type Connection struct {
	ConnectionID int `db:"connection_id"`
	Follower     User
	Following    User
	CreatedAt    time.Time `db:"created_at"`
}
