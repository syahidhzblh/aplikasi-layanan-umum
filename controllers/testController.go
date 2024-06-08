package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	fp := filepath.Join("views", "features", "index.html")
	content, err := os.ReadFile(fp)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error read file")
		return
	}
	c.Data(http.StatusOK, "text/html", content)
}
