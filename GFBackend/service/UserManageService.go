package service

import (
	"GFBackend/logger"
	"GFBackend/middleware/auth"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"fmt"
)

func Register(username, password string) error {
	salt := utils.GetRandomString(6)
	newUser := dao.User{
		Username: username,
		Password: utils.EncodeInMD5(password + salt),
		Salt:     salt,
	}

	userDao := dao.User{}
	err := userDao.CreateUser(newUser)
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Create User Error: %s", err.Error()))
		return err
	}
	_, err = auth.CasbinEnforcer.AddPolicy(username, "regular")
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Add New User Policy Error: %s", err.Error()))
		return err
	}

	return nil
}

func Login() {

}

func Logout() {

}

func UpdatePassword() {

}

func Delete() {

}
