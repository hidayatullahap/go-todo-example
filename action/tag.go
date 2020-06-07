package action

import (
	"github.com/hidayatullahap/go-todo-example/core"
	"github.com/hidayatullahap/go-todo-example/repo"
)

type Tag struct {
	app     *core.App
	tagRepo *repo.TagRepo
}

func NewTag(app *core.App) *Tag {
	return &Tag{
		app:     app,
		tagRepo: repo.NewTagRepo(app.Db),
	}
}
