package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleLikeDAOLock sync.Mutex
var articleLikeDAO *ArticleLikeDAO

type IArticleLikeDAO interface {
}

type ArticleLikeDAO struct {
	db *gorm.DB
}

func NewArticleLikeDAO() *ArticleLikeDAO {
	if articleLikeDAO == nil {
		articleLikeDAOLock.Lock()
		if articleLikeDAO == nil {
			articleLikeDAO = &ArticleLikeDAO{
				db: model.NewDB(),
			}
		}
		articleLikeDAOLock.Unlock()
	}
	return articleLikeDAO
}