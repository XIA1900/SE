package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
)

type IUserDAO interface {
	CreateUser(user model.User, tx *gorm.DB) error
	GetUserByUsername(username string) model.User
	DeleteUserByUsername(username string, tx *gorm.DB) error
}

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return new(UserDAO)
}

func (userDAO *UserDAO) CreateUser(user model.User, tx *gorm.DB) error {
	// strings in Select() must be as same as User field variables name
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Select("Username", "Password", "Salt").Create(&user)
	} else {
		result = tx.Select("Username", "Password", "Salt").Create(&user)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userDAO *UserDAO) GetUserByUsername(username string) model.User {
	var user model.User
	model.DB.Where("username = ?", username).First(&user)
	return user
}

func (userDAO *UserDAO) DeleteUserByUsername(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Where("Username = ?", username).Delete(&model.User{})
	} else {
		result = tx.Where("Username = ?", username).Delete(&model.User{})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}