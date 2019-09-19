package cache

import (
	"fmt"
	"os"
	"strconv"
	"github.com/go-redis/redis"
)
// RedisClient Redis缓存客户端单例
var Redis *redis.Client

// Init 在中间件中初始化redis链接
func Init() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(fmt.Sprintf("Redis 连接异常！ 错误信息: %s", err))

	}

	Redis = client
}
