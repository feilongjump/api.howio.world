package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return Client.Set(context.Background(), key, value, expiration)
}

func Get(key string) string {
	return Client.Get(context.Background(), key).Val()
}
