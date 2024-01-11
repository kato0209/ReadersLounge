package models

import "time"

type PostLike struct {
	PostLikeID int       `db:"post_like_id"`
	User       User      `db:"user"`
	CreatedAt  time.Time `db:"created_at"`
}

type CommentLike struct {
	CommentLikeID int       `db:"comment_like_id"`
	User          User      `db:"user"`
	CreatedAt     time.Time `db:"created_at"`
}
