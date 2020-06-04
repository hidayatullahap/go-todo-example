package repo

import (
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/jinzhu/gorm"
)

type TagRepo struct {
	db *gorm.DB
}

func (r *TagRepo) FindAll() (tags []model.Tag, err error) {
	err = r.db.Find(&tags).Error
	return
}

func (r *TagRepo) FindOne(id string) (tag model.Tag, err error) {
	err = r.db.Where("id = ?", id).First(&tag).Error
	return
}

func (r *TagRepo) FindTagsIn(tagIDs []string) (tags []model.Tag, err error) {
	if len(tagIDs) > 0 {
		err = r.db.Where("id IN (?)", tagIDs).Find(&tags).Error
	}

	return
}

func (r *TagRepo) Create(tag model.Tag) (err error) {
	err = r.db.Create(&tag).Error
	return
}

func (r *TagRepo) Update(id string, tag model.Tag) (err error) {
	err = r.db.Model(&tag).Where("id = ?", id).Updates(&tag).Error
	return
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{
		db: db,
	}
}
