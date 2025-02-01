package main

import (
	"net/http"

	"github.com/OtaviOuu/exatas-estante-api/internal/handlers"
	"github.com/OtaviOuu/exatas-estante-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	r.Get("/books", bookHandler.GetAllWithPagination)

	http.ListenAndServe(":8080", r)
}
