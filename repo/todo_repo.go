package repo

import (
	"github.com/hidayatullahap/go-todo-example/core"
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
	if err != nil {
		return
	}

	// Populate tag ids in todo_tags
	todoTags, err := r.FindTodoTags(id)
	if err != nil {
		return
	}

	tagIds := r.PluckTodoTagsId(todoTags)

	var tags []model.Tag
	if len(todoTags) > 0 {
		tagRepo := NewTagRepo(r.db)

		// skip error on not found row
		tagsFind, _ := tagRepo.FindTagsIn(tagIds)
		tags = tagsFind
	}

	todo.Tags = &tags
	return
}

func (r *TodoRepo) FindTodoTags(todoID string) ([]model.TodoTag, error) {
	var todoTags []model.TodoTag
	err := r.db.Where("todo_id = ?", todoID).Find(&todoTags).Error
	if err != nil {
		return todoTags, err
	}

	return todoTags, nil
}

func (r *TodoRepo) PluckTodoTagsId(todoTags []model.TodoTag) (ids []string) {
	for _, tag := range todoTags {
		ids = append(ids, core.Int32ToString(tag.TagID))
	}

	return
}

func (r *TodoRepo) Create(todo model.Todo) (err error) {
	tx := r.db.Begin()

	err = tx.Create(&todo).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if todo.TodoTags != nil {
		for _, tag := range *todo.TodoTags {
			todoID := todo.ID
			tag.TodoID = todoID

			err = tx.Create(&tag).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}

	tx.Commit()
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
