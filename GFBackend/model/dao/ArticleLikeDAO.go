package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleLikeDAOLock sync.Mutex
var articleLikeDAO *ArticleLikeDAO

type IArticleLikeDAO interface {
	CreateLike(username string, articleID int, likeDay string) error
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

func (articleLikeDAO *ArticleLikeDAO) CreateLike(username string, articleID int, likeDay string) error {
	result := articleLikeDAO.db.Omit("ID").Create(&model.ArticleLike{
		Username:  username,
		ArticleID: articleID,
		LikeDay:   likeDay,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
