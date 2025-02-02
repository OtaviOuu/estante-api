package main

import (
	"net/http"
	"os"

	"github.com/OtaviOuu/estante-api/internal/handlers"
	"github.com/OtaviOuu/estante-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/books", bookHandler.GetAllWithPagination)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	http.ListenAndServe("0.0.0.0:"+port, r)
}
