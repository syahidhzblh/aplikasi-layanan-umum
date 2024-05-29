package main

import (
	"aplikasi-layanan-umum/controllers"
	"aplikasi-layanan-umum/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.NewIndex)
	// r.GET("/index", controllers.NewIndex)

	r.GET("/run-query", func(c *gin.Context) {
		db := database.GetConnection()
		rows, err := db.Query("SELECT * FROM your_table")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var results []map[string]interface{}
		columns, _ := rows.Columns()
		for rows.Next() {
			row := make(map[string]interface{})
			columnPointers := make([]interface{}, len(columns))
			for i, _ := range columns {
				columnPointers[i] = new(interface{})
			}
			rows.Scan(columnPointers...)
			for i, colName := range columns {
				row[colName] = *(columnPointers[i].(*interface{}))
			}
			results = append(results, row)
		}

		c.JSON(http.StatusOK, results)
	})

	http.ListenAndServe(":9000", r) // listen and serve on 0.0.0.0:9000
}
