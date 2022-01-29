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

func UpdateUser(c *gin.Context) {
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

	model.UpdateUser(newUser)

	successMsg := RegisterResult{
		Message: "Successfully",
	}
	c.JSON(http.StatusOK, successMsg)
	return
}
