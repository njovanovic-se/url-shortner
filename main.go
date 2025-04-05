package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handlers "github.com/njovanovic-se/url-shortner/handler"
	"github.com/njovanovic-se/url-shortner/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortner !",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	r.GET("/:short-url", func(c *gin.Context) {
		handlers.HandlerShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
