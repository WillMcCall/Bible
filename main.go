package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/WillMcCall/Bible/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")  // Serves static files
	r.LoadHTMLGlob("templates/**/*") // Serves templates

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home", gin.H{
			"data": "test",
		})
	})

	r.GET("/:word/:page", func(c *gin.Context) {
		word := c.Param("word")
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": api.GetWordJSON(word, page),
		})
	})
	r.Run(":8081")
}
