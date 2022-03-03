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
	UpdateUsed(username string, remainingSize float64, tx *gorm.DB) error
	UpdateCapacity(username string, newCapacity float64, tx *gorm.DB) error
	GetSpaceInfo(username string) (model.Space, error)
}

type SpaceDAO struct {
	db *gorm.DB
}

func NewSpaceDAO() *SpaceDAO {
	if spaceDAO == nil {
		spaceDAOLock.Lock()
		if spaceDAO == nil {
			spaceDAO = &SpaceDAO{
				db: model.NewDB(),
			}
		}
		spaceDAOLock.Unlock()
	}
	return spaceDAO
}

func (spaceDAO *SpaceDAO) CreateSpaceInfo(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = spaceDAO.db.Select("Username").Create(&model.Space{
			Username: username,
		})
	} else {
		result = tx.Select("Username").Create(&model.Space{
			Username: username,
		})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) DeleteSpaceInfo(username string, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = spaceDAO.db.Where("Username = ?", username).Delete(&model.Space{})
	} else {
		result = tx.Where("Username = ?", username).Delete(&model.Space{})
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) UpdateUsed(username string, usedSize float64, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = spaceDAO.db.Model(&model.Space{}).Where("username = ?", username).Update("Remaining", usedSize)
	} else {
		result = tx.Model(&model.Space{}).Where("username = ?", username).Update("Remaining", usedSize)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (spaceDAO *SpaceDAO) UpdateCapacity(username string, newCapacity float64, tx *gorm.DB) error {
	var result *gorm.DB
	if tx == nil {
		result = spaceDAO.db.Model(&model.Space{}).Where("username = ?", username).Update("Capacity", newCapacity)
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
	result := spaceDAO.db.Where("username = ?", username).First(&spaceInfo)
	if result.Error != nil {
		return spaceInfo, result.Error
	}
	return spaceInfo, nil
}
