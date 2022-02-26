package dao

import (
	"GFBackend/model"
	"sync"
)

var spaceDAOLock sync.Mutex
var spaceDAO *SpaceDAO

type ISpaceDAO interface {
	CreateSpaceInfo(username string) error
	DeleteSpaceInfo(username string) error
	UpdateRemaining(username string, remainingSize float64) error
	UpdateCapacity(username string, newCapacity float64) error
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

func (spaceDAO *SpaceDAO) CreateSpaceInfo(username string) error {
	return nil
}

func (spaceDAO *SpaceDAO) DeleteSpaceInfo(username string) error {
	return nil
}

func (spaceDAO *SpaceDAO) UpdateRemaining(username string, remainingSize float64) error {
	return nil
}

func (spaceDAO *SpaceDAO) UpdateCapacity(username string, newCapacity float64) error {
	return nil
}

func (spaceDAO *SpaceDAO) GetSpaceInfo(username string) (model.Space, error) {
	return model.Space{}, nil
}
