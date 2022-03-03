package service

import (
	"GFBackend/logger"
	"GFBackend/model"
	"GFBackend/model/dao"
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
	"sync"
	"time"
)

var communityManageServiceLock sync.Mutex
var communityManageService *CommunityManageService

type ICommunityManageService interface {
	CreateCommunity(creator string, name string, description string, createTime time.Time) error
	GetCommunityByName(name string) (model.Community, model.User, error)
}

type CommunityManageService struct {
	communityDAO dao.ICommunityDAO
	userDAO      dao.IUserDAO
}

func NewCommunityManageService(communityDAO dao.ICommunityDAO, userDAO dao.IUserDAO) *CommunityManageService {
	if communityManageService == nil {
		communityManageServiceLock.Lock()
		if communityManageService == nil {
			communityManageService = &CommunityManageService{
				communityDAO: communityDAO,
				userDAO:      userDAO,
			}
		}
		communityManageServiceLock.Unlock()
	}
	return communityManageService
}

var CommunityServiceSet = wire.NewSet(
	dao.NewUserDAO,
	wire.Bind(new(dao.IUserDAO), new(*dao.UserDAO)),
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	NewCommunityManageService,
)

func (communityManageService *CommunityManageService) CreateCommunity(creator string, name string, description string, createTime time.Time) error {
	newCommunity := model.Community{
		Creator:     creator,
		Name:        name,
		Description: description,
		Create_Time: createTime,
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		createCommunityError := communityManageService.communityDAO.CreateCommunity(newCommunity, tx)
		if createCommunityError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Create Community Error: %s", createCommunityError.Error()))
			return createCommunityError
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (communityManageService *CommunityManageService) GetCommunityByName(name string) (model.Community, model.User, error) {
	newCommunity := model.Community{
		Name: name,
	}
	resCommunity, err := communityManageService.communityDAO.GetCommunityByName(newCommunity)
	resUser := userManageService.userDAO.GetUserByUsername(resCommunity.Creator)
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Get Community By Name Error: %s", err.Error()))
		return model.Community{}, model.User{}, err
	}
	return resCommunity, resUser, nil
}
