package main

import (
	"net/http"

	"github.com/WillMcCall/Bible/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": api.GetWordJSON("God", 3),
		})
	})
	r.Run(":8081")
}
