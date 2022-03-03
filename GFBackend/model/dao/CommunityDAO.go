package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
)

type ICommunityDAO interface {
	CreateCommunity(community model.Community, tx *gorm.DB) error
	GetCommunityByName(community model.Community) (model.Community, error)
}

type CommunityDAO struct{}

func NewCommunityDAO() *CommunityDAO {
	return new(CommunityDAO)
}

func (communityDAO *CommunityDAO) CreateCommunity(community model.Community, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Select("Creator", "Name", "Description", "Create_Time").Create(&community)
	} else {
		result = tx.Select("Creator", "Name", "Description", "Create_Time").Create(&community)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (communityDAO *CommunityDAO) GetCommunityByName(community model.Community) (model.Community, error) {
	result := model.DB.Select("Creator", "Name", "Description", "Create_Time").Where("Name = ?", community.Name).First(&community)
	if result.Error != nil {
		return community, result.Error
	} else {
		dbCommunity := model.Community{}
		model.DB.Where("Name = ?", community.Name).First(&dbCommunity)
		return dbCommunity, nil
	}
}
