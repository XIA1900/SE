package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleCommentDAOLock sync.Mutex
var articleCommentDAO *ArticleCommentDAO

type IArticleCommentDAO interface {
}

type ArticleCommentDAO struct {
	db *gorm.DB
}

func NewArticleCommentDAO() *ArticleCommentDAO {
	if articleCommentDAO == nil {
		articleCommentDAOLock.Lock()
		if articleCommentDAO == nil {
			articleCommentDAO = &ArticleCommentDAO{
				db: model.NewDB(),
			}
		}
		articleCommentDAOLock.Unlock()
	}
	return articleCommentDAO
}
