package main

import (
	"aplikasi-layanan-umum/controllers"
	"aplikasi-layanan-umum/services"
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

	//Main Route
	r.GET("/", controllers.LoginController)
	r.GET("/index", controllers.IndexController)
	r.GET("/users", services.GetUser)

	http.ListenAndServe(":9000", r) // listen and serve on 0.0.0.0:9000
}
