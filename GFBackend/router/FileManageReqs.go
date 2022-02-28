package router

import "github.com/gin-gonic/gin"

func InitFileManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	fileManageController, _ := InitializeFileManageController()

	fileManageReqsGroup := baseGroup.Group("/file")
	{
		fileManageReqsGroup.POST("/scan", fileManageController.ScanFiles)
	}

	return nil
}
