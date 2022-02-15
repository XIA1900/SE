package controller

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username    string `json:"Username" example:"yingjiechen21"`
	Password    string `json:"Password" example:"f9ae5f68ae1e7f7f3fc06053e9b9b539"`
	NewPassword string `json:"NewPassword" example:"3ecb441b741bcd433288f5557e4b9118"`
}

// UserRegister godoc
// @Summary Register a new User
// @Description get strings by username & password
// @Tags user_manage
// @Accept json
// @Produce json
// @Param UserInfo body controller.UserInfo true "User Register only needs Username & Password(encoded by md5)"
// @Success 201 {object} controller.HTTPError "<b>Success</b>. User Register Successfully"
// @Failure 406 {object} controller.HTTPError "<b>Failure</b>. User Has Existed"
// @Failure 500 {object} controller.HTTPError "<b>Failure</b>. Server Internal Error."
// @Router /user/register [post]
func UserRegister(context *gin.Context) {
}

func UserLogin(context *gin.Context) {
}

func UserLogout(context *gin.Context) {
}

func UserUpdatePassword(context *gin.Context) {
}
