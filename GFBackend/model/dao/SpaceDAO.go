package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var spaceDAOLock sync.Mutex
var spaceDAO *SpaceDAO

type ISpaceDAO interface {
	CreateSpaceInfo(username string, tx *gorm.DB) error
	DeleteSpaceInfo(username string, tx *gorm.DB) error
	UpdateRemaining(username string, remainingSize float64, tx *gorm.DB) error
	UpdateCapacity(username string, newCapacity float64, tx *gorm.DB) error
	GetSpaceInfo(username string) (model.Space, error)
}

type SpaceDAO struct{}

func NewSpaceDAO() *SpaceDAO {
	if spaceDAO == nil {
		spaceDAOLock.Lock()
		if spaceDAO == nil {
			spaceDAO = new(SpaceDAO)
		}
		spaceDAOLock.Unlock()
	}
	return spaceDAO
}

func (spaceDAO *SpaceDAO) CreateSpaceInfo(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Select("Username").Create(&model.Space{})
	} else {
		result = tx.Select("Username").Create(&model.Space{})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) DeleteSpaceInfo(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Where("Username = ?", username).Delete(&model.Space{})
	} else {
		result = tx.Where("Username = ?", username).Delete(&model.Space{})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) UpdateRemaining(username string, remainingSize float64, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Model(&model.Space{}).Where("username = ?", username).Update("Remaining", remainingSize)
	} else {
		result = tx.Model(&model.Space{}).Where("username = ?", username).Update("Remaining", remainingSize)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) UpdateCapacity(username string, newCapacity float64, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = model.DB.Model(&model.Space{}).Where("username = ?", username).Update("Capacity", newCapacity)
	} else {
		result = tx.Model(&model.Space{}).Where("username = ?", username).Update("Capacity", newCapacity)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) GetSpaceInfo(username string) (model.Space, error) {
	spaceInfo := model.Space{}
	result := model.DB.Where("username = ?", username).First(&spaceInfo)
	if result.Error != nil {
		return spaceInfo, result.Error
	}
	return spaceInfo, nil
}
