package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/njovanovic-se/url-shortner/handler"
	"github.com/njovanovic-se/url-shortner/store"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("environment file not loaded: %w", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set.")
	}

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

	dbConfig := store.DatabaseConfig{
		MaxOpenConnections:  25,
		MaxIdleConnections:  10,
		ConnMaxLifetime:     10 * time.Minute,
		ConnMaxIdleLifetime: 5 * time.Minute,
	}
	dbConfig.ApplyDefaults()

	db, err := store.NewPostgresDB(dsn, dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	store.NewUrlShortenerRepositoryImpl(db)

	err = r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
