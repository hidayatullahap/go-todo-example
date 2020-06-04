package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/labstack/echo/v4"
)

func (a *Tag) FindList(c echo.Context) error {
	list, err := a.tagRepo.FindAll()
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, list)
}

func (a *Tag) FindDetail(c echo.Context) error {
	id := c.Param("id")

	detail, err := a.tagRepo.FindOne(id)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, detail)
}
