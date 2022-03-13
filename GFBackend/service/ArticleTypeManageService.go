package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var articleTypeManageServiceLock sync.Mutex
var articleTypeManageService *ArticleTypeManageService

type IArticleTypeManageService interface {
}

type ArticleTypeManageService struct {
	articleTypeDAO dao.IArticleTypeDAO
}

func NewArticleTypeManageService(articleTypeDAO dao.IArticleTypeDAO) *ArticleTypeManageService {
	if articleTypeManageService == nil {
		articleTypeManageServiceLock.Lock()
		if articleTypeManageService == nil {
			articleTypeManageService = &ArticleTypeManageService{
				articleTypeDAO: articleTypeDAO,
			}
		}
		articleTypeManageServiceLock.Unlock()
	}
	return articleTypeManageService
}

var ArticleTypeManageServiceSet = wire.NewSet(
	dao.NewArticleTypeDAO,
	wire.Bind(new(dao.IArticleTypeDAO), new(*dao.ArticleTypeDAO)),
	NewArticleTypeManageService,
)
