package dao

import (
	"GFBackend/model"
	"errors"
	"sync"
	"time"
)

var followDAOLock sync.Mutex
var followDAO *FollowDAO

type IFollowDAO interface {
	UserFollow(followee, follower string) error
}

type FollowDAO struct{}

func NewFollowDAO() *FollowDAO {
	if followDAO == nil {
		followDAOLock.Lock()
		if followDAO == nil {
			followDAO = new(FollowDAO)
		}
		followDAOLock.Unlock()
	}
	return followDAO
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
