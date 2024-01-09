package models

import "time"

type PostLike struct {
	PostLikeID int       `db:"post_like_id"`
	CreatedAt  time.Time `db:"created_at"`
}
