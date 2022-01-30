package model

import "time"

type User struct {
	ID         uint
	USERNAME   string
	PASSWORD   string
	SALT       string
	NICKNAME   string
	BIRTHDAY   time.Time
	GENDER     int8
	DEPARTMENT string
}

func (u User) TableName() string {
	//return "TEST"
	return "USER"
}

func AddUser(user User) {
	db := getDB()
	db.Create(&user)
}

func ChangePassword(user User) {
	db := getDB()
	db.Model(&user).Where("USERNAME", user.USERNAME).Update("PASSWORD", user.PASSWORD)
}

func DeleteUser(user User) {
	db := getDB()
	db.Where("USERNAME", user.USERNAME).Delete(&user)
}
