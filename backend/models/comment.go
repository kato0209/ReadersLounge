package models

import (
	"time"
)

type Comment struct {
	CommentID int `db:"comment_id"`
	User      User
	Post      Post
	Content   string `db:"content"`
	Likes     []CommentLike
	CreatedAt time.Time `db:"created_at"`
}
