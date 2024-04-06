package main

import (
	"github.com/books/server"
	"net/http"
)

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if the request is for the OPTIONS method (pre-flight request)
		// and return with OK status if it is
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next handler/middleware in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/add-book", server.AddBookHandler)
	mux.HandleFunc("/api/remove-book", server.RemoveBookHandler)
	mux.HandleFunc("/api/update-book", server.UpdateBookHandler)
	mux.HandleFunc("/api/search", server.SearchBookHandler)

	// Wrap the mux router with the corsMiddleware
	wrappedMux := corsMiddleware(mux)

	// Use wrappedMux with the CORS middleware applied
	http.ListenAndServe(":8080", wrappedMux)
}
