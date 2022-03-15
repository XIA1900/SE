package reqs

import "github.com/gin-gonic/gin"

func InitArticleManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	articleManageController, _ := InitializeArticleManageController()

	articleTypeManageReqsGroup := baseGroup.Group("/article")
	{
		articleTypeManageReqsGroup.POST("/create", articleManageController.CreateArticle)
		articleTypeManageReqsGroup.GET("/search/:SearchWord", articleManageController.SearchArticle)
	}

	return articleTypeManageReqsGroup
}
