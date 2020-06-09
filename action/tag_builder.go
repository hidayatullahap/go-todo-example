package action

import (
	"github.com/hidayatullahap/go-todo-example/action/request"
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/model"
	"github.com/labstack/echo/v4"
)

func (a *Tag) buildTagModel(c echo.Context) (model.Tag, error) {
	req := new(request.TagCreateRequest)
	tag := model.Tag{}

	err := core.BindValidate(c, req)
	if err != nil {
		return tag, err
	}

	tag.Name = req.Name
	return tag, err
}
