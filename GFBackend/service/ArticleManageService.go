package service

import (
	"GFBackend/elasticsearch"
	"GFBackend/entity"
	"GFBackend/logger"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"github.com/google/wire"
	"strings"
	"sync"
)

var articleManageServiceLock sync.Mutex
var articleManageService *ArticleManageService

type IArticleManageService interface {
	CreateArticle(username string, articleInfo entity.ArticleInfo) error
	DeleteArticleByID(id int, operator string) error
	UpdateArticleTitleOrContentByID(articleInfo entity.ArticleInfo, operator string) error
}

type ArticleManageService struct {
	articleDAO     dao.IArticleDAO
	articleTypeDAO dao.IArticleTypeDAO
	communityDAO   dao.ICommunityDAO
}

func NewArticleManageService(articleDAO dao.IArticleDAO, articleTypeDAO dao.IArticleTypeDAO, communityDAO dao.ICommunityDAO) *ArticleManageService {
	if articleManageService == nil {
		articleManageServiceLock.Lock()
		if articleManageService == nil {
			articleManageService = &ArticleManageService{
				articleDAO:     articleDAO,
				articleTypeDAO: articleTypeDAO,
				communityDAO:   communityDAO,
			}
		}
		articleManageServiceLock.Unlock()
	}
	return articleManageService
}

var ArticleManageServiceSet = wire.NewSet(
	dao.NewArticleDAO,
	wire.Bind(new(dao.IArticleDAO), new(*dao.ArticleDAO)),
	dao.NewArticleTypeDAO,
	wire.Bind(new(dao.IArticleTypeDAO), new(*dao.ArticleTypeDAO)),
	dao.NewCommunityDAO,
	wire.Bind(new(dao.ICommunityDAO), new(*dao.CommunityDAO)),
	NewArticleManageService,
)

func (articleManageService *ArticleManageService) CreateArticle(username string, articleInfo entity.ArticleInfo) error {
	article := entity.Article{
		Username:    username,
		Title:       articleInfo.Title,
		TypeID:      articleInfo.TypeID,
		CommunityID: articleInfo.CommunityID,
		CreateDay:   utils.GetCurrentDate(),
		Content:     articleInfo.Content,
	}

	_, typeErr := articleManageService.articleTypeDAO.GetArticleTypeByID(article.TypeID)
	if typeErr != nil {
		if strings.Contains(typeErr.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(typeErr.Error())
		return errors.New("500")
	}

	_, communityErr := articleManageService.communityDAO.GetOneCommunityByID(article.CommunityID)
	if communityErr != nil {
		if strings.Contains(communityErr.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(communityErr.Error())
		return errors.New("500")
	}

	articleID, err1 := articleManageService.articleDAO.CreateArticle(article)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return err1
	}

	res := elasticsearch.CreateDocument(entity.ArticleOfES{
		ID:       articleID,
		Username: username,
		Title:    articleInfo.Title,
		Content:  articleInfo.Content,
	})
	if !res {
		return errors.New("article cannot be searched")
	}

	return nil
}

func (articleManageService *ArticleManageService) DeleteArticleByID(id int, operator string) error {
	article, err1 := articleManageService.articleDAO.GetArticleByID(id)
	if err1 != nil {
		if !strings.Contains(err1.Error(), "not found") {
			logger.AppLogger.Error(err1.Error())
		}
		return err1
	}
	if article.Username == operator {
		err2 := articleManageService.articleDAO.DeleteArticleByID(id)
		if err2 != nil {
			logger.AppLogger.Error(err2.Error())
			return err2
		}

		elasticsearch.DeleteDocument(entity.ArticleOfES{
			ID: id,
		})
	}
	return nil
}

func (articleManageService *ArticleManageService) UpdateArticleTitleOrContentByID(articleInfo entity.ArticleInfo, operator string) error {
	article, err1 := articleManageService.articleDAO.GetArticleByID(articleInfo.ID)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}

	if article.Username == operator {
		err2 := articleManageService.articleDAO.UpdateArticleTitleOrContentByID(articleInfo.ID, articleInfo.Title, articleInfo.Content)
		if err2 != nil {
			logger.AppLogger.Error(err1.Error())
			return errors.New("500")
		}

		flag := elasticsearch.UpdateDocument(entity.ArticleOfES{
			ID:       articleInfo.ID,
			Username: article.Username,
			Title:    articleInfo.Title,
			Content:  articleInfo.Content,
		})
		if !flag {
			return errors.New("500")
		}
	} else {
		return errors.New("400")
	}

	return nil
}
