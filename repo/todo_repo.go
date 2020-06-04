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

func (r *TodoRepo) FindOne(id string) (todo model.Todo, err error) {
	err = r.db.Where("id = ?", id).First(&todo).Error
	return
}

func (r *TodoRepo) Create(todo model.Todo) (err error) {
	err = r.db.Create(&todo).Error
	return
}

// Using map for update nil values, struct will skip those values
func (r *TodoRepo) Update(id string, todo model.Todo) (err error) {
	updateField := map[string]interface{}{
		"message":     todo.Message,
		"note":        todo.Note,
		"custom_date": todo.CustomDate}

	err = r.db.Model(&todo).Where("id = ?", id).Updates(updateField).Error
	return
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}
