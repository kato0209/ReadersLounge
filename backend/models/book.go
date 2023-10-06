package models

import "time"

type Book struct {
	BookID      int       `db:"book_id"`
	ISBNcode    string    `db:"ISBNcode"`
	Title       string    `db:"title"`
	Author      string    `db:"author"`
	Price       string    `db:"price"`
	Publisher   string    `db:"publisher"`
	PublishedAt time.Time `db:"published_at"`
	ItemURL     string    `db:"item_url"`
	Image       string    `db:"image"`
}
