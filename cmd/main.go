package main

import (
	"net/http"
	"os"

	"github.com/OtaviOuu/estante-api/internal/handlers"
	"github.com/OtaviOuu/estante-api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	r := chi.NewRouter()

	r.Use(cors.Default().Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/books", bookHandler.GetAllWithPagination)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe("0.0.0.0:"+port, r)
	if err != nil {
		panic(err)
	}
}
