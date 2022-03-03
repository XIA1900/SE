package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
	"time"
)

var followDAOLock sync.Mutex
var followDAO *FollowDAO

type IFollowDAO interface {
	GetOneFollow(followee, follower string) (model.Follow, error)
	UserFollow(followee, follower string) error
	UserUnfollow(followee, follower string) error
	DeleteFollow(username string, tx *gorm.DB) error
}

type FollowDAO struct {
	db *gorm.DB
}

func NewFollowDAO() *FollowDAO {
	if followDAO == nil {
		followDAOLock.Lock()
		if followDAO == nil {
			followDAO = &FollowDAO{
				db: model.NewDB(),
			}
		}
		followDAOLock.Unlock()
	}
	return followDAO
}

func (followDAO *FollowDAO) GetOneFollow(followee, follower string) (model.Follow, error) {
	follow := model.Follow{}
	result := followDAO.db.Where("Followee = ? AND Follower = ?", followee, follower).First(&follow)
	if result.Error != nil {
		return follow, result.Error
	}
	return follow, nil
}

func (followDAO *FollowDAO) UserFollow(followee, follower string) error {
	follow := model.Follow{
		Followee:   followee,
		Follower:   follower,
		Create_Day: time.Now().Format("2006-01-02"),
	}

	result := followDAO.db.Create(&follow)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (followDAO *FollowDAO) UserUnfollow(followee, follower string) error {
	result := followDAO.db.Where("Followee = ? and Follower = ?", followee, follower).Delete(&model.Follow{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (followDAO *FollowDAO) DeleteFollow(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx != nil {
		result = tx.Where("Follower = ?", username).Delete(&model.Follow{})
	} else {
		result = followDAO.db.Where("Follower = ?", username).Delete(&model.Follow{})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
