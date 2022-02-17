package dao

import (
	"GFBackend/model"
)

type IUserDAO interface {
	CreateUser(user User) error
}

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return new(UserDAO)
}

func (userDAO *UserDAO) CreateUser(user User) error {
	// strings in Select() must be as same as User field variables name
	result := model.DB.Select("Username", "Password", "Salt").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
