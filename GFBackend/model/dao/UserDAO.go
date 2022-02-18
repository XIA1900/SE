package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
)

type IUserDAO interface {
	CreateUser(user User, tx *gorm.DB) error
	GetUserByUsername(username string) User
}

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return new(UserDAO)
}

func (userDAO *UserDAO) CreateUser(user User, tx *gorm.DB) error {
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

func (userDAO *UserDAO) GetUserByUsername(username string) User {
	var user User
	model.DB.Where("username = ?", username).First(&user)
	return user
}
