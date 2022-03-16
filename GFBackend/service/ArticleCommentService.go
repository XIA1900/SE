package service

import (
	"GFBackend/model/dao"
	"github.com/google/wire"
	"sync"
)

var articleCommentServiceLock sync.Mutex
var articleCommentService *ArticleCommentService

type IArticleCommentService interface {
}

type ArticleCommentService struct {
	articleCommentDAO dao.IArticleCommentDAO
}

func NewArticleCommentService(articleCommentDAO dao.IArticleCommentDAO) *ArticleCommentService {
	if articleCommentService == nil {
		articleCommentServiceLock.Lock()
		if articleCommentService == nil {
			articleCommentService = &ArticleCommentService{
				articleCommentDAO: articleCommentDAO,
			}
		}
		articleCommentServiceLock.Unlock()
	}
	return articleCommentService
}

var ArticleCommentServiceSet = wire.NewSet(
	dao.NewArticleCommentDAO,
	wire.Bind(new(dao.IArticleCommentDAO), new(*dao.ArticleCommentDAO)),
	NewArticleCommentService,
)
