package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var articleLikeServiceLock sync.Mutex
var articleLikeService *ArticleLikeService

type IArticleLikeService interface {
}

type ArticleLikeService struct {
	articleLikeDAO dao.IArticleLikeDAO
}

func NewArticleLikeService(articleLikeDAO dao.IArticleLikeDAO) *ArticleLikeService {
	if articleManageService == nil {
		articleLikeServiceLock.Lock()
		if articleLikeService == nil {
			articleLikeService = &ArticleLikeService{
				articleLikeDAO: articleLikeDAO,
			}
		}
		articleLikeServiceLock.Unlock()
	}
	return articleLikeService
}

var ArticleLikeServiceSet = wire.NewSet(
	dao.NewArticleLikeDAO,
	wire.Bind(new(dao.IArticleLikeDAO), new(*dao.ArticleLikeDAO)),
	NewArticleLikeService,
)
