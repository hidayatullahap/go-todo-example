package action

import (
	"net/http"

	stdErr "errors"

	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/labstack/echo/v4"
)

func (a *Todo) Create(c echo.Context) error {
	todo, err := a.buildTodoModel(c)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = a.todoRepo.Create(todo)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Todo successfully created"})
}

func (a *Todo) Update(c echo.Context) error {
	id := c.Param("id")

	if !a.isExist(id) {
		return errors.NotFound(c, stdErr.New("todo not found"))
	}

	todo, err := a.buildTodoModel(c)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = a.todoRepo.Update(id, todo)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo successfully updated"})
}

func (a *Todo) buildTodoModel(c echo.Context) (model.Todo, error) {
	req := new(request.TodoCreateRequest)
	todo := model.Todo{}

	err := c.Bind(req)
	if err != nil {
		return todo, err
	}

	err = c.Validate(req)
	if err != nil {
		return todo, err
	}

	todo.Message = req.Message
	todo.Note = req.Note

	if req.CustomDate != nil {
		customDate, err := core.FormatDate(*req.CustomDate)
		if err != nil {
			return todo, err
		}

		todo.CustomDate = &customDate
	}

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
