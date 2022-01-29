package router

import (
	"backend/service"
	"github.com/gin-gonic/gin"
)

func InitAdminGroup(r *gin.Engine) {
	userGroup := r.Group("/admin")
	{
		userGroup.POST("/login", service.AdminLogin)
		userGroup.POST("/register", service.AddAdmin)
		userGroup.POST("/changePassword", service.UpdateAdmin)
		userGroup.POST("/deleteAdmin", service.DeleteAdmin)
	}
}
