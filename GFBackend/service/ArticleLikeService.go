package service

import (
	"GFBackend/logger"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"github.com/google/wire"
	"strings"
	"sync"
)

var articleLikeServiceLock sync.Mutex
var articleLikeService *ArticleLikeService

type IArticleLikeService interface {
	CreateLike(username string, articleID int) error
}

type ArticleLikeService struct {
	articleLikeDAO dao.IArticleLikeDAO
}

func NewArticleLikeService(articleLikeDAO dao.IArticleLikeDAO) *ArticleLikeService {
	if articleManageService == nil {
		articleLikeServiceLock.Lock()
		if articleLikeService == nil {
			articleLikeService = &ArticleLikeService{
				articleLikeDAO: articleLikeDAO,
			}
		}
		articleLikeServiceLock.Unlock()
	}
	return articleLikeService
}

var ArticleLikeServiceSet = wire.NewSet(
	dao.NewArticleLikeDAO,
	wire.Bind(new(dao.IArticleLikeDAO), new(*dao.ArticleLikeDAO)),
	NewArticleLikeService,
)

func (articleLikeService *ArticleLikeService) CreateLike(username string, articleID int) error {
	err1 := articleLikeService.articleLikeDAO.CreateLike(username, articleID, utils.GetCurrentDate())
	if err1 != nil {
		if strings.Contains(err1.Error(), "Duplicate") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}
	return nil
}
