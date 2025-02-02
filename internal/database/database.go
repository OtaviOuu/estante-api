package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
