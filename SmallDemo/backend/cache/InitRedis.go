package cache

import (
	"backend/config"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis() {
	appConfig := config.GetAppConfig()

	rdb = redis.NewClient(&redis.Options{
		Addr:     appConfig.Redis.IP + ":" + appConfig.Redis.Port,
		Password: appConfig.Redis.Password,
		DB:       appConfig.Redis.DB,
	})
}

func GetRDB() *redis.Client {
	return rdb
}
