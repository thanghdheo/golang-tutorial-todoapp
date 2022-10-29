package todostorage

import "gorm.io/gorm"

type sqlStorage struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
