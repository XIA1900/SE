package router

import (
	"GFBackend/controller"
	"github.com/gin-gonic/gin"
)

func InitUserManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	userManageReqsGroup := baseGroup.Group("/user")
	{
		userManageReqsGroup.POST("/register", controller.RegularRegister)
		userManageReqsGroup.POST("/login", controller.UserLogin)
		userManageReqsGroup.POST("/logout", controller.UserLogout)
		userManageReqsGroup.POST("/password", controller.UserUpdatePassword)
		userManageReqsGroup.POST("/update", controller.UserUpdate)

		adminReqsGroup := userManageReqsGroup.Group("/admin")
		{
			adminReqsGroup.POST("/register", controller.AdminRegister)
			adminReqsGroup.POST("/delete", controller.UserDelete)
		}
	}

	return userManageReqsGroup
}
