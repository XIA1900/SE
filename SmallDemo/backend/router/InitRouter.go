package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Router Groups Init
	InitUserGroup(r)

	return r
}
