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

func (user User) TableName() string {
	return "USER"
}

func CreateUser(user User) {
	model.DB.Create(&user)
}
