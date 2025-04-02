package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deepakburi/go-bookstore/pkg/models"
	"github.com/deepakburi/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get all books from the database
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Fprintf(w, "Get all books")

}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["bookID"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}

	book, db := models.GetBookByID(bookID)
	if db.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}

	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Fprintf(w, "Get book by ID: %s", vars["bookID"])
	// Implementation for getting a book by ID
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := utils.ParseBody(r, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}
	bookCreated := book.CreateBook()
	res, err := json.Marshal(bookCreated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating book"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	fmt.Fprintf(w, "Book created successfully")
	// Implementation for creating a new book
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var updatedBook models.Book
	err := utils.ParseBody(r, &updatedBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	bookID, err := strconv.ParseInt(vars["bookID"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}

	bookDetails, db := models.GetBookByID(bookID)
	if db.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found"))
		return
	}
	if bookDetails.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookID, err := strconv.ParseInt(vars["bookID"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid book ID"))
		return
	}

	book := models.DeleteBook(bookID)
	res, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Fprintf(w, "Book deleted successfully")
	// Implementation for deleting a book
}
