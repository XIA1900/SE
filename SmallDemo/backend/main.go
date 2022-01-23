package main

import (
	"backend/config"
	"backend/model"
	"backend/router"
)

func main() {
	// Importance: first step is to load configuration parameters!!!
	config.LoadAppConfig()

	// necessary components init
	model.InitDB()
	appConfig := config.GetAppConfig()
	r := router.InitRouter()

	// run system
	err := r.Run(":" + appConfig.Server.Port)
	if err != nil {
		panic("server error")
	}
}
