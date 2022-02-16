package router

import (
	"GFBackend/controller"
	"github.com/gin-gonic/gin"
)

func InitUserManageReqs() *gin.RouterGroup {
	appRouter := AppRouter

	userManageReqsGroup := appRouter.Group("/user")
	{
		userManageReqsGroup.POST("/register", controller.UserRegister)
		userManageReqsGroup.POST("/login", controller.UserLogin)
		userManageReqsGroup.POST("/logout", controller.UserLogout)
		userManageReqsGroup.POST("/password", controller.UserUpdatePassword)
	}

	return userManageReqsGroup
}
