package services

import (
	"aplikasi-layanan-umum/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	users, err := controllers.GetAllUser()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error Fetching User")
	}
	c.JSON(http.StatusOK, users)
}
