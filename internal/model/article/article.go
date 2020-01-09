package article

import (
	"fmt"
	"gopkg.in/jeevatkm/go-model.v1"
	utils "gosharp/library/type"
	"time"
)

type Article struct {
	Id        int `gorm:"primary_key;auto_increment"`
	Title     string
	ImageUrl  string
	Content   string
	Sort      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ArticleReq struct {
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
	Article *Article
}

type ArticleResp struct {
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
	Model  *Article
	Models []*Article
}

func (s *ArticleSerializer) SingleResponse() *ArticleResp {
	response := &ArticleResp{}
	err := model.Copy(response, s.Model)
	fmt.Println(err)
	response.ImageUrl = utils.FormatUrl(s.Model.ImageUrl)
	response.CreatedAt = s.Model.CreatedAt.Unix()
	return response
}

func (s *ArticleSerializer) ListResponse() []*ArticleResp {
	var response []*ArticleResp
	if len(s.Models) == 0 {
		return make([]*ArticleResp, 0)
	}
	for i := range s.Models {
		item := s.Models[i]
		serializer := ArticleSerializer{Model: item}
		response = append(response, serializer.SingleResponse())
	}
	return response
}
