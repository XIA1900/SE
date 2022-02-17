package controller

import (
	"GFBackend/model/dao"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strings"
)

//type IUserManageController interface {
//	RegularRegister(context *gin.Context)
//	AdminRegister(context *gin.Context)
//	UserLogin(context *gin.Context)
//	UserLogout(context *gin.Context)
//	UserUpdatePassword(context *gin.Context)
//	UserDelete(context *gin.Context)
//	UserUpdate(context *gin.Context)
//}

type UserManageController struct {
	userManageService service.IUserManageService
}

var UserManageSet = wire.NewSet(
	dao.NewUserDAO,
	wire.Bind(new(dao.IUserDAO), new(*dao.UserDAO)),
	service.NewUserManageService,
	wire.Bind(new(service.IUserManageService), new(*service.UserManageService)),
	NewUserManageController,
)

func NewUserManageController(userManageService service.IUserManageService) *UserManageController {
	return &UserManageController{userManageService: userManageService}
}

// RegularRegister godoc
// @Summary Register a new Regular User
// @Description only need strings username & password
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "Regular User Register only needs Username, Password(encoded by md5) & ForAdmin with false."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. User Register Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User Has Existed"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/register [post]
func (userManageController *UserManageController) RegularRegister(context *gin.Context) {
	var registerInfo UserInfo
	if err := context.ShouldBindJSON(&registerInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err := userManageController.userManageService.Register(registerInfo.Username, registerInfo.Password)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "User Has Existed.",
			}
			context.JSON(http.StatusBadRequest, er)
		} else {
			er := ResponseMsg{
				Code:    http.StatusInternalServerError,
				Message: "Server Internal Error.",
			}
			context.JSON(http.StatusInternalServerError, er)
		}
		return
	}

	context.JSON(http.StatusCreated, ResponseMsg{
		Code:    http.StatusCreated,
		Message: "Create User Successfully",
	})
}

func (userManageController *UserManageController) AdminRegister(context *gin.Context) {

}

// UserLogin godoc
// @Summary Admin / Regular User login
// @Description only need strings username & password
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "only needs username and password"
// @Success 200 {object} controller.ResponseMsg "<b>Success</b>. User Login Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Username / Password incorrect"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/login [post]
func (userManageController *UserManageController) UserLogin(context *gin.Context) {
	var userInfo UserInfo
	if err := context.ShouldBindJSON(&userInfo); err != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

}

func (userManageController *UserManageController) UserLogout(context *gin.Context) {
}

func (userManageController *UserManageController) UserUpdatePassword(context *gin.Context) {
}

func (userManageController *UserManageController) UserDelete(context *gin.Context) {

}

func (userManageController *UserManageController) UserUpdate(context *gin.Context) {

}
