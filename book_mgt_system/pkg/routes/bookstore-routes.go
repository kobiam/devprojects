package routes

import (
	"github.com/gorilla/mux"
	"github.com/kobiam/devprojects/book_mgt_system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}