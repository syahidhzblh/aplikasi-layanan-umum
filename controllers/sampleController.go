package controllers

import "github.com/gin-gonic/gin"

func SampleController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Test Page",
	})
}
