package routes

import (
	"github.com/deepakburi/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// BookstoreRoutes sets up the routes for the bookstore API
func BookstoreRoutes(r *mux.Router) {
	r.HandleFunc("/bookstore", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/bookstore/{bookID}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/bookstore", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/bookstore/{bookID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/bookstore/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
