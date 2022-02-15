package service

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLogin(c *gin.Context) {
	var loginInfo LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		errorMsg := LoginResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	user := model.Admin{
		USERNAME: loginInfo.Username,
	}
	dbUser := model.GetAdminInfo(user)

	if loginInfo.Username == dbUser.USERNAME && loginInfo.Password == dbUser.PASSWORD {
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

func AddAdmin(c *gin.Context) {
	var registerInfo RegisterInfo
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		errorMsg := RegisterResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	newAdmin := model.Admin{
		USERNAME: registerInfo.Username,
		PASSWORD: registerInfo.Password,
	}

	model.AddAdmin(newAdmin)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
}

func UpdateAdmin(c *gin.Context) {
	var changePasswordInfo ChangePasswordInfo
	if err := c.ShouldBindJSON(&changePasswordInfo); err != nil {
		errorMsg := ChangePasswordResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}
	updateAdmin := model.Admin{
		USERNAME: changePasswordInfo.Username,
		PASSWORD: "",
	}
	updateAdmin = model.GetAdminInfo(updateAdmin)
	if updateAdmin.PASSWORD == changePasswordInfo.OldPassword {
		updateAdmin.PASSWORD = changePasswordInfo.NewPassword
		model.UpdateAdmin(updateAdmin)
		successMsg := ChangePasswordResult{
			Message: "Successfully",
		}
		c.JSON(http.StatusOK, successMsg)
		return
	} else {
		print(updateAdmin.PASSWORD)
		errorMsg := ChangePasswordResult{
			Message: "Username or Password is not correct",
		}
		c.JSON(http.StatusInternalServerError, errorMsg)
		return
	}
}

func DeleteAdmin(c *gin.Context) {
	var registerInfo RegisterInfo
	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		errorMsg := RegisterResult{
			Message: "Request Parameters Error",
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	newAdmin := model.Admin{
		USERNAME: registerInfo.Username,
		PASSWORD: registerInfo.Password,
	}

	model.DeleteAdmin(newAdmin)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
}
