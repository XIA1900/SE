package dao

import (
	"GFBackend/controller"
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var communityDAO *CommunityDAO
var communityDAOLock sync.Mutex

type ICommunityDAO interface {
	CreateCommunity(community model.Community) error
	GetCommunityByName(community model.Community) (model.Community, error)
	UpdateCommunity(community controller.CommunityInfo) error
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

<<<<<<< HEAD
func (communityDAO *CommunityDAO) UpdateCommunity(communityInfo controller.CommunityInfo) error {
	OldCommunity := model.DB.Where("Creator = ? AND Name = ?", communityInfo.Creator, communityInfo.Name).First(&model.Community{})
	if OldCommunity.Error != nil {
		return OldCommunity.Error
	} else {
		result := model.DB.Model(&model.Community{}).Updates(map[string]interface{}{
			"Name":        communityInfo.NewName,
			"Description": communityInfo.Description,
		})
		if result.Error != nil {
			return result.Error
		}
=======
func (communityDAO *CommunityDAO) UpdateCommunity(communityInfo model.Community) error {
	result := communityDAO.db.Model(&model.Community{}).Where("Creator = ?", communityInfo.Creator).Updates(model.Community{
		Name:        communityInfo.Name,
		Description: communityInfo.Description,
	})
	if result.Error != nil {
		return result.Error
>>>>>>> 731b2a190b5a69a94b1b2754e65e4a6d5fc7d70e
	}
	return nil
}
