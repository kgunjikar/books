package server

import (
	"encoding/json"
	"fmt"
	"github.com/books/DB"
	"github.com/books/schema"
	"log"
	"net/http"
)

var db DB.DBHandle

func init() {
	db = DB.GetDB()
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello World\n")
}

func SearchBookHandler(w http.ResponseWriter, r *http.Request) {
	var searchReq schema.SearchRequest
	log.Printf("Searching Book \n")
	err := json.NewDecoder(r.Body).Decode(&searchReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(searchReq)

	// Perform search logic here. For now, just log and send a response.
	log.Printf("Searching for book: %s\n", searchReq.BookName)
	DB.SearchBooks(db, searchReq.BookName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Search completed"})
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book schema.Book
	log.Printf("Adding Book \n")
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add book to the database or storage. For now, just log the details.
	log.Printf("Adding book: %+v\n", book)
	DB.AddBook(db, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book added successfully"})
}

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		BookId int `json:"bookId"`
	}
	log.Printf("Removing Book \n")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Logic to remove the book from the database using request.BookId.
	// For now, log the action.
	log.Printf("Removing book with ID: %s\n", request.BookId)
	DB.RemoveBook(db, request.BookId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book removed successfully"})
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book schema.Book
	fmt.Println("Updating book")
	// Assuming Book struct has a BookId field for identification.
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Logic to update the book in the database using book details.
	// For now, log the action.
	log.Printf("Updating book: %+v\n", book)
	DB.UpdateBook(db, book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated successfully"})
}
