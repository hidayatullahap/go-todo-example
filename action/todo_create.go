package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/labstack/echo/v4"
)

func (a *Todo) CreateTodo(c echo.Context) error {
	todo, err := a.getTodoModelCreate(c)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = a.todoRepo.Create(todo)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo successfully created"})
}

func (a *Todo) getTodoModelCreate(c echo.Context) (model.Todo, error) {
	req := new(request.TodoCreateRequest)
	todo := model.Todo{}

	err := c.Bind(req)
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

	return todo, err
}
