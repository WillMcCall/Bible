package main

import (
	"net/http"

	"github.com/WillMcCall/Bible/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")  // Serves static files
	r.LoadHTMLGlob("templates/**/*") // Serves templates

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home", gin.H{
			"data": api.GetWordJSON("abide"),
		})
	})

	r.GET("/:word", func(c *gin.Context) {
		word := c.Param("word")
		c.JSON(http.StatusOK, gin.H{
			"data": api.GetWordJSON(word),
		})
	})
	r.Run(":8081")
}
