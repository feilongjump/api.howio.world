package redis

import (
	"context"
	"github.com/feilongjump/api.howio.world/internal/config"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func ConnectRedis() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.GetString("database.redis.addr"),
		Password: config.GetString("database.redis.password"),
		DB:       config.GetInt("database.redis.db"),
	})

	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return Client
}
