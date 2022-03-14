package controller

import (
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"sync"
)

var articleLikeControllerLock sync.Mutex
var articleLikeController *ArticleLikeController

type ArticleLikeController struct {
	articleLikeService service.IArticleLikeService
}

func NewArticleLikeController(articleLikeService service.IArticleLikeService) *ArticleLikeController {
	if articleLikeController == nil {
		articleLikeControllerLock.Lock()
		if articleLikeController == nil {
			articleLikeController = &ArticleLikeController{
				articleLikeService: articleLikeService,
			}
		}
		articleLikeControllerLock.Unlock()
	}
	return articleLikeController
}

var ArticleLikeControllerSet = wire.NewSet(
	service.ArticleLikeServiceSet,
	wire.Bind(new(service.IArticleLikeService), new(*service.ArticleLikeService)),
	NewArticleLikeController,
)

func (articleLikeController ArticleLikeController) CreateLike(context *gin.Context) {

}
