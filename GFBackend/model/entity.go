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

type Community struct {
	ID          int
	Creator     string
	Name        string
	Description string
	Num_member  int
	//Create_time *time.Time
}

func (c Community) TableName() string {
	return "Community"
}
