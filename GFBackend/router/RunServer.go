package router

import (
	"GFBackend/config"
	"GFBackend/docs"
	"GFBackend/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
)

var AppRouter *gin.Engine

func RunServer() {
	AppRouter = gin.Default()

	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	AppRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := AppRouter.Run(":" + strconv.Itoa(config.AppConfig.Server.Port))
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
		panic("server error")
	}
}
