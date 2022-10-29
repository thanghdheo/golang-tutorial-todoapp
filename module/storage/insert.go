package todostorage

import (
	"context"
	todomodel "todo-app/module/model"
)

func (store *sqlStorage) CreateTodo(ctx context.Context, data *todomodel.TodoCreate) error{
	if err := store.db.Create(data).Error; err != nil{
		return err	
	}

	return nil
}