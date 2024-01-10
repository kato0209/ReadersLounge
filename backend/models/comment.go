package models

import (
	"time"
)

type Comment struct {
	CommentID int `db:"comment_id"`
	User      User
	Post      Post
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
