package action

import (
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
)

type Todo struct {
	app      *core.App
	todoRepo *repo.TodoRepo
}

func NewTodo(app *core.App) *Todo {
	return &Todo{
		app:      app,
		todoRepo: repo.NewTodoRepo(app.Db),
	}
}
