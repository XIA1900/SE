package service

import (
	"GFBackend/entity"
	"GFBackend/logger"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"github.com/google/wire"
	"strings"
	"sync"
)

var articleFavoriteServiceLock sync.Mutex
var articleFavoriteService *ArticleFavoriteService

type IArticleFavoriteService interface {
	CreateFavorite(username string, articleID int) error
	DeleteFavorite(username string, articleID int) error
	GetUserFavorites(username string, pageNO, pageSize int) (entity.ArticleFavoritesInfo, error)
}

type ArticleFavoriteService struct {
	articleDAO         dao.IArticleDAO
	articleFavoriteDAO dao.IArticleFavoriteDAO
}

func NewArticleFavoriteService(articleFavoriteDAO dao.IArticleFavoriteDAO, articleDAO dao.IArticleDAO) *ArticleFavoriteService {
	if articleFavoriteService == nil {
		articleFavoriteServiceLock.Lock()
		if articleFavoriteService == nil {
			articleFavoriteService = &ArticleFavoriteService{
				articleDAO:         articleDAO,
				articleFavoriteDAO: articleFavoriteDAO,
			}
		}
		articleFavoriteServiceLock.Unlock()
	}
	return articleFavoriteService
}

var ArticleFavoriteServiceSet = wire.NewSet(
	dao.NewArticleDAO,
	wire.Bind(new(dao.IArticleDAO), new(*dao.ArticleDAO)),
	dao.NewArticleFavoriteDAO,
	wire.Bind(new(dao.IArticleFavoriteDAO), new(*dao.ArticleFavoriteDAO)),
	NewArticleFavoriteService,
)

func (articleFavoriteService *ArticleFavoriteService) CreateFavorite(username string, articleID int) error {
	_, err1 := articleFavoriteService.articleDAO.GetArticleByID(articleID)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}

	err2 := articleFavoriteService.articleFavoriteDAO.CreateFavorite(username, articleID, utils.GetCurrentDate())
	if err2 != nil {
		if strings.Contains(err2.Error(), "Duplicate") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}

	return nil
}

func (articleFavoriteService *ArticleFavoriteService) DeleteFavorite(username string, articleID int) error {
	articleFavorite, err1 := articleFavoriteService.articleFavoriteDAO.GetOne(username, articleID)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}

	if articleFavorite.Username == username {
		err2 := articleFavoriteService.articleFavoriteDAO.DeleteFavorite(username, articleID)
		if err2 != nil {
			logger.AppLogger.Error(err1.Error())
			return errors.New("500")
		}
	}

	return nil
}

func (articleFavoriteService *ArticleFavoriteService) GetUserFavorites(username string, pageNO, pageSize int) (entity.ArticleFavoritesInfo, error) {
	favorites, err1 := articleFavoriteService.articleFavoriteDAO.GetFavoritesByUsername(username, (pageNO-1)*pageSize, pageSize)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return entity.ArticleFavoritesInfo{}, errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return entity.ArticleFavoritesInfo{}, errors.New("500")
	}

	count, err2 := articleFavoriteService.articleFavoriteDAO.CountFavoritesByUsername(username)
	if err2 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return entity.ArticleFavoritesInfo{}, errors.New("400")
		}
		logger.AppLogger.Error(err2.Error())
		return entity.ArticleFavoritesInfo{}, errors.New("500")
	}

	totalPageNO := count / int64(pageSize)
	if count%int64(pageSize) != 0 {
		totalPageNO += 1
	}

	return entity.ArticleFavoritesInfo{
		PageNO:           pageNO,
		PageSize:         pageSize,
		TotalPageNO:      totalPageNO,
		ArticleFavorites: favorites,
	}, nil
}
