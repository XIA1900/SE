package service

import (
	"GFBackend/model"
	"GFBackend/model/dao"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"sync"
)

var fileManageServiceLock sync.Mutex
var fileManageService *FileManageService

type IFileManageService interface {
	GetSpaceInfo(username string) (model.Space, error)
	RegisterSpaceInfo(username string) error
	UpdateRemaining(username string) error
	ExpandSize(username string, newSize float64) error
	FreeSpace(username string) error
	Upload(context *gin.Context) error
	Download(context *gin.Context) error
}

type FileManageService struct {
	spaceDAO dao.ISpaceDAO
}

func NewFileManageService(spaceDAO dao.ISpaceDAO) *FileManageService {
	if fileManageService == nil {
		fileManageServiceLock.Lock()
		if fileManageService == nil {
			fileManageService = &FileManageService{
				spaceDAO: spaceDAO,
			}
		}
		fileManageServiceLock.Unlock()
	}
	return fileManageService
}

var FileManageServiceSet = wire.NewSet(
	dao.NewSpaceDAO,
	wire.Bind(new(dao.ISpaceDAO), new(*dao.SpaceDAO)),
	NewFileManageService,
)

func (fileManageService FileManageService) GetSpaceInfo(username string) (model.Space, error) {
	return model.Space{}, nil
}

func (fileManageService FileManageService) RegisterSpaceInfo(username string) error {
	return nil
}

func (fileManageService FileManageService) UpdateRemaining(username string) error {
	return nil
}

func (fileManageService FileManageService) ExpandSize(username string, newSize float64) error {
	return nil
}

func (fileManageService FileManageService) FreeSpace(username string) error {
	return nil
}

func (fileManageService FileManageService) Upload(context *gin.Context) error {
	return nil
}

func (fileManageService FileManageService) Download(context *gin.Context) error {
	return nil
}
