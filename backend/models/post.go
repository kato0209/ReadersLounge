package models

import (
	"time"
)

type Post struct {
	PostID    int `db:"post_id"`
	User      User
	Content   string `db:"content"`
	Rating    int    `db:"rating"`
	Image     *PostImage
	CreatedAt time.Time `db:"created_at"`
	Book      Book
	Like      []PostLike
}

type PostImage struct {
	Source       *[]byte
	FileName     *string `db:"image"`
	EncodedImage *string
}
