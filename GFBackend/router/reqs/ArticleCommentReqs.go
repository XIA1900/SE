package reqs

import "github.com/gin-gonic/gin"

func InitArticleCommentReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	//articleCommentController, _ := InitializeArticleCommentController()

	articleCommentReqsGroup := baseGroup.Group("/articlecomment")
	{

	}

	return articleCommentReqsGroup
}
