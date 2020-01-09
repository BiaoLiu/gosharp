package service

import (
	"fmt"
	"gosharp/internal/model/article"
	"gosharp/library/database/orm"
	"gosharp/library/log"
	"gosharp/library/type"
)

func (s *Service) GetArticleList(title, createdStart, createdEnd string, offset, limit int) ([]*article.Article, int) {
	query := s.dao.DB().Model(&article.Article{})
	//查询条件
	query = orm.WhereIgnoreBlank(query, "title = ? ", title)
	query = orm.WhereIgnoreBlank(query, "created_at >= ? ", createdStart)
	query = orm.WhereIgnoreBlank(query, "created_at < ? ", utils.CastEndTime(createdEnd))

	var articles []*article.Article
	total, _ := orm.PaginateWithCount(query, "sort, created_at desc", offset, limit, &articles)
	return articles, total
}

func (s *Service) GetAllArticleList() []*article.Article {
	var articles []*article.Article
	s.dao.DB().Order("created_at desc").Find(&articles)
	return articles
}

func (s *Service) GetArticleById(ArticleId int) *article.Article {
	article := new(article.Article)
	s.dao.DB().Where("id = ?", ArticleId).First(article)
	return article
}

func (s *Service) CreateArticle(arg *article.ArticleReq) (*article.Article, error) {
	article := &article.Article{
		Title:    arg.Title,
		ImageUrl: utils.ParseUrl(arg.ImageUrl),
		Content:  arg.Content,
		Sort:     arg.Sort,
	}

	if err := s.dao.DB().Create(article).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("CreateArticle error: %s", err))
		return nil, err
	}
	return article, nil
}

func (s *Service) UpdateArticle(arg *article.ArticleReq) error {
	arg.ImageUrl = utils.ParseUrl(arg.ImageUrl)

	attrs := utils.Struct2Map(*arg)
	if err := s.dao.DB().Model(&article.Article{}).Update(attrs).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("UpdateArticle error: %s", err))
		return err
	}
	return nil
}

func (s *Service) DeleteArticle(ids []int) (int, error) {
	if err := s.dao.DB().Where("id in (?)", ids).Delete(&article.Article{}).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("DeleteArticle error: %s", err))
		return 0, err
	}
	return len(ids), nil
}
