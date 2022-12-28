package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	CoverImageUrl string `json:"cover_image_url"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id=? and deleted_on = 0", id).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []*Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	return
}

func GetArticle(id int) (*Article, error) {
	var article Article

	err := db.Where("id=? and deleted_on = 0", id).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=? and deleted_on = 0", id).Updates(data)
	return true
}

func AddArticle(article *Article) (bool, error) {
	err := db.Create(article).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteArticle(id int) bool {
	db.Where("id=? and deleted_on = 0", id).Delete(Article{})
	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
