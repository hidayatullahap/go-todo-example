package action

import (
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
)

type Tag struct {
	env     *core.Env
	tagRepo *repo.TagRepo
}

func NewTag(env *core.Env) *Tag {
	return &Tag{
		env:     env,
		tagRepo: repo.NewTagRepo(env.Db),
	}
}
