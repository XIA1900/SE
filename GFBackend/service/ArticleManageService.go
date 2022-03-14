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
	articleDAO     dao.IArticleDAO
	articleTypeDAO dao.IArticleTypeDAO
	communityDAO   dao.ICommunityDAO
}

func NewArticleManageService(articleDAO dao.IArticleDAO, articleTypeDAO dao.IArticleTypeDAO, communityDAO dao.ICommunityDAO) *ArticleManageService {
	if articleManageService == nil {
		articleManageServiceLock.Lock()
		if articleManageService == nil {
			articleManageService = &ArticleManageService{
				articleDAO:     articleDAO,
				articleTypeDAO: articleTypeDAO,
				communityDAO:   communityDAO,
			}
		}
		articleManageServiceLock.Unlock()
	}
	return articleManageService
}

var ArticleManageServiceSet = wire.NewSet(
	dao.NewArticleDAO,
	wire.Bind(new(dao.IArticleDAO), new(*dao.ArticleDAO)),
	dao.NewArticleTypeDAO,
	wire.Bind(new(dao.IArticleTypeDAO), new(*dao.ArticleTypeDAO)),
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	NewArticleManageService,
)
