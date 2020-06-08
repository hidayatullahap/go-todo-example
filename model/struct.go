package model

import "time"

type Todo struct {
	ID        int32      `gorm:"column:id" json:"id"`
	Message   string     `gorm:"column:message" json:"message"`
	Note      *string    `gorm:"column:note" json:"note"`
	IsDone    bool       `gorm:"column:is_done" json:"is_done"`
	TodoTags  *[]TodoTag `gorm:"-" json:"todo_tags,omitempty"`
	Tags      *[]Tag     `gorm:"-" json:"tags"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Tag struct {
	ID        int32      `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type TodoTag struct {
	ID     int32 `gorm:"column:id" json:"id"`
	TodoID int32 `gorm:"column:todo_id" json:"todo_id"`
	TagID  int32 `gorm:"column:tag_id" json:"tag_id"`
}
