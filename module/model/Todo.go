package todomodel

import (
	"errors"
	"strings"
	"time"
)

type Todo struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	Deleted   bool       `json:"deleted" gorm:"column:deleted"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Todo) TableName() string {
	return "todo"
}

type TodoCreate struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type TodoUpdate struct {
	Id        *int       `json:"id" gorm:"column:id;"`
	Title     *string    `json:"title" gorm:"column:title;"`
	Deleted   bool       `json:"deleted" gorm:"deleted;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (TodoCreate) TableName() string {
	return Todo{}.TableName()
}

func (TodoUpdate) TableName() string {
	return Todo{}.TableName()
}

func (res *TodoCreate) Validate() error {
	res.Id = 0
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Title can't be blank")
	}
	return nil
}
