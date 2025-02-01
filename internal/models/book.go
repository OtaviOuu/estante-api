package models

type Condition int

const (
	Usado Condition = iota
	Novo
)

type IBookBuilder interface {
	SetID(id string) *BookBuilder
	SetName(name string) *BookBuilder
	SetAuthor(author string) *BookBuilder
	SetDescription(description string) *BookBuilder
	SetPricing(pricing float64) *BookBuilder
	SetCondition(condition Condition) *BookBuilder
	SetLanguages(languages string) *BookBuilder
	SetPublisher(publisher string) *BookBuilder
	SetYear(year int) *BookBuilder
	SetLink(link string) *BookBuilder
	SetISBN(isbn string) *BookBuilder
	Build() *Book
}

type Book struct {
	ID          string
	Name        string
	Author      string
	Description string
	Pricing     float64
	Condition   Condition
	Languages   string
	Publisher   string
	Year        int
	Link        string
	ISBN        string
}

type BookBuilder struct {
	book Book
}

func NewBookBuilder() *BookBuilder {
	return &BookBuilder{}
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

func (b *BookBuilder) SetCondition(condition Condition) *BookBuilder {
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

func (b *BookBuilder) SetYear(year int) *BookBuilder {
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

func (b *BookBuilder) Build() Book {
	return b.book
}
