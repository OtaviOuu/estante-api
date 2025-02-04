package services

import (
	"github.com/OtaviOuu/estante-api/internal/models"
	"github.com/OtaviOuu/estante-api/internal/repositories"
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
	if page <= 0 {
		page = 1
	}

	if pageSize < 16 {
		pageSize = 16
	}

	foundBooks, err := s.repo.GetAllWithPagination(page, pageSize)
	if err != nil {
		return nil, err
	}

	return foundBooks, nil
}

func (s *BookService) GetByKeyWord(page int, pageSize int, keyword string) ([]*models.Book, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize < 16 {
		pageSize = 16
	}

	foundBooks, err := s.repo.GetByKeyWord(page, pageSize, keyword)
	if err != nil {
		return nil, err
	}

	return foundBooks, nil
}
