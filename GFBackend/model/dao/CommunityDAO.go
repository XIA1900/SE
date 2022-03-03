package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var communityDAO *CommunityDAO
var communityDAOLock sync.Mutex

type ICommunityDAO interface {
	CreateCommunity(community model.Community, tx *gorm.DB) error
	GetCommunityByName(community model.Community) (model.Community, error)
}

type CommunityDAO struct {
	db *gorm.DB
}

func NewCommunityDAO() *CommunityDAO {
	if communityDAO == nil {
		communityDAOLock.Lock()
		if communityDAO == nil {
			communityDAO = &CommunityDAO{
				db: model.NewDB(),
			}
		}
		communityDAOLock.Unlock()
	}
	return communityDAO
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
