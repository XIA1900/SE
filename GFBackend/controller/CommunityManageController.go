package controller

import (
	"GFBackend/middleware/auth"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strings"
	"sync"
)

var communityManageControllerLock sync.Mutex
var communityManageController *CommunityManageController

type CommunityManageController struct {
	communityManageService service.ICommunityManageService
}

func NewCommunityManageController(communityManageService service.ICommunityManageService) *CommunityManageController {
	if communityManageController == nil {
		communityManageControllerLock.Lock()
		if communityManageController == nil {
			communityManageController = &CommunityManageController{
				communityManageService: communityManageService,
			}
		}
		communityManageControllerLock.Unlock()
	}
	return communityManageController
}

var CommunityManageSet = wire.NewSet(
	service.CommunityManageServiceSet,
	wire.Bind(new(service.ICommunityManageService), new(*service.CommunityManageService)),
	NewCommunityManageController,
)

// CreateCommunity godoc
// @Summary Create a new Community
// @Description need token in cookie, need community name & description only
// @Tags Community Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param CommunityInfo body controller.CommunityInfo true "Create a new community needs Creator, Name & Description."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Create Community Success"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Community already exists"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /community/create [post]
func (communityManageController *CommunityManageController) CreateCommunity(context *gin.Context) {
	respMsg := ResponseMsg{
		Code:    400,
		Message: "Bad Parameters or Community already exists",
	}

	var communityInfo CommunityInfo
	if err1 := context.ShouldBindJSON(&communityInfo); err1 != nil {
		context.JSON(respMsg.Code, respMsg)
		return
	}

	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)

	err2 := communityManageController.communityManageService.CreateCommunity(username, communityInfo.Name, communityInfo.Description)
	if err2 != nil {
		if strings.Contains(err2.Error(), "400") {
			context.JSON(respMsg.Code, respMsg)
			return
		}
		respMsg.Code = 500
		respMsg.Message = "Internal Server Error"
		context.JSON(respMsg.Code, respMsg)
		return
	}

	respMsg.Code = 200
	respMsg.Message = "Create Community Success"
	context.JSON(respMsg.Code, respMsg)
	return
}

// GetCommunityByName godoc
// @Summary Get the Community by Name
// @Description need strings community name
// @Tags Community Manage
// @Accept json
// @Produce json
// @Param CommunityInfo body controller.CommunityInfo true "Create a new community needs Creator, Name & Description."
// @Success 201 {object} controller.CommunityResponseMsg "<b>Success</b>. Create Community Success"
// @Failure 400 {object} controller.CommunityResponseMsg "<b>Failure</b>. Bad Parameters or Community already exists"
// @Failure 500 {object} controller.CommunityResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /community/getcommunity [get]
func (communityManageController *CommunityManageController) GetCommunityByName(context *gin.Context) {

}

// UpdateCommunity godoc
// @Summary Update community information including Name, Description
// @Description need ID, Name, Description
// @Tags Community Manage
// @Accept json
// @Produce json
// @Param communityInfo body controller.CommunityInfo true "need ID, Name, Description"
// @Success 201 {object} controller.CommunityResponseMsg "<b>Success</b>. Update Password Successfully"
// @Failure 400 {object} controller.CommunityResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} controller.CommunityResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /community/updatecommunitybyid [post]
func (communityManageController *CommunityManageController) UpdateCommunity(context *gin.Context) {

}

// DeleteCommunity godoc
// @Summary Delete community information
// @Description need ID
// @Tags Community Manage
// @Accept json
// @Produce json
// @Param communityInfo body controller.CommunityInfo true "need ID"
// @Success 201 {object} controller.CommunityResponseMsg "<b>Success</b>. Update Password Successfully"
// @Failure 400 {object} controller.CommunityResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} controller.CommunityResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /community/deletecommunitybyid [post]
func (communityManageController *CommunityManageController) DeleteCommunity(context *gin.Context) {

}
