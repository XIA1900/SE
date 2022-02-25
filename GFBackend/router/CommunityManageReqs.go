package router

import "github.com/gin-gonic/gin"

func InitCommunityManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	communityManageController, _ := InitializeCommunityManageController()

	communityManageReqsGroup := baseGroup.Group("/community")
	{
		communityManageReqsGroup.POST("/create", communityManageController.CreateCommunity)
		communityManageReqsGroup.GET("/getcommunity", communityManageController.GetCommunityByName)
	}
	return communityManageReqsGroup
}
