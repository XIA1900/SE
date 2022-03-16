package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var articleFavoriteServiceLock sync.Mutex
var articleFavoriteService *ArticleFavoriteService

type IArticleFavoriteService interface {
}

type ArticleFavoriteService struct {
	articleFavoriteDAO dao.IArticleFavoriteDAO
}

func NewArticleFavoriteService(articleFavoriteDAO dao.IArticleFavoriteDAO) *ArticleFavoriteService {
	if articleFavoriteService == nil {
		articleFavoriteServiceLock.Lock()
		if articleFavoriteService == nil {
			articleFavoriteService = &ArticleFavoriteService{
				articleFavoriteDAO: articleFavoriteDAO,
			}
		}
		articleFavoriteServiceLock.Unlock()
	}
	return articleFavoriteService
}

var ArticleFavoriteServiceSet = wire.NewSet(
	dao.NewArticleFavoriteDAO,
	wire.Bind(new(dao.IArticleFavoriteDAO), new(*dao.ArticleFavoriteDAO)),
	NewArticleFavoriteService,
)
