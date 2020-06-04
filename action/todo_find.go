package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/labstack/echo/v4"
)

func (a *Todo) FindTodoList(c echo.Context) error {
	list, err := a.todoRepo.FindAll()
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, list)
}

func (a *Todo) FindTodoDetail(c echo.Context) error {
	id := c.Param("id")

	detail, err := a.todoRepo.FindOne(id)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, detail)
}
