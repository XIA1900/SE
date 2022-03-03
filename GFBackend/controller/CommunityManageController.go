package controller

import (
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
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
	service.CommunityServiceSet,
	wire.Bind(new(service.ICommunityManageService), new(*service.CommunityManageService)),
	NewCommunityManageController,
)

// CreateCommunity godoc
// @Summary Create a new Community
// @Description need strings creator & community name & description & create time
// @Tags Community Manage
// @Accept json
// @Produce json
// @Param CommunityInfo body controller.CommunityInfo true "Create a new community needs Creator, Name & Description."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Create Community Success"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Community already exists"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /community/create [post]
func (communityManageController *CommunityManageController) CreateCommunity(context *gin.Context) {
	var communityInfo CommunityInfo
	if err := context.ShouldBindJSON(&communityInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err := communityManageController.communityManageService.CreateCommunity(communityInfo.Creator, communityInfo.Name, communityInfo.Description, communityInfo.Create_Time)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "Community already exists",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := ResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	}
	context.JSON(http.StatusCreated, ResponseMsg{
		Code:    http.StatusCreated,
		Message: "Create Community Success",
	})
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
// @Router /community/create [get]
func (communityManageController *CommunityManageController) GetCommunityByName(context *gin.Context) {
	var communityInfo CommunityInfo
	if err := context.ShouldBindJSON(&communityInfo); err != nil {
		er := CommunityResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	resCommunity, resUser, err := communityManageController.communityManageService.GetCommunityByName(communityInfo.Name)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			er := CommunityResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "Community not found",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := CommunityResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	}
	context.JSON(http.StatusOK, CommunityResponseMsg{
		Code:        http.StatusOK,
		Message:     "Get Community Success",
		Creator:     resUser.Nickname,
		Name:        resCommunity.Name,
		Description: resCommunity.Description,
		Num_Member:  resCommunity.Num_Member,
		Create_Time: resCommunity.Create_Time,
	})
}

func (communityManageController *CommunityManageController) UpdateCommunity(context *gin.Context) {
	var communityInfo CommunityInfo
	if err := context.ShouldBindJSON(&communityInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err := communityManageController.communityManageService.UpdateCommunity(communityInfo)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "Community not found",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := ResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	}
	context.JSON(http.StatusOK, ResponseMsg{
		Code:    http.StatusOK,
		Message: "Update Community Success",
	})
}
