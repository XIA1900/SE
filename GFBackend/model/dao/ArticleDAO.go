package dao

import (
	"GFBackend/entity"
	"GFBackend/model"
	"gorm.io/gorm"
	"sync"
)

var articleDAOLock sync.Mutex
var articleDAO *ArticleDAO

type IArticleDAO interface {
	CreateArticle(article entity.Article) (int, error)
	DeleteArticleByID(id int) error
	UpdateArticleTitleOrContentByID(id int, newTitle, newContent string) error
	GetArticleByID(id int) (entity.Article, error)
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

func (articleDAO *ArticleDAO) CreateArticle(article entity.Article) (int, error) {
	result := articleDAO.db.Create(&article)
	if result.Error != nil {
		return -1, result.Error
	}
	return article.ID, nil
}

func (articleDAO *ArticleDAO) DeleteArticleByID(id int) error {
	result := articleDAO.db.Where("ID = ?", id).Delete(&entity.Article{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleDAO *ArticleDAO) UpdateArticleTitleOrContentByID(id int, newTitle, newContent string) error {
	result := articleDAO.db.Model(&entity.Article{}).Where("ID = ?", id).
		Update("Title", newTitle).Update("Content", newContent)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleDAO *ArticleDAO) GetArticleByID(id int) (entity.Article, error) {
	var article entity.Article
	result := articleDAO.db.Where("ID = ?", id).First(&article)
	if result.Error != nil {
		return entity.Article{}, result.Error
	}
	return article, nil
}
