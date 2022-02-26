package controller

import (
	"GFBackend/service"
	"github.com/gin-gonic/gin"
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

// StaticResourcesReqs godoc
// @Summary Request User Files
// @Description Static files request, need to claim the username and filename in the url
// @Tags Static Resource
// @Accept json
// @Produce json
// @Router /resources/userfiles/{username}/{filename} [get]
func (fileManageController *FileManageController) StaticResourcesReqs() {}

// UploadFile godoc
// @Summary User Uploads files including images, video etc.
// @Description need token in cookie, html file type input element include name attribute with value "uploadFile"
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Upload Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or No Enough Space"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/upload [post]
func (fileManageController *FileManageController) UploadFile(context *gin.Context) {

}

// ScanFiles godoc
// @Summary Scan User files
// @Description need token in cookie, only get self files
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Scan Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or No Enough Space"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/upload [post]
func (fileManageController *FileManageController) ScanFiles(context *gin.Context) {
	// define struct including list in DTO as return data type
}
