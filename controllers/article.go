package controllers

import (
	"github.com/gin-gonic/gin"
	"gosharp/forms/article_form"
	"gosharp/serializers/article_serializer"
	"gosharp/services/article_service"
	"gosharp/utils/app"
	string_util "gosharp/utils/string"
)

// swagger:operation GET /articles Article 文章列表
// ---
// summary: 文章列表
// description: 文章列表
// parameters:
// - name: title
//   in: query
//   type: string
//   description: 标题
// - name: created_start
//   in: query
//   type: string
//   description: 创建开始时间 如：2018-12-05
// - name: created_end
//   in: query
//   type: string
//   description: 创建结束时间 如：2018-12-05
// - name: offset
//   in: query
//   type: number
// - name: limit
//   in: query
//   type: number
// responses:
//     "200":
//       "$ref": "#/responses/ArticleResponseWrap"
func ArticleList(c *gin.Context) {
	title := c.Query("title")
	createdStart := c.Query("created_start")
	createdEnd := c.Query("created_end")

	offset := string_util.SafeInt(c.Query("offset"), 0)
	limit := string_util.SafeInt(c.Query("limit"), 15)

	articles, total := article_service.GetArticleList(title, createdStart, createdEnd, offset, limit)

	serializer := article_serializer.ArticlesSerializer{articles}
	//设置header
	app.SetPagerHeader(c, offset, limit, int(total))
	app.APIResponse(c, true, serializer.Response(), "")

}

// swagger:operation GET /articles/all Article 所有文章列表
// ---
// summary: 所有文章列表
// description: 所有文章列表
// responses:
//     "200":
//       "$ref": "#/responses/ArticleResponseWrap"
func ArticleAllList(c *gin.Context) {
	articles := article_service.GetAllArticleList()

	serializer := article_serializer.ArticlesSerializer{articles}

	app.APIResponse(c, true, serializer.Response(), "")
}

// swagger:operation GET /articles/{id} Article 文章详情
// ---
// summary: 文章详情
// description: 文章详情
// parameters:
// - name: id
//   in: path
//   description: 文章ID
//   type: number
//   required: true
// responses:
//     "200":
//       "$ref": "#/responses/ArticleResponseWrap"
func ArticleRetrieve(c *gin.Context) {
	id := string_util.SafeInt(c.Param("id"), 0)

	article := article_service.GetArticleById(id)
	if article == nil {
		app.APIResponse(c, false, nil, "不存在。")
		return
	}

	serializer := article_serializer.ArticleSerializer{article}

	app.APIResponse(c, true, serializer.Response(), "")
}

// swagger:operation POST /articles Article 新增文章
// ---
// summary: 新增文章
// description: 新增文章
// parameters:
// - name: Body
//   in: body
//   schema:
//       "$ref": "#/definitions/ArticleForm"
// responses:
//   "200":
//       "$ref": ""
func ArticleCreate(c *gin.Context) {
	var form article_form.ArticleForm
	if err := app.BindAndValidate(c, &form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	r, err := article_service.CreateArticle(&form)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, gin.H{"id": r.Id}, "")
}

// swagger:operation PUT /articles/{id} Article 修改文章
// ---
// summary: 修改文章
// description: 修改文章
// parameters:
// - name: id
//   in: path
//   description: 文章ID
//   type: number
//   required: true
// - name: Body
//   in: body
//   schema:
//       "$ref": "#/definitions/ArticleForm"
// responses:
//   "200":
//       "$ref": ""
func ArticleUpdate(c *gin.Context) {
	id := string_util.SafeInt(c.Param("id"), 0)

	var form article_form.ArticleForm
	if err := app.BindAndValidate(c, &form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	form.Article = article_service.GetArticleById(id)
	if form.Article == nil {
		app.APIResponse(c, false, nil, "不存在。")
		return
	}

	err := article_service.UpdateArticle(&form)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, nil, "")
}

// swagger:operation DELETE /articles Article 删除文章
// ---
// summary: 删除文章
// description: 删除文章
// parameters:
// - name: Body
//   in: body
//   schema:
//       "$ref": "#/definitions/IdsForm"
// responses:
//   "200":
//       "$ref": ""
func ArticleDelete(c *gin.Context) {
	var form article_form.IdsForm
	if err := app.BindAndValidate(c, &form); err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}

	i, err := article_service.DeleteArticle(form.Ids)
	if err != nil {
		app.APIResponse(c, false, nil, err.Error())
		return
	}
	app.APIResponse(c, true, gin.H{"count": i}, "")
}
