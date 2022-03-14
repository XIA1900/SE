package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var articleManageServiceLock sync.Mutex
var articleManageService *ArticleManageService

type IArticleManageService interface {
}

type ArticleManageService struct {
	articleDAO dao.IArticleDAO
}

func NewArticleManageService(articleDAO dao.IArticleDAO) *ArticleManageService {
	if articleManageService == nil {
		articleManageServiceLock.Lock()
		if articleManageService == nil {
			articleManageService = &ArticleManageService{
				articleDAO: articleDAO,
			}
		}
		articleManageServiceLock.Unlock()
	}
	return articleManageService
}

var ArticleManageServiceSet = wire.NewSet(
	dao.NewArticleDAO,
	wire.Bind(new(dao.IArticleDAO), new(*dao.ArticleDAO)),
	NewArticleManageService,
)