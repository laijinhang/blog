package redis

import (
	"blog/model/config"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	//连接服务器
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host + ":" + config.RedisConfig.Port, // use default Addr
		Password: config.RedisConfig.Pwd,                                  // no password set
		DB:       config.RedisConfig.DB,                                   // use default DB
	})
	redisClient.Ping()
}
