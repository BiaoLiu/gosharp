package http

import (
	"github.com/gin-gonic/gin"
	"gosharp/internal/model/article"
	"gosharp/internal/model/common"
	"gosharp/library/app"
	"gosharp/library/type"
)

// @Summary 文章列表
// @Description 文章列表
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param title query string false "标题"
// @Param created_start query string false "创建开始时间"
// @Param created_end query string false "创建结束时间"
// @Param offset query int false "偏移量"
// @Param limit query int false "页容量"
// @Success 200 {array} article.ArticleResp
// @Router /articles [get]
func ArticleList(c *gin.Context) {
	title := c.Query("title")
	createdStart := c.Query("created_start")
	createdEnd := c.Query("created_end")

	offset := utils.SafeInt(c.Query("offset"), 0)
	limit := utils.SafeInt(c.Query("limit"), 15)

	articles, total := svc.GetArticleList(title, createdStart, createdEnd, offset, limit)

	serializer := article.ArticleSerializer{Models: articles}
	//设置header
	app.SetPagerHeader(c, offset, limit, int(total))
	app.APIResponse(c, true, serializer.ListResponse(), "")
}

// @Summary 所有文章列表
// @Description 所有文章列表
// @Tags 文章
// @Accept  json
// @Produce  json
// @Success 200 {array} article.ArticleResp
// @Router /articles/all [get]
func ArticleAllList(c *gin.Context) {
	articles := svc.GetAllArticleList()

	serializer := article.ArticleSerializer{Models: articles}

	app.APIResponse(c, true, serializer.ListResponse(), "")
}

// @Summary 文章详情
// @Description 文章详情
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {object} article.ArticleResp
// @Router /articles/{id} [get]
func ArticleRetrieve(c *gin.Context) {
	id := utils.SafeInt(c.Param("id"), 0)

	art := svc.GetArticleById(id)
	if art == nil {
		app.APIResponse(c, false, nil, "不存在。")
		return
	}

	serializer := article.ArticleSerializer{Model: art}

	app.APIResponse(c, true, serializer.SingleResponse(), "")
}

// @Summary 新增文章
// @Description 新增文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param account body article.ArticleReq true "文章"
// @Success 200
// @Router /articles [post]
func ArticleCreate(c *gin.Context) {
	arg := new(article.ArticleReq)
	if err := app.BindAndValidate(c, arg); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	r, err := svc.CreateArticle(arg)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, gin.H{"id": r.Id}, "")
}

// @Summary 修改文章
// @Description 修改文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param account body article.ArticleReq true "文章"
// @Success 200
// @Router /articles [put]
func ArticleUpdate(c *gin.Context) {
	id := utils.SafeInt(c.Param("id"), 0)

	arg := new(article.ArticleReq)
	if err := app.BindAndValidate(c, arg); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	arg.Article = svc.GetArticleById(id)
	if arg.Article == nil {
		app.APIResponse(c, false, nil, "不存在。")
		return
	}

	err := svc.UpdateArticle(arg)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, nil, "")
}

// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param account body common.IdsReq true "文章id"
// @Success 200
// @Router /articles [delete]
func ArticleDelete(c *gin.Context) {
	arg := new(common.IdsReq)
	if err := app.BindAndValidate(c, arg); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	i, err := svc.DeleteArticle(arg.Ids)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, gin.H{"count": i}, "")
}
