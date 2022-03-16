package controller

import (
	"GFBackend/service"
	"github.com/google/wire"
	"sync"
)

var articleCommentControllerLock sync.Mutex
var articleCommentController *ArticleCommentController

type ArticleCommentController struct {
	articleCommentService service.IArticleCommentService
}

func NewArticleCommentController(articleCommentService service.IArticleCommentService) *ArticleCommentController {
	if articleCommentController == nil {
		articleCommentControllerLock.Lock()
		if articleCommentController == nil {
			articleCommentController = &ArticleCommentController{
				articleCommentService: articleCommentService,
			}
		}
		articleCommentControllerLock.Unlock()
	}
	return articleCommentController
}

var ArticleCommentControllerSet = wire.NewSet(
	service.ArticleCommentServiceSet,
	wire.Bind(new(service.IArticleCommentService), new(*service.ArticleCommentService)),
	NewArticleCommentController,
)
