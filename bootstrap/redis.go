package bootstrap

import "github.com/feilongjump/api.howio.world/internal/redis"

// SetupRedisDB 初始化 redis
func SetupRedisDB() {

	redis.ConnectRedis()
}
