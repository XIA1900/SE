package model

import (
	"time"
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
	Create_Time time.Time
}

func (c Community) TableName() string {
	return "Community"
}
