package article_service

import (
	"encoding/json"
	"github.com/GoldenLeeK/go-gin-blog/models"
	"github.com/GoldenLeeK/go-gin-blog/pkg/e"
	"github.com/GoldenLeeK/go-gin-blog/pkg/gredis"
	"github.com/GoldenLeeK/go-gin-blog/pkg/logging"
	"github.com/GoldenLeeK/go-gin-blog/service/cache_service"
)

type Article struct {
	ID       int
	PageSize int
	PageNum  int
	Maps     map[string]interface{}
	Article  *models.Article
}

func (a *Article) ExistByID() (bool, error) {
	exists := models.ExistArticleById(a.ID)
	return exists, nil
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, article, 3600)
	return article, nil

}

func (a *Article) GetAll() ([]*models.Article, error) {
	var cacheArticles = []*models.Article{}

	var tagId = -1
	var state = -1
	if value, ok := a.Maps["tag_id"]; ok {
		tagId = value.(int)
	}
	if value, ok := a.Maps["state"]; ok {
		state = value.(int)
	}

	cache := cache_service.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}

	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.Maps)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, articles, 3600)
	return articles, nil

}

func (a *Article) Total() int {
	return models.GetArticleTotal(a.Maps)
}

func (a *Article) Add() (bool, error) {
	err := gredis.LikeDeletes(e.CACHE_ARTICLE)
	if err != nil {
		return false, err
	}
	return models.AddArticle(a.Article)
}
