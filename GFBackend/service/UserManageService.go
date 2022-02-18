package service

import (
	"GFBackend/logger"
	"GFBackend/middleware/auth"
	"GFBackend/model"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type IUserManageService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	Logout(username, token string) error
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
	newUser := model.User{
		Username: username,
		Password: utils.EncodeInMD5(password + salt),
		Salt:     salt,
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		createUserError := userManageService.userDAO.CreateUser(newUser, tx)
		if createUserError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Create User Error: %s", createUserError.Error()))
			return createUserError
		}

		_, CasbinAddPolicyError := auth.CasbinEnforcer.AddPolicy(username, "regular")
		if CasbinAddPolicyError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Add New User Policy Error: %s", CasbinAddPolicyError.Error()))
			return CasbinAddPolicyError
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (userManageService *UserManageService) Login(username, password string) (string, error) {
	dbUser := userManageService.userDAO.GetUserByUsername(username)
	if dbUser.Username == "" {
		return "", errors.New("400")
	}

	inputPassword := utils.EncodeInMD5(password + dbUser.Salt)
	if inputPassword != dbUser.Password {
		return "", errors.New("400")
	}

	token, err := auth.TokenGenerate(username)
	if err != nil {
		return "", errors.New("500")
	}

	return token.Token, nil
}

func (userManageService *UserManageService) Logout(username, token string) error {
	sign := auth.TokenVerify(token)
	if !sign {
		return errors.New("")
	}
	return nil
}

func (userManageService *UserManageService) UpdatePassword() {

}

func (userManageService *UserManageService) Delete() {

}
