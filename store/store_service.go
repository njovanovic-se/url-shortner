package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StoreService struct {
	redisClient *redis.Client
}

var (
	storeService = &StoreService{}
	ctx          = context.Background()
)

func InitializeStore() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("\nFailed to save shortened url: error message: {%s} - original url {%s}; shortened url {%s}", err, originalUrl, shortUrl))
	}
	fmt.Print("\nSuccessfully saved shortened url")
}

func GetInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	fmt.Printf("\n{%s}", err)
	if err != nil {
		panic(fmt.Sprintf("\nFailed to retrieve original url: error message: {%s} - shortened url {%s}", err, shortUrl))
	}
	return result
}

const CacheDuration = 6 * time.Hour
