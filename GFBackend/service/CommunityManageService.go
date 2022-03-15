package service

import (
	"GFBackend/logger"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"github.com/google/wire"
	"strings"
	"sync"
)

var communityManageServiceLock sync.Mutex
var communityManageService *CommunityManageService

type ICommunityManageService interface {
	//CreateCommunity(creator string, name string, description string, createTime string) error
	//GetCommunityByName(name string) (model.Community, model.User, error)
	//UpdateCommunity(updateInfo model.Community) error
	//DeleteCommunity(id int) error

	CreateCommunity(creator string, name string, description string) error
	DeleteCommunityByID(id int) error
}

type CommunityManageService struct {
	communityDAO       dao.ICommunityDAO
	communityMemberDAO dao.ICommunityMemberDAO
}

func NewCommunityManageService(communityDAO dao.ICommunityDAO, communityMemberDAO dao.ICommunityMemberDAO) *CommunityManageService {
	if communityManageService == nil {
		communityManageServiceLock.Lock()
		if communityManageService == nil {
			communityManageService = &CommunityManageService{
				communityDAO:       communityDAO,
				communityMemberDAO: communityMemberDAO,
			}
		}
		communityManageServiceLock.Unlock()
	}
	return communityManageService
}

var CommunityManageServiceSet = wire.NewSet(
	dao.NewCommunityMemberDAO,
	wire.Bind(new(dao.ICommunityMemberDAO), new(*dao.CommunityMemberDAO)),
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	NewCommunityManageService,
)

func (communityManageService *CommunityManageService) CreateCommunity(creator, name, description string) error {
	newCommunityID, err1 := communityManageService.communityDAO.CreateCommunity(name, creator, description, utils.GetCurrentDate())
	if err1 != nil {
		if strings.Contains(err1.Error(), "Duplicate") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}

	err2 := communityManageService.communityMemberDAO.Create(newCommunityID, creator, utils.GetCurrentDate())
	if err2 != nil {
		logger.AppLogger.Error(err2.Error())
		return errors.New("500")
	}

	return nil
}

func (communityManageService *CommunityManageService) DeleteCommunityByID(id int) error {
	return nil
}
