package DB

import (
	"database/sql"
	"github.com/books/schema"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DBHandle *sql.DB

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

// AddBook inserts a new book into the database.
func AddBook(db *sql.DB, book schema.Book) error {
	stmt, err := db.Prepare("INSERT INTO books(name, author, publish_date, description) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Name, book.Author, book.PublishDate, book.Description)
	return err
}

// RemoveBook deletes a book from the database by ID.
func RemoveBook(db *sql.DB, bookID int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", bookID)
	return err
}

// SearchBooks finds books that match the given name.
func SearchBooks(db *sql.DB, bookName string) ([]schema.Book, error) {
	rows, err := db.Query("SELECT id, name, author, publish_date, description FROM books WHERE name LIKE ?", "%"+bookName+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []schema.Book
	for rows.Next() {
		var b schema.Book
		if err := rows.Scan(&b.ID, &b.Name, &b.Author, &b.PublishDate, &b.Description); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// UpdateBook modifies details of an existing book.
func UpdateBook(db *sql.DB, book schema.Book) error {
	stmt, err := db.Prepare("UPDATE books SET name = ?, author = ?, publish_date = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Name, book.Author, book.PublishDate, book.Description, book.ID)
	return err
}
