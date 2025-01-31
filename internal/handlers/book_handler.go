package handlers

type BookHandler struct {
	service BookService
}

type BookService interface {
	getAllWithPagination(page int, pageSize int) ([]string, error)
	getByName(name string) ([]string, error)
}

func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}
