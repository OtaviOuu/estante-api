package repositories

import (
	"database/sql"

	"github.com/OtaviOuu/estante-api/internal/database"
	"github.com/OtaviOuu/estante-api/internal/models"
)

type IBookRepository interface{}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() (*BookRepository, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &BookRepository{
		db: db.DB,
	}, nil
}

func (r *BookRepository) GetAllWithPagination(page int, pageSize int) ([]*models.Book, error) {
	query := `
	SELECT book_title, book_description, book_price, condition, category, author, language, publisher, year, isbn, id 
	FROM books
	LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(query, pageSize, (page*pageSize)-pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var (
			title, desc, author, language, category, publisher, isbn, id sql.NullString
			year                                                         sql.NullString
			price                                                        sql.NullFloat64
			condition                                                    sql.NullString
		)

		if err := rows.Scan(&title, &desc, &price, &condition, &category, &author, &language, &publisher, &year, &isbn, &id); err != nil {
			return nil, err
		}

		// n sei como fazer melhor q isso
		if !year.Valid {
			year.String = ""
		}

		if !isbn.Valid {
			isbn.String = ""
		}

		book := models.NewBookBuilder().
			SetName(title.String).
			SetDescription(desc.String).
			SetPricing(price.Float64).
			SetCondition(condition.String).
			SetAuthor(author.String).
			SetLanguages(language.String).
			SetPublisher(publisher.String).
			SetYear(year.String).
			SetISBN(isbn.String).
			SetID(id.String).
			Build()

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
