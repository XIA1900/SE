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

	if loginInfo.Username == "jake_admin" && loginInfo.Password == "12345" {
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

	model.UpdateAdmin(newAdmin)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
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
