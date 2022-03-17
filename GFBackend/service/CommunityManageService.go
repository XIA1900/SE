package service

import (
	"GFBackend/entity"
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
	CreateCommunity(creator string, name string, description string) error
	DeleteCommunityByID(id int, operator string) error
	UpdateDescriptionByID(id int, newDescription, operator string) error
	GetNumberOfMembersByID(id int) (int64, error)
	GetOneCommunityByID(id int) (entity.Community, error)
	GetCommunitiesByNameFuzzyMatch(name string, pageNO, pageSize int) ([]entity.Community, int64, error)
	GetCommunities(pageNO, pageSize int) ([]entity.Community, int64, error)
	JoinCommunityByID(id int, username string) error
	LeaveCommunityByID(id int, username string) error
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

func (communityManageService *CommunityManageService) DeleteCommunityByID(id int, operator string) error {
	community, err1 := communityManageService.communityDAO.GetOneCommunityByID(id)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return err1
	}
	if community.Creator == operator {
		err2 := communityManageService.communityDAO.DeleteCommunityByID(id)
		if err2 != nil {
			logger.AppLogger.Error(err2.Error())
		}

		err3 := communityManageService.communityMemberDAO.DeleteByCommunityID(id)
		if err3 != nil {
			logger.AppLogger.Error(err3.Error())
		}
	}

	return nil
}

func (communityManageService *CommunityManageService) UpdateDescriptionByID(id int, newDescription, operator string) error {
	community, err1 := communityManageService.communityDAO.GetOneCommunityByID(id)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}
	if community.Creator == operator {
		err2 := communityManageService.communityDAO.UpdateDescriptionByID(id, newDescription)
		if err2 != nil {
			logger.AppLogger.Error(err2.Error())
			return errors.New("500")
		}
	} else {
		return errors.New("400")
	}
	return nil
}

func (communityManageService *CommunityManageService) GetNumberOfMembersByID(id int) (int64, error) {
	count, err1 := communityManageService.communityMemberDAO.CountMemberByCommunityID(id)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return 0, err1
	}
	return count, nil
}

func (communityManageService *CommunityManageService) GetOneCommunityByID(id int) (entity.Community, error) {
	community, err1 := communityManageService.communityDAO.GetOneCommunityByID(id)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return entity.Community{}, err1
	}
	return community, nil
}

func (communityManageService *CommunityManageService) GetCommunitiesByNameFuzzyMatch(name string, pageNO, pageSize int) ([]entity.Community, int64, error) {
	communities, err1 := communityManageService.communityDAO.GetCommunitiesByNameFuzzyMatch(name, (pageNO-1)*pageSize, pageSize)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return nil, -1, err1
	}

	count, err := communityManageService.communityDAO.CountByNameFuzzyMatch(name)
	if err != nil {
		logger.AppLogger.Error(err1.Error())
		return nil, -1, err1
	}
	totalPageNO := count / int64(pageSize)
	if count%int64(pageSize) != 0 {
		totalPageNO += 1
	}

	return communities, totalPageNO, nil
}

func (communityManageService *CommunityManageService) GetCommunities(pageNO, pageSize int) ([]entity.Community, int64, error) {
	communities, err1 := communityManageService.communityDAO.GetCommunities((pageNO-1)*pageSize, pageSize)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return nil, -1, err1
	}

	count, err := communityManageService.communityDAO.CountCommunities()
	if err != nil {
		logger.AppLogger.Error(err1.Error())
		return nil, -1, err1
	}
	totalPageNO := count / int64(pageSize)
	if count%int64(pageSize) != 0 {
		totalPageNO += 1
	}

	return communities, totalPageNO, nil
}

func (communityManageService *CommunityManageService) JoinCommunityByID(id int, username string) error {
	_, err1 := communityManageService.communityDAO.GetOneCommunityByID(id)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return err1
	}

	err2 := communityManageService.communityMemberDAO.Create(id, username, utils.GetCurrentDate())
	if err2 != nil {
		if strings.Contains(err2.Error(), "Duplicate") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err2.Error())
		return err2
	}
	return nil
}

func (communityManageService *CommunityManageService) LeaveCommunityByID(id int, username string) error {
	err2 := communityManageService.communityMemberDAO.DeleteByCommunityIDAndMember(id, username)
	if err2 != nil {
		logger.AppLogger.Error(err2.Error())
	}
	return nil
}
