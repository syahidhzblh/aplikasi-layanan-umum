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

	//Include static file
	r.Static("/scss", "./views/scss")
	r.Static("/css", "./views/css")
	r.Static("/js", "./views/js")

	//Controller
	r.GET("/", controllers.LoginController)
	r.GET("/index", controllers.IndexController)

	http.ListenAndServe(":9000", r) // listen and serve on 0.0.0.0:9000
}
