package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	db string
}

func (a *Todo) FindTodoList(c echo.Context) error {
	return c.String(http.StatusOK, "Todo list goes here")
}

func NewTodo() *Todo {
	return &Todo{
		db: "asd",
	}
}
