package service

import (
	"GFBackend/cache"
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
	Register(username, password string, forAdmin bool) error
	Login(username, password string) (string, error)
	Logout(username, token string) error
	UpdatePassword(username, password, newPassword string) error
	Delete(username string) error
}

type UserManageService struct {
	userDAO dao.IUserDAO
}

func NewUserManageService(userDAO dao.IUserDAO) *UserManageService {
	return &UserManageService{userDAO: userDAO}
}

func (userManageService *UserManageService) Register(username, password string, forAdmin bool) error {
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

		role := "regular"
		if forAdmin {
			role = "admin"
		}
		_, CasbinAddPolicyError := auth.CasbinEnforcer.AddGroupingPolicy(username, role)
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
		logger.AppLogger.Error(err.Error())
		return "", errors.New("500")
	}

	sign, _ := auth.GetTokenSign(token.Token)
	err = cache.AddLoginUserWithSign(username, sign)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return "", errors.New("500")
	}

	return token.Token, nil
}

func (userManageService *UserManageService) Logout(username, token string) error {
	sign := auth.TokenVerify(token)
	err := cache.DelLoginUserSign(username)
	if !sign || err != nil {
		if err != nil {
			logger.AppLogger.Error(err.Error())
		}
		return errors.New("")
	}
	return nil
}

func (userManageService *UserManageService) UpdatePassword(username, password, newPassword string) error {
	user := userManageService.userDAO.GetUserByUsername(username)

	if utils.EncodeInMD5(password+user.Salt) != user.Password {
		return errors.New("400")
	}

	err := userManageService.userDAO.UpdateUserPassword(username, newPassword)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}

	return nil
}

func (userManageService *UserManageService) Delete(username string) error {
	user := userManageService.userDAO.GetUserByUsername(username)
	if user.Username == "" {
		return errors.New("user does not exist")
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		deleteUserError := userManageService.userDAO.DeleteUserByUsername(username, tx)
		if deleteUserError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Delete User Error: %s", deleteUserError.Error()))
			return deleteUserError
		}

		_, CasbinAddPolicyError := auth.CasbinEnforcer.DeleteUser(username)
		if CasbinAddPolicyError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Delete User Policy Error: %s", CasbinAddPolicyError.Error()))
			return CasbinAddPolicyError
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}
