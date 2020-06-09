package action

import (
	stdErr "errors"
	"net/http"

	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/labstack/echo/v4"
)

func (a *Tag) Create(c echo.Context) error {
	todo, err := a.buildTagModel(c)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = a.tagRepo.Create(todo)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Tag successfully created"})
}

func (a *Tag) Update(c echo.Context) error {
	id := c.Param("id")

	if !a.isExist(id) {
		return errors.NotFound(c, stdErr.New("tag not found"))
	}

	todo, err := a.buildTagModel(c)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	err = a.tagRepo.Update(id, todo)
	if err != nil {
		return errors.BadRequest(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Tag successfully updated"})
}
