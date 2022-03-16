package controller

import (
	"GFBackend/service"
	"github.com/google/wire"
	"sync"
)

var articleFavoriteControllerLock sync.Mutex
var articleFavoriteController *ArticleFavoriteController

type ArticleFavoriteController struct {
	articleFavoriteService service.IArticleFavoriteService
}

func NewArticleFavoriteController(articleFavoriteService service.IArticleFavoriteService) *ArticleFavoriteController {
	if articleFavoriteController == nil {
		articleFavoriteControllerLock.Lock()
		if articleFavoriteController == nil {
			articleFavoriteController = &ArticleFavoriteController{
				articleFavoriteService: articleFavoriteService,
			}
		}
		articleFavoriteControllerLock.Unlock()
	}
	return articleFavoriteController
}

var ArticleFavoriteControllerSet = wire.NewSet(
	service.ArticleFavoriteServiceSet,
	wire.Bind(new(service.IArticleFavoriteService), new(*service.ArticleFavoriteService)),
	NewArticleFavoriteController,
)
