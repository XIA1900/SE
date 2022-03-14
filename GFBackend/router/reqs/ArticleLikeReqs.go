package reqs

import "github.com/gin-gonic/gin"

func InitArticleLikeReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	articleLikeController, _ := InitializeArticleLikeController()

	articleLikeReqsGroup := baseGroup.Group("/articlelike")
	{
		articleLikeReqsGroup.POST("/create", articleLikeController.CreateLike)
	}

	return articleLikeReqsGroup

}
