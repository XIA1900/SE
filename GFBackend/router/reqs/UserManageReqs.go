package reqs

import (
	"github.com/gin-gonic/gin"
)

func InitUserManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	userManageController, _ := InitializeUserManageController()

	userManageReqsGroup := baseGroup.Group("/user")
	{
		userManageReqsGroup.POST("/register", userManageController.RegularRegister)
		userManageReqsGroup.POST("/login", userManageController.UserLogin)
		userManageReqsGroup.POST("/logout", userManageController.UserLogout)
		userManageReqsGroup.POST("/password", userManageController.UserUpdatePassword)
		userManageReqsGroup.POST("/update", userManageController.UserUpdate)
		userManageReqsGroup.POST("/follow", userManageController.UserFollow)
		userManageReqsGroup.POST("/unfollow", userManageController.UserUnfollow)

		adminReqsGroup := userManageReqsGroup.Group("/admin")
		{
			adminReqsGroup.POST("/register", userManageController.AdminRegister)
			adminReqsGroup.POST("/delete", userManageController.UserDelete)
		}
	}

	return userManageReqsGroup
}
