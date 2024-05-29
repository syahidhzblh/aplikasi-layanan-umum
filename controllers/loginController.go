package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	fp := filepath.Join("views", "login.html")
	content, err := os.ReadFile(fp)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error Reading File")
		return
	}
	c.Data(http.StatusOK, "text/html", content)
}
