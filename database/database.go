package database

import (
	"database/sql"
	"log"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://test:test123@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to Connect to Database")
	}
	return db
}
