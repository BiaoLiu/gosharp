package article_serializer

import (
	"fmt"
	"gopkg.in/jeevatkm/go-model.v1"
	"gosharp/library/type"
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
	Model  *models.Article
	Models []*models.Article
}

func (s *ArticleSerializer) SingleResponse() *ArticleResponse {
	response := &ArticleResponse{}
	err := model.Copy(response, s.Model)
	fmt.Println(err)
	response.ImageUrl = utils.FormatUrl(s.Model.ImageUrl)
	response.CreatedAt = s.Model.CreatedAt.Unix()
	return response
}

func (s *ArticleSerializer) ListResponse() []*ArticleResponse {
	var response []*ArticleResponse
	if len(s.Models) == 0 {
		return make([]*ArticleResponse, 0)
	}
	for i := range s.Models {
		item := s.Models[i]
		serializer := ArticleSerializer{Model: item}
		response = append(response, serializer.SingleResponse())
	}
	return response
}
