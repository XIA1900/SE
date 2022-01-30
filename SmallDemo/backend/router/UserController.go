package router

import (
	"backend/service"
	"github.com/gin-gonic/gin"
)

func InitUserGroup(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", service.Login)
		userGroup.POST("/register", service.Register)
		userGroup.POST("/changePassword", service.ChangePassword)
		userGroup.POST("/deleteUser", service.DeleteUser)
	}
}
