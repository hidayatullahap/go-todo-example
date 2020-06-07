package repo

import (
	"fmt"
	"strings"

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

// Using map for update nil/non-zero values, gorm will skip those values
func (r *TodoRepo) Update(todoID string, todo model.Todo) (err error) {
	updateField := map[string]interface{}{
		"message":     todo.Message,
		"note":        todo.Note,
		"custom_date": todo.CustomDate}

	tx := r.db.Begin()
	err = tx.Model(&todo).Where("id = ?", todoID).Updates(updateField).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if todo.TodoTags != nil {
		err = r.UpdateTodoTags(tx, todoID, *todo.TodoTags)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	return
}

// Update many-to-many relationship; upsert tags then delete unrelated tags id(s)
// example sql:
// INSERT IGNORE INTO `todo_tags` ( `todo_id`, `tag_id`) VALUES ( 1, 1 ), (1, 5), (1, 6);
// DELETE a FROM todo_tags a inner join todo_tags b ON a.id = b.id AND a.todo_id = 1 AND a.tag_id NOT IN (1,5,6);
func (r *TodoRepo) UpdateTodoTags(tx *gorm.DB, todoID string, tags []model.TodoTag) (err error) {
	if len(tags) == 0 {
		return
	}

	var tagIds []string
	for _, tag := range tags {
		tagIds = append(tagIds, core.Int32ToString(tag.TagID))
	}

	var todoTagIds []string
	for _, tagID := range tagIds {
		todoTagIds = append(todoTagIds, fmt.Sprintf("( %s,%s )", todoID, tagID))
	}

	insertValues := strings.Join(todoTagIds, ",")

	err = tx.Exec(`INSERT IGNORE INTO todo_tags ( todo_id, tag_id) VALUES ` + insertValues).Error
	if err != nil {
		return
	}

	err = tx.Exec(`DELETE a FROM todo_tags a inner join todo_tags b ON a.id = b.id AND a.todo_id = (?) AND a.tag_id NOT IN (?)`, todoID, tagIds).Error
	if err != nil {
		return
	}

	return
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}
