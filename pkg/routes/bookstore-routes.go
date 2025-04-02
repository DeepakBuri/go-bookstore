package routes

import (
	"github.com/deepakburi/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// BookstoreRoutes sets up the routes for the bookstore API
func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
