package action

import (
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
)

type Todo struct {
	app      *core.App
	todoRepo *repo.TodoRepo
}

func (a *Todo) isExist(id string) bool {
	_, err := a.todoRepo.FindOne(id)
	if err != nil {
		return false
	}

	return true
}

func NewTodo(app *core.App) *Todo {
	return &Todo{
		app:      app,
		todoRepo: repo.NewTodoRepo(app.Db),
	}
}
