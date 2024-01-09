package models

import "time"

type PostLike struct {
	PostLikeID int       `db:"post_like_id"`
	User       User      `db:"user"`
	CreatedAt  time.Time `db:"created_at"`
}
