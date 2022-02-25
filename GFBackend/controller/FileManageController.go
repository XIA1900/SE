package controller

import (
	"GFBackend/service"
	"github.com/google/wire"
	"sync"
)

var fileManageControllerLock sync.Mutex
var fileManageController *FileManageController

type FileManageController struct {
	fileManageService service.IFileManageService
}

func NewFileManageController(fileManageService service.IFileManageService) *FileManageController {
	if fileManageController == nil {
		fileManageControllerLock.Lock()
		if fileManageController == nil {
			fileManageController = &FileManageController{
				fileManageService: fileManageService,
			}
		}
		fileManageControllerLock.Unlock()
	}
	return fileManageController
}

var FileManageControllerSet = wire.NewSet(
	service.FileManageServiceSet,
	wire.Bind(new(service.IFileManageService), new(*service.FileManageService)),
	NewFileManageController,
)
