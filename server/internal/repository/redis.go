package repository

import (
	"context"
	"fmt"
	"healtech-backend/server/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis(cfg *config.Config) {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // No password set
		DB:       0,  // Use default DB
	})

	// Test connection
	ctx := context.Background()
	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Redis connection established")
}
