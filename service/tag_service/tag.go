package tag_service

import (
	"encoding/json"
	"github.com/GoldenLeeK/go-gin-blog/models"
	"github.com/GoldenLeeK/go-gin-blog/pkg/e"
	"github.com/GoldenLeeK/go-gin-blog/pkg/gredis"
	"github.com/GoldenLeeK/go-gin-blog/pkg/logging"
	"github.com/GoldenLeeK/go-gin-blog/service/cache_service"
)

type Tag struct {
	ID       int
	PageSize int
	PageNum  int
	Maps     map[string]interface{}
	Tag      *models.Tag
}

func (t *Tag) ExistByID() (bool, error) {
	exists := models.ExistTagById(t.ID)
	return exists, nil
}

func (t *Tag) ExistByName() (bool, error) {
	exists := models.ExistTagByName(t.Tag.Name)
	return exists, nil
}

func (t *Tag) Get() (*models.Tag, error) {
	var cacheTag *models.Tag
	cache := cache_service.Tag{ID: t.ID}
	key := cache.GetTagKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTag)
			return cacheTag, nil
		}
	}

	tag, err := models.GetTag(t.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tag, 3600)
	return tag, nil

}

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var cacheTags = []*models.Tag{}

	var name = ""
	var state = -1
	if value, ok := t.Maps["name"]; ok {
		name = value.(string)
	}
	if value, ok := t.Maps["state"]; ok {
		state = value.(int)
	}

	cache := cache_service.Tag{
		State:    state,
		Name:     name,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}

	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.Maps)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, tags, 3600)
	return tags, nil

}

func (t *Tag) Total() int {
	return models.GetTagTotal(t.Maps)
}

func (t *Tag) Add() (bool, error) {
	err := gredis.LikeDeletes(e.CACHE_TAG)
	if err != nil {
		return false, err
	}
	return models.AddTag(t.Tag)
}

func (t *Tag) Update() (bool, error) {
	err := gredis.LikeDeletes(e.CACHE_TAG)
	if err != nil {
		return false, err
	}
	return models.EditTag(t.Tag)
}

func (t *Tag) Delete() (bool, error) {
	err := gredis.LikeDeletes(e.CACHE_ARTICLE)
	if err != nil {
		return false, err
	}
	return models.DeleteTag(t.ID)
}
