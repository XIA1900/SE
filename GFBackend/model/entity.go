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

type Follow struct {
	Followee   string
	Follower   string
	Create_Day string
}

func (follow Follow) TableName() string {
	return "Follow"
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

type Space struct {
	ID        int
	Username  string
	Capacity  float32
	Remaining float32
}

func (space Space) TableName() string {
	return "Space"
}
