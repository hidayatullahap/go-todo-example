package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core/errors"
	"github.com/hidayatullahap/go-todo-example/model"
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

func (a *Tag) buildTagModel(c echo.Context) (model.Tag, error) {
	req := new(request.TagCreateRequest)
	tag := model.Tag{}

	err := c.Bind(req)
	if err != nil {
		return tag, err
	}

	err = c.Validate(req)
	if err != nil {
		return tag, err
	}

	tag.Name = req.Name

	return tag, err
}
