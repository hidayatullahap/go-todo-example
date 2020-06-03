package repo

import (
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/jinzhu/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func (r *TodoRepo) FindAll() (todos []model.Todo, err error) {
	err = r.db.Find(&todos).Error
	return
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}