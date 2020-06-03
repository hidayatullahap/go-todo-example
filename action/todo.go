package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	env *core.Env
}

func (a *Todo) FindTodoList(c echo.Context) error {
	return c.String(http.StatusOK, "Todo list goes here")
}

func NewTodo(env *core.Env) *Todo {
	return &Todo{
		env: env,
	}
}
