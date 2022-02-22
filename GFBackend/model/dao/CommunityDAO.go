package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
)

type ICommunityDAO interface {
	CreateCommunity(community model.Community, tx *gorm.DB) error
}

type CommunityDAO struct{}

func NewCommunityDAO() *CommunityDAO {
	return new(CommunityDAO)
}

func (communityDAO *CommunityDAO) CreateCommunity(community model.Community, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Select("Creator", "Name", "Description").Create(&community)
	} else {
		result = tx.Select("Creator", "Name", "Description").Create(&community)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
