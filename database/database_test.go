package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://test:test123@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	defer db.Close()
}

func TestSelectTable(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, role FROM users;"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, role string
		err = rows.Scan(&id, &name, &role)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name: ", name)
		fmt.Println("Roles: ", role)
	}
}
