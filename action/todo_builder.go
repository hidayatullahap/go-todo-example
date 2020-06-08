package action

import (
	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/labstack/echo/v4"
)

func (a *Todo) buildTodoModel(c echo.Context) (model.Todo, error) {
	req := new(request.TodoCreateRequest)
	todo := model.Todo{}

	err := core.BindValidate(c, req)
	if err != nil {
		return todo, err
	}

	todo.Message = req.Message
	todo.Note = req.Note

	if req.Tags != nil {
		tags := buildTodoTagsModel(req)
		todo.TodoTags = &tags
	}

	return todo, err
}

func buildTodoTagsModel(req *request.TodoCreateRequest) []model.TodoTag {
	var tags []model.TodoTag
	for _, tagID := range *req.Tags {
		tag := model.TodoTag{
			TagID: tagID,
		}

		tags = append(tags, tag)
	}

	return tags
}
