package article_form

import (
	"gosharp/library/app"
	"gosharp/models"
)

type ArticleForm struct {
	app.BaseForm

	//标题
	// required: true
	Title string `json:"title" valid:"Required;MaxSize(50)"`
	//主图url
	// required: true
	ImageUrl string `json:"image_url" valid:"Required"`
	//内容
	// required: true
	Content string `json:"content" valid:"Required"`
	//排序号
	// required: true
	Sort int `json:"sort" valid:"Required;Min(1)"`

	//swagger:ignore
	Article *models.Article
}

type IdsForm struct {
	app.BaseForm

	Ids []int `json:"ids"`
}
