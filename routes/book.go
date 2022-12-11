package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/mysql"
	"waysbook/repositories"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	BookRepository := repositories.RepositoryBook(mysql.DB)
	h := handlers.HandlerBook(BookRepository)

	r.HandleFunc("/books", h.FindBook).Methods("GET")
	r.HandleFunc("/book/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/book", middleware.Auth(middleware.UploadPdf(middleware.UploadFile(h.CreateBook)))).Methods("POST")
	r.HandleFunc("/book/{id}", middleware.Auth(middleware.UploadPdf(middleware.UploadFile(h.UpdateBook)))).Methods("PATCH")
	r.HandleFunc("/book/{id}", middleware.Auth(h.DeleteBook)).Methods("DELETE")
	r.HandleFunc("/books-promo", h.FindBookPromo).Methods("GET")
	r.HandleFunc("/books-regular", h.FindBookRegular).Methods("GET")
	r.HandleFunc("/book-promo/{id}", h.UpdateBookPromo ).Methods("PATCH")
	
}