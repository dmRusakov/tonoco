package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return value, nil
}
