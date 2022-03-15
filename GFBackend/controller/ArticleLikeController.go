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

// CreateLike godoc
// @Summary User like Article
// @Description need token in cookie, need article id
// @Tags Article Like Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param ArticleID query integer true "233333"
// @Success 200 {object} entity.ResponseMsg "<b>Success</b>. Create Successfully"
// @Failure 400 {object} entity.ResponseMsg "<b>Failure</b>. Bad Parameters"
// @Failure 500 {object} entity.ResponseMsg "<b>Failure</b>. Server Internal Error."
// @Router /articlelike/create/:ArticleID [get]
func (articleLikeController ArticleLikeController) CreateLike(context *gin.Context) {

}
