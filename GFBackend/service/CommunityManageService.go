package service

import (
	"GFBackend/logger"
	"GFBackend/model"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"fmt"
	"gorm.io/gorm"
)

type ICommunityManageService interface {
	CreateCommunity(creator string, name string, description string, createTime *utils.LocalTime) error
	GetCommunityByName(name string) (model.Community, error)
}

type CommunityManageService struct {
	communityDAO dao.ICommunityDAO
}

func NewCommunityManageService(communityDAO dao.ICommunityDAO) *CommunityManageService {
	return &CommunityManageService{communityDAO: communityDAO}
}

func (communityManageService *CommunityManageService) CreateCommunity(creator string, name string, description string, createTime *utils.LocalTime) error {
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

func (communityManageService *CommunityManageService) GetCommunityByName(name string) (model.Community, error) {
	newCommunity := model.Community{
		Name: name,
	}
	resCommunity, err := communityManageService.communityDAO.GetCommunityByName(newCommunity)
	if err != nil {
		logger.AppLogger.Error(fmt.Sprintf("Get Community By Name Error: %s", err.Error()))
		return model.Community{}, err
	}
	return resCommunity, nil
}
