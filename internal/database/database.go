package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Definindo a estrutura Database
type Database struct {
	*sql.DB
}

// Nova função para conectar ao PostgreSQL
func NewDatabase() (*Database, error) {
	connStr := "postgresql://postgres:ZmaQuCHiRxfYVXhmdynawjYolUfamNHO@viaduct.proxy.rlwy.net:30025/railway"

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
