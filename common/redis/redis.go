package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    context.Context = context.Background()
)

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "my-redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := Client.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	log.Println("Connected to Redis")
}
