package dao

import (
	"GFBackend/entity"
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleLikeDAOLock sync.Mutex
var articleLikeDAO *ArticleLikeDAO

type IArticleLikeDAO interface {
	CreateLike(username string, articleID int, likeDay string) error
	DeleteLike(username string, articleID int) error
	GetLike(username string, articleID int) (entity.ArticleLike, error)
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
	result := articleLikeDAO.db.Omit("ID").Create(&entity.ArticleLike{
		Username:  username,
		ArticleID: articleID,
		LikeDay:   likeDay,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleLikeDAO *ArticleLikeDAO) DeleteLike(username string, articleID int) error {
	result := articleLikeDAO.db.Where("Username = ? AND ArticleID = ?", username, articleID).Delete(&entity.ArticleLike{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleLikeDAO *ArticleLikeDAO) GetLike(username string, articleID int) (entity.ArticleLike, error) {
	var articleLike entity.ArticleLike
	result := articleLikeDAO.db.Where("Username = ? AND ArticleID = ?", username, articleID).First(&articleLike)
	if result.Error != nil {
		return entity.ArticleLike{}, result.Error
	}
	return articleLike, nil
}
