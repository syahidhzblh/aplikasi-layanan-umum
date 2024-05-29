package main

import (
	"aplikasi-layanan-umum/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//Include static css file
	r.Static("/css", "./views/css")

	//Controller
	r.GET("/", controllers.LoginController)
	r.GET("/index", controllers.IndexController)

	http.ListenAndServe(":9000", r) // listen and serve on 0.0.0.0:9000
}
