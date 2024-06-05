package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://test:test123@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("Failed Connect to DB")
	}
	return db
}
