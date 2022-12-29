package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTag(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id=? and deleted_on = 0", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tag, nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []*Tag, err error) {
	err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Where("deleted_on = 0").Count(&count)
	return
}
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=? and deleted_on = 0", name).First(&tag)
	return tag.ID > 0
}
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id=? and deleted_on = 0", id).First(&tag)
	return tag.ID > 0
}
func AddTag(tag *Tag) (bool, error) {
	err := db.Create(tag).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func EditTag(tag *Tag) (bool, error) {
	err := db.Model(tag).Updates(tag).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func DeleteTag(id int) (bool, error) {
	err := db.Where("id = ? and deleted_on = 0", id).Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
