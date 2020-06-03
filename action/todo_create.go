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
	req := new(request.TodoCreateRequest)

	err := c.Bind(req)
	if err != nil {
		return err
	}

	todo := model.Todo{
		Message: req.Message,
		Note:    req.Note,
	}

	if req.CustomDate != nil {
		customDate, err := core.FormatDate(*req.CustomDate)
		if err != nil {
			return errors.BadRequest(c, err)
		}

		todo.CustomDate = &customDate
	}

	return c.String(http.StatusOK, "Create Todo endpoint")
}
