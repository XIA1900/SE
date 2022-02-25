package dao

import (
	"GFBackend/model"
	"errors"
	"time"
)

type IFollowDAO interface {
	UserFollow(followee, follower string) error
}

type FollowDAO struct{}

func NewFollowDAO() *FollowDAO {
	return new(FollowDAO)
}

func (followDAO *FollowDAO) UserFollow(followee, follower string) error {
	follow := model.Follow{
		Followee:   followee,
		Follower:   follower,
		Create_Day: time.Now().Format("2006-01-02"),
	}

	result := model.DB.Create(&follow)
	if result.Error != nil {
		return errors.New("500")
	}

	return nil
}
