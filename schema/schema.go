package schema

import "time"

// Book represents a book with its details.
type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publishDate"`
	Description string    `json:"description"`
}

// SearchRequest represents the search request payload.
type SearchRequest struct {
	BookName string `json:"name"`
}
