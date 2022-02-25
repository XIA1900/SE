package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var fileManageServiceLock sync.Mutex
var fileManageService *FileManageService

type IFileManageService interface {
}

type FileManageService struct {
	followDAO dao.IFollowDAO
}

func NewFileManageService(followDAO dao.IFollowDAO) *FileManageService {
	if fileManageService == nil {
		fileManageServiceLock.Lock()
		if fileManageService == nil {
			fileManageService = &FileManageService{
				followDAO: followDAO,
			}
		}
		fileManageServiceLock.Unlock()
	}
	return fileManageService
}

var FileManageServiceSet = wire.NewSet(
	dao.NewFollowDAO,
	wire.Bind(new(dao.IFollowDAO), new(*dao.FollowDAO)),
	NewFileManageService,
)
