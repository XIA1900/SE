package router

import (
	"GFBackend/config"
	"GFBackend/docs"
	"GFBackend/logger"
	"GFBackend/middleware/interceptor"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
)

var AppRouter *gin.Engine

func RunServer() {
	appConfig := config.AppConfig

	interceptor.InitNonAuthReq()
	AppRouter = gin.Default()
	AppRouter.Use(interceptor.AuthInterceptor())

	docs.SwaggerInfo_swagger.BasePath = appConfig.Server.BasePath
	AppRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseGroup := AppRouter.Group(appConfig.Server.BasePath)
	{
		InitUserManageReqs(baseGroup)
	}

	err := AppRouter.Run(":" + strconv.Itoa(appConfig.Server.Port))
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
		panic("server error")
	}
}
