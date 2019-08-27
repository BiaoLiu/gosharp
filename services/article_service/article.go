package article_service

import (
	"fmt"
	"gosharp/forms/article_form"
	"gosharp/library/db"
	"gosharp/library/log"
	string_util "gosharp/library/string"
	"gosharp/models"
)

func GetArticleList(title, createdStart, createdEnd string, offset, limit int) ([]*models.Article, int) {
	query := db.Gorm.Model(&models.Article{})
	//查询条件
	query = db.WhereIgnoreBlank(query, "title = ? ", title)
	query = db.WhereIgnoreBlank(query, "created_at >= ? ", createdStart)
	query = db.WhereIgnoreBlank(query, "created_at < ? ", string_util.CastEndTime(createdEnd))

	var articles []*models.Article
	total, _ := db.PaginateWithCount(query, "sort, created_at desc", offset, limit, &articles)
	return articles, total
}

func GetAllArticleList() []*models.Article {
	var articles []*models.Article
	db.Gorm.Order("created_at desc").Find(&articles)
	return articles
}

func GetArticleById(ArticleId int) *models.Article {
	article := new(models.Article)
	db.Gorm.Where("id = ?", ArticleId).First(article)
	return article
}

func CreateArticle(form *article_form.ArticleForm) (*models.Article, error) {
	article := &models.Article{
		Title:    form.Title,
		ImageUrl: string_util.ParseUrl(form.ImageUrl),
		Content:  form.Content,
		Sort:     form.Sort,
	}

	if err := db.Gorm.Create(article).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("CreateArticle error: %s", err))
		return nil, err
	}
	return article, nil
}

func UpdateArticle(form *article_form.ArticleForm) error {
	form.ImageUrl = string_util.ParseUrl(form.ImageUrl)

	attrs := string_util.Struct2Map(*form)
	if err := db.Gorm.Model(&models.Article{}).Update(attrs).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("UpdateArticle error: %s", err))
		return err
	}
	return nil
}

func DeleteArticle(ids []int) (int, error) {
	if err := db.Gorm.Where("id in (?)", ids).Delete(&models.Article{}).Error; err != nil {
		log.Logger.Error(fmt.Sprintf("DeleteArticle error: %s", err))
		return 0, err
	}
	return len(ids), nil
}
