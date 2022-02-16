package dao

import (
	"GFBackend/model"
	"time"
)

type User struct {
	ID         int
	Username   string
	Password   string
	Salt       string
	Nickname   string
	Birthday   time.Time
	Gender     string
	Department string
}

func (u User) TableName() string {
	return "USER"
}

func (u User) CreateUser(user User) error {
	// strings in Select() must be as same as User field variables name
	result := model.DB.Select("Username", "Password", "Salt").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
