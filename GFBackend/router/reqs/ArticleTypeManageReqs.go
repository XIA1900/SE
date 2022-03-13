package reqs

import "github.com/gin-gonic/gin"

func InitArticleTypeManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	articleTypeManageController, _ := InitializeArticleTypeManageController()

	articleTypeManageReqsGroup := baseGroup.Group("/articletype")
	{
		articleTypeManageReqsGroup.POST("/create", articleTypeManageController.CreateArticleType)
	}

	return articleTypeManageReqsGroup
}
