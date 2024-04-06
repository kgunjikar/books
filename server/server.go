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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Write the HTML content to the response writer
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var creds schema.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the username and password are "root"
    if creds.UserName == "root" && creds.PassWord == "root" {
        // Respond with a success message and load the allinone.html page
	log.Printf("Logged in successfully")
	// w.Header().Set("Content-Type", "text/html")
        http.ServeFile(w, r, "html/allinone.html")
        return
    }

    // Respond with an error message for incorrect credentials
    http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "html/login.html")
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
