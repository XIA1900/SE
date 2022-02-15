package router

import (
	"GFBackend/controller"
	"github.com/gin-gonic/gin"
)

func InitUserManageReqs() *gin.RouterGroup {
	appRouter := AppRouter

	userReqsGroup := appRouter.Group("/user")
	{
		userReqsGroup.POST("/register", controller.UserRegister)
		userReqsGroup.POST("/login", controller.UserLogin)
		userReqsGroup.POST("/logout", controller.UserLogout)
		userReqsGroup.POST("/password", controller.UserUpdatePassword)
	}

	return userReqsGroup
}
