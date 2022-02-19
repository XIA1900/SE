package controller

import (
	"GFBackend/config"
	"GFBackend/middleware/auth"
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

func NewUserManageController(userManageService service.IUserManageService) *UserManageController {
	return &UserManageController{userManageService: userManageService}
}

var UserManageSet = wire.NewSet(
	dao.NewUserDAO,
	wire.Bind(new(dao.IUserDAO), new(*dao.UserDAO)),
	service.NewUserManageService,
	wire.Bind(new(service.IUserManageService), new(*service.UserManageService)),
	NewUserManageController,
)

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

	err := userManageController.userManageService.Register(registerInfo.Username, registerInfo.Password, registerInfo.ForAdmin)
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

// AdminRegister godoc
// @Summary Register a new Admin User
// @Description only need strings username & password & ForAdmin, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "Admin User Register only needs Username, Password(encoded by md5) & ForAdmin with true."
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. User Register Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User Has Existed"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /user/admin/register [post]
func (userManageController *UserManageController) AdminRegister(context *gin.Context) {
	userManageController.RegularRegister(context)
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

	if token, err := userManageController.userManageService.Login(userInfo.Username, userInfo.Password); err != nil {
		if strings.Contains(err.Error(), "400") {
			er := ResponseMsg{
				Code:    http.StatusBadRequest,
				Message: "Username or Password is not correct",
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
	} else {
		success := ResponseMsg{
			Code:    http.StatusOK,
			Message: token,
		}
		context.SetCookie("token", token, config.AppConfig.JWT.Expires*60, config.AppConfig.Server.BasePath, "localhost", false, true)
		context.JSON(http.StatusOK, success)
		return
	}
}

// UserLogout godoc
// @Summary Admin / Regular User logout
// @Description need strings username in post request, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Router /user/logout [post]
func (userManageController *UserManageController) UserLogout(context *gin.Context) {
	context.SetCookie("token", "", -1, config.AppConfig.Server.BasePath, "localhost", false, true)

	type Info struct {
		Username string `json:"username"`
	}
	var info Info
	err := context.ShouldBind(&info)
	if err != nil {
		return
	}

	token, err := context.Cookie("token")
	if err != nil {
		return
	}

	err = userManageController.userManageService.Logout(info.Username, token)
	if err != nil {
		return
	}
}

func (userManageController *UserManageController) UserUpdatePassword(context *gin.Context) {
}

// UserDelete godoc
// @Summary Admin delete Users, cannot self delete
// @Description need strings username in post request, need token in cookie
// @Tags User Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Router /user/admin/delete [post]
func (userManageController *UserManageController) UserDelete(context *gin.Context) {
	type Info struct {
		Username string `json:"username"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	token, _ := context.Cookie("token")
	currentUsername, _ := auth.GetTokenUsername(token)
	if err1 != nil || info.Username == currentUsername {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters or Current User cannot delete self.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	err2 := userManageController.userManageService.Delete(info.Username)
	if err2 != nil {
		er := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "User not exist.",
		}
		if strings.Contains(err2.Error(), "user Policy") {
			er.Code = http.StatusInternalServerError
			er.Message = "Internal Server Error"
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

	context.JSON(http.StatusCreated, ResponseMsg{
		Code:    http.StatusCreated,
		Message: "Delete User Successfully",
	})
}

func (userManageController *UserManageController) UserUpdate(context *gin.Context) {

}
