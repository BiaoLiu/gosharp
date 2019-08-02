package article_serializer

import (
	"gosharp/models"
	string_util "gosharp/utils/string"
)

type ArticleResponse struct {
	Id int `json:"id"`
	//标题
	Title string `json:"title"`
	//主图
	ImageUrl string `json:"image_url"`
	//内容
	Content string `json:"content"`
	//排序号
	Sort int `json:"sort"`
	//创建时间
	CreatedAt int64 `json:"created_at"`
}

type ArticleSerializer struct {
	*models.Article
}

func (s ArticleSerializer) Response() ArticleResponse {
	return ArticleResponse{
		Id:        s.Id,
		Title:     s.Title,
		ImageUrl:  string_util.FormatUrl(s.ImageUrl),
		Content:   s.Content,
		Sort:      s.Sort,
		CreatedAt: s.CreatedAt.Unix(),
	}
}

type ArticlesSerializer struct {
	Articles []models.Article
}

func (s ArticlesSerializer) Response() []ArticleResponse {
	var response []ArticleResponse

	if s.Articles == nil || len(s.Articles) == 0 {
		return make([]ArticleResponse, 0)
	}

	for i, _ := range s.Articles {
		restaurant := &s.Articles[i]
		serializer := ArticleSerializer{restaurant}
		response = append(response, serializer.Response())
	}
	return response
}
