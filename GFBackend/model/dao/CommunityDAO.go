package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var communityDAO *CommunityDAO
var communityDAOLock sync.Mutex

type ICommunityDAO interface {
	CreateCommunity(communityName, username, description, createDay string) (int, error)
	DeleteCommunityByID(id int) error
	UpdateDescriptionByID(id int, newDescription string) error
	AddNumMemberByID(id int) error
	GetCommunityByID(id int) (model.Community, error)
	GetCommunities() ([]model.Community, error)
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

func (communityDAO *CommunityDAO) CreateCommunity(communityName, username, description, createDay string) (int, error) {
	newCommunity := model.Community{
		Creator:     username,
		Name:        communityName,
		Description: description,
		CreateDay:   createDay,
	}
	result := communityDAO.db.Create(&newCommunity)
	if result.Error != nil {
		return -1, result.Error
	}
	return newCommunity.ID, nil
}

func (communityDAO *CommunityDAO) DeleteCommunityByID(id int) error {
	return nil
}

func (communityDAO *CommunityDAO) UpdateDescriptionByID(id int, newDescription string) error {
	return nil
}

func (communityDAO *CommunityDAO) AddNumMemberByID(id int) error {
	return nil
}

func (communityDAO *CommunityDAO) GetCommunityByID(id int) (model.Community, error) {
	return model.Community{}, nil
}

func (communityDAO *CommunityDAO) GetCommunities() ([]model.Community, error) {
	return nil, nil
}
