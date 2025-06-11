package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/njovanovic-se/url-shortner/models"
	"github.com/njovanovic-se/url-shortner/shortener"
	"github.com/njovanovic-se/url-shortner/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	err := store.Save(c.Request.Context(), &models.Shortener{
		ShortUrl:    shortUrl,
		OriginalUrl: creationRequest.LongUrl,
		UserId:      creationRequest.UserId,
	})

	if err != nil {
		panic(fmt.Errorf("failed to save url for userID: %s: %w", creationRequest.UserId, err))
	}

	err = store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	if err != nil {
		panic(fmt.Errorf("failed to save url to temporary cache: %w", err))
	}

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func HandlerShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("short-url")
	original_url, err := store.GetInitialUrl(shortUrl)

	if err != nil {
		fmt.Printf("failed to load data for short URL provided: %v", err)
	}

	c.Redirect(302, original_url)
	original_url, err = store.Load(c.Request.Context(), shortUrl)
	if err != nil {
		fmt.Printf("failed to load data for short URL provided: %v", err)
	}
	c.Redirect(302, original_url)
}
