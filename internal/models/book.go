package models

type Condition int

type IBookBuilder interface {
	SetID(id string) *BookBuilder
	SetName(name string) *BookBuilder
	SetAuthor(author string) *BookBuilder
	SetDescription(description string) *BookBuilder
	SetPricing(pricing float64) *BookBuilder
	SetCondition(condition string) *BookBuilder
	SetLanguages(languages string) *BookBuilder
	SetPublisher(publisher string) *BookBuilder
	SetYear(year string) *BookBuilder
	SetLink(link string) *BookBuilder
	SetISBN(isbn string) *BookBuilder
	Build() *Book
}
type Book struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Pricing     float64 `json:"pricing"`
	Condition   string  `json:"condition"`
	Languages   string  `json:"languages"`
	Publisher   string  `json:"publisher"`
	Year        string  `json:"year"`
	Link        string  `json:"link"`
	ISBN        string  `json:"isbn"`
}

type BookBuilder struct {
	book *Book
}

func NewBookBuilder() IBookBuilder {
	return &BookBuilder{
		book: &Book{},
	}
}

func (b *BookBuilder) SetID(id string) *BookBuilder {
	b.book.ID = id
	return b
}

func (b *BookBuilder) SetName(name string) *BookBuilder {
	b.book.Name = name
	return b
}

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.book.Author = author
	return b
}

func (b *BookBuilder) SetDescription(description string) *BookBuilder {
	b.book.Description = description
	return b
}

func (b *BookBuilder) SetPricing(pricing float64) *BookBuilder {
	b.book.Pricing = pricing
	return b
}

func (b *BookBuilder) SetCondition(condition string) *BookBuilder {
	b.book.Condition = condition
	return b
}

func (b *BookBuilder) SetLanguages(languages string) *BookBuilder {
	b.book.Languages = languages
	return b
}

func (b *BookBuilder) SetPublisher(publisher string) *BookBuilder {
	b.book.Publisher = publisher
	return b
}

func (b *BookBuilder) SetYear(year string) *BookBuilder {
	b.book.Year = year
	return b
}

func (b *BookBuilder) SetLink(link string) *BookBuilder {
	b.book.Link = link
	return b
}

func (b *BookBuilder) SetISBN(isbn string) *BookBuilder {
	b.book.ISBN = isbn
	return b
}

func (b *BookBuilder) Build() *Book {
	return b.book
}
