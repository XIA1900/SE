package main

import (
	"backend/config"
	"backend/log"
	"backend/model"
	"backend/router"
	"go.uber.org/zap"
)

func main() {
	// Importance: first step is to load configuration parameters!!!
	config.LoadAppConfig()

	// necessary components init
	log.InitLog()
	defer log.Logger.Sync()
	log.Logger.Debug("debug log test", zap.String("param", "http"))
	log.Logger.Info("info log test")
	log.Logger.Warn("warn test")
	log.Logger.Error("error test")

	model.InitDB()
	appConfig := config.GetAppConfig()
	r := router.InitRouter()

	// run system
	err := r.Run(":" + appConfig.Server.Port)
	if err != nil {
		panic("server error")
	}
}
