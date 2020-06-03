package core

import "github.com/jinzhu/gorm"

type Env struct {
	Db *gorm.DB
}
