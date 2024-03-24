package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NewIndex(c *gin.Context) {
	fp := filepath.Join("views", "index.html")
	content, err := os.ReadFile(fp)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error Reading File")
		return
	}
	c.Data(http.StatusOK, "text/html", content)
}

// func NewIndex() func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fp := filepath.Join("views", "index.html")
// 		tmpl, err := template.ParseFiles(fp)

// 		if err != nil {
// 			w.Write([]byte(err.Error()))
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		err = tmpl.Execute(w, nil)
// 		if err != nil {
// 			w.Write([]byte(err.Error()))
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
