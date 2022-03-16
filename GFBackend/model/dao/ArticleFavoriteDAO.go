package dao

import (
	"GFBackend/entity"
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleFavoriteDAOLock sync.Mutex
var articleFavoriteDAO *ArticleFavoriteDAO

type IArticleFavoriteDAO interface {
	CreateFavorite(username string, articleID int) error
	DeleteFavorite(username string, articleID int) error
	GetOne(username string, articleID int) (entity.ArticleFavorite, error)
	GetFavoritesByUsername(username string, offset, limit string) ([]entity.ArticleFavorite, error)
}

type ArticleFavoriteDAO struct {
	db *gorm.DB
}

func NewArticleFavoriteDAO() *ArticleFavoriteDAO {
	if articleFavoriteDAO == nil {
		articleFavoriteDAOLock.Lock()
		if articleFavoriteDAO == nil {
			articleFavoriteDAO = &ArticleFavoriteDAO{
				db: model.NewDB(),
			}
		}
		articleFavoriteDAOLock.Unlock()
	}
	return articleFavoriteDAO
}

func (articleFavoriteDAO *ArticleFavoriteDAO) CreateFavorite(username string, articleID int) error {
	return nil
}

func (articleFavoriteDAO *ArticleFavoriteDAO) DeleteFavorite(username string, articleID int) error {
	return nil
}

func (articleFavoriteDAO *ArticleFavoriteDAO) GetOne(username string, articleID int) (entity.ArticleFavorite, error) {
	return entity.ArticleFavorite{}, nil
}

func (articleFavoriteDAO *ArticleFavoriteDAO) GetFavoritesByUsername(username string, offset, limit string) ([]entity.ArticleFavorite, error) {
	return nil, nil
}
