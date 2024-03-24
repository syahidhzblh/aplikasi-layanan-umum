package database

import (
	"database/sql"
	"log"
	"os"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Failed to Connect to Database")
	}
	return db
}
