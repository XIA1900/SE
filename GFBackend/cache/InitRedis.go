package cache

import (
	"GFBackend/config"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var RDB *redis.Client

func InitRedis() {
	appConfig := config.AppConfig

	RDB = redis.NewClient(&redis.Options{
		Addr:     appConfig.Redis.IP + ":" + strconv.Itoa(appConfig.Redis.Port),
		Password: appConfig.Redis.Password,
		DB:       appConfig.Redis.DB,
	})
}
