package routes

import (
	"github.com/azad-ali786/bookmanagementsystem/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreROutes = func(router *mux.Router){
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByBookId).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}