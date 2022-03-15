package controller

import (
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"sync"
)

var articleManageControllerLock sync.Mutex
var articleManageController *ArticleManageController

type ArticleManageController struct {
	articleManageService service.IArticleManageService
}

func NewArticleManageController(articleManageService service.IArticleManageService) *ArticleManageController {
	if articleManageController == nil {
		articleManageControllerLock.Lock()
		if articleManageController == nil {
			articleManageController = &ArticleManageController{
				articleManageService: articleManageService,
			}
		}
		articleManageControllerLock.Unlock()
	}
	return articleManageController
}

var ArticleManageControllerSet = wire.NewSet(
	service.ArticleManageServiceSet,
	wire.Bind(new(service.IArticleManageService), new(*service.ArticleManageService)),
	NewArticleManageController,
)

// CreateArticle godoc
// @Summary Create a new article
// @Description need token in cookie, need new article info
// @Tags Article Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param ArticleInfo body entity.ArticleInfo true "Create New Article"
// @Success 200 {object} entity.ResponseMsg "<b>Success</b>. Create Successfully"
// @Failure 400 {object} entity.ResponseMsg "<b>Failure</b>. Bad Parameters / Info Error"
// @Failure 500 {object} entity.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /article/create [post]
func (articleManageController ArticleManageController) CreateArticle(context *gin.Context) {

}

// SearchArticle godoc
// @Summary Create a new article
// @Description need token in cookie, need new article search info
// @Tags Article Manage
// @Accept json
// @Produce json
// @Param SearchWord query string true "Search Word"
// @Success 200 {object} entity.ResponseMsg "<b>Success</b>. Search Successfully"
// @Failure 400 {object} entity.ResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} entity.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /article/search/:SearchWord [post]
func (articleManageController ArticleManageController) SearchArticle(context *gin.Context) {

}
