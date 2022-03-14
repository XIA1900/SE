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

func (articleManageController ArticleManageController) CreateArticle(context *gin.Context) {

}
