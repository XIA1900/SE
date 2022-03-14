package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleDAOLock sync.Mutex
var articleDAO *ArticleDAO

type IArticleDAO interface {
}

type ArticleDAO struct {
	db *gorm.DB
}

func NewArticleDAO() *ArticleDAO {
	if articleDAO == nil {
		articleDAOLock.Lock()
		if articleDAO == nil {
			articleDAO = &ArticleDAO{
				db: model.NewDB(),
			}
		}
		articleDAOLock.Unlock()
	}
	return articleDAO
}
