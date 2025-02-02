package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/OtaviOuu/estante-api/internal/services"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetAllWithPagination(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	pageIndex, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		http.Error(w, "Invalid size", http.StatusBadRequest)
		return
	}

	booksFound, err := h.service.GetAllWithPagination(pageIndex, pageSize)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error fetching books: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(booksFound)
}
