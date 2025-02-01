package services

import (
	"github.com/OtaviOuu/exatas-estante-api/internal/models"
	"github.com/OtaviOuu/exatas-estante-api/internal/repositories"
)

type IBookService interface {
	getAllWithPagination(page int, pageSize int) ([]string, error)
	getByName(name string) ([]string, error)
}

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService() *BookService {
	db, err := repositories.NewBookRepository()
	if err != nil {
		return nil
	}
	return &BookService{
		repo: db,
	}
}

func (s *BookService) GetAllWithPagination(page int, pageSize int) ([]*models.Book, error) {
	foundBooks, err := s.repo.GetAllWithPagination(page, pageSize)
	if err != nil {
		return nil, err
	}

	return foundBooks, nil
}

func (s *BookService) getByName(name string) ([]string, error) {
	return nil, nil
}
