package controller

import (
	"GFBackend/model/dao"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strings"
)

type CommunityManageController struct {
	communityManageService service.ICommunityManageService
}

func NewCommunityManageController(communityManageService service.ICommunityManageService) *CommunityManageController {
	return &CommunityManageController{communityManageService: communityManageService}
}

var CommunityManageSet = wire.NewSet(
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	service.NewCommunityManageService,
	wire.Bind(new(service.ICommunityManageService), new(*service.CommunityManageService)),
	NewCommunityManageController,
)

// CreateCommunity godoc
// @Summary Create a new Community
// @Description need strings creator & community name & description
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
