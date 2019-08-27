package article_serializer

import (
	"fmt"
	"gopkg.in/jeevatkm/go-model.v1"
	string_util "gosharp/library/string"
	"gosharp/models"
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
	Article  *models.Article
	Articles []*models.Article
}

func (s *ArticleSerializer) SingleResponse() *ArticleResponse {
	response := &ArticleResponse{}
	err := model.Copy(response, s.Article)
	fmt.Println(err)
	response.ImageUrl = string_util.FormatUrl(s.Article.ImageUrl)
	response.CreatedAt = s.Article.CreatedAt.Unix()
	return response
}

func (s *ArticleSerializer) ListResponse() []*ArticleResponse {
	var response []*ArticleResponse
	if s.Articles == nil {
		return make([]*ArticleResponse, 0)
	}
	for i := range s.Articles {
		item := s.Articles[i]
		serializer := ArticleSerializer{Article: item}
		response = append(response, serializer.SingleResponse())
	}
	return response
}
