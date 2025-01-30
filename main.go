package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func db_connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getBooks(page int, pageSize int, db *sql.DB) []string {
	rows, err := db.Query("SELECT book_title FROM books LIMIT ? OFFSET ?", pageSize, (page-1)*pageSize)
	if err != nil {
		return nil
	}
	defer rows.Close()

	titles := []string{}
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil
		}
		titles = append(titles, title)
	}
	return titles
}

func getBookByName(name string, db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT book_title FROM books WHERE book_title LIKE ? OR author LIKE ? OR book_description LIKE ?", "%"+name+"%", "%"+name+"%", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := []string{}
	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		if err != nil {
			return nil, err
		}

		matches = append(matches, title)
	}
	return matches, nil
}

func main() {
	db, err := db_connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	books, err := getBookByName("IMPA", db)
	if err != nil {
		log.Println(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
