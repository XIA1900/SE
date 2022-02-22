package service

import (
	"GFBackend/logger"
	"GFBackend/model"
	"GFBackend/model/dao"
	"fmt"
	"gorm.io/gorm"
)

type ICommunityManageService interface {
	CreateCommunity(creator string, name string, description string) error
}

type CommunityManageService struct {
	communityDAO dao.ICommunityDAO
}

func NewCommunityManageService(communityDAO dao.ICommunityDAO) *CommunityManageService {
	return &CommunityManageService{communityDAO: communityDAO}
}

func (communityManageService *CommunityManageService) CreateCommunity(creator string, name string, description string) error {
	newCommunity := model.Community{
		Creator:     creator,
		Name:        name,
		Description: description,
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
