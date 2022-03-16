package controller

import (
	"GFBackend/entity"
	"GFBackend/middleware/auth"
	"GFBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
	"strings"
	"sync"
)

var articleManageControllerLock sync.Mutex
var articleManageController *ArticleManageController

type ArticleManageController struct {
	articleManageService service.IArticleManageService
}

func NewArticleManageController(articleManageService service.IArticleManageService) *ArticleManageController {
	if articleManageController == nil {
		articleManageControllerLock.Lock()
		if articleManageController == nil {
			articleManageController = &ArticleManageController{
				articleManageService: articleManageService,
			}
		}
		articleManageControllerLock.Unlock()
	}
	return articleManageController
}

var ArticleManageControllerSet = wire.NewSet(
	service.ArticleManageServiceSet,
	wire.Bind(new(service.IArticleManageService), new(*service.ArticleManageService)),
	NewArticleManageController,
)

// CreateArticle godoc
// @Summary Create a new article
// @Description need token in cookie, need new article info
// @Tags Article Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param ArticleInfo body entity.ArticleInfo true "Create New Article"
// @Success 200 {string} string "<b>Success</b>. Create Successfully"
// @Failure 400 {string} string "<b>Failure</b>. Bad Parameters / Info Error"
// @Failure 500 {string} string "<b>Failure</b>. Server Internal Error."
// @Router /article/create [post]
func (articleManageController ArticleManageController) CreateArticle(context *gin.Context) {
	var articleInfo entity.ArticleInfo
	err1 := context.ShouldBindJSON(&articleInfo)
	if err1 != nil {
		context.JSON(400, "Bad Parameters")
		return
	}

	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)

	err2 := articleManageController.articleManageService.CreateArticle(username, articleInfo)
	if err2 != nil {
		if strings.Contains(err2.Error(), "400") {
			context.JSON(400, "Info Error")
		} else {
			context.JSON(500, "Server Internal Error")
		}
		return
	}

	context.JSON(200, "Create Successfully")
}

// DeleteArticle godoc
// @Summary Delete Article By ID
// @Description need token in cookie, need new article id
// @Tags Article Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param id query string true "Article ID"
// @Success 200 {string} string "<b>Success</b>. Delete Successfully no matter what"
// @Failure 400 {string} string "<b>Failure</b>. Bad Parameters"
// @Router /article/delete/:id [get]
func (articleManageController ArticleManageController) DeleteArticle(context *gin.Context) {
	id, err1 := strconv.Atoi(context.Param("id"))
	if err1 != nil {
		context.JSON(400, "Bad Parameters")
		return
	}

	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)

	_ = articleManageController.articleManageService.DeleteArticleByID(id, username)

	context.JSON(200, "Delete Successfully")
}

// UpdateArticleTitleOrContentByID godoc
// @Summary Update Article Title or Content By ID
// @Description need token in cookie, need ID, Title, Content in article info only
// @Tags Article Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param ArticleInfo body entity.ArticleInfo true "Update Article Info"
// @Success 200 {string} string "<b>Success</b>. Update Successfully"
// @Failure 400 {string} string "<b>Failure</b>. Bad Parameters / Not Found"
// @Failure 500 {string} string "<b>Failure</b>. Server Internal Error."
// @Router /article/update [post]
func (articleManageController ArticleManageController) UpdateArticleTitleOrContentByID(context *gin.Context) {
	var articleInfo entity.ArticleInfo
	err1 := context.ShouldBindJSON(&articleInfo)
	if err1 != nil {
		context.JSON(400, "Bad Parameters")
		return
	}

	token, _ := context.Cookie("token")
	username, _ := auth.GetTokenUsername(token)

	err2 := articleManageController.articleManageService.UpdateArticleTitleOrContentByID(articleInfo, username)
	if err2 != nil {
		if strings.Contains(err2.Error(), "400") {
			context.JSON(400, "Not Found")
		} else {
			context.JSON(500, "Server Internal Error")
		}
		return
	}

	context.JSON(200, "Update Successfully")
}

// GetOneArticleByID godoc
// @Summary Get One Article By ID
// @Description need token in cookie, need ID
// @Tags Article Manage
// @Accept json
// @Produce json
// @Security ApiAuthToken
// @Param id query integer true "Article ID"
// @Success 200 {object} entity.ArticleDetail "<b>Success</b>. Get Successfully"
// @Failure 400 {string} string "<b>Failure</b>. Bad Parameters / Not Found"
// @Failure 500 {string} string "<b>Failure</b>. Server Internal Error."
// @Router /article/getone/:id [get]
func (articleManageController ArticleManageController) GetOneArticleByID(context *gin.Context) {
	//id, err1 := strconv.Atoi(context.Param("id"))
	//if err1 != nil {
	//	context.JSON(400, "Bad Parameters")
	//	return
	//}

}
