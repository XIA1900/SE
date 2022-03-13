package controller

import (
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"sync"
)

var articleTypeManageControllerLock sync.Mutex
var articleTypeManageController *ArticleTypeManageController

type ArticleTypeManageController struct {
	articleTypeManageService service.IArticleTypeManageService
}

func NewArticleTypeManageController(articleTypeManageService service.IArticleTypeManageService) *ArticleTypeManageController {
	if articleTypeManageController == nil {
		articleTypeManageControllerLock.Lock()
		if articleTypeManageController == nil {
			articleTypeManageController = &ArticleTypeManageController{
				articleTypeManageService: articleTypeManageService,
			}
		}
		articleTypeManageControllerLock.Unlock()
	}
	return articleTypeManageController
}

var ArticleTypeManageControllerSet = wire.NewSet(
	service.ArticleTypeManageServiceSet,
	wire.Bind(new(service.IArticleTypeManageService), new(*service.ArticleTypeManageService)),
	NewArticleTypeManageController,
)

// CreateArticleType godoc
// @Summary Create a new article type by admin user
// @Description need token in cookie, need new article type information
// @Tags Article Type Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param ArticleTypeInfo body controller.ArticleTypeInfo true "New Article Type Information"
// @Success 200 {object} controller.ResponseMsg "<b>Success</b>. User Login Successfully"
// @Failure 400 {object} controller.ResponseMsg "<b>Failure</b>. Bad Parameters or Username / Password incorrect"
// @Failure 500 {object} controller.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /articletype/create [post]
func (articleTypeManageController *ArticleTypeManageController) CreateArticleType(context *gin.Context) {

}
