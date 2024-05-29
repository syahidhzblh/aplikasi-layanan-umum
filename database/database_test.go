package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabase(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://test:test123@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("Failed Connect to DB")
	}
	defer db.Close()
}

func TestQuerySql(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://test:test123@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("Failed Connect to DB")
	}
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, username, password, role FROM users;"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, password, role string

		err = rows.Scan(&id, &username, &password, &role)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Username: ", username)
		fmt.Println("Password: ", password)
		fmt.Println("Role: ", role)
	}
}
