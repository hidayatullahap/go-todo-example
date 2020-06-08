package action

import (
	"net/http"

	stdErr "errors"

	"github.com/hidayatullahap/go-todo-example/core/errors"
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
