package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var communityDAO *CommunityDAO
var communityDAOLock sync.Mutex

type ICommunityDAO interface {
	CreateCommunity(community model.Community) error
	GetCommunityByName(community model.Community) (model.Community, error)
	UpdateCommunity(community model.Community) error
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

func (communityDAO *CommunityDAO) CreateCommunity(community model.Community) error {
	var result *gorm.DB
	result = communityDAO.db.Select("Creator", "Name", "Description", "Create_Time").Create(&community)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (communityDAO *CommunityDAO) GetCommunityByName(community model.Community) (model.Community, error) {
	result := communityDAO.db.Select("Creator", "Name", "Description", "Create_Time").Where("Name = ?", community.Name).First(&community)
	if result.Error != nil {
		return community, result.Error
	} else {
		dbCommunity := model.Community{}
		communityDAO.db.Where("Name = ?", community.Name).First(&dbCommunity)
		return dbCommunity, nil
	}
}

func (communityDAO *CommunityDAO) UpdateCommunity(communityInfo model.Community) error {
	result := communityDAO.db.Model(&model.Community{}).Where("Creator = ?", communityInfo.Creator).Updates(model.Community{
		Name:        communityInfo.Name,
		Description: communityInfo.Description,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
