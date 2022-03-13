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
	RemoveArticleType(typeName string) error
	UpdateDescription(typeName, newDescription string) error
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
	result := articleTypeDAO.db.Omit("ID").Create(&articleType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleTypeDAO *ArticleTypeDAO) RemoveArticleType(typeName string) error {
	result := articleTypeDAO.db.Where("typeName = ?", typeName).Delete(&model.ArticleType{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleTypeDAO *ArticleTypeDAO) UpdateDescription(typeName, newDescription string) error {
	result := articleTypeDAO.db.Model(&model.ArticleType{}).Where("typeName", typeName).Update("Description", newDescription)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (articleTypeDAO *ArticleTypeDAO) GetArticleTypes() ([]model.ArticleType, error) {
	var articleTypes []model.ArticleType
	result := articleTypeDAO.db.Find(&articleTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return articleTypes, nil
}
