package model

import (
	"GFBackend/utils"
)

type User struct {
	ID         int
	Username   string
	Password   string
	Salt       string
	Nickname   string
	Birthday   string
	Gender     string
	Department string
}

func (u User) TableName() string {
	return "User"
}

type Community struct {
	ID          int
	Creator     string
	Name        string
	Description string
	Num_Member  int
	Create_Time *utils.LocalTime
}

func (c Community) TableName() string {
	return "Community"
}
