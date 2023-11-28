package models

type Book struct {
	BookID      int    `db:"book_id"`
	ISBNcode    string `db:"ISBNcode"`
	Title       string `db:"title"`
	Author      string `db:"author"`
	Price       int    `db:"price"`
	Publisher   string `db:"publisher"`
	PublishedAt string `db:"published_at"`
	ItemURL     string `db:"item_url"`
	Image       string `db:"image"`
}

type BooksGenre struct {
	ID             int    `db:"id"`
	BooksGenreID   string `db:"books_genre_id"`
	BooksGenreName string `db:"books_genre_name"`
	GenreLevel     int    `db:"genre_level"`
	ParentGenreID  string `db:"parent_genre_id"`
}

type RakutenApiBooksResponse struct {
	Items []struct {
		Item struct {
			ISBNcode    string `json:"isbn"`
			Title       string `json:"title"`
			Author      string `json:"author"`
			Price       int    `json:"itemPrice"`
			Publisher   string `json:"publisherName"`
			PublishedAt string `json:"salesDate"`
			ItemURL     string `json:"itemUrl"`
			Image       string `json:"largeImageUrl"`
		} `json:"Item"`
	} `json:"Items"`
}

type RakutenApiBooksGenreResponse struct {
	Children []struct {
		Child struct {
			BooksGenreID   string `json:"booksGenreId"`
			BooksGenreName string `json:"booksGenreName"`
			GenreLevel     int    `json:"genreLevel"`
		} `json:"child"`
	}
}
