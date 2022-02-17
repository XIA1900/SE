package service

import (
	"GFBackend/logger"
	"GFBackend/middleware/auth"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"fmt"
)

type IUserManageService interface {
	Register(username, password string) error
	Login(username, password string) error
	Logout()
	UpdatePassword()
	Delete()
}

type UserManageService struct {
	userDAO dao.IUserDAO
}

func NewUserManageService(userDAO dao.IUserDAO) *UserManageService {
	return &UserManageService{userDAO: userDAO}
}

func (userManageService *UserManageService) Register(username, password string) error {
	salt := utils.GetRandomString(6)
	newUser := dao.User{
		Username: username,
		Password: utils.EncodeInMD5(password + salt),
		Salt:     salt,
	}

	err := userManageService.userDAO.CreateUser(newUser)
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

func (userManageService *UserManageService) Login(username, password string) error {
	return nil
}

func (userManageService *UserManageService) Logout() {

}

func (userManageService *UserManageService) UpdatePassword() {

}

func (userManageService *UserManageService) Delete() {

}
