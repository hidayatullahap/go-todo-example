package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *Todo) FindTodoList(c echo.Context) error {
	list, err := a.todoRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, list)
}
