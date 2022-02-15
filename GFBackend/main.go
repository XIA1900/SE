package main

import (
	"GFBackend/config"
	"GFBackend/elasticsearch"
	"GFBackend/logger"
	"GFBackend/middleware/auth"
	"GFBackend/model"
	"GFBackend/router"
)

func main() {
	// Components Initialization
	config.InitConfig()
	logger.InitAppLogger()
	defer logger.AppLogger.Sync()
	model.InitDB()
	// cache.InitRedis()
	auth.InitCasbin()
	elasticsearch.InitES()
	router.RunServer()
}
