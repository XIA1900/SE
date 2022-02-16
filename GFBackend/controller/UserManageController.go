package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserInfo struct {
	Username    string `json:"Username" example:"yingjiechen21"`
	Password    string `json:"Password" example:"f9ae5f68ae1e7f7f3fc06053e9b9b539"`
	NewPassword string `json:"NewPassword" example:"3ecb441b741bcd433288f5557e4b9118"`
	ForAdmin    bool   `json:"ForAdmin" example:true`
}

type NewUserInfo struct {
	Username   string    `json:"Username" example:"yingjiechen21"`
	Nickname   string    `json:"Nickname" example:"Peter Park"`
	Birthday   time.Time `json:"Birthday" example:"2022-02-30"`
	Gender     string    `json:"Gender" example:"male/female/unknown"`
	Department string    `json:"Department" example:"CS:GO"`
}

// RegularRegister godoc
// @Summary Register a new Regular User
// @Description only need strings username & password
// @Tags User Manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "Regular User Register only needs Username, Password(encoded by md5) & ForAdmin with false."
// @Success 201 {object} controller.HTTPError "<b>Success</b>. User Register Successfully"
// @Failure 400 {object} controller.HTTPError "<b>Failure</b>. Bad Parameters or User Has Existed"
// @Failure 500 {object} controller.HTTPError "<b>Failure</b>. Server Internal Error."
// @Router /user/register [post]
func RegularRegister(context *gin.Context) {
	var registerInfo UserInfo
	if err := context.ShouldBindJSON(&registerInfo); err != nil {
		er := HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Bad Parameters or User Has Existed or User has no permission.",
		}
		context.JSON(http.StatusBadRequest, er)
		return
	}

}

func AdminRegister(context *gin.Context) {

}

func UserLogin(context *gin.Context) {
}

func UserLogout(context *gin.Context) {
}

func UserUpdatePassword(context *gin.Context) {
}

func UserDelete(context *gin.Context) {

}

func UserUpdate(context *gin.Context) {

}
