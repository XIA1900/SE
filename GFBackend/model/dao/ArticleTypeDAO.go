package dao

import (
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleTypeDAOLock sync.Mutex
var articleTypeDAO *ArticleTypeDAO

type IArticleTypeDAO interface {
	CreateArticleType(articleType model.ArticleType) error
	RemoveArticleType(articleTypeID int) error
	GetArticleTypes() ([]model.ArticleType, error)
}

type ArticleTypeDAO struct {
	db *gorm.DB
}

func NewArticleTypeDAO() *ArticleTypeDAO {
	if articleTypeDAO == nil {
		articleTypeDAOLock.Lock()
		if articleTypeDAO == nil {
			articleTypeDAO = &ArticleTypeDAO{
				db: model.NewDB(),
			}
		}
		articleTypeDAOLock.Unlock()
	}
	return articleTypeDAO
}

func (articleTypeDAO *ArticleTypeDAO) CreateArticleType(articleType model.ArticleType) error {
	return nil
}

func (articleTypeDAO *ArticleTypeDAO) RemoveArticleType(articleTypeID int) error {
	return nil
}

func (articleTypeDAO *ArticleTypeDAO) GetArticleTypes() ([]model.ArticleType, error) {
	return nil, nil
}
