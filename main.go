package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

var books []Book

// gets all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//gets one book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// creats one booke
func createBooks(w http.ResponseWriter, r *http.Request) {

}

// updates book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// deletes book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "a902l", Title: "book 1", Author: &Author{Firstname: "jake", Lastname: "Luecke"}})
	books = append(books, Book{ID: "2", Isbn: "gsdgw", Title: "book 2", Author: &Author{Firstname: "steve", Lastname: "smith"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
