package controller

import (
	"GFBackend/middleware/auth"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strings"
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
// @Description need token in cookie, html file type input element include name attribute with value "uploadFilename"
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Upload Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or No Enough Space"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/upload [post]
func (fileManageController *FileManageController) UploadFile(context *gin.Context) {
	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)
	file, err1 := context.FormFile("uploadFilename")
	errMsg := ResponseMsg{
		Code:    http.StatusBadRequest,
		Message: "Bad Parameters",
	}
	if err1 != nil {
		context.JSON(400, errMsg)
		return
	}
	err2 := fileManageController.fileManageService.Upload(context, username, file)
	if err2 != nil {
		if strings.Contains(err2.Error(), "400") {
			errMsg.Code = 400
			errMsg.Message = "Not enough space"
		} else if strings.Contains(err2.Error(), "500") {
			errMsg.Code = 500
			errMsg.Message = "Internal Server Error"
		}
		context.JSON(errMsg.Code, errMsg)
		return
	}

	context.JSON(200, ResponseMsg{
		Code:    200,
		Message: "Upload Success",
	})

}

// DownloadFile godoc
// @Summary User Downloads File, only self data for now
// @Description need token in cookie, need filename in json
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param filename body string true "filename in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Upload Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or No Enough Space"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/download [post]
func (fileManageController *FileManageController) DownloadFile(context *gin.Context) {
	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)
	errMsg := ResponseMsg{
		Code:    http.StatusBadRequest,
		Message: "Bad Parameters",
	}
	type Info struct {
		Filename string `json:"filename"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	if err1 != nil || info.Filename == "" {
		context.JSON(400, errMsg)
		return
	}

	err2 := fileManageController.fileManageService.Download(context, username, info.Filename)
	if err2 != nil {
		return
	}
}

// UserDeleteFile godoc
// @Summary Delete User File, only have permission to delete self data
// @Description need token in cookie, need filename in json
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param filename body string true "filename in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Delete Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Other"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/delete [post]
func (fileManageController *FileManageController) UserDeleteFile(context *gin.Context) {
	type Info struct {
		Filename string `json:"filename"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	if err1 != nil || info.Filename == "" {
		errMsg := ResponseMsg{
			Code:    400,
			Message: "No Filename",
		}
		context.JSON(400, errMsg)
		return
	}

	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)
	err2 := fileManageController.fileManageService.DeleteUserFile(username, info.Filename)
	if err2 != nil {
		errMsg := ResponseMsg{
			Code:    400,
			Message: "File not Exists",
		}
		context.JSON(400, errMsg)
		return
	}

	context.JSON(200, ResponseMsg{
		Code:    200,
		Message: "Delete Successfully",
	})
}

// ScanFiles godoc
// @Summary Scan User files
// @Description need token in cookie, only get self files
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Success 201 {object} controller.UserFiles "<b>Success</b>. Scan Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters."
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/scan [post]
func (fileManageController *FileManageController) ScanFiles(context *gin.Context) {
	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)
	userFiles, err1 := fileManageController.fileManageService.GetUserFiles(username)
	if err1 != nil {
		errMsg := ResponseMsg{
			Code:    400,
			Message: "Bad Parameters",
		}
		if strings.Contains(err1.Error(), "500") {
			errMsg.Code = http.StatusInternalServerError
			errMsg.Message = "Internal Server Error"
		}
		context.JSON(errMsg.Code, errMsg)
		return
	}
	respMsg := UserFiles{
		ResponseMsg: ResponseMsg{
			Code:    200,
			Message: "Scan Successfully",
		},
		Filenames: userFiles,
	}
	context.JSON(200, respMsg)
}

// UserSpaceInfo godoc
// @Summary Browse User Space Info
// @Description need token in cookie
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Success 201 {object} model.Space "<b>Success</b>. Get User Space Info Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or User not exists."
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/space/info [post]
func (fileManageController *FileManageController) UserSpaceInfo(context *gin.Context) {
	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)

	spaceInfo, err1 := fileManageController.fileManageService.GetSpaceInfo(username)
	if err1 != nil {
		errMsg := ResponseMsg{
			Code:    http.StatusBadRequest,
			Message: "User not exists",
		}
		if strings.Contains(err1.Error(), "500") {
			errMsg.Code = http.StatusInternalServerError
			errMsg.Message = "Server Internal Error"
		}
		context.JSON(errMsg.Code, errMsg)
	}

	context.JSON(200, spaceInfo)
}

// UpdateUserCapacity godoc
// @Summary Expand User Capacity, only admin user can do this
// @Description need token in cookie, need target user and  new capacity in json
// @Tags Static Resource
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param username body string true "username in post request body"
// @Param capacity body number true "capacity(e.g. 16.8) in post request body"
// @Success 201 {object} controller.ResponseMsg "<b>Success</b>. Update Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Other"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /file/space/update [post]
func (fileManageController *FileManageController) UpdateUserCapacity(context *gin.Context) {
	type Info struct {
		Username string  `json:"username"`
		Capacity float64 `json:"capacity"`
	}
	var info Info
	err1 := context.ShouldBind(&info)
	if err1 != nil {
		errMsg := ResponseMsg{
			Code:    400,
			Message: "Bad Parameters",
		}
		context.JSON(400, errMsg)
		return
	}

	err2 := fileManageController.fileManageService.UpdateCapacity(info.Username, info.Capacity)
	if err2 != nil {
		errMsg := ResponseMsg{
			Code:    400,
			Message: "Internal Server Error",
		}
		context.JSON(500, errMsg)
		return
	}

	context.JSON(200, ResponseMsg{
		Code:    200,
		Message: "Update Successfully",
	})
}
