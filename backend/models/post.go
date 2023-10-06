package models

type Post struct {
	PostID    int `db:"post_id"`
	User      User
	Content   string `db:"content"`
	Rating    int    `db:"rating"`
	Image     string `db:"image"`
	CreatedAt string `db:"created_at"`
	Book      Book
}
