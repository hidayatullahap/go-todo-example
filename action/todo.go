package action

import (
	"net/http"

	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
	"github.com/labstack/echo/v4"
)

type Todo struct {
	env      *core.Env
	todoRepo *repo.TodoRepo
}

func (a *Todo) FindTodoList(c echo.Context) error {
	list, err := a.todoRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, list)
}

func NewTodo(env *core.Env) *Todo {
	return &Todo{
		env:      env,
		todoRepo: repo.NewTodoRepo(env.Db),
	}
}
