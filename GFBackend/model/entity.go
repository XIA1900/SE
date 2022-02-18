package model

import "time"

type User struct {
	ID         int
	Username   string
	Password   string
	Salt       string
	Nickname   string
	Birthday   *time.Time
	Gender     string
	Department string
}

func (u User) TableName() string {
	return "User"
}
