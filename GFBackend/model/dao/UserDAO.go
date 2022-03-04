package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var userDAOLock sync.Mutex
var userDAO *UserDAO

type IUserDAO interface {
	CreateUser(user model.User) error
	GetUserByUsername(username string) model.User
	DeleteUserByUsername(username string) error
	UpdateUserPassword(username string, newPassword string) error
	UpdateUserByUsername(userInfo model.User) error
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO() *UserDAO {
	if userDAO == nil {
		userDAOLock.Lock()
		if userDAO == nil {
			userDAO = &UserDAO{
				db: model.NewDB(),
			}
		}
		userDAOLock.Unlock()
	}
	return userDAO
}

func (userDAO *UserDAO) CreateUser(user model.User) error {
	// strings in Select() must be as same as User field variables name
	var result *gorm.DB
	result = userDAO.db.Select("Username", "Password", "Salt").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userDAO *UserDAO) GetUserByUsername(username string) model.User {
	var user model.User
	userDAO.db.Where("username = ?", username).First(&user)
	return user
}

func (userDAO *UserDAO) DeleteUserByUsername(username string) error {
	var result *gorm.DB
	result = userDAO.db.Where("Username = ?", username).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userDAO *UserDAO) UpdateUserPassword(username string, newPassword string) error {
	result := userDAO.db.Model(&model.User{}).Where("Username = ?", username).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (userDAO *UserDAO) UpdateUserByUsername(userInfo model.User) error {
	result := userDAO.db.Model(&model.User{}).Where("Username = ?", userInfo.Username).Updates(model.User{
		Nickname:   userInfo.Nickname,
		Birthday:   userInfo.Birthday,
		Gender:     userInfo.Gender,
		Department: userInfo.Department,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
