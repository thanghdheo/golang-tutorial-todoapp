package todostorage

import (
	"context"
	todomodel "todo-app/module/model"
)

func (store *sqlStorage) UpdateTodoAsDelete(ctx context.Context, cond map[string]interface{}, data *todomodel.TodoUpdate) error {

	if err := store.db.Where(cond).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
