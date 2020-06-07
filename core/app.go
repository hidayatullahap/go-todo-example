package core

import "github.com/jinzhu/gorm"

type App struct {
	Db *gorm.DB
}
