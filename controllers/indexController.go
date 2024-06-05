package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {

	//Index View
	fp := filepath.Join("views", "index.html")
	content, err := os.ReadFile(fp)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error Reading File")
		return
	}
	c.Data(http.StatusOK, "text/html", content)
}
