package todostorage

import (
	"context"
	todomodel "todo-app/module/model"
)

func (storage *sqlStorage) UpdateTodo(ctx context.Context, cond map[string]interface{}, data *todomodel.TodoUpdate) error{
	if err := storage.db.Where(cond).Updates(&data).Error; err != nil{
		return err
	}

	return nil
}