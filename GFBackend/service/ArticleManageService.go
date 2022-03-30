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
	GetOneArticleByID(id int) (entity.ArticleDetail, error)
	GetArticlesBySearchWords(articleSearchInfo entity.ArticleSearchInfo) (entity.ArticlesForSearching, error)
	GetArticleList(pageNO, pageSize int) ([]entity.Article, error)
}

type ArticleManageService struct {
	articleDAO         dao.IArticleDAO
	articleTypeDAO     dao.IArticleTypeDAO
	articleLikeDAO     dao.IArticleLikeDAO
	articleFavoriteDAO dao.IArticleFavoriteDAO
	articleCommentDAO  dao.IArticleCommentDAO
	communityDAO       dao.ICommunityDAO
}

func NewArticleManageService(articleDAO dao.IArticleDAO, articleTypeDAO dao.IArticleTypeDAO,
	communityDAO dao.ICommunityDAO, articleCommentDAO dao.IArticleCommentDAO,
	articleLikeDAO dao.IArticleLikeDAO, articleFavoriteDAO dao.IArticleFavoriteDAO) *ArticleManageService {
	if articleManageService == nil {
		articleManageServiceLock.Lock()
		if articleManageService == nil {
			articleManageService = &ArticleManageService{
				articleDAO:         articleDAO,
				articleTypeDAO:     articleTypeDAO,
				articleCommentDAO:  articleCommentDAO,
				articleLikeDAO:     articleLikeDAO,
				articleFavoriteDAO: articleFavoriteDAO,
				communityDAO:       communityDAO,
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
	dao.NewArticleCommentDAO,
	wire.Bind(new(dao.IArticleCommentDAO), new(*dao.ArticleCommentDAO)),
	dao.NewArticleLikeDAO,
	wire.Bind(new(dao.IArticleLikeDAO), new(*dao.ArticleLikeDAO)),
	dao.NewArticleFavoriteDAO,
	wire.Bind(new(dao.IArticleFavoriteDAO), new(*dao.ArticleFavoriteDAO)),
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

		err3 := articleManageService.articleCommentDAO.DeleteCommentByArticleID(id)
		if err3 != nil {
			logger.AppLogger.Error(err3.Error())
			return err3
		}

		err4 := articleManageService.articleLikeDAO.DeleteLikeByArticleID(id)
		if err4 != nil {
			logger.AppLogger.Error(err4.Error())
			return err4
		}
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

func (articleManageService *ArticleManageService) GetOneArticleByID(id int) (entity.ArticleDetail, error) {
	article, err1 := articleManageService.articleDAO.GetArticleByID(id)
	if err1 != nil {
		if strings.Contains(err1.Error(), "not found") {
			return entity.ArticleDetail{}, errors.New("400")
		}
		logger.AppLogger.Error(err1.Error())
		return entity.ArticleDetail{}, errors.New("500")
	}

	articleType, err2 := articleManageService.articleTypeDAO.GetArticleTypeByID(article.TypeID)
	if err2 != nil {
		if strings.Contains(err2.Error(), "not found") {
			return entity.ArticleDetail{}, errors.New("400")
		}
		logger.AppLogger.Error(err2.Error())
		return entity.ArticleDetail{}, errors.New("500")
	}

	community, err3 := articleManageService.communityDAO.GetOneCommunityByID(article.CommunityID)
	if err3 != nil {
		if strings.Contains(err3.Error(), "not found") {
			return entity.ArticleDetail{}, errors.New("400")
		}
		logger.AppLogger.Error(err3.Error())
		return entity.ArticleDetail{}, err3
	}

	countLikeOfArticle, err4 := articleManageService.articleLikeDAO.CountLikeOfArticle(article.ID)
	if err4 != nil {
		logger.AppLogger.Error(err4.Error())
		return entity.ArticleDetail{}, err4
	}

	countFavoriteOfArticle, err5 := articleManageService.articleFavoriteDAO.CountFavoriteOfArticle(article.ID)
	if err5 != nil {
		logger.AppLogger.Error(err5.Error())
		return entity.ArticleDetail{}, err5
	}

	countCommentsOfArticle, err6 := articleManageService.articleCommentDAO.CountCommentsOfArticle(article.ID)
	if err6 != nil {
		logger.AppLogger.Error(err6.Error())
		return entity.ArticleDetail{}, err6
	}

	return entity.ArticleDetail{
		ID:            article.ID,
		Title:         article.Title,
		TypeName:      articleType.TypeName,
		CommunityName: community.Name,
		Content:       article.Content,
		NumLike:       countLikeOfArticle,
		NumFavorite:   countFavoriteOfArticle,
		NumComment:    countCommentsOfArticle,
	}, nil
}

func (articleManageService *ArticleManageService) GetArticlesBySearchWords(articleSearchInfo entity.ArticleSearchInfo) (entity.ArticlesForSearching, error) {
	searchWords := articleSearchInfo.SearchWords
	from := (articleSearchInfo.PageNO - 1) * articleSearchInfo.PageSize
	size := articleSearchInfo.PageSize
	documents, count := elasticsearch.MixSearchDocuments(searchWords, from, size)

	totalPageNO := count / int64(size)
	if count%int64(size) != 0 {
		totalPageNO += 1
	}

	return entity.ArticlesForSearching{
		PageNO:      articleSearchInfo.PageNO,
		PageSize:    articleSearchInfo.PageSize,
		TotalPageNO: totalPageNO,
		Articles:    documents,
	}, nil
}

func (articleManageService *ArticleManageService) GetArticleList(pageNO, pageSize int) ([]entity.Article, error) {
	articles, err1 := articleManageService.articleDAO.GetArticleList((pageNO-1)*pageSize, pageSize)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return nil, err1
	}
	return articles, nil
}
