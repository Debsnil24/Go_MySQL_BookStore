package routes

import (
	"github.com/gorilla/mux"
	"github.com/Debsnil24/Go_MySQL_BookStore.git/pkg/controller"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.getBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controller.getBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}",controller.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controller.deleteBook).Methods("DELETE")

}