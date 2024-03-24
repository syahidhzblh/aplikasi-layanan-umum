package main

import (
	"aplikasi-layanan-umum/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.SampleController)
	r.GET("/index", controllers.NewIndex)

	http.ListenAndServe(":9000", r) // listen and serve on 0.0.0.0:9000
}
