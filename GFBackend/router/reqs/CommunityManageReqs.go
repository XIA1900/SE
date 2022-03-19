package reqs

import "github.com/gin-gonic/gin"

func InitCommunityManageReqs(baseGroup *gin.RouterGroup) *gin.RouterGroup {

	communityManageController, _ := InitializeCommunityManageController()

	communityManageReqsGroup := baseGroup.Group("/community")
	{
		communityManageReqsGroup.POST("/create", communityManageController.CreateCommunity)
		communityManageReqsGroup.GET("/delete/:id", communityManageController.DeleteCommunityByID)
		communityManageReqsGroup.POST("/update", communityManageController.UpdateDescriptionByID)
		communityManageReqsGroup.GET("/numberofmember/:id", communityManageController.GetNumberOfMemberByID)
		communityManageReqsGroup.GET("/getone/:id", communityManageController.GetOneCommunityByID)
		communityManageReqsGroup.GET("/getbyname", communityManageController.GetCommunitiesByNameFuzzyMatch)
		communityManageReqsGroup.GET("/get", communityManageController.GetCommunities)
		communityManageReqsGroup.GET("/join/:id", communityManageController.JoinCommunityByID)
		communityManageReqsGroup.GET("/leave/:id", communityManageController.LeaveCommunityByID)
		communityManageReqsGroup.GET("/getmember", communityManageController.GetMembersByCommunityIDs)

	}
	return communityManageReqsGroup
}
