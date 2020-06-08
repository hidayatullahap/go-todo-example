package repo

import (
	"fmt"
	"strings"

	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/jinzhu/gorm"
	col "github.com/thoas/go-funk"
)

type TodoRepo struct {
	db *gorm.DB
}

func (r *TodoRepo) FindAll() (todos []model.Todo, err error) {
	err = r.db.Order("created_at").Find(&todos).Error
	if err != nil {
		return
	}

	err = r.MapTagsToTodo(&todos)
	if err != nil {
		return
	}

	return
}

func (r *TodoRepo) FindOne(id string) (todo model.Todo, err error) {
	err = r.db.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return
	}

	// find many-to-many tag relation
	tags, err := r.FindTodoTags(id)
	if err != nil {
		return
	}

	todo.Tags = &tags
	return
}

func (r *TodoRepo) FindTodoTags(todoID string) ([]model.Tag, error) {
	var todoTags []model.TodoTag
	var tags []model.Tag

	err := r.db.Where("todo_id = ?", todoID).Find(&todoTags).Error
	if err != nil {
		return tags, err
	}

	tagIds := r.PluckTodoTagsId(todoTags)

	if len(todoTags) > 0 {
		tagRepo := NewTagRepo(r.db)

		// skip error on not found row
		tagsFind, _ := tagRepo.FindTagsIn(tagIds)
		tags = tagsFind
	}

	return tags, nil
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
		"message": todo.Message,
		"note":    todo.Note}

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

func (r *TodoRepo) MapTagsToTodo(todos *[]model.Todo) (err error) {
	if todos == nil {
		return
	}

	var todoIds []int32
	for _, todo := range *todos {
		todoIds = append(todoIds, todo.ID)
	}

	// get tag ids then map it into todos
	var todoTags []model.TodoTag
	var tagIds []int32
	err = r.db.Where(" todo_id IN (?)", todoIds).Find(&todoTags).Error
	if err != nil {
		return
	}

	for _, todoTag := range todoTags {
		tagIds = append(tagIds, todoTag.TagID)
	}

	tagIds = col.UniqInt32(tagIds)

	var tags []model.Tag
	err = r.db.Where("id IN (?)", tagIds).Find(&tags).Error
	if err != nil {
		return
	}

	// map key for tag_id
	var mapTag = make(map[int32]model.Tag)
	for _, tag := range tags {
		mapTag[tag.ID] = tag
	}

	// map tags to todo
	mapTodos := *todos
	for i, todo := range mapTodos {
		var tmpTags []model.Tag
		mapTodos[i].Tags = &tmpTags

		// search tag_id key in todo_tags from map tag
		for _, todoTag := range todoTags {
			if todo.ID == todoTag.TodoID {
				if tag, ok := mapTag[todoTag.TagID]; ok {
					tmpTags = append(tmpTags, tag)
				}
			}
		}
	}

	return
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		db: db,
	}
}
