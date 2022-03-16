package dao

import (
	"GFBackend/entity"
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleCommentDAOLock sync.Mutex
var articleCommentDAO *ArticleCommentDAO

type IArticleCommentDAO interface {
	CountCommentsOfArticle(articleID int) (int64, error)
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

func (articleCommentDAO *ArticleCommentDAO) CountCommentsOfArticle(articleID int) (int64, error) {
	var count int64
	result := articleCommentDAO.db.Model(&entity.ArticleLike{}).Where("ArticleID = ?", articleID).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
