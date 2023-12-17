package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host      string
	Port      string
	Password  string
	MaxMemory string
}

func Connect(ctx context.Context, cfg *Config) (*redis.Client, error) {
	// Create a Redis client.
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	// Ping the Redis server to verify the connection.
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	// Return the Redis client.
	return client, nil
}
