package action

import (
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
)

type Todo struct {
	env      *core.Env
	todoRepo *repo.TodoRepo
}

func NewTodo(env *core.Env) *Todo {
	return &Todo{
		env:      env,
		todoRepo: repo.NewTodoRepo(env.Db),
	}
}
