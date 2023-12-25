package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func Set(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) error {
	err := client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
