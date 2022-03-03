package service

import (
	"GFBackend/logger"
	"GFBackend/model"
	"GFBackend/model/dao"
	"errors"
	"fmt"
	"github.com/google/wire"
	"sync"
)

var communityManageServiceLock sync.Mutex
var communityManageService *CommunityManageService

type ICommunityManageService interface {
	CreateCommunity(creator string, name string, description string, createTime string) error
	GetCommunityByName(name string) (model.Community, model.User, error)
	UpdateCommunity(updateInfo model.Community) error
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

func (communityManageService *CommunityManageService) CreateCommunity(creator string, name string, description string, createTime string) error {
	newCommunity := model.Community{
		Creator:     creator,
		Name:        name,
		Description: description,
		Create_Time: createTime,
	}

	createCommunityError := communityManageService.communityDAO.CreateCommunity(newCommunity)
	if createCommunityError != nil {
		logger.AppLogger.Error(fmt.Sprintf("Create Community Error: %s", createCommunityError.Error()))
		return createCommunityError
	}
	return nil
}

func (communityManageService *CommunityManageService) GetCommunityByName(name string) (model.Community, model.User, error) {
	newCommunity := model.Community{
		Name: name,
	}
	resCommunity, err := communityManageService.communityDAO.GetCommunityByName(newCommunity)
	resCommunity.Create_Time = resCommunity.Create_Time[:10]
	resUser := userManageService.userDAO.GetUserByUsername(resCommunity.Creator)
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Get Community By Name Error: %s", err.Error()))
		return model.Community{}, model.User{}, err
	}
	return resCommunity, resUser, nil
}

func (communityManageService *CommunityManageService) UpdateCommunity(updateInfo model.Community) error {
	err := communityManageService.communityDAO.UpdateCommunity(updateInfo)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}
	return nil
}
