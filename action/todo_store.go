package action

import (
	"net/http"
	"strconv"

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

func (a *Todo) UpdateStatus(c echo.Context) error {
	id := c.Param("id")
	param := c.QueryParam("is_done")

	if len(param) == 0 {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
	}

	isDone, err := strconv.ParseBool(param)
	if err != nil {
		return errors.BadRequest(c, stdErr.New("param only accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False"))
	}

	if !a.isExist(id) {
		return errors.NotFound(c, stdErr.New("todo not found"))
	}

	err = a.todoRepo.UpdateStatus(id, isDone)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Todo update status success"})
}
