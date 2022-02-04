package service

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResult struct {
	Message string
}

func Login(c *gin.Context) {
	var loginInfo LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		errorMsg := LoginResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	if loginInfo.Username == "jake16" && loginInfo.Password == "12345" {
		successMsg := LoginResult{
			Message: "Successfully",
		}
		c.JSON(http.StatusOK, successMsg)
		return
	}

	errorMsg := LoginResult{
		Message: "Username or Password is not correct",
	}
	c.JSON(http.StatusInternalServerError, errorMsg)
}

type RegisterInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterResult struct {
	Message string
}

func Register(c *gin.Context) {
	var registerInfo RegisterInfo
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		errorMsg := RegisterResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	newUser := model.User{
		USERNAME: registerInfo.Username,
		PASSWORD: registerInfo.Password,
	}

	model.AddUser(newUser)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
}

type ChangePasswordInfo struct {
	Username    string `form:"username" json:"username" binding:"required"`
	OldPassword string `form:"OldPassword" json:"OldPassword" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}
type ChangePasswordResult struct {
	Message string
}

func ChangePassword(c *gin.Context) {
	var changePasswordInfo ChangePasswordInfo
	if err := c.ShouldBindJSON(&changePasswordInfo); err != nil {
		errorMsg := ChangePasswordResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}
	updateUser := model.User{
		USERNAME: changePasswordInfo.Username,
	}
	model.GetUserInfo(updateUser)
	//model.ChangePassword(updateUser)
	//successMsg := ChangePasswordResult{
	//	Message: "Successfully",
	//}
	//c.JSON(http.StatusOK, successMsg)
	if updateUser.PASSWORD == changePasswordInfo.OldPassword {
		updateUser.PASSWORD = changePasswordInfo.NewPassword
		model.ChangePassword(updateUser)
		successMsg := ChangePasswordResult{
			Message: "Successfully",
		}
		c.JSON(http.StatusOK, successMsg)
		return
	} else {
		print(updateUser.PASSWORD)
		errorMsg := ChangePasswordResult{
			Message: "Username or Password is not correct",
		}
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
}

func DeleteUser(c *gin.Context) {
	var registerInfo RegisterInfo
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		errorMsg := RegisterResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	newUser := model.User{
		USERNAME: registerInfo.Username,
		PASSWORD: registerInfo.Password,
	}

	model.DeleteUser(newUser)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
}
